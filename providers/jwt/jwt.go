package jwt

import (
	"atom/container"
	"atom/providers/config"
	"atom/providers/log"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/sync/singleflight"
)

const (
	CtxKey     = "claims"
	HttpHeader = "Authorization"
)

func init() {
	if err := container.Container.Provide(NewJWT); err != nil {
		log.Fatal(err)
	}
}

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

const TOKEN_PREFIX = "Bearer "

type BaseClaims struct {
	UUID     string
	UserID   uint64
	Username string
	NickName string
	RoleID   uint64
}

type JWT struct {
	config       *config.Config
	singleflight *singleflight.Group
	SigningKey   []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT(config *config.Config) (*JWT, error) {
	return &JWT{
		config:     config,
		SigningKey: []byte(config.Http.JWT.SigningKey),
	}, nil
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	bf, _ := time.ParseDuration(j.config.Http.JWT.BufferTime)
	ep, _ := time.ParseDuration(j.config.Http.JWT.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Second * 10)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),                // 过期时间 7天  配置文件
			Issuer:    j.config.Http.JWT.Issuer,                              // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := j.singleflight.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, TOKEN_PREFIX)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

func (j *JWT) GetClaims(c *gin.Context) (*CustomClaims, error) {
	token := c.Request.Header.Get(HttpHeader)
	claims, err := j.ParseToken(token)
	if err != nil {
		log.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在 Authorization 且 Claims 为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func (j *JWT) GetUserID(c *gin.Context) uint64 {
	if claims, exists := c.Get(CtxKey); !exists {
		if cl, err := j.GetClaims(c); err != nil {
			return 0
		} else {
			return cl.UserID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.UserID
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func (j *JWT) GetUserUuid(c *gin.Context) string {
	if claims, exists := c.Get(CtxKey); !exists {
		if cl, err := j.GetClaims(c); err != nil {
			return uuid.UUID{}.String()
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.UUID
	}
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func (j *JWT) GetRoleId(c *gin.Context) uint64 {
	if claims, exists := c.Get(CtxKey); !exists {
		if cl, err := j.GetClaims(c); err != nil {
			return 0
		} else {
			return cl.RoleID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.RoleID
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func (j *JWT) GetUserInfo(c *gin.Context) *CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := j.GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse
	}
}
