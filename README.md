# Configuration files

Database connection config: ./cmd/gin-pgx-api/config.json; The JSON file is passed to the program itself, which performs the connection.

Database container config: ./.env; The environment file is passed to the docker container to configure the database.

**If you are using a database container, you need to make sure that these files have the same parameters, with the container host always being "pg" and the port being 5432.**

Ideally, they should be added to .gitignore. I didn't add them to make it easier for you to run.

# Run

Docker:

```
$ docker-compose up --build
```

After that, you can use swagger via `http://localhost:8000/swagger/index.html`.

# Structure brief

./cmd/gin-pgx-api: A program that parses a database configuration file, opens a connection to the database based on the config and starts the service. In short, it is a something like launcher.

./internal/domain: Domain models.

./internal/logic: Business logic, service interface, repository interface.

./internal/ginapi: API based on Gin framework.

./internal/pgx: Repository implementation. `pgxpool` is used.
