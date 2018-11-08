#!/bin/bash

env GOOS=linux GOARCH=amd64 go build -o rabbitmq-data-copier *.go