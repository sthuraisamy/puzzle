#!/bin/bash
# Script to build puzzle binaries for multiple platforms including MacOS, Linux and Windows
# The build command compile write the binaries into bin folder
# MacOSX - puzzle_osx
# Linux - puzzle_lnx
# Windows - puzzle_win
env GOOS=darwin GOARCH=amd64 go build -o ./bin/puzzle_osx puzzle.go
env GOOS=linux GOARCH=amd64 go build -o ./bin/puzzle_lnx puzzle.go
env GOOS=windows GOARCH=amd64 go build -o ./bin/puzzle_win puzzle.go