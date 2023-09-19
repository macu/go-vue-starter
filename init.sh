#!/bin/sh

# NOTE The database and user role must already have been created (see sql/init.pgsql).
# Local DB connection info is extracted from env.json

# Prompt for username and password
read -p "Enter a test username [test]: " username
read -sp "Enter a test password [testpass]: " password
echo
read -p "Enter a test email [test@test.com]: " email

# Default each to "test"
username=${username:-test}
password=${password:-testpass}
email=${email:-"test@test.com"}

# Record build date
sh update-build-date.sh

# Build frontend
npm run prod || { echo 'Client code compilation failed.' ; exit 1; }

# Initialize database
dbuser=$(jq -r '.dbUser' env.json)
dbpass=$(jq -r '.dbPass' env.json)
dbname=$(jq -r '.dbName' env.json)
PGPASSWORD=$dbpass psql -U $dbuser -d $dbname -f ./sql/init.pgsql

# Run backend and init db and user
go run ./ -createUser -username="$username" -password="$password" -email="$email"
