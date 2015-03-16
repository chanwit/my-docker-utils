package main

import "fmt"
import "os"
import "io/ioutil"
import "strings"
import "path/filepath"

func main() {
	filename := os.Args[1]
	if filepath.Ext(filename) == "" {
		filename = filename + ".profile"
	}
	bytes, err := ioutil.ReadFile(filename)
	fileContent := strings.TrimSpace(string(bytes))
	if err == nil {
		r := strings.NewReplacer("\r\n", " ", "\r", " ", "\n", " ", "\t", " ")
		fmt.Print(r.Replace(fileContent))
	} else {
		os.Exit(1)
	}
}
