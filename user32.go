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
