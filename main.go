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
		commitText := os.Args[1]
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

			formatBranch := "#" + string(currentBranch)
			runGitCommit := exec.Command("git", "commit", "-m "+string(commitText)+" "+formatBranch)
			_, err = runGitCommit.Output()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			// Print the output
			fmt.Println(string("✅: Excuted Git Command: git commit -m " + string(commitText) + " " + formatBranch))
			fmt.Println("☕️: Just push")

			return
		}
	} else {
		fmt.Println("No Commit Message")
	}
}
