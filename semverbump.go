// Package semverbump provides a function to bump a semVer tag, extending golang.org/x/mod/semver library.
//
// The function Bump() takes a semVer tag and a bump type (major, minor, patch) and returns the new semVer tag.
// For example if the current semVer tag is v1.0.0 and the bump type is patch, the new semVer tag will be v1.0.1.
// If the current semVer tag is v1.0.1 and the bump type is minor, the new semVer tag will be v1.1.0
// (note that the patch version is reset to 0).
//
// Shortened semVer tags are supported, for example if the current semVer tag is v1.0 and the bump type is patch,
// the bumped semVer tag will be v1.0.1.
//
// If the current semVer tag does not have a leading v, it will be added.
// If it is an empty string, it will behave as if the current semVer tag is v0.0.0, e.g. if the bump type is patch,
// the bumped semVer tag will be v0.0.1.
//
// If the current semVer tag is invalid, such as "x.y.z", an error will be returned.
//
// The library currently does not handle pre-release or build metadata.
package semverbump

import (
	"fmt"
	"golang.org/x/mod/semver"
	"strconv"
	"strings"
)

// Bump takes a semVer tag and a bump type (major, minor, patch) and returns a new semVer tag
// whose value is incremented according to the bump type. If the current semVer tag does not have a leading v,
// it will be added. If it is an empty string, it will behave as if the current semVer tag is v0.0.0
// Invalid semVer tags will return an error.
func Bump(currentVersion, majorMinorPatch string) (string, error) {
	if currentVersion == "" {
		currentVersion = "v0.0.0"
	} else {
		if !semver.IsValid(currentVersion) {
			if semver.IsValid("v" + currentVersion) {
				currentVersion = "v" + currentVersion
			} else {
				return "", fmt.Errorf("invalid semVer tag '%s'", currentVersion)
			}
		}
		// semVer lib allows v1 and v1.0, so check for that
		if currentVersion == semver.Major(currentVersion) {
			currentVersion = currentVersion + ".0.0"
		}
		if currentVersion == semver.MajorMinor(currentVersion) {
			currentVersion = currentVersion + ".0"
		}
	}
	// now we have a valid currentVersion
	newMajor, newMinor, newPatch := majorMinorPatchFromSemVer(currentVersion)
	if majorMinorPatch == "major" {
		newMajor++
		newMinor = 0
		newPatch = 0
	} else if majorMinorPatch == "minor" {
		newMinor++
		newPatch = 0
	} else if majorMinorPatch == "patch" {
		newPatch++
	} else {
		return "", fmt.Errorf("invalid arg for majorMinorPatch, expect 'major', 'minor' or 'patch', got '%s'", majorMinorPatch)
	}
	return semVerFromMajorMinorPatch(newMajor, newMinor, newPatch), nil
}

func majorMinorPatchFromSemVer(semVer string) (int, int, int) {
	parts := strings.Split(semVer, ".")
	major, _ := strconv.Atoi(strings.TrimPrefix(parts[0], "v"))
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])
	return major, minor, patch
}

func semVerFromMajorMinorPatch(major, minor, patch int) string {
	return "v" + strconv.Itoa(major) + "." + strconv.Itoa(minor) + "." + strconv.Itoa(patch)
}
