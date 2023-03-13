package info

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func bash(cmd string) string {
	job := exec.Command("/bin/sh", "-c", cmd)
	var stdout, stderr bytes.Buffer
	job.Stdout = &stdout
	job.Stderr = &stderr
	err := job.Run()
	outStr := stdout.String()
	if err != nil {
		log.Printf("job.Run() failed with %s\n", err)
		return ""
	}
	return outStr
}

func Controllers() []string {
	cons := bash(`lspci | grep "^[0-9,a-z]" | grep -E 'Fusion-MPT|MegaRAID|Adaptec' | awk -F ':' '{print $NF}' | awk -F '[(|[]' '{print $1}' | awk '{gsub(/^\s+|\s+$/, "");print}'`)
	con_slice := []string{}
	if cons != "" {
		con_slice = strings.Split(cons, "\n")
	}
	return con_slice
}
