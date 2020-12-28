#!/usr/bin/env bash

# NCP variables
APPLICATION_NAME="pv"
DOCKER_REPOSITORY_NAME=$(echo ${APPLICATION_NAME} | tr '[:upper:]' '[:lower:]' | tr -d '[:space:]')
DOCKER_TAG="${DOCKER_REPOSITORY_NAME}:latest"

# Other variables
DIRECTORY_WORKING=$(pwd)
DIRECTORY_ENVIRONMENTS="${DIRECTORY_WORKING}/scripts"

## build: build the Docker image like in production (from scratch)
function task_build {
  echo "Building Docker image of ${APPLICATION_NAME} ..."
  build_docker_image Dockerfile
}

# build-debug: build the Docker image for remote debugging (with Delve)
#function task_build_debug {
#  echo "Building Docker image of ${APPLICATION_NAME} for remote-debugging ..."
#  build_docker_image Dockerfile.debug
#}

## chain: chain the desired steps ((b)uild and further steps) by entering their first letters, e.g. "b..." runs all steps sequentially
function task_chain {
  parameter=$1

  if [[ $parameter =~ b ]]; then    # Regexp --> contains "b"
    task_build
  fi
}

## call-graph: open the call graph of the project
function task_call_graph {
  # Calculate environment variable "GOPATH" if it's not already set
  if [[ -z "$GOPATH" ]]; then                                     # Check if variable GOPATH is empty (--> not set)
    echo "Environment variable \"GOPATH\" is empty"
    if [[ $(pwd | grep -oP "\/src\/") == "/src/" ]]; then         # Check if working dir has a substring "/src/"
      echo "Working directory contains a substring \"/src/\""
      export GOPATH=$(pwd | grep -oP ".*(?=\/src\/)")             # GOPATH is probably everything before "/src/"
      echo "Environment variable \"GOPATH\" is set to ${GOPATH}"
    else
      echo "GOPATH is empty and can't get calculated because the working directory"
      echo "contains no substring \"/src/\" that is used to identify the GOPATH"
      echo 'Please set environment variable \"GOPATH\" or execute the script in any directory inside $GOPATH/src/'
      exit 1
    fi
  fi

  # Install go-callvis if necessary
  if [[ ! -e $GOPATH/bin/go-callvis ]]; then
    echo "Installing go-callvis ..."
    cd ~                                      # Change directory to avoid adding go-callvis to go.mod and go.sum
    go get github.com/ofabry/go-callvis
    cd $DIRECTORY_WORKING
    echo "Finished installing go-callvis."
  fi

  # Open the call graph
  root=$(pwd | grep -oP "(?<=\/src\/).*")     # Positive Lookbehind --> get everything after "/src/"
  echo "Rendering the call graph for ${root} ..."
  echo "Stop serving by pressing \"Ctrl + C\""
  $GOPATH/bin/go-callvis -group=pkg,type -nostd ${root}
}

## mocks: generate and update mocks (for testing purpose)
function task_mocks {
  echo "Generating mocks ..."
  docker pull vektra/mockery
  docker run -v "$PWD":/src -w /src --rm vektra/mockery --all --output=./domain/mocks
  echo "Finished generating mocks."
}

# build_docker_image builds and tags the Docker image based on the provided Dockerfile
function build_docker_image {
  dockerfile=$1
  echo "Building Docker image ... (tag = ${DOCKER_TAG})"
  docker build -t "${DOCKER_TAG}" -f "${dockerfile}" .
  echo "Finished building ${APPLICATION_NAME}."
}

function task_usage {
    echo "Usage: $0"
    sed -n 's/^##//p' <$0 | column -t -s ':' |  sed -E $'s/^/\t/'
}

CMD=${1:-}
shift || true
RESOLVED_COMMAND=$(echo "task_"$CMD | sed 's/-/_/g')
if [ "$(LC_ALL=C type -t $RESOLVED_COMMAND)" == "function" ]; then
    pushd $(dirname "${BASH_SOURCE[0]}") >/dev/null
    $RESOLVED_COMMAND "$@"
else
    task_usage
fi