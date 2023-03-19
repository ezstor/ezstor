package info

import (
	"os/exec"
	"strings"
)

func Controllers() ([]string, error) {
	outb, err := exec.Command(
		"/bin/sh", "-c", `lspci`,
	).Output()
	if err != nil {
		return nil, err
	}

	var conts []string
	for _, line := range strings.Split(strings.Trim(string(outb), "\n"), "\n") {
		switch {
		case strings.Contains(line, "Mass storage controller") ||
			strings.Contains(line, "RAID bus controller") ||
			strings.Contains(line, "Serial Attached SCSI controller"):
			var cont string
			elems := strings.Split(line, ":")
			cont = elems[len(elems)-1]
			cont = strings.Trim(cont, " ")
			if strings.Contains(cont, "(") {
				cont = strings.Split(cont, "(")[0]
				cont = strings.Trim(cont, " ")
			}
			if strings.Contains(cont, "[") {
				cont = strings.Split(cont, "[")[0]
				cont = strings.Trim(cont, " ")
			}
			conts = append(conts, cont)
		}
	}

	return conts, nil
}
