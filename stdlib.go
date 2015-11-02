package sys

import (
    "os"
    "path/filepath"
)

func FullPath(path string) (p string) {
    p, _ = filepath.Abs(path)
    return
}

func FileExist(path string) bool {
    _, err := os.Stat(path)
    if err != nil {
        return false
    } else {
        return true
    }

}
