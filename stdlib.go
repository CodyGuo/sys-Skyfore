/* Copyright (C) Skyfore 2015 *\
   <- 本库仅供个人学习交流之用 ->
  <- 未经允许,严禁用于商业软件 ->
   <- 许可条约见 License.txt ->
\* Copyright (C) Skyfore 2015 */

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
