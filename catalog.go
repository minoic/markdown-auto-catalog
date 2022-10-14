package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	DocumentPath string
	ContentPath  string
	Filter       string
)

func init() {
	flag.StringVar(&DocumentPath, "doc", "", "")
	flag.StringVar(&ContentPath, "content", "", "")
	flag.StringVar(&Filter, "filter", "", "")
	flag.Parse()
}

func dfs(ctl *strings.Builder, folderPath string, depth int) {
	fmt.Println(folderPath)
	gap := strings.Repeat("  ", depth)
	files, _ := ioutil.ReadDir(folderPath)
	for _, file := range files {
		if file.IsDir() {
			ctl.WriteString(fmt.Sprintf("%s- %s\n", gap, file.Name()))
			dfs(ctl, path.Join(folderPath, file.Name()), depth+1)
			continue
		}
		if ok, _ := regexp.MatchString(Filter, file.Name()); !ok {
			continue
		}
		ctl.WriteString(fmt.Sprintf("%s- [%s](%s)\n", gap, strings.Trim(file.Name(), "[]"), strings.ReplaceAll(path.Join(folderPath, file.Name()), " ", "%20")))
	}
}

func main() {
	readme, err := os.OpenFile(DocumentPath, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	readmeB, err := io.ReadAll(readme)
	if err != nil {
		panic(err)
	}
	readmeS := string(readmeB)
	flag := "<!-- catalog -->"
	splits := strings.Split(readmeS, flag)
	ctl := strings.Builder{}
	ctl.WriteString("\n\n")
	dfs(&ctl, ContentPath, 0)
	readmeS = splits[0] + flag + ctl.String() + "\n" + flag + splits[2]
	err = readme.Truncate(0)
	if err != nil {
		panic(err)
	}
	_, err = readme.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	_, err = readme.WriteString(readmeS)
	if err != nil {
		panic(err)
	}
}
