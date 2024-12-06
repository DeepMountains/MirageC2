package main

import (
	"MirageC2_V1.0/public/common"
	"MirageC2_V1.0/public/initbase"
	"runtime"
)

func main() {
	common.C2Platform = runtime.GOOS
	common.LoadConfig()
	initbase.Display()
	initbase.GetInput()
}
