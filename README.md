# Internal Transfers API

A Go-based HTTP API that processes internal financial transfers between accounts, persisting account state and transaction logs in PostgreSQL.

The system supports:

* Account creation
* Account balance query
* Internal transfers between accounts
* Database migrations via golang-migrate
* Unit tests for service layer
* Dockerized setup (API + PostgreSQL)
  
---
## 🏗 Architecture Overview

The project follows clean architecture principles with clear separation of concerns:
```

internal
 ├── database        # DB connection and migrations
 ├── domain          # Domain errors and mappings
 ├── handler         # HTTP handlers (entrypoints)
 ├── model           # Domain models
 ├── repository      # Database implementations
 └── service         # Business logic + unit tests
```

### Design Principles

* Clear layering: handler → service → repository
* Dependency inversion via repository interfaces
* Explicit error mapping from domain to HTTP
* Transaction-safe database operations
* Testable business logic (unit tests with mocked repositories)

# 
## 🚀 Features
**1. Create Account**

**POST** `/accounts`

Creates a new account with an initial balance.

**Request**

```
{
  "account_id": 123,
  "initial_balance": "100.23344"
}
```

**Response**

* `201 Created` (empty body)
* Proper error response if invalid input or duplicate account
---
**2. Get Account Balance**

**GET** `/accounts/{account_id}`

**Response**
```
{
  "account_id": 123,
  "balance": "100.23344"
}
```

* `200 OK` if found

* `404 Not Found` if account does not exist

---
**3. Submit Transaction**

**POST** `/transactions`

Transfers funds between two accounts.

**Request**
```
{
  "source_account_id": 123,
  "destination_account_id": 456,
  "amount": "100.12345"
}
```

**Behavior**

* Validates input

* Ensures source account has sufficient balance

* Executes atomic database transaction

* Updates balances

* Persists transaction log

**Response**

* `201 Created` (empty body)

* Error response if:

  * Invalid amount
  
  * Same source/destination
  
  * Insufficient funds
  
  * Account not found

  * Processing failure
---
## **💾 Database**

PostgreSQL 16 is used.

**Migrations**

Migrations are managed using `golang-migrate`.

They run automatically during application startup.

Tables:

* `accounts`

* `transactions`
---
## **🐳 Running with Docker**

**1. Requirements**

* Docker

* Docker Compose

**2. Start the application**

```docker-compose up --build```

Services:

* API: http://localhost:8080

* PostgreSQL: Can be accessed through terminalby typing
  ``` docker exec -it transfers-db psql -U postgres -d transfers ```

Database credentials:

```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=transfers
```
---
## **🧪 Running Tests**

Unit tests exist for:

* account_service

* transaction_service

Run locally:

```
go test ./internal/service/
```

Tests cover:

* Account creation

* Balance retrieval

* Transfer validation

* Insufficient funds

* Invalid inputs

* Transaction integrity rules
---

## **🧠 Assumptions**

* All accounts operate in the same currency.

* No authentication or authorization is required.

* Balances and amounts are stored using string decimal representation to preserve precision.

* All monetary operations are validated before persistence.

* Transfers are executed inside a single database transaction to guarantee atomicity.

* Concurrent transfers maintain consistency through database-level guarantees.

---
## **🔐 Data Integrity**

The system guarantees:

* Atomic transactions (BEGIN/COMMIT/ROLLBACK)

* No partial balance updates

* Consistent account states

* Proper error propagation

---
## **📌 Error Handling**

Domain errors are centralized and mapped to HTTP status codes.

Examples:

* ErrInvalidAmountFormat

* ErrInsufficientFunds

* ErrAccountNotFound

* ErrSameAccountTransfer

This keeps handlers clean and business rules encapsulated.

---
## **🧼 Code Quality**

* Clean code principles applied

* Small, focused functions

* Explicit interfaces for testability

* Clear package separation

* Unit-tested service layer

* No framework lock-in

---
## **📈 Future Improvements**

* Integration tests

* Structured logging

* Observability (metrics/tracing)

* Idempotency keys for transfers

* Pagination for transaction history

* Rate limiting

---
## **📄 License**

MIT License
