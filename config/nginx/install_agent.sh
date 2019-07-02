#! /usr/bin/env bash

set -e

token="$1"
ref="${2:-master}"

function run {
	echo $ "$@" >&2
	$@
}

cd /tmp

echo "downloading latest agent release..."
curl > agent.zip \
	-fLH "Private-Token: ${token}" \
	"https://gitlab.enix.io/api/v4/projects/166/jobs/artifacts/${ref}/download?job=build-agent-linux"


run apt update
run apt install -y python-boto duplicity zip

run unzip -o agent.zip

run cp bananactl-linux /usr/local/bin/bananactl
run cp config/systemd/* /etc/systemd/system/

run systemctl start banana.timer
run systemctl enable banana.timer

echo "success!"
