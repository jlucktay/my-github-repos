package main

import (
	"fmt"
	"os"

	"go.jlucktay.dev/version"
	"golang.org/x/term"
)

func prettyPrintTerminal(input []string, repoType string) error {
	fmt.Printf("%d repo %s:\n", len(input), repoType)

	// get terminal width
	tw, _, errTGS := term.GetSize(int(os.Stdout.Fd()))
	if errTGS != nil {
		return fmt.Errorf("couldn't get terminal size: %w", errTGS)
	}

	// get longest repo name
	longestRepoName := 0
	for i := range input {
		if len(input[i]) > longestRepoName {
			longestRepoName = len(input[i])
		}
	}

	// do math to divide lines evenly across width
	longestRepoName++ // add a single padding space
	reposPerLine := tw / longestRepoName

	// space out repo names in columns and pretty print
	for i := 0; i < len(input); i += reposPerLine {
		for j := 0; j < reposPerLine && i+j < len(input); j++ {
			fmt.Printf("%-[1]*[2]s", longestRepoName, input[i+j])
		}

		fmt.Println()
	}

	fmt.Println()

	return nil
}

type jsonOutput struct {
	Version string   `json:"version"`
	Sources []string `json:"sources,omitempty"`
	Forks   []string `json:"forks,omitempty"`
}

var jsonBuffer jsonOutput

const (
	printSources = "sources"
	printForks   = "forks"
)

func prettyPrintJSON(input []string, repoType string) error {
	ver, err := version.Details()
	if err != nil {
		return fmt.Errorf("could not get version details: %w", err)
	}

	jsonBuffer.Version = ver

	switch repoType {
	case printSources:
		jsonBuffer.Sources = append(jsonBuffer.Sources, input...)
	case printForks:
		jsonBuffer.Forks = append(jsonBuffer.Forks, input...)
	}

	return nil
}
