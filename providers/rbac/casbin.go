package rbac

import (
	"atom/container"
	"atom/database/query"
	"atom/providers/log"
	"context"
	"errors"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(NewCasbin); err != nil {
		log.Fatal(err)
	}
}

type Casbin struct {
	query    *query.Query
	enforcer *casbin.CachedEnforcer
}

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

func NewCasbin(query *query.Query, db *gorm.DB) (IRbac, error) {
	cb := &Casbin{query: query}

	a, _ := gormadapter.NewAdapterByDB(db)
	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	m, err := model.NewModelFromString(text)
	if err != nil {
		log.Error(err, "字符串加载模型失败!")
		return nil, err
	}
	cb.enforcer, _ = casbin.NewCachedEnforcer(m, a)

	cb.enforcer.SetExpireTime(60 * 60)
	_ = cb.enforcer.LoadPolicy()
	return cb, nil
}

func (cb *Casbin) Can(role, method, path string) bool {
	return false
}

func (cb *Casbin) Reload() error {
	return nil
}

func (cb *Casbin) Update(roleID uint, infos []CasbinInfo) error {
	roleIdStr := strconv.Itoa(int(roleID))
	cb.Clear(0, roleIdStr)

	rules := [][]string{}
	for _, v := range infos {
		rules = append(rules, []string{roleIdStr, v.Path, v.Method})
	}

	success, _ := cb.enforcer.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}

	err := cb.enforcer.InvalidateCache()
	if err != nil {
		return err
	}

	return nil
}

func (cb *Casbin) UpdateApi(before, after CasbinInfo) error {
	rule := cb.query.CasbinRule

	_, err := rule.WithContext(context.Background()).
		Where(rule.V1.Eq(before.Path)).
		Where(rule.V1.Eq(before.Method)).
		UpdateSimple(
			rule.V1.Value(after.Path),
			rule.V2.Value(after.Method),
		)
	if err != nil {
		return err
	}

	return cb.enforcer.InvalidateCache()
}

// 获取权限列表
func (cb *Casbin) GetPolicyPathByRoleID(roleID uint) (pathMaps []CasbinInfo) {
	roleIdStr := strconv.Itoa(int(roleID))
	list := cb.enforcer.GetFilteredPolicy(0, roleIdStr)
	for _, v := range list {
		pathMaps = append(pathMaps, CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// 清除匹配的权限
func (cb *Casbin) Clear(v int, p ...string) bool {
	success, _ := cb.enforcer.RemoveFilteredPolicy(v, p...)
	return success
}
