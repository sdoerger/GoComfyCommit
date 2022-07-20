package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	// -----------------
	// Command line Args
	// -----------------

	if os.Args != nil && len(os.Args) > 1 {
		changeType := os.Args[1]
		commitText := os.Args[2]
		if commitText != "" && changeType != "" {

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

			fullCommitText := changeType + ": [" + string(trimmedBranch) + "] " + string(commitText)
			fmt.Println(fullCommitText)
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
