# Plutus API

This is a RESTful API built with Go, Echo, and Gorm based on problem definition

```
Code Test Details:

Task: Develop a API using Go and integrate GORM for database operations.

Requirements:

Implement CRUD (Create, Read, Update, Delete) operations for a basic resource (e.g., User, Transactions, etc.).
Use GORM for database operations.
Include proper error handling.
Provide endpoints to interact with the API (e.g., /transactions, /transactions/{id}).
Ensure the API follows RESTful principles.
```


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Download the dependencies:
```
go mod download
```

### Running the API

You can run the API using the `Makefile`:
```
make run
```
This will start the API on `localhost:8080`.

### Running the tests

You can run the tests using the `Makefile`:
```
make test
```
### Docker

You can also build and run the API using Docker:

```
make docker-run

```

This will build a Docker image and start a container, running the API on `localhost:8080`.

## API Endpoints

I didn't use swagger as wanted to keep it simple by using curl commands, the API provides the following endpoints:

- `POST /api/v1/users`: Create a new user
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"Ahmad Berahman", "email":"ahmad.berahman@hotmail.com"}' http://localhost:8080/api/v1/users
```
- `GET /api/v1/users/:id`: Get a user by ID
```
curl -X GET http://localhost:8080/api/v1/users/{id}
```
- `PUT /api/v1/users/:id`: Update a user by ID
```
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Pouria Berahman"}' http://localhost:8080/api/v1/users/{id}
```
- `DELETE /api/v1/users/:id`: Delete a user by ID
```
curl -X DELETE http://localhost:8080/api/v1/users/{id}
```

- `POST /api/v1/transactions`: Create a new transaction
```
curl -X POST -H "Content-Type: application/json" -d '{"userId":1, "amount":"100.50", "currency":"usd", "type":"credit"}' http://localhost:8080/api/v1/transactions
```

- `GET /api/v1/transactions/:id`: Get a transaction by ID
```
curl -X GET http://localhost:8080/api/v1/transactions/{id}
```

- `GET /api/v1/transactions`: Get all transactions
```
curl -X GET "http://localhost:8080/api/v1/transactions?page=1&pageSize=10"
```

- `PUT /api/v1/transactions/:id`: Update a transaction by ID
```
curl -X PUT -H "Content-Type: application/json" -d '{"amount":"2000.50", "type":"debit"}' http://localhost:8080/api/v1/transactions/{id}
```