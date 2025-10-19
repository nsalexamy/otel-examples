#!/bin/bash


REPO_NAME="otel-go-example"



usage() {
  echo ''
  echo ''
  echo "Usage: ${0} [-u dockerhub-username] [-p dockerhub-password] [-t tag]"
  echo ''
  echo 'options'
  echo '  -u : Docker username'
  echo '  -p : Docker password'
  echo '  -t : Tag for the docker image'
  echo '  -h : Display this help message'

  exit 1
}

DOCKERHUB_USERNAME="${DOCKERHUB_USERNAME:-}"
DOCKERHUB_PASSWORD="${DOCKERHUB_PASSWORD:-}"
IMAGE_TAG="${IMAGE_TAG:-}"

while getopts "hu:p:t:" OPTION
do
  case ${OPTION} in
    u) DOCKERHUB_USERNAME="${OPTARG}" ;;
    p) DOCKERHUB_PASSWORD="${OPTARG}" ;;
    t) IMAGE_TAG="${OPTARG}" ;;
    h) usage ;;
    ?) usage ;;
  esac
done

# This script builds and pushes the Docker image to Docker Hub.

# check if two arguments are passed
#if [ "$#" -ne 2 ]; then
#    echo "Usage: $0 <docker-username> <docker-password>"
#    exit 1
#fi
#
#DOCKERHUB_USERNAME=$1
#DOCKERHUB_PASSWORD=$2

IMAGE_TAG=${IMAGE_TAG:-"0.1.0"}

# check all arguments are provided
if [ -z "${DOCKERHUB_USERNAME}" ] || [ -z "${DOCKERHUB_PASSWORD}" ] || [ -z "${IMAGE_TAG}" ]; then
  echo "All arguments are required"
  usage
fi

echo "Docker Hub Username: $DOCKERHUB_USERNAME and Image Tag: $IMAGE_TAG"



docker login -u $DOCKERHUB_USERNAME -p $DOCKERHUB_PASSWORD


# get working directory name
WORKING_DIR=$(basename "$PWD")

BACKEND_DIR="."


echo "Backend directory: $BACKEND_DIR"


docker buildx build --platform linux/amd64,linux/arm64  -t $DOCKERHUB_USERNAME/$REPO_NAME:$IMAGE_TAG $BACKEND_DIR --push