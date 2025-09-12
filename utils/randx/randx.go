package randx

import (
    crand "crypto/rand"
    "encoding/hex"
    "io"
)

// Bytes returns n cryptographically secure random bytes.
func Bytes(n int) ([]byte, error) {
    if n <= 0 {
        return []byte{}, nil
    }
    b := make([]byte, n)
    _, err := io.ReadFull(crand.Reader, b)
    return b, err
}

var urlSafe = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")

// String returns a URL-safe random string of length n using a cryptographically secure source.
func String(n int) (string, error) {
    if n <= 0 {
        return "", nil
    }
    out := make([]byte, n)
    // rejection sampling to avoid modulo bias
    // we want uniform selection from urlSafe (len m)
    m := len(urlSafe)
    threshold := 256 - (256 % m)
    var buf [1]byte
    for i := 0; i < n; i++ {
        for {
            if _, err := io.ReadFull(crand.Reader, buf[:]); err != nil {
                return "", err
            }
            b := int(buf[0])
            if b < threshold {
                out[i] = urlSafe[b%m]
                break
            }
        }
    }
    return string(out), nil
}

// MustString is like String but panics on error.
func MustString(n int) string {
    s, err := String(n)
    if err != nil {
        panic(err)
    }
    return s
}

// UUIDv4 generates a RFC 4122 version 4 UUID string.
func UUIDv4() (string, error) {
    b := make([]byte, 16)
    if _, err := io.ReadFull(crand.Reader, b); err != nil {
        return "", err
    }
    // Set version (4) and variant (10)
    b[6] = (b[6] & 0x0f) | 0x40
    b[8] = (b[8] & 0x3f) | 0x80
    // Format 8-4-4-4-12
    dst := make([]byte, 36)
    hex.Encode(dst[0:8], b[0:4])
    dst[8] = '-'
    hex.Encode(dst[9:13], b[4:6])
    dst[13] = '-'
    hex.Encode(dst[14:18], b[6:8])
    dst[18] = '-'
    hex.Encode(dst[19:23], b[8:10])
    dst[23] = '-'
    hex.Encode(dst[24:36], b[10:16])
    return string(dst), nil
}
