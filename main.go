// package vip allows you to run vim in a pipeline
package main

import (
	"io/ioutil"
	"os"
	"os/exec"

	"fmt"
)

func main() {
	//fmt.Fprint(os.Stderr, os.Args[1:])

	tmpFile, err := createTempFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	//fmt.Fprint(os.Stderr, tmpFile)

	vip := exec.Command("vi",
		append(os.Args[1:],
			"-n",
			"-i", "NONE",
			"-u", "NONE",
			"-c", "silent! w! "+tmpFile,
			"-c", "q",
			"-")...)
	vip.Stdin, vip.Stderr = os.Stdin, os.Stderr
	//vip.Stdout = os.Stdout
	err = vip.Run()
	//fmt.Fprint(os.Stderr, err)
	out, _ := ioutil.ReadFile(tmpFile)
	//fmt.Fprint(os.Stderr, err)
	fmt.Fprintf(os.Stdout, "%s", out)
}

func createTempFile() (string, error) {
	tmp, err := ioutil.TempFile("", "vip")
	if err != nil {
		return "", err
	}

	name := tmp.Name()
	err = tmp.Close()
	if err != nil {
		return "", err
	}

	return name, nil
}
