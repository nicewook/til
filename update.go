package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const Header = `# Today I Learned

![https://unsplash.com/photos/s9CC2SKySJM?utm_source=unsplash&utm_medium=referral&utm_content=creditShareLink](img/green-chameleon-s9CC2SKySJM-unsplash.jpg)
---

`

func markdownFiles(files []string) []string {
	var result []string
	for _, filename := range files {
		if strings.Contains(filename, "md") {
			result = append(result, filepath.Base(filename))
		}
	}
	return result
}

func main() {
	content := Header

	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	sort.Strings(dirs)
	fmt.Println(dirs)
	for _, dirName := range dirs {
		files, err := ioutil.ReadDir(dirName)
		if err != nil {
			log.Fatal(err)
		}

		var fNames []string
		for _, file := range files {
			fNames = append(fNames, file.Name())
		}
		mdFiles := markdownFiles(fNames)
		if len(mdFiles) == 0 {
			continue
		}

		content += fmt.Sprintf("## %s\n\n", dirName)
		for _, mdFile := range mdFiles {
			content += fmt.Sprintf("- [%s](%s)\n", mdFile, filepath.Join(dirName, mdFile))
		}
	}

	f, err := os.OpenFile("README.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.Write([]byte(content)); err != nil {
		log.Fatal(err)
	}
}
