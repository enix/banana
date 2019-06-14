#! /usr/bin/env bash

set -e

token="$1"

function run {
	echo $ "$@" >&2
	$@
}

run cd /tmp

echo "downloading latest agent release..."
curl > agent.zip \
	-fLH "Private-Token: $token" \
	"https://gitlab.enix.io/api/v4/projects/166/jobs/artifacts/develop/download?job=build-agent-linux"

run unzip -o agent.zip

run cp bananactl-linux /usr/bin/bananactl
run cp config/systemd/* /etc/systemd/system/

run systemctl daemon-reload
run systemctl start banana
run systemctl enable banana

echo "success!"
