package strx

import (
    "strings"
)

// Coalesce returns the first non-empty string from vals, or "" if all are empty.
func Coalesce(vals ...string) string {
    for _, v := range vals {
        if v != "" {
            return v
        }
    }
    return ""
}

// ContainsFold reports whether substr is within s, case-insensitively.
func ContainsFold(s, substr string) bool {
    if substr == "" {
        return true
    }
    return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// SplitAndTrim splits s by sep, trims spaces for each part, and filters empty items.
func SplitAndTrim(s, sep string) []string {
    if s == "" {
        return nil
    }
    parts := strings.Split(s, sep)
    out := make([]string, 0, len(parts))
    for _, p := range parts {
        p = strings.TrimSpace(p)
        if p != "" {
            out = append(out, p)
        }
    }
    return out
}

// JoinNonEmpty joins parts using sep, skipping empty strings.
func JoinNonEmpty(sep string, parts ...string) string {
    out := make([]string, 0, len(parts))
    for _, p := range parts {
        if p != "" {
            out = append(out, p)
        }
    }
    return strings.Join(out, sep)
}

// TrimAll collapses any whitespace runs to single spaces and trims both ends.
func TrimAll(s string) string {
    if s == "" {
        return ""
    }
    fields := strings.Fields(s)
    return strings.Join(fields, " ")
}

