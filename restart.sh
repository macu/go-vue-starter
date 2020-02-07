#!/bin/sh

# Exit when any command fails
set -e

# Build frontend
npm run prod

# Run backend
go run *.go
