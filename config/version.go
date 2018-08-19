package config

import (
	"bytes"
	"fmt"
)

const AwsshAsciiLogo = `
 █████╗  ██╗    ██╗ ██████╗ ██████╗ ██╗
██╔══██╗ ██║    ██║ ██╔═══╝ ██╔═══╝ ██║   
███████║ ██║ █╗ ██║ ██████  ██████  ███████╗
██╔══██║ ██║███╗██║     ██╗     ██╗ ██╔══██║
██║  ██║ ╚███╔███╔╝ ██████║ ██████║ ██║  ██║
╚═╝  ╚═╝  ╚══╝╚══╝  ╚═════╝ ╚═════╝ ╚═╝  ╚═╝
`

var (
	Version  = "v0.0.1"
	BuildFor string

	buildSha, buildDate, buildArch, buildOS string
)

type BuildInfo struct {
	Version, Sha, Date, Arch, OS, For string
}

func (b BuildInfo) String() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("version=%s", b.Version))

	if b.Sha != "" {
		buff.WriteString(fmt.Sprintf(", commit=%s", b.Sha))
	}
	if b.Date != "" {
		buff.WriteString(fmt.Sprintf(", build-date=%s", b.Date))
	}
	if b.Arch != "" {
		buff.WriteString(fmt.Sprintf(", build-arch=%s", b.Arch))
	}
	if b.OS != "" {
		buff.WriteString(fmt.Sprintf(", build-os=%s", b.OS))
	}
	if b.For != "" {
		buff.WriteString(fmt.Sprintf(", build-for=%s", b.For))
	}
	return buff.String()
}

var CurrentBuildInfo = BuildInfo{
	Version: Version,
	For:     BuildFor,
	Sha:     buildSha,
	Date:    buildDate,
	Arch:    buildArch,
	OS:      buildOS,
}
