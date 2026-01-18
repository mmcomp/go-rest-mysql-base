# GO Base Restful with Gin and Mysql

Here is a sample Restful API backend with the following features:
- mysql
- migration
- JWT authentication
- rate limit
- user menu ACL

The `env.example` file is a sample environment variables needed by the code to function correctly.
You may copy it and rename it to `.env` and place it in the root path of the code and then run the code.
```
go run .
```

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

## Deploy
I added a simple `sh` file that simply runs all the commands needed to run a new version of the code.
The `deploy.sh` is the file, you are free to use it.