#!/bin/bash

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker build -t evearakaki/spliff:"$TRAVIS_COMMIT"
docker push  evearakaki/spliff:"$TRAVIS_COMMIT"
