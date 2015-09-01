/* Copyright (C) Skyfore 2015 *\
   <- 本库仅供个人学习交流之用 ->
  <- 未经允许,严禁用于商业软件 ->
   <- 许可条约见 License.txt ->
\* Copyright (C) Skyfore 2015 */

package sys

import (
	"syscall"
	"unsafe"
)

type DLLClass struct {
	Handle uintptr
}

func TEXT(str string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}

func MustLoadLibrary(name string) uintptr {

	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic("函数库" + name + " 载入失败 ...")
	}

	return uintptr(lib)
}

func MustGetProcAddress(lib uintptr, name string) uintptr {

	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic("函数" + name + " 地址获取失败 ...")
	}

	return uintptr(addr)
}

func Dll(dllName string) *DLLClass {
	DllCls := &DLLClass{}
	DllCls.Handle = MustLoadLibrary(dllName)
	return DllCls
}

func (dllCls *DLLClass) Free() {
	syscall.FreeLibrary(syscall.Handle(dllCls.Handle))
}

func (dllCls *DLLClass) Call(apiName interface{}, argvs ...uintptr) uintptr {
	var nArgvs = uintptr(len(argvs))
	var apiFunc uintptr

	switch apiName.(type) {
	case string:
		apiFunc = MustGetProcAddress(dllCls.Handle, apiName.(string))
	case uintptr:
		apiFunc = apiName.(uintptr)
	default:
		panic("函数参数 apiName 必须是 uintptr 或 string 类型 ...")
	}

	if nArgvs <= 3 {
		cArgs := []uintptr{0, 0, 0}
		for k, v := range argvs {
			cArgs[k] = v
		}
		apiRet, _, _ := syscall.Syscall(apiFunc, nArgvs, cArgs[0], cArgs[1], cArgs[2])
		return apiRet
	} else if nArgvs <= 6 {
		cArgs := []uintptr{0, 0, 0, 0, 0, 0}
		for k, v := range argvs {
			cArgs[k] = v
		}
		apiRet, _, _ := syscall.Syscall6(apiFunc, nArgvs, cArgs[0], cArgs[1], cArgs[2], cArgs[3], cArgs[4], cArgs[5])
		return apiRet
	} else if nArgvs <= 9 {
		cArgs := []uintptr{0, 0, 0, 0, 0, 0, 0, 0, 0}
		for k, v := range argvs {
			cArgs[k] = v
		}
		apiRet, _, _ := syscall.Syscall9(apiFunc, nArgvs, cArgs[0], cArgs[1], cArgs[2], cArgs[3], cArgs[4], cArgs[5], cArgs[6], cArgs[7], cArgs[8])
		return apiRet
	} else if nArgvs <= 12 {
		cArgs := []uintptr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for k, v := range argvs {
			cArgs[k] = v
		}
		apiRet, _, _ := syscall.Syscall12(apiFunc, nArgvs, cArgs[0], cArgs[1], cArgs[2], cArgs[3], cArgs[4], cArgs[5], cArgs[6], cArgs[7], cArgs[8], cArgs[9], cArgs[10], cArgs[11])
		return apiRet
	} else if nArgvs <= 15 {
		cArgs := []uintptr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for k, v := range argvs {
			cArgs[k] = v
		}
		apiRet, _, _ := syscall.Syscall15(apiFunc, nArgvs, cArgs[0], cArgs[1], cArgs[2], cArgs[3], cArgs[4], cArgs[5], cArgs[6], cArgs[7], cArgs[8], cArgs[9], cArgs[10], cArgs[11], cArgs[12], cArgs[13], cArgs[14])
		return apiRet
	} else {
		panic("过多的 DLL 调用函数 ...")
	}
}
