# GO Base Restful with Gin and Mysql

Here is a sample Restful API backend with the following features:
- mysql
- migration
- JWT authentication
- rate limit
- user menu ACL
- open document

The `env.example` file is a sample environment variables needed by the code to function correctly.
You may copy it and rename it to `.env` and place it in the root path of the code and then run the code.
```
go run .
```

## MySQL
The mysql connection is set in the `.env` or environment vaiables as below:
```
DB_HOST=
DB_NAME=
DB_PORT=
DB_USER=
DB_PASSWORD=
```
`You need to have a working MySql Server and create a database for this to work.`

## Migration
If you want to create a migration you can run the following command in the rooy path of the code:
```
./migrate.new.sh MIGRATION-NAME
```
This way it will create two files in the `migrations` folder including the name that you provided and a number which should be unique.
One file is the `UP` migration and the other is `DOWN`.
You can edit these files and create you migration.

In order to run the migration you can call the following command:
```
go run . migration run
```
I added intial database schema as first migration so if you run above command and you sat the correct database configurations in environments, you will have the intial tables and a base Admin group and one menu (`users`) and one user with following credentials:
```
Username: admin
Password: admin
```

In order to run one step down in migration you can use the following command:
```
go run . migrate down
```

## Authentication
When user logs in, we generate two JWT tokens for him/her. One access and one refresh token.
You need to use access token to access the requied api and refresh for refreshing to new tokens.

## Rate limit
We have a simple in memory Rate Limit Middleware  which you can modify in the `routes.go` file.

## ACL
We designed a simple User Group Menu Access ACL. You may change it the way you see fit.

## Open Document (Swagger)
If you want to generate Open Doc for you api and if you provided correct docs for it, like in the existing controllers, you need to install https://github.com/swaggo/swag first.
The you may run the following commands:
```
swag init
swag fmt
```
This will update the existing `docs` folder and you can see this after running the code in the following address:
```
[YOUR ADDRESS]/swagger/index.html
```

## Deploy
I added a simple `sh` file that simply runs all the commands needed to run a new version of the code.
The `deploy.sh` is the file, you are free to use it.
