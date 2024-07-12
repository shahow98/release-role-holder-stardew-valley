package main

import (
	"fmt"

	core "shahow.top/crhsv-core"
)

func main() {
	ReleaseRole()
}

func ReleaseRole() {
	baseConf := core.ReadConfig(core.BaseConf)
	releaseConf := core.ReadConfig(core.ReleaseConf)

	auto := false
	if baseConf["detectMode"] == "AUTO" {
		auto = true
	}
	gameDir, err := core.DetectDir(auto, baseConf["gameDir"])
	if err != nil {
		panic(err)
	}
	farm := releaseConf["farm"]
	farmer := releaseConf["farmer"]
	fmt.Printf("release %s -> %s\n", farm, farmer)
	src, err := core.GetFarmInfo(gameDir, farm)
	if err != nil {
		panic(err)
	}
	b := core.Release(src, farmer)
	core.SaveFarmInfo(gameDir, farm, src, b, true)
	fmt.Println(core.ReleaseConf + " is released.")
}
