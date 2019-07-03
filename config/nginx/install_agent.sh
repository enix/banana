#! /usr/bin/env bash

set -e

token="$1"
ref="${2:-master}"

function run {
	echo $ "$@" >&2
	$@
}

cd /tmp

set +e
which bananactl
if [[ $? -eq 0 ]]; then
	isUpgrade="yes"
fi
set -e

echo "downloading latest agent release..."
curl > agent.zip \
	-fLH "Private-Token: ${token}" \
	"https://gitlab.enix.io/api/v4/projects/166/jobs/artifacts/${ref}/download?job=build-agent-linux"

if [[ -z ${isUpgrade} ]]; then
	run apt update
	run apt install -y python-boto duplicity zip
fi

run unzip -o agent.zip

run cp bananactl-linux /usr/local/bin/bananactl
run cp config/systemd/* /etc/systemd/system/

if [[ -z ${isUpgrade} ]]; then
	run systemctl start banana.timer
	run systemctl enable banana.timer
else
	run systemctl daemon-reload
fi

echo "success!"
