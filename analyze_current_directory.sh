#!/usr/bin/env bash

# This script will run a set of Go static analysis tools against a local project and collect results into an HTML file.

# Ensure we are on a valid Go project
if [ -z "${GOPATH}" ]; then
    echo "ERROR: GOPATH environment variable not defined"
    exit 1
fi

project_dir=$(pwd)

if ! [[ ${project_dir} == ${GOPATH}/src/* ]]; then
    echo "ERROR: The current directory does not correspond to a valid Go project"
    exit 1
fi

# Get the analysis done
project_name=${project_dir#${GOPATH}/src/}
echo "Analyzing local checkout of \"${project_name}\"... this will take a minute"

docker run -v ${project_dir}:/go/src/${project_name} nearmap/goreportcard ${project_name}

# Report result
if [[ $? == 0 ]]; then
    echo "Report was successfully generated -> goreportcard.html"
else
    echo "ERROR: Something went wrong."
fi
