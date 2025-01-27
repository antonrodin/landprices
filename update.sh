#!/bin/bash

git pull

env GOOS=linux CGO_ENABLED=1 go build -o landApp ./cmd/api

sudp supervisorctl restart landtitles