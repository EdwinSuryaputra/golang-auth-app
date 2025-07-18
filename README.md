# iam-api (Identity and Access Management)
Originally made by Iwibowo Team.

This repository provides an Identity and Access Management (IAM) system for authentication, authorization, and user management. It includes features like role-based access control (RBAC), JWT authentication with Redis storage, and integration with GoFiber and Fx for scalable and maintainable architecture.    

## Migrations
We use github.com/golang-migrate/migrate for migrations. All SQL migrations are stored in the `migrations` directory. To create a new migration file:
```
make generate-sql-migration name=<input_the_file_name>
```

Up files are used to construct the latest database schema version. Down files are used to deconstruct or rollback to previous schema version.

To execute SQL migrations in the `migrations` directory, run the following command:
```
make migrate-sql
```

To generate `gorm.io/gen` files into `/app/adapters/sql/gorm/`, run the following command:

```
make generate-gorm-model
```

This will generate all gorm gen files based on the migration files in `migrations` directory. 
This is done automatically using ephemeral mysql docker container to run the migration and generate the gorm files.

## Build & Run
To run this app, run the following command:
```
make run
```