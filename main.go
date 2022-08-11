package main

import (
	"fmt"
	"os"
	"os/exec"

	helpers "GoCommit/helpers"
)

func main() {

	// -----------------
	// Command line Args
	// -----------------

	if os.Args != nil && len(os.Args) >= 1 {

		changeType := ""
		commitText := "Update"
		setupPath := "./config.json"

		// Check if there is a config.json file
		_, noFile := helpers.OpenFileRead(setupPath)
		if noFile == nil {

			// // Get setup json
			setupProfiles, err := helpers.SetupJson(setupPath)

			if err == nil {
				fmt.Println(setupProfiles.Profiles)

			}

		}

		if len(os.Args) == 2 {
			commitText = os.Args[1]
			changeType = ""
		}
		if len(os.Args) == 3 {
			changeType = os.Args[1]
			commitText = os.Args[2]
		}

		if commitText != "" {

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
			// Trimm branchname to 8, if needed
			trimmedBranch := string(currentBranch[:len(currentBranch)-1])
			if len(currentBranch) > 8 {
				trimmedBranch = string(currentBranch[:8])
			}

			fullCommitText := ""

			if os.Args != nil && len(os.Args) == 1 {
				fullCommitText = string(commitText)
			}

			if os.Args != nil && len(os.Args) == 2 {
				fullCommitText = string(commitText)
			}

			if os.Args != nil && len(os.Args) == 3 && changeType != "" {
				fullCommitText = changeType + ": [" + string(trimmedBranch) + "] " + string(commitText)
			}

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
	} else {
		fmt.Println("No Commit Message")
	}
}
