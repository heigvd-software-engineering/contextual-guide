# Contextual-guide

## Configuration

_(!) The entrypoint is in `src/cmd/server/main.go`_
### Http server

You can set the http port with the command `-port=<the port>`

### Database

The database configuration is set with some env variables

| Key     | Exemple             |
| ------- | ------------------- |
| DB_HOST | localhost           |
| DB_PORT | 5432                |
| DB_NAME | contextual-guide    |
| DB_USER | postgresadmin       |
| DB_PASS | admin123            |
