/* Copyright (C) Skyfore 2015 *\
   <- 本库仅供个人学习交流之用 ->
  <- 未经允许,严禁用于商业软件 ->
   <- 许可条约见 License.txt ->
\* Copyright (C) Skyfore 2015 */

package sys

var User32 *DLLClass

func init() {
    User32 = Dll("User32.dll")
}

/*
func MessageBox(hwnd int, szText string, szTitle string, uFlag int) uintptr {
	r := User32.Call("MessageBoxW", uintptr(hwnd), TEXT(szText), TEXT(szTitle), uintptr(uFlag))
	return r
}
*/
