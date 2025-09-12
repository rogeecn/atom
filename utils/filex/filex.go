package filex

import (
    "errors"
    "io/fs"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"
    "time"
)

// Exists reports whether the given path exists (file or directory).
func Exists(path string) bool {
    if path == "" {
        return false
    }
    _, err := os.Stat(path)
    return err == nil
}

// IsFile reports whether the path exists and is a regular file.
func IsFile(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return info.Mode().IsRegular()
}

// IsDir reports whether the path exists and is a directory.
func IsDir(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return info.IsDir()
}

// IsReadable reports whether the target can be opened for reading.
func IsReadable(path string) bool {
    f, err := os.Open(path)
    if err != nil {
        return false
    }
    _ = f.Close()
    return true
}

// IsWritable reports whether a file can be written to. If the file does not exist,
// it checks writability of the parent directory by attempting to create a temp file.
func IsWritable(path string) bool {
    if path == "" {
        return false
    }
    info, err := os.Stat(path)
    if err == nil {
        if info.IsDir() {
            // Try writing a temp file in the directory
            f, err := os.CreateTemp(path, ".wtmp-*")
            if err != nil {
                return false
            }
            name := f.Name()
            _ = f.Close()
            _ = os.Remove(name)
            return true
        }
        // Try opening for write
        f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0)
        if err != nil {
            return false
        }
        _ = f.Close()
        return true
    }
    if !errors.Is(err, os.ErrNotExist) {
        return false
    }
    // Not exist: check parent dir is writable by creating a temp file
    dir := filepath.Dir(path)
    if dir == "." || dir == "" {
        dir = "."
    }
    f, err := os.CreateTemp(dir, ".wtmp-*")
    if err != nil {
        return false
    }
    name := f.Name()
    _ = f.Close()
    _ = os.Remove(name)
    return true
}

// IsExecutable reports whether a file is considered executable on the current OS.
// On Unix, checks the executable bits; on Windows, checks the extension against PATHEXT.
func IsExecutable(path string) bool {
    info, err := os.Stat(path)
    if err != nil || info.IsDir() {
        return false
    }
    if runtime.GOOS == "windows" {
        ext := strings.ToLower(filepath.Ext(path))
        if ext == ".exe" || ext == ".bat" || ext == ".cmd" || ext == ".com" || ext == ".ps1" {
            return true
        }
        pathext := strings.ToLower(os.Getenv("PATHEXT"))
        for _, e := range strings.Split(pathext, ";") {
            if e != "" && ext == strings.ToLower(e) {
                return true
            }
        }
        return false
    }
    // Unix permissions: any execute bit set
    return info.Mode()&0o111 != 0
}

// Size returns the size of the file at path.
func Size(path string) (int64, error) {
    info, err := os.Stat(path)
    if err != nil {
        return 0, err
    }
    return info.Size(), nil
}

// Touch creates the file if it does not exist, or updates its modification time if it does.
func Touch(path string) error {
    now := time.Now()
    if Exists(path) {
        return os.Chtimes(path, now, now)
    }
    f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0o644)
    if err != nil {
        return err
    }
    return f.Close()
}

// EnsureDir ensures that a directory exists with the specified permission bits.
func EnsureDir(path string, perm fs.FileMode) error {
    if path == "" {
        return nil
    }
    if Exists(path) {
        if !IsDir(path) {
            return errors.New("path exists and is not a directory")
        }
        return nil
    }
    return os.MkdirAll(path, perm)
}

// FindInPath searches for an executable named name in the system PATH.
func FindInPath(name string) (string, bool) {
    p, err := exec.LookPath(name)
    if err != nil {
        return "", false
    }
    return p, true
}

// ReadFileString reads a whole file into a string.
func ReadFileString(path string) (string, error) {
    b, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(b), nil
}

// WriteFileString writes a string to a file with the given permissions, creating or truncating it.
func WriteFileString(path, s string, perm fs.FileMode) error {
    return os.WriteFile(path, []byte(s), perm)
}

// WriteFileAtomic writes data to a temporary file in the same directory and renames it into place.
func WriteFileAtomic(path string, data []byte, perm fs.FileMode) error {
    dir := filepath.Dir(path)
    if err := EnsureDir(dir, 0o755); err != nil {
        return err
    }
    f, err := os.CreateTemp(dir, ".atomic-*")
    if err != nil {
        return err
    }
    tmp := f.Name()
    // Best-effort cleanup on error
    defer func() { _ = os.Remove(tmp) }()
    if _, err := f.Write(data); err != nil {
        _ = f.Close()
        return err
    }
    if err := f.Chmod(perm); err != nil {
        _ = f.Close()
        return err
    }
    if err := f.Sync(); err != nil {
        _ = f.Close()
        return err
    }
    if err := f.Close(); err != nil {
        return err
    }
    return os.Rename(tmp, path)
}

