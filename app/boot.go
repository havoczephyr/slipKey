package app

import (
	"os"
)

func Boot() (bool, string, error) {
	fixBool := false

	wd, err := os.Getwd()
	if err != nil {
		return fixBool, wd, err
	}
	if os.Args[1] == "-f" {
		fixBool = true
	}
	return fixBool, wd, nil
}
