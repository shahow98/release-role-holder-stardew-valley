package main

import (
	"fmt"
	"os"

	core "shahow.top/rrhsv-core"
)

func main() {
	ReadDetailInfo()
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}

func ReadDetailInfo() {
	baseConf := core.ReadConfig(core.BaseConf)
	auto := false
	if baseConf["detectMode"] == "AUTO" {
		auto = true
	}
	gameDir, err := core.DetectDir(auto, baseConf["gameDir"])
	if err != nil {
		panic(err)
	}
	farms := core.FindAllFarms(gameDir)
	var detailInfo string = ""
	for _, farm := range farms {
		farmInfo, err := core.GetFarmInfo(gameDir, farm)
		if err != nil {
			continue
		}
		farmhands := core.FindAllFarmer(farmInfo)
		detailInfo += "farm: " + farm + "\n"
		for _, farmer := range farmhands {
			detailInfo += "\tfarmer: " + string(farmer.Name) + "\n"
		}
	}
	os.WriteFile(core.DetailConf, []byte(detailInfo), 0644)
	fmt.Println(core.DetailConf + " is refreshed.")
}
