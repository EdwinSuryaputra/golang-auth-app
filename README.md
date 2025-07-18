#  golang-auth-app

This repository provides an Identity and Access Management (IAM) system for authentication, authorization, and user management. It includes features like role-based access control (RBAC), JWT authentication with Redis storage, and integration with GoFiber and Fx for scalable and maintainable architecture.    

## Documentation
### 📃 Swagger
```
http://localhost:3500/api-docs
```

## Makefile Commands

### ⚓ Docker Deployment 
```
make deploy-docker
```

### 🔧 Generate Gorm Model:
```bash
make generate-model
```

### 🧱 Create New Migration

```
make generate-migration name=your_migration_name
```
Generates a new migration file. Replace your_migration_name with a descriptive name.


### 📦 Apply Migrations
```
make migrate-sql
```
Applies all pending migrations to the database (production-safe).

### 🚀 Run Server
```
make run
```
Starts the development server using <i>go run main.go</i>

