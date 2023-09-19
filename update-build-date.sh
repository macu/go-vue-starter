#!/bin/sh

DATETIME=$(date +'%Y%m%d%H%M')

# Replace versionStamp in local env file
sed -i.previous -e "s/\"versionStamp\": \"[0-9]*\"/\"versionStamp\": \"$DATETIME\"/" env.json

rm env.json.previous
