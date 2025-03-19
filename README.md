# Expense Tracker API

This project is a RESTful API for an expense tracker application built with Go. It demonstrates basic CRUD operations for expenses, including creating, reading, updating, and deleting expenses. The API also supports user authentication using JWT.

## Features

- Sign up as a new user.
- Generate and validate JWTs for handling authentication and user sessions.
- List and filter your past expenses.
- Add a new expense.
- Remove existing expenses.
- Update existing expenses.

## API Endpoints

### Sign Up

Sign up as a new user using the POST method.

```
POST /signup
{
  "username": "john_doe",
  "password": "password123"
}
```

### Login

Login as an existing user using the POST method.

```
POST /login
{
  "username": "john_doe",
  "password": "password123"
}
```

The endpoint returns a JWT token to be used for authenticated requests.

### Create Expense

Create a new expense using the POST method.

```
POST /expenses
{
  "amount": 50.0,
  "category": "Groceries",
  "note": "Bought groceries",
  "date": "2025-03-18T12:00:00Z"
}
```

### Get Expenses

Get all expenses for the authenticated user using the GET method.

```
GET /expenses
```

### Get Expense

Get a single expense by its ID using the GET method.

```
GET /expenses/{id}
```

### Update Expense

Update an existing expense using the PUT method.

```
PUT /expenses/{id}
{
  "amount": 60.0,
  "category": "Groceries",
  "note": "Updated groceries expense",
  "date": "2025-03-18T12:00:00Z"
}
```

### Delete Expense

Delete an existing expense using the DELETE method.

```
DELETE /expenses/{id}
```

## Installation

1. Clone the repository:
```
git clone https://github.com/your-username/expense-tracker-api.git
```

2. Change to the project directory:
```
cd expense-tracker-api
```

3. Install the dependencies:
```
go mod tidy
```

4. Start the server:
```
go run main.go
```

## Usage

Once the server is running, you can use an API client like Postman to test the endpoints.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

https://roadmap.sh/projects/expense-tracker-api