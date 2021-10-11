# Contextual-guide

## Configuration

_(!) The entrypoint is in `src/cmd/server/main.go`_

### Http server

You can set app url in the `.app.env` file.

### Database

The database configuration is set with some env variables

| Key     | Exemple             |
| ------- | ------------------- |
| DB_HOST | localhost           |
| DB_PORT | 5432                |
| DB_NAME | contextual-guide    |
| DB_USER | postgresadmin       |
| DB_PASS | admin123            |

## OpenAPI

 1. [Install](https://goswagger.io/install.html) go-swagger

### Swagger



```bash
  swagger generate spec -o ./swagger.yaml --scan-models
```

### Serve specification
```bash
  swagger serve -F=swagger swagger.yaml 
```