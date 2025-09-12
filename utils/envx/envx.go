package envx

import (
    "os"
    "strconv"
    "strings"
    "time"
)

// Get returns the environment variable or the provided default if unset.
func Get(key, def string) string {
    if v, ok := os.LookupEnv(key); ok {
        return v
    }
    return def
}

// MustGet returns the environment variable or panics if it is not set.
func MustGet(key string) string {
    if v, ok := os.LookupEnv(key); ok {
        return v
    }
    panic("required env var not set: " + key)
}

// Bool parses a boolean-like env var with a default fallback.
// Recognized true values: 1, t, true, y, yes, on (case-insensitive)
// Recognized false values: 0, f, false, n, no, off
func Bool(key string, def bool) bool {
    v, ok := os.LookupEnv(key)
    if !ok {
        return def
    }
    s := strings.TrimSpace(strings.ToLower(v))
    switch s {
    case "1", "t", "true", "y", "yes", "on":
        return true
    case "0", "f", "false", "n", "no", "off":
        return false
    default:
        return def
    }
}

// Int parses an int env var with a default fallback.
func Int(key string, def int) int {
    v, ok := os.LookupEnv(key)
    if !ok {
        return def
    }
    n, err := strconv.Atoi(strings.TrimSpace(v))
    if err != nil {
        return def
    }
    return n
}

// Int64 parses an int64 env var with a default fallback.
func Int64(key string, def int64) int64 {
    v, ok := os.LookupEnv(key)
    if !ok {
        return def
    }
    n, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
    if err != nil {
        return def
    }
    return n
}

// Float64 parses a float64 env var with a default fallback.
func Float64(key string, def float64) float64 {
    v, ok := os.LookupEnv(key)
    if !ok {
        return def
    }
    n, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
    if err != nil {
        return def
    }
    return n
}

// Duration parses a time.Duration env var with a default fallback.
func Duration(key string, def time.Duration) time.Duration {
    v, ok := os.LookupEnv(key)
    if !ok {
        return def
    }
    d, err := time.ParseDuration(strings.TrimSpace(v))
    if err != nil {
        return def
    }
    return d
}

// Expand expands ${var} or $var in the string according to the values of the current environment.
func Expand(s string) string {
    return os.ExpandEnv(s)
}

