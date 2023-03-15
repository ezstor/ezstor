package info

import (
	"os/exec"
	"strings"
)

func Controllers() ([]string, error) {
	outb, err := exec.Command(
		"/bin/sh", "-c",
		`lspci | grep "^[0-9,a-z]" | grep -E 'Fusion-MPT|MegaRAID|Adaptec' | awk -F ':' '{print $NF}' | awk -F '[(|[]' '{print $1}' | awk '{gsub(/^\s+|\s+$/, "");print}'`,
	).Output()
	if err != nil {
		return nil, err
	}
	output := string(outb)

	var conts []string
	if output != "" {
		conts = strings.Split(strings.Trim(output, "\n"), "\n")
	}
	return conts, nil
}
