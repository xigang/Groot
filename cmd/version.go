package cmd

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli"
)

var (
	version      = "0.0.0"
	buildData    = "1970-01-01T00:00:00Z"
	gitCommit    = ""
	gitTag       = ""
	gitTreeState = ""
)

type Version struct {
	Version      string
	BuildData    string
	GitCommit    string
	GitTag       string
	GitTreeState string
	GoVersion    string
	Compiler     string
	Platform     string
}

var GrootVersion = cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "print version information",
	Action: func(c *cli.Context) error {
		version := GetVersion()
		fmt.Printf("%s: %s\n", c.App.Name, version.Version)
		fmt.Printf(" BuildDate: %s\n", version.BuildData)
		fmt.Printf(" GitCommit: %s\n", version.GitCommit)
		fmt.Printf(" GitTag: %s\n", version.GitTag)
		fmt.Printf(" GitTreeState: %s\n", version.GitTreeState)
		fmt.Printf(" GoVersion: %s\n", version.GoVersion)
		fmt.Printf(" Complier: %s\n", version.Compiler)
		fmt.Printf(" Platform: %s\n", version.Platform)

		return nil
	},
}

func (v *Version) GetVersion() string {
	return v.Version
}

func GetVersion() Version {
	var versionStr string
	if gitCommit != "" && gitTag != "" && gitTreeState == "clean" {
		versionStr = gitTag
	} else {
		versionStr = "v" + version
		if len(gitCommit) >= 7 {
			versionStr += "+" + gitCommit[0:7]
			if gitTreeState != "clean" {
				versionStr += ".dirty"
			}
		}
	}

	return Version{
		Version:      versionStr,
		BuildData:    buildData,
		GitCommit:    gitCommit,
		GitTag:       gitTag,
		GitTreeState: gitTreeState,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
