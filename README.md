# Go Vue Starter

Create env.json with database and server configuration:

```
{
	"dbUser": "root",
	"dbPass": "qazqaz",
	"dbName": "test",
	"httpPort": "2020"
}
```

Run `sh init.sh` to build the frontend, start the server, and initialize the database.

Run `sh restart.sh` to rebuild the frontend and start the server.

You must create a testing database and update the connection info in main.go:
`db, err := sql.Open("mysql", "root:qazqaz@/test?charset=utf8&parseTime=true")`

You may have to delete package-lock.json and re-install the dev dependencies
listed in package.json at their latest versions to get a working build.

All files in css, img, and js are public.
