package core

import (
	"os"
	"path"
	"time"
)

const saveDir string = "Saves"

func FindAllFarms(dir string) []string {
	dirs, err := os.ReadDir(GetSaveDir(dir))
	if err != nil {
		panic(err)
	}
	farms := make([]string, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			farms = append(farms, dir.Name())
		}
	}
	return farms
}

func GetSaveDir(dir string) string {
	return path.Join(dir, saveDir)
}

func GetFarmFile(dir string, f string, ft string) string {
	return path.Join(GetSaveDir(dir), f, f+ft)
}

func GetFarmInfo(dir string, f string) ([]byte, error) {
	data, err := os.ReadFile(GetFarmFile(dir, f, ""))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SaveFarmInfo(dir string, f string, src []byte, rpl []byte, isBak bool) {
	if isBak {
		t := time.Now().Format("20060102150405")
		err := os.WriteFile(path.Join(GetFarmFile(dir, f, "."+t+".bak")), src, 0644)
		if err != nil {
			panic(err)
		}
	}
	err := os.WriteFile(GetFarmFile(dir, f, ""), rpl, 0644)
	if err != nil {
		panic(err)
	}
}
