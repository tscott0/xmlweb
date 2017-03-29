package main

import (
	"html/template"
)

var (
	homeTempl = template.Must(template.New("home").Parse(tmplText))
)

const tmplText = `
{{define "home"}}
<!DOCTYPE html>
<html lang="en">
    <head>
        <title>WebSocket Example</title>
    </head>
    <style>
	div {
        font-family: monospace;
        font-size: 14px;
		color: #202020;
	}
    .master {
        background-color: lightgreen;
        padding: 20px;
		border-style: solid;
		border-color: white;
		border-radius: 10px;
    }
    .title {
        font-weight: bold;
        font-size: 20px;
        padding-bottom: 10px;
    }
    .label {
        width: 120px;
        display: inline-block;
        font-weight: bold;
        padding-left: 20px;
    }
    .art {
        background-color: powderblue;
        padding: 20px;
        font-size: 16px;
		border-style: solid;
		border-color: white;
		border-radius: 10px;
    }
    .art-sub {
        background-color: #c6ecf0;
		padding: 20px;
		margin-top: 10px;
		border-radius: 10px;
	}
    .sub-title {
        font-weight: bold;
        font-size: 18px;
        padding-bottom: 10px;
    }
    .meta {
        background-color: khaki;
        padding: 20px;
        font-size: 16px;
		border-style: solid;
		border-color: white;
		border-radius: 10px;
    }
    .fileData {
        background-color: thistle;
        padding: 10px;
		white-space: pre-wrap;
        font-size: 14px;
		border-style: solid;
		border-color: white;
		border-radius: 10px;
    </style>
    <body>
        <div class="master">
			<div class="title">{{.Artifacts.Project}}</div>
			<div class="master-prop"><div class="label">Version</div>{{.Artifacts.Version}}</div>
			<div class="master-prop"><div class="label">Created</div>{{.Artifacts.Created}}</div>
		</div>
        {{- range .Artifacts.Artifact -}}
        <div class="art">
			<div class="title">{{.ProjectName}}</div>
			<div class="art-prop"><div class="label">Version</div>{{.Version}}</div>
			<div class="art-prop"><div class="label">Created</div>{{.Created}}</div>
			<div class="art-prop"><div class="label">GroupID</div>{{.GroupID}}</div>
			<div class="art-prop"><div class="label">ArtifactID</div>{{.ArtifactID}}</div>
			<div class="art-prop"><div class="label">Type</div>{{.Type}}</div>
			<div class="art-sub">
				<div class="sub-title">SCM</div>
				<div class="art-sub-prop"><div class="label">Module</div>{{.SCM.Module}}</div>
				<div class="art-sub-prop"><div class="label">Server</div>{{.SCM.Server}}</div>
				<div class="art-sub-prop"><div class="label">Tag</div>{{.SCM.Tag}}</div>
				<div class="art-sub-prop"><div class="label">Type</div>{{.SCM.Type}}</div>
				<div class="art-sub-prop"><div class="label">Branch</div>{{.SCM.Branch}}</div>
				<div class="art-sub-prop"><div class="label">Repo</div>{{.SCM.Repo}}</div>
			</div>
			<div class="art-sub">
				<div class="sub-title">Jenkins</div>
				<div class="art-sub-prop"><div class="label">Job Name</div>{{.Jenkins.JobName}}</div>
				<div class="art-sub-prop"><div class="label">ID</div>{{.Jenkins.ID}}</div>
				<div class="art-sub-prop"><div class="label">Slave</div>{{.Jenkins.Slave}}</div>
				<div class="art-sub-prop"><div class="label">URL</div><a href="{{.Jenkins.URL}}">{{.Jenkins.URL}}</a></div>
			</div>
		</div>
        {{end}}
        <div class="meta">
			<div class="title">Meta</div>
			<div class="meta-prop"><div class="label">Hostname</div>{{.Host}}</div>
			<div class="meta-prop"><div class="label">Timestamp</div>{{.LastMod}}</div>
		</div>
        <div class="fileData">{{.Data}}</div>
    </body>
</html>
{{end}}
`
