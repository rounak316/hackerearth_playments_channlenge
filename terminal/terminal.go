// Exposed module for terminal, used as the initalization point
package terminal

import (
	"bufio"
	"fmt"
	"os"
)

// Start the dummy terminal
func Start() {
	fmt.Println("Starting the application")
	fmt.Println("Welcome to linux terminal\n")
	ROOT = createRootDirectory()
	CWD = ROOT
	for {

		str1 := CWD.findPath("")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(TAG_USER, str1, "$>")
		text, _ := reader.ReadString('\n')
		commandParser(text)
	}
}
