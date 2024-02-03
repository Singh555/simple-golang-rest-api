### Go RESTful API with PostgreSQL Integration

This Go program demonstrates the setup of a simple RESTful API using the Gin web framework and PostgreSQL as the database. The application includes database initialization, migrations, and a user creation endpoint.
How it Works

    The User struct represents the model for users with ID, Username, and Email fields.

    The initDB function initializes the PostgreSQL database connection, tests the connection, and runs migrations using the golang-migrate library.

    The runMigrations function uses golang-migrate to apply database migrations located in the "migrations" directory.

    The main function initializes the database, creates a Gin router, defines a route for creating a new user (/users), and starts the server.

    The createUser function handles the creation of a new user. It binds JSON data from the request body to the User struct, inserts the user into the database, and returns the created user with the inserted ID.

How to Run

To run the API, execute the following steps:

```bash

# Clone the repository
git clone https://github.com/your-username/your-repository.git
cd your-repository

# Install dependencies
go get -u github.com/gin-gonic/gin
go get -u github.com/golang-migrate/migrate/v4
go get -u github.com/golang-migrate/migrate/v4/database/postgres
go get -u github.com/jmoiron/sqlx
go get -u github.com/lib/pq

# Run the code
go run main.go
```
Ensure that you have Go installed on your machine.
Access the API

The API is accessible at localhost:8080. You can use tools like curl or Postman to interact with the endpoints.

    Create User Endpoint:

        URL: POST /users

        Request Body Example:

        json

{
"username": "example_user",
"email": "user@example.com"
}

Response Example:

json

        {
          "id": 1,
          "username": "example_user",
          "email": "user@example.com"
        }

Dependencies

The code relies on the following external dependencies:

    github.com/gin-gonic/gin: The Gin web framework for building web applications.
    github.com/golang-migrate/migrate/v4: A database migration tool.
    github.com/golang-migrate/migrate/v4/database/postgres: Database driver for PostgreSQL.
    github.com/jmoiron/sqlx: A set of extensions on top of database/sql for working with databases.
    github.com/lib/pq: PostgreSQL driver for Go's database/sql package.