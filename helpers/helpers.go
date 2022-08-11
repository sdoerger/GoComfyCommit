package pwhelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func AssertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

type tCredentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// type tSubTargets struct {
//     SubTargets []string `json:"subTargets"`
// }

// User struct which contains a name
// a type and a list of social links
type tProfile struct {
	CommitMessage string `json:"commitMessage"`
	TruncBranchAt int    `json:"truncBranchAt"`
	Alias         string `json:"alias"`
}
type tSetup struct {
	Profiles []tProfile `json:"profiles"`
}

func OpenFileRead(path string) ([]byte, error) {
	importFile, err := os.Open(path)
	if err != nil {
		// log.Fatalf("Error opening setup file: %v", err)
	}

	byteValue, err := ioutil.ReadAll(importFile)
	if err != nil {
		// log.Fatalf("Error at reading setup file: %v", err)
	}

	return byteValue, err
}

func SetupJson(setupPath string) (tSetup, error) {

	file, err := OpenFileRead(setupPath)
	if err != nil {
		log.Fatalf("Error at reading setup file: %v", err)
	}

	var setup tSetup
	json.Unmarshal(file, &setup)

	return setup, err

}

func FindProfile(profiles []string, profile string) /*(string, error)*/ {
	fmt.Println(profiles)
	fmt.Println(profile)
}

func ParseMessagePattern(patternString string) {
	fmt.Println(patternString)
}

// func RemoveDuplicateStr(strSlice []string) []string {
// 	allKeys := make(map[string]bool)
// 	list := []string{}
// 	for _, item := range strSlice {
// 		if _, value := allKeys[item]; !value {
// 			allKeys[item] = true
// 			list = append(list, item)
// 		}
// 	}
// 	return list
// }
