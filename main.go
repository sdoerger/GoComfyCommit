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
	alias := flag.String("p", "default", "alias")
	crop := flag.String("c", "0", "crop")
	flag.Parse()

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
	_, noConfig := helpers.OpenFileRead(setupPath)
	if noConfig == nil {

		// // Get setup json
		setupProfiles, err := helpers.SetupJson(setupPath)

		// TODO: DO STUFF with config
		if err == nil {

			profile, err := helpers.FindProfile(setupProfiles.Profiles, *alias)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(profile)

			// string to int
			cropToStr, err := strconv.Atoi(*crop)
			if err != nil {
				fmt.Println(err.Error())
			}

			// If profiles has crop item and no crop is set (default: 0), assign it from profile
			if profile.CropBranchFromTo > 0 && cropToStr <= 0 {
				fmt.Println("RUNS")
				*crop = strconv.Itoa(profile.CropBranchFromTo)
			}

			setFullCommitTest := helpers.SetFullCommitTest()
			fmt.Println(setFullCommitTest)

			/* TODO: RM */
			// return

			// fmt.Println(helpers.FindProfile(setupProfiles.Profiles, *alias))

			// TODO
			// Set crop length (if not flag on crop is set)
			// *crop =

		}

	}

	fmt.Println("CROP")
	fmt.Println(*crop)
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

	// string to int
	toCrop, err := strconv.Atoi(*crop)
	if err != nil {
		// ... handle error
		panic(err)
	}

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
