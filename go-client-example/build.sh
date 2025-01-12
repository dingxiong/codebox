#!/bin/bash

GOOS=linux GOARCH=amd64 go build -tags netgo -ldflags '-extldflags "-static"' && kubectl cp goclient-example xiong-go-client-test-5fb4c98bf5-6ftgn:/