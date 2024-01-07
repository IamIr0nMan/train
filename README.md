# Train Ticket Booking System (gRPC)

This project implements a train ticket booking system using Golang and gRPC. It allows users to purchase tickets, allocate seats, view receipts, manage seat allocations, and more.

## How to run this project?

1. Clone this repository:

```bash
git clone https://github.com/IamIr0nMan/train.git
```

2. Download dependencies:

```go
go mod download
```

3. Running the server:

```go
go run server/server.go
```

4. Running the client:

```go
go run client/client.go
```

## Unit Testing

> Current test coverage: ~**90**%

- Running the unit tests with coverage report:

  ```go
  go test -coverprofile=coverage.out && go tool cover -html=coverage.out
  ```
