#!/bin/sh

sh update-build-date.sh

npm run prod || { echo 'Client code compilation failed.' ; exit 1; }

go run *.go
