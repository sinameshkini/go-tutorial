# Go Tutorial Microservice

This project is a simple microservice built with Go, following Clean Architecture principles. It utilizes the following technologies:

- **GORM**: For interacting with the database
- **Echo**: A fast and minimalist web framework
- **Logrus**: A structured logger for Go
- **Cobra**: For building command-line applications
- **Viper**: For configuration management
- **JWT**: For handling authentication and authorization

## Table of Contents

- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [API Design](#api-design)
- [Contributing](#contributing)
- [License](#license)

## Project Structure

The project follows the Clean Architecture pattern, separating business logic from delivery mechanisms. Here's an overview of the folder structure:

```plaintext
.
├── cmd                         # Command-line related files for the service
│   └── main.go                 # Main entry point of the application
├── configs                     # Configuration files (YAML, JSON, etc.)
│   └── config.yaml             # Main configuration file
├── internal                    # Internal application code (business logic, domain, etc.)
│   ├── domain                  # Domain layer (entities, repository interfaces, etc.)
│   │   └── user.go             # Example domain entity
│   ├── usecase                 # Application use cases (business rules)
│   │   └── auth_usecase.go     # Example use case for authentication
│   ├── repository              # Data access layer (database interaction)
│   │   └── user_repository.go  # Example repository implementation for user
│   ├── delivery                # Delivery layer (controllers, handlers)
│   │   └── http                # HTTP handlers using Echo
│   │       └── auth_handler.go # Example HTTP handler for authentication
│   └── service                 # Service layer (integration with external services)
│       └── auth_service.go     # Example service implementation for authentication
├── migrations                  # Database migration files
│   └── 202309010001_init.sql   # Example SQL migration file
├── pkg                         # External packages that can be reused (optional)
│   ├── logger                  # Custom logger using Logrus
│   │   └── logger.go           # Logger setup
│   └── config                  # Viper configuration setup
│       └── config.go           # Viper setup
├── scripts                     # Utility scripts (e.g., setup, deployment)
│   └── migrate.sh              # Example migration script
├── go.mod                      # Go module definition
├── go.sum                      # Go dependencies lock file
└── README.md                   # This README file
```

### Key Folders

- **cmd/**: Entry point for the application, contains the `main.go` file and Cobra command definitions.
- **configs/**: Configuration files for different environments (e.g., development, production). Managed using **Viper**.
- **internal/**: Core application code, divided into several layers:
    - **domain/**: Defines domain entities and repository interfaces.
    - **usecase/**: Contains business logic and application use cases (e.g., authentication use cases like sign-up, sign-in, reset password).
    - **repository/**: Implements repository interfaces for interacting with the database using GORM.
    - **delivery/**: Contains HTTP handlers using Echo, including authentication APIs.
    - **service/**: Handles external service integrations (e.g., sending email for password reset).
- **pkg/**: Optional, but can include reusable packages such as the custom **Logrus** logger or **Viper** configuration setup.
- **migrations/**: Database migration files.
- **scripts/**: Utility scripts for setting up or managing the service.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/sinameshkini/go-tutorial.git
   cd go-tutorial
   ```

2. Install the dependencies:

   ```bash
   go mod tidy
   ```

3. Set up your database and run the migrations:

   ```bash
   ./scripts/migrate.sh
   ```

## Usage

To run the application, simply use the following command:

```bash
go run cmd/main.go
```

Alternatively, you can build the application and run the binary:

```bash
go build -o go-tutorial cmd/main.go
./go-tutorial
```

### Running with Cobra

Cobra allows you to define and use CLI commands. For example, to run a specific command:

```bash
go run cmd/main.go <command>
```

## Configuration

The application uses **Viper** for configuration management, which supports environment variables, configuration files, and more. The primary configuration file is located at `configs/config.yaml`.

Example `config.yaml`:

```yaml
server:
  port: 8080

database:
  host: localhost
  port: 5432
  user: user
  password: password
  dbname: go_tutorial

jwt:
  secret_key: your-secret-key
  token_expiration: 3600  # Token expiration in seconds

log:
  level: info
```

## API Design

The Authentication service exposes three main APIs: **Sign Up**, **Sign In**, and **Reset Password**. These APIs use JWT for authentication and authorization.

### 1. Sign Up API

**Endpoint**: `/auth/signup`  
**Method**: `POST`  
**Request Body**:

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response**:

- **201 Created**: User successfully registered.
- **400 Bad Request**: Invalid input (e.g., email already exists).

### 2. Sign In API

**Endpoint**: `/auth/signin`  
**Method**: `POST`  
**Request Body**:

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response**:

- **200 OK**: Successful login, returns JWT token.

  ```json
  {
    "token": "jwt-token-here"
  }
  ```

- **401 Unauthorized**: Invalid email or password.

### 3. Reset Password API

**Endpoint**: `/auth/reset-password`  
**Method**: `POST`  
**Request Body**:

```json
{
  "email": "user@example.com"
}
```

**Response**:

- **200 OK**: Email sent with reset instructions.
- **404 Not Found**: Email not found.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes.

### Steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/my-feature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/my-feature`).
5. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
