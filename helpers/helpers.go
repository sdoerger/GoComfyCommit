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
type tSetup struct {
	Credentials tCredentials `json:"credentials"`
	TargetUrl   string       `json:"targetUrl"`
	LoginUrl    string       `json:"loginUrl"`
	SubTargets  []string     `json:"subTargets"`
	DownloadDir string       `json:"downloadDir"`
	Content     string       `json:"content"`
}

func OpenFileRead(path string) ([]byte, error) {
	importFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening setup file: %v", err)
	}

	byteValue, err := ioutil.ReadAll(importFile)
	if err != nil {
		log.Fatalf("Error at reading setup file: %v", err)
	}

	return byteValue, err
}

func SetupJson(setupPath string) tSetup {

	fmt.Println("JSON RUNS")

	file, err := OpenFileRead(setupPath)
	if err != nil {
		log.Fatalf("Error at reading setup file: %v", err)
	}

	var setup tSetup
	json.Unmarshal(file, &setup)

	return setup

}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
