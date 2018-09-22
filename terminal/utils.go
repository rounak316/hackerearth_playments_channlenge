// This module contains util functions
package terminal

import (
	"fmt"
	"strings"
)

// Command Parser: Parses input and directs to specific command executor
func commandParser(cmd string) {

	cmd = strings.Trim(cmd, "\n")
	cmds := strings.Split(cmd, " ")

	if len(cmds) == 1 {

		switch cmds[0] {
		case "ls":
			Enlist(CWD, 0)
			break
		case "pwd":
			pwd := CWD.findPath("")
			fmt.Println("PATH:", pwd)
			break

		default:
			fmt.Println("ERR:", "CANNOT RECOGNIZE INPUT.")
			break

		}
	} else if len(cmds) > 1 {

		switch cmds[0] {

		case "mkdir":
			for j := 1; j < len(cmds); j++ {
				CWD.createDirectory(cmds[j])
			}
			break

		case "cd":

			if cmds[1][0] == '/' {
				cdInto := cmds[1]
				CWD.moveToDir(cdInto)
			} else {
				cdInto := CWD.findPath("") + cmds[1]
				CWD.moveToDir(cdInto)
			}

			break

		case "rm":
			for j := 1; j < len(cmds); j++ {

				if cmds[j][0] == '/' {
					cdInto := cmds[j]
					CWD.deleteDir(cdInto)
				} else {
					cdInto := CWD.findPath("") + cmds[j]
					CWD.deleteDir(cdInto)
				}
			}

			break

		default:
			fmt.Println("ERR:", "CANNOT RECOGNIZE INPUT.")
			break

		}

	} else {

		fmt.Println("ERR:", "CANNOT RECOGNIZE INPUT.")

	}

}

// Space seperated graph display of dirctory listing: Used to debug the app by nlisting the directory struture from root
func prepareSpace(level int) (space string) {

	for i := 0; i < level; i++ {
		space += "   "
	}
	return
}

// Enlist the files/dirs in "dir", recursively
func enlist(dir *Dir, level int) {

	for key, val := range dir.files {
		if val == CWD {
			fmt.Println(prepareSpace(level), key, ("(CWD)"))
		} else {
			fmt.Println(prepareSpace(level), key)
		}

		enlist(val, level+1)
	}
}

// Enlist wrapper exposed to allow access for directory traversal
func Enlist(dir *Dir, level int) {

	if ROOT == CWD {
		// fmt.Println(prepareSpace(level), dir.dirName, ("(CWD)"))
	} else {
		// fmt.Println(prepareSpace(level), dir.dirName)
	}
	fmt.Println("DIRS:")
	enlist(dir, level+1)

}

// Bootstrap Function, used to create root directory and mimic the file system
func createRootDirectory() (rootDir *Dir) {

	rootDir = &Dir{
		"ROOT",
		make(map[string]*Dir),
		true,
		nil,
	}

	return

}

// Change current directory to root directory
func moveToRoot() {
	CWD = ROOT
}
