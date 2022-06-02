package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func ReadArgs() []string {
	res := make([]string, 0)
	scan := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scan.Scan() {
		res = append(res, scan.Text())
	}
	return res
}
func main() {
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		args := make([]string, 0)
		args = append(args, os.Args[2:]...)
		args = append(args, ReadArgs()...)
		exCmd := exec.Command(cmd, args...)
		res, _ := exCmd.CombinedOutput()
		fmt.Print(string(res))
	}
}
