#!/bin/sh

go build -o bin/paperrulers-linux-amd64 -ldflags="-s -w" 
strip bin/paperrulers-linux-amd64

