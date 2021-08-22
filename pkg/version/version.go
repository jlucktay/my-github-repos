// Package version has some global strings that should be set with ldflags at compile time, and will attempt to derive
// some (hopefully) sensible default values as a fallback if the appropriate ldflags are not set.
package version

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Take ldflags from GoReleaser.
var (
	// Executable is the name of the currently executing binary.
	// Defaults to the base path of the string returned by calling 'os.Executable()'.
	Executable string

	// Version is the semver-compatible git tag that this binary was built from.
	// Defaults to 'v0.0.0'.
	Version string

	// Commit is the short hash of the commit that this binary was built from.
	// Defaults to the value returned by running 'git describe --always --dirty'.
	Commit string

	// BuildDate is the build timestamp of the currently executing binary.
	// Defaults to the modification time (from calling 'os.Stat') on the path returned by calling 'os.Executable()'.
	BuildDate string

	// BuiltBy is the name of the user that built the currently executing binary.
	// Defaults to the username returned by calling 'user.Current()'.
	BuiltBy string
)

// Details returns a string describing the current binary.
func Details() (string, error) {
	var exePath string

	if Executable == "" || BuildDate == "" {
		var err error

		exePath, err = os.Executable()
		if err != nil {
			return "", fmt.Errorf("could not look up current executable: %w", err)
		}
	}

	if Executable == "" {
		Executable = filepath.Base(exePath)
	}

	if Version == "" {
		Version = "v0.0.0"
	}

	if Commit == "" {
		cmd := exec.Command("git", "describe", "--always", "--dirty")

		output, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("could not run '%s': %w", cmd, err)
		}

		Commit = strings.TrimSpace(string(output))
	}

	if BuildDate == "" {
		osfi, err := os.Stat(exePath)
		if err != nil {
			return "", fmt.Errorf("could not stat current executable: %w", err)
		}

		BuildDate = osfi.ModTime().Format(time.RFC3339)
	}

	if BuiltBy == "" {
		currUser, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("could not get current user: %w", err)
		}

		BuiltBy = currUser.Username
	}

	return fmt.Sprintf("%s %s built from commit %s with %s on %s by %s.",
		Executable, Version, Commit, runtime.Version(), BuildDate, BuiltBy), nil
}
