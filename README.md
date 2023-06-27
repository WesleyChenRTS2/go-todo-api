# TODO App

This is a TODO app implemented in Go using the Echo web framework and a SQL database. It provides RESTful endpoints to manage todo items.

## Installation

1. Clone the repository

2. Install dependencies: `go mod download`

3. Set up the database:

- Create a PostgreSQL database.
- Update the database connection details in the `dbconf.yml` & `main.go` file.
<!-- TODO: Make one source of truth -->
- Run migrations: `goose -dir db/migrations postgres "$(go run db/readconf.go)" up`

4. Build and run the app: `air`

The app will start running on http://localhost:8080.

## Endpoints

### Create a new todo

- URL: `POST /todo`
- Request body:

`{
"title": "New Todo Item",
"description": "This is a new todo item."
}`

### Get a single todo

- URL: `GET /todo/{id}`

### Get all todos

- URL: `GET /todos`

### Update a todo

- URL: `PUT /todo/{id}`
- Request body:

`{
"title": "Updated Todo Item",
"description": "This is an updated todo item."
}`

### Delete a todo

- URL: `DELETE /todo/{id}`

## Database

The app uses a PostgreSQL database to store the todo items. You can find the database schema in the `db.sql` file.

## Dependencies

The following dependencies are used in this project:

- `github.com/labstack/echo/v4`: Web framework for building the RESTful API.
- `github.com/lib/pq`: PostgreSQL driver for Go's database/sql package.
- `github.com/cosmtrek/air`: Live reloading for Go in developement

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
