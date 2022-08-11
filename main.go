package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"

	helpers "GoCommit/helpers"
)

func main() {

	// -----------------
	// FLAGS
	// -----------------

	commitText := flag.String("m", "Update", "commit message")
	changeType := flag.String("t", "Update", "*changeType")
	alias := flag.String("a", "default", "alias")
	crop := flag.String("c", "0", "crop")
	flag.Parse()

	fmt.Println("Alias : ", *alias)
	fmt.Println("crop : ", *crop)
	fmt.Println("*commitText : ", *commitText)
	fmt.Println("*changeType : ", *changeType)

	setupPath := "./config.json"

	// Check if there is a config.json file
	_, noConfig := helpers.OpenFileRead(setupPath)
	if noConfig == nil {

		// // Get setup json
		setupProfiles, err := helpers.SetupJson(setupPath)

		// TODO: DO STUFF with config
		if err == nil {
			fmt.Println(setupProfiles.Profiles)

		}

	}

	// return

	if len(*commitText) == 0 {
		fmt.Println("No Commit Message")
		return

	}
	// -------------------------------------------
	// ------------  GET CURRENT BRANCH  ---------
	// -------------------------------------------
	runCurrentBranch := exec.Command("git", "branch", "--show-current")
	currentBranch, err := runCurrentBranch.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// -------------------------------------------
	// ------------  ADD ALL FILES  --------------
	// -------------------------------------------

	runGitAddAll := exec.Command("git", "add", ".")
	_, err = runGitAddAll.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// -------------------------------------------
	// ------------  COMMIT CMD MESSAGE + BRANCH
	// -------------------------------------------
	// Trimm branchname if set

	// string to int
	toCrop, err := strconv.Atoi(*crop)
	if err != nil {
		// ... handle error
		panic(err)
	}

	fmt.Println(toCrop)

	trimmedBranch := string(currentBranch[:len(currentBranch)-1])
	if len(currentBranch) > toCrop {
		trimmedBranch = string(currentBranch[:toCrop])
	}

	fullCommitText := ""

	fullCommitText = *changeType + ": [" + string(trimmedBranch) + "] " + string(*commitText)

	runGitCommit := exec.Command("git", "commit", "-m"+fullCommitText)
	_, err = runGitCommit.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string("✅: Commited:" + string(fullCommitText)))
	fmt.Println("☕️: Just push")

	return
}
