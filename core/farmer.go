package core

import (
	"fmt"
	"regexp"
)

type Farmer struct {
	Name []byte
	Id   []byte
}

func FindAllFarmer(b []byte) []Farmer {
	matchFarmer, err := regexp.Compile("<Farmer>.*?</Farmer>")
	if err != nil {
		panic(err)
	}
	match := matchFarmer.FindAll(b, -1)
	farmhands := make([]Farmer, 0)
	for _, item := range match {
		name := FindFarmerName(item)
		if len(name) > 0 {
			id := FindFarmerId(item)
			farmhands = append(farmhands, Farmer{name, id})
		}
	}
	return farmhands
}

func FindFarmerName(b []byte) []byte {
	matchFarmerName, err := regexp.Compile("<Farmer><name>(.*?)</name>")
	if err != nil {
		panic(err)
	}
	submatch := matchFarmerName.FindSubmatch(b)
	if len(submatch) > 1 {
		return submatch[1]
	}
	return []byte{}
}

func FindFarmerId(b []byte) []byte {
	matchFarmerId, err := regexp.Compile("<Farmer>.*?<userID>(.*?)</userID>")
	if err != nil {
		panic(err)
	}
	submatch := matchFarmerId.FindSubmatch(b)
	if len(submatch) > 1 {
		return submatch[1]
	}
	return []byte{}
}

func Release(b []byte, farmer string) []byte {
	matchFarmerContent, err := regexp.Compile(fmt.Sprintf("(<Farmer><name>%s</name>.*?)(<userID>\\S*?</userID>)(.*?</Farmer>)", farmer))
	if err != nil {
		panic(err)
	}
	return matchFarmerContent.ReplaceAll(b, []byte("$1<userID />$3"))
}
