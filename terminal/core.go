// Core module, that mimics the filesystem and sets up the directory listing and manipulation functionalty
package terminal

import (
	"fmt"
	"strings"
)

type Dir struct {
	dirName string
	files   map[string]*Dir
	isRoot  bool
	parent  *Dir
}

func (dir *Dir) findPath(comingFrom string) string {
	if dir.isRoot {

		return ("/" + comingFrom)

	} else {
		return dir.parent.findPath(dir.dirName + "/" + comingFrom)
	}

}

func (dir *Dir) travelToRoot(comingFrom string) string {
	if dir.isRoot {
		CWD = dir
		return ("/" + dir.dirName + comingFrom)

	} else {

		return dir.parent.travelToRoot(dir.dirName + "/" + comingFrom)
	}

}

func (dir *Dir) createDirectory(dirName string) {
	if _, ok := dir.files[dirName]; ok {
		fmt.Println("ERR:", "DIRECTORY ALREADY EXISTS")
	} else {
		fmt.Println("SUCC: CREATED", dirName)
		dir.files[dirName] = &Dir{
			dirName,
			make(map[string]*Dir),
			false,
			dir,
		}
	}
}

func (dir *Dir) enlist() {
	// fmt.Println("TMP ENLIST", dir.dirName)
	// for k, v := range dir.files {
	// 	fmt.Println(k, v)
	// }

}

func (dir *Dir) deleteDir(path string) {

	oldPath := CWD

	if path == "/" {
		fmt.Println("Can't remove root")
		return
	}

	splittedPath := strings.Split(path, "/")

	directories := splittedPath[1 : len(splittedPath)-1]
	dirToDelete := splittedPath[len(splittedPath)-1]
	CWD = ROOT

	for _, dir := range directories {

		if CWD.moveTo(dir) {
			continue
		} else {
			fmt.Println("Invalid Path", CWD)
			CWD = oldPath
			return
		}
	}

	// -----

	if _, ok := CWD.files[dirToDelete]; ok {
		// File exists
		CWD.files[dirToDelete] = nil
		delete(CWD.files, dirToDelete)

	} else {
		fmt.Println("Invalid Dir", dirToDelete)

	}

	CWD = oldPath
}

func (dir *Dir) changeDir(dirName string) bool {
	if _dir, ok := CWD.files[dirName]; ok {
		// File exists
		CWD = _dir
		return true

	} else {
		fmt.Println("ERR:", "INVALID PATH")
		return false
	}
}

func (cwd *Dir) moveTo(dirName string) bool {
	return cwd.changeDir(dirName)

}

func (cwd *Dir) moveToDir(path string) {

	if path == "/" {
		moveToRoot()
		return
	}

	oldPath := cwd

	splittedPath := strings.Split(path, "/")

	directories := splittedPath[1:]
	CWD = ROOT

	for _, dir := range directories {

		if CWD.moveTo(dir) {
			continue
		} else {
			CWD = oldPath
			break
		}
	}

}
