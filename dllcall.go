package sys

import (
    "runtime"
    "syscall"
    "unsafe"
)

func init() {
    runtime.LockOSThread()
}

type DLLClass struct {
    Handle uintptr
}

func TEXT(str string) uintptr {
    return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}

func MustLoadLibrary(name string) uintptr {
    lib, err := syscall.LoadLibrary(name)
    if err != nil {
        panic("DLL < " + name + " > failed to load," + err.Error())
    }

    return uintptr(lib)
}

func MustGetProcAddress(lib uintptr, name string) uintptr {
    addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
    if err != nil {
        panic("Failed to get the address function < " + name + " >," + err.Error())
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
        panic("The parameters of the function < " + apiName.(string) + " > must be uintptr or string type")
    }

    if nArgvs <= 3 {
        cArgs := []uintptr{0, 0, 0}
        for k, v := range argvs {
            cArgs[k] = v
        }
        apiRet, _, _ := syscall.Syscall(apiFunc, nArgvs,
            cArgs[0],
            cArgs[1],
            cArgs[2])
        return apiRet
    } else if nArgvs <= 6 {
        cArgs := []uintptr{0, 0, 0, 0, 0, 0}
        for k, v := range argvs {
            cArgs[k] = v
        }
        apiRet, _, _ := syscall.Syscall6(apiFunc, nArgvs,
            cArgs[0],
            cArgs[1],
            cArgs[2],
            cArgs[3],
            cArgs[4],
            cArgs[5])
        return apiRet
    } else if nArgvs <= 9 {
        cArgs := []uintptr{0, 0, 0, 0, 0, 0, 0, 0, 0}
        for k, v := range argvs {
            cArgs[k] = v
        }
        apiRet, _, _ := syscall.Syscall9(apiFunc, nArgvs,
            cArgs[0],
            cArgs[1],
            cArgs[2],
            cArgs[3],
            cArgs[4],
            cArgs[5],
            cArgs[6],
            cArgs[7],
            cArgs[8])
        return apiRet
    } else if nArgvs <= 12 {
        cArgs := []uintptr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
        for k, v := range argvs {
            cArgs[k] = v
        }
        apiRet, _, _ := syscall.Syscall12(apiFunc, nArgvs,
            cArgs[0],
            cArgs[1],
            cArgs[2],
            cArgs[3],
            cArgs[4],
            cArgs[5],
            cArgs[6],
            cArgs[7],
            cArgs[8],
            cArgs[9],
            cArgs[10],
            cArgs[11])
        return apiRet
    } else if nArgvs <= 15 {
        cArgs := []uintptr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
        for k, v := range argvs {
            cArgs[k] = v
        }
        apiRet, _, _ := syscall.Syscall15(apiFunc, nArgvs,
            cArgs[0],
            cArgs[1],
            cArgs[2],
            cArgs[3],
            cArgs[4],
            cArgs[5],
            cArgs[6],
            cArgs[7],
            cArgs[8],
            cArgs[9],
            cArgs[10],
            cArgs[11],
            cArgs[12],
            cArgs[13],
            cArgs[14])
        return apiRet
    } else {
        panic("Function < " + apiName.(string) + " > Call too many parameters.")
    }
}
