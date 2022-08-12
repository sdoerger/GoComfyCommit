package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"
	"regexp"

	helpers "GoCommit/helpers"
)

func main() {

	// -----------------
	// FLAGS
	// -----------------

	commitText := flag.String("m", "Update", "commit message")
	changeType := flag.String("t", "", "*changeType")
	alias := flag.String("p", "", "alias")
	crop := flag.String("c", "[0,0]", "crop")
	flag.Parse()

	// Convert crop string form flag to slice
	var cropRange []int
	if err := json.Unmarshal([]byte(*crop), &cropRange); err != nil {
		panic(err)
	}
	var commitMsgPattern string

	setupPath := "./config.json"

	// #################################################
	// IF SETUP FILE ###################################
	// #################################################

	// Check if there is a config.json file
	_, hasNoConfig := helpers.OpenFileRead(setupPath)
	if hasNoConfig == nil && len(*alias) > 0 {

		// // Get setup json
		setupProfiles, err := helpers.SetupJson(setupPath)

		// TODO: DO STUFF with config
		if err == nil {

			// #######################
			// FIND PROFILE ##########
			// #######################

			profile, err := helpers.FindProfile(setupProfiles.Profiles, *alias)
			if err != nil {
				fmt.Println(err.Error())
				// return
			}

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
			// SET COMMIT MESSAGE PATTERN ############
			// ##########################################
			if len(profile.CommitMessage) > 0 {
				commitMsgPattern = profile.CommitMessage
			}

			// ##########################################
			// SET DEFAULT CHANGE TYPE ##################
			// ##########################################
			if len(profile.DefaultCommitType) > 0 && len(*changeType) <= 0 {
				*changeType = profile.DefaultCommitType
			}
		}
	}

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

	// SET PATTERN TO COMMIT TEXT (IF SETUP)
	if hasNoConfig == nil && len(commitMsgPattern) > 0 {
		fullCommitText = helpers.CommitMessageByPattern(commitMsgPattern, *changeType, trimmedBranch, string(*commitText))
	} else {
		// DEFAULT IF NO TYPE
		if len(*changeType) >= 0 {
			fullCommitText = "[" + string(trimmedBranch) + "] " + string(*commitText)
		}

		// DEFAULT
		fullCommitText = string(*commitText)
	}

	// RM duplicate spaces
	space := regexp.MustCompile(`\s+`)
	fullCommitText = space.ReplaceAllString(fullCommitText, " ")

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
