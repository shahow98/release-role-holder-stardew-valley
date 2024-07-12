package core

import (
	"os"
	"path"
	"regexp"
	"strings"
)

const BaseConf = "conf/base.conf"
const DetailConf = "conf/detail.conf"
const ReleaseConf = "conf/release.conf"

const LF string = "\n"
const CRLF string = "\r\n"

func ReadConfig(file string) map[string]string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	content := string(data)
	separator := LF
	if strings.Contains(content, CRLF) {
		separator = CRLF
	}
	lines := strings.Split(content, separator)
	conf := make(map[string]string)
	for _, line := range lines {
		matchConf, err := regexp.Compile("^(\\S+?)=([^#]+).*$")
		if err != nil {
			panic(err)
		}
		submatchs := matchConf.FindStringSubmatch(line)
		if len(submatchs) == 3 {
			conf[submatchs[1]] = strings.Trim(submatchs[2], " ")
		}
	}
	return conf
}

func DetectDir(auto bool, dir string) (string, error) {
	if auto || dir == "" {
		dir = path.Join(os.Getenv("appdata"), "StardewValley")
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) || err != nil {
		return "", err
	}
	_, err = os.Stat(GetSaveDir(dir))
	if os.IsNotExist(err) || err != nil {
		return "", err
	}
	return dir, nil
}
