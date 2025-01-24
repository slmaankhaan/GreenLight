# Greenlight API

Greenlight is a JSON API for retrieving and managing information about movies. It provides endpoints to manage movies, users, authentication, and more. The API supports CRUD operations for movie data, user registration and activation, password reset functionality, and authentication via tokens.

## API Endpoints

### Movies
- `GET /v1/movies` - Retrieve all movies
- `POST /v1/movies` - Create a new movie
- `GET /v1/movies/:id` - Retrieve a specific movie by ID
- `PATCH /v1/movies/:id` - Update a specific movie by ID
- `DELETE /v1/movies/:id` - Delete a specific movie by ID

### Users
- `POST /v1/users` - Register a new user
- `PUT /v1/users/activated` - Activate a specific user
- `PUT /v1/users/password` - Update a user’s password

### Tokens
- `POST /v1/tokens/authentication` - Generate a new authentication token
- `POST /v1/tokens/password-reset` - Generate a password reset token

### Health Check
- `GET /v1/healthcheck` - Check the health status and version of the application

### Metrics
- `GET /debug/vars` - Display application metrics

## Project Setup

### Prerequisites
- Go 1.16 or higher
- PostgreSQL for the database
- SMTP server for sending emails (for user activation and welcome emails)

### Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/greenlight-api.git
   cd greenlight-api
   ```

2. Set up your environment variables:
   - `GREENLIGHT_DB_DSN` - Connection string for PostgreSQL (e.g., `postgres://user:password@localhost:5432/greenlight`)
   - SMTP configuration for sending emails

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run ./cmd/api
   ```

   The API will be running on `http://localhost:4000`. You can access the endpoints by sending requests to this address.

## Database Setup

Make sure you have PostgreSQL running locally or remotely. Configure the connection string in the `.env` file or set the `GREENLIGHT_DB_DSN` environment variable. Use the provided migrations to set up the necessary tables:

```bash
make db/migrations/up
```

## Quality Control

To run quality checks on the code (tidy dependencies, format, vet, staticcheck, etc.):

```bash
make audit
```

## Vendoring Dependencies

To vendor dependencies and ensure they are available locally:

```bash
make vendor
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
