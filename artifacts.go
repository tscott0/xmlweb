package main

import (
	"encoding/xml"
	"fmt"
)

type Artifacts struct {
	Created  string     `xml:"created,attr" json:"created"`
	Project  string     `xml:"project,attr" json:"project"`
	Version  string     `xml:"version,attr" json:"version"`
	Artifact []Artifact `xml:"artifact"`
}

type Artifact struct {
	Created     string      `xml:"created,attr" json:"created"`
	ProjectName string      `xml:"projectName,attr"`
	Type        string      `xml:"type,attr"`
	GroupID     string      `xml:"groupId,attr"`
	ArtifactID  string      `xml:"artifactId,attr"`
	Version     string      `xml:"version,attr"`
	SCM         SCM         `xml:"scm"`
	Jenkins     Jenkins     `xml:"jenkins"`
	Artifactory Artifactory `xml:"artifactor"`
}

type SCM struct {
	Module string `xml:"module,attr"`
	Server string `xml:"server,attr"`
	Tag    string `xml:"tag,attr"`
	Type   string `xml:"type,attr"`
	Branch string `xml:"branch,attr"`
	Repo   string `xml:"repo,attr"`
}

type Jenkins struct {
	ID      string `xml:"id,attr"`
	JobName string `xml:"jobName,attr"`
	Slave   string `xml:"slave,attr"`
	URL     string `xml:"url,attr"`
}

type Artifactory struct {
	URL      string `xml:"url,attr"`
	BuildTag string `xml:"buildTag,attr"`
}

func newArtifacts(x []byte) Artifacts {
	artifacts := Artifacts{}
	err := xml.Unmarshal(x, &artifacts)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return artifacts
	}
	return artifacts
}
