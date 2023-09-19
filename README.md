# Go Vue Starter

Create env.json with database and server configuration for local development:

```
{
	"dbUser": "root",
	"dbPass": "qazqaz",
	"dbName": "test",
	"httpPort": "2020",
	"versionStamp": ""
}
```

Version stamp is updated automatically when running `restart.sh`.

Run `sh init.sh` to build the frontend, start the server, and initialize the database.

Run `sh restart.sh` to rebuild the frontend and start the server.

You may have to delete package-lock.json and re-install the dev dependencies
listed in package.json at their latest versions to get a working build.

All files in css, img, and js are public.

## Add libraries

```
go mod init github.com/macu/go-vue-starter
go get github.com/jackc/pgx/v4/stdlib
```
## Mac setup

A postgres password must be set before the app can connect.

```
$ brew install postgresql go node jq
$ psql postgres
postgres=# ALTER USER matt WITH PASSWORD 'somepassword';
postgres=# \q
$ cd starterdemo
// Update env.json
$ npm install
$ sh init.sh
```

## Create Postgres database

```
$ psql postgres
postgres=# CREATE DATABASE starterdemo;
```

### Access database from command line

```
$ psql starterdemo
postgres=# \dt
postgres=# \d+ user_account
```

## Build and run

On first use, or to re-initialize database:
```
$ sh init.sh
```

To rebuild client and server, and run:
```
$ sh restart.sh
```
