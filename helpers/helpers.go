package helpers

import (
	"encoding/json"
	"errors"
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

// type tSubTargets struct {
//     SubTargets []string `json:"subTargets"`
// }

// User struct which contains a name
// a type and a list of social links
type TProfile struct {
	CommitMessage string `json:"commitMessage"`
	TruncBranchAt int    `json:"truncBranchAt"`
	Alias         string `json:"alias"`
}
type TSetup struct {
	Profiles []TProfile `json:"profiles"`
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

func SetupJson(setupPath string) (TSetup, error) {

	file, err := OpenFileRead(setupPath)
	if err != nil {
		log.Fatalf("Error at reading setup file: %v", err)
	}

	var setup TSetup
	json.Unmarshal(file, &setup)

	return setup, err

}

func FindProfile(profiles []TProfile, alias string) (TProfile, error) {

	match, matchedProfile := false, TProfile{}

	for _, profile := range profiles {
		if profile.Alias == alias {
			match = true
			matchedProfile = profile

		}
	}
	if match {
		return matchedProfile, nil
	}
	return matchedProfile, errors.New("Profile is not in config file (-p " + alias + ")")

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
