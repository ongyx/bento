//go:build discretegpu

package bento

/*
#cgo LDFLAGS: -L${SRCDIR}/discretegpu_windows.def

__declspec(dllexport) unsigned long NvOptimusEnablement = 0x00000001;
__declspec(dllexport) int AmdPowerXpressRequestHighPerformance = 1;
*/
import "C"
