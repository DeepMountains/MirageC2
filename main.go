package main

import (
	"MirageC2_V1.0/public/common"
	"MirageC2_V1.0/public/initbase"
)

func main() {
	common.LoadConfig()
	initbase.Display()
	initbase.GetInput()
}
