#!/bin/bash
echo "Downloading dependencies."
go get .
echo "Building cligpt"
go build -ldflags "-s -w" -o cligpt main.go
