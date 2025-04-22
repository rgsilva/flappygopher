package sifrpc

/*
#define _EE
#include <stdlib.h>
#include <kernel.h>
#include <sifrpc.h>
#include <loadfile.h>
#include <sbv_patches.h>
#include <iopcontrol.h>

static int resetAndPatchIOP()
{
	SifInitRpc(0);
	while(!SifIopReset("", 0)){};
	while(!SifIopSync()){};
	SifInitRpc(0);

	int ret = sbv_patch_enable_lmb();
	if (ret != 0) {
		return ret;
	}

	return sbv_patch_disable_prefix_check();
}
*/
import "C"
import (
	"ps2go/debug"
	"unsafe"
)

func ResetAndPatchIOP() {
	ret := C.resetAndPatchIOP()
	debug.Printf("ResetAndPatchIOP: %d\n", int(ret))
}

func LoadModule(path string) {
	cPath := C.CString(path)

	ret := int(C.SifLoadModule(cPath, 0, nil))
	debug.Printf("Load module: %s, %d\n", path, ret)

	C.free(unsafe.Pointer(cPath))
}

func LoadModuleBuffer(ptr unsafe.Pointer, size int) {
	ret := int(C.SifExecModuleBuffer(ptr, C.uint(size), C.uint(0), nil, nil))
	debug.Printf("Load module buffer:%d\n", ret)
}
