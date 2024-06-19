package main

import (
	"bufio"
	cp "github.com/otiai10/copy"
	"os"
	"strings"
)

func main() {
	var folderName string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	folderName = scanner.Text()
	folderName = strings.Replace(folderName, ". ", "_", -1)
	folderName = strings.Replace(folderName, " ", "_", -1)
	_ = cp.Copy("aNewLessonFolderGenerator\\sampleNewTask", folderName)
}
