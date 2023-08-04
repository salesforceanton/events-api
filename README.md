# events-api
# Web-server to perform CRUD operations with Event object

### Implemented Endpoints:

```
1. auth/sign-up    POST   - create user in db
2. auth/sing-in    POST   - check is user with defined credential exist and return 12h acceess token
3. api/events/     GET    - get all events which orhanizer - current user
4. api/events/:id  GET    - get event by id if current user is organizer
5. api/events/:id  POST   - update event record
6. api/events/     POST   - create event record and organizer will be current user automatically
7. api/events/:id  DELETE - delete event record

All CRUD operations via events-api check user-record access (only organizer can change/delete event record)
Auth middleware is also included - check user via token and persist it to execution context
```

### Stack:

```
1. GO
2. PostgresSQL
3. Gin
4. Migrate
```

### Setting up:

Set on your `.env` file variables according with `.env.example` file
To configure db - we can use migrate - files to up/down migration is included:

```migrate -path schema/ -database postgres://${EVENTSAPI_DB_USERNAME}:${EVENTSAPI_DB_PASSWORD}@${EVENTSAPI_DB_HOST}/${EVENTSAPI_DB_NAME} up```

** In this example I have used elephantSQL free Postrgres service

To run server use comman

```go run cmd/main.go```

### Swager

Swagger is also included
To test via this tool - open swagger on PORT which web-server is working on 

```http://${HOST}:${PORT}/swagger/index.html```