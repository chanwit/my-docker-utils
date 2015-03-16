package main

import (
	"os"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func split(s string, delim string) []string {
	re, _ := regexp.Compile(`"([\d|\w|\s]+)"`)
	for _, match := range re.FindAllStringSubmatch(s, -1) {
		rp := strings.NewReplacer(delim, "%20")
		escape := rp.Replace(match[1])
		s = strings.Replace(s, match[1], escape, -1)
	}
	r := strings.Split(s, delim)
	for i, _ := range r {
		if os.Getenv("DEBUG") != "" {
			fmt.Println("DEBUG: " + r[i])
		}
		r[i] = strings.Replace(r[i], "%20", delim, -1)
		r[i] = strings.Replace(r[i], "\"", "", -1)
	}
	return r
}

func main() {
	args := strings.Join(os.Args[1:], " ")
	r, _ := regexp.Compile(`\$\((.*)\)`)
	result := r.FindStringSubmatch(args)
	var cmd *exec.Cmd
	if len(result) >= 2 {
		subCmd := strings.Split(result[1], " ")
		cmd = exec.Command(subCmd[0], subCmd[1:]...)
		output, _ := cmd.Output()
		subArgs := r.ReplaceAllString(args, string(output))
		newCmd := split(subArgs, " ")
		cmd = exec.Command(newCmd[0], newCmd[1:]...)
	} else {
		cmd = exec.Command(os.Args[1], os.Args[2:]...)
	}
	if os.Getenv("DEBUG") != "" {
		fmt.Printf("%s\n", cmd)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if os.Getenv("DEBUG") == "" {
		cmd.Run()
	}
}
