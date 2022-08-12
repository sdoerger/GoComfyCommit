package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"

	helpers "GoCommit/helpers"
)

func main() {

	// -----------------
	// FLAGS
	// -----------------

	commitText := flag.String("m", "Update", "commit message")
	changeType := flag.String("t", "Update", "*changeType")
	alias := flag.String("p", "default", "alias")
	crop := flag.String("c", "[0,0]", "crop")
	flag.Parse()

	// var cropRange []int

	// cropToString := []byte(*crop)
	// cropRange := cropToString[1]

	// Convert crop string form flag to slice
	var cropRange []int
	if err := json.Unmarshal([]byte(*crop), &cropRange); err != nil {
		panic(err)
	}

	fmt.Println("cropRange")
	fmt.Println(cropRange)

	fmt.Println("Alias : ", *alias)
	fmt.Println("crop : ", *crop)
	fmt.Println("*commitText : ", *commitText)
	fmt.Println("*changeType : ", *changeType)
	fmt.Println("\n")

	setupPath := "./config.json"

	// #################################################
	// IF SETUP FILE ###################################
	// #################################################

	// Check if there is a config.json file
	_, hasNoConfig := helpers.OpenFileRead(setupPath)
	if hasNoConfig == nil {

		// // Get setup json
		setupProfiles, err := helpers.SetupJson(setupPath)

		// TODO: DO STUFF with config
		if err == nil {

			// Find profile from config array
			profile, err := helpers.FindProfile(setupProfiles.Profiles, *alias)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(profile)

			// #######################
			// ADD CROP RANGE ########
			// #######################

			// IF profiles has crop item and no crop is set (default: 0), assign it from profile
			if len(profile.CropBranchFromTo) == 2 && *crop == "[0,0]" {
				// Convert crop string form flag to slice
				var convertProfileCrop []int
				if err := json.Unmarshal([]byte(*crop), &convertProfileCrop); err != nil {
					panic(err)
				}
				cropRange = profile.CropBranchFromTo
			}

			// ##########################################
			// SET COMMIT MESSAGE BY PATTERN ############
			// ##########################################

			// IF profiles has crop item and no crop is set (default: 0), assign it from profile
			if len(profile.CommitMessage) > 0 {
				fmt.Println(profile.CommitMessage)
				// var convertProfileCrop []int
				// if err := json.Unmarshal([]byte(*crop), &convertProfileCrop); err != nil {
				// 	panic(err)
				// }
			}

			setFullCommitTest := helpers.SetFullCommitTest()
			fmt.Println(setFullCommitTest)
		}
	}

	/* TODO: RM */
	return

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
	// Crop branchname if set by flag or config

	trimmedBranch := string(currentBranch[:len(currentBranch)-1])
	if len(currentBranch) > cropRange[1] {
		trimmedBranch = string(currentBranch[cropRange[0]:cropRange[1]])
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
