
# Icepeak

**Icepeak** is a lightweight, modular web framework written in Go. It provides a clean and flexible structure for building web applications with simplicity, performance, and ease of use in mind.

## Features

- **Routing**: Simple and intuitive routing system with support for dynamic parameters, middleware, and named routes.
- **Middleware**: Easy-to-use middleware support for extending request/response handling.
- **Modular Design**: Organized project structure with clear separation of concerns.
- **High Performance**: Built on Go, leveraging its speed and concurrency model.

## Getting Started

### Prerequisites

- Go 1.17 or higher

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/icepeak.git
   cd icepeak
   ```

2. **Initialize Go Modules:**

   ```bash
   go mod tidy
   ```

3. **Run the Application:**

   ```bash
   go run main.go
   ```

4. **Test the Routes:**

   - Open your browser or use `curl` to test the routes:

     ```bash
     curl http://localhost:8080/
     ```
     Should return: `Welcome to Icepeak!`

     ```bash
     curl http://localhost:8080/hello
     ```
     Should return: `Hello, Icepeak!`

## Project Structure

```plaintext
icepeak/
├── app/
│   ├── controllers/     # Application controllers
│   ├── models/          # Application models
│   ├── middlewares/     # Custom middlewares
│   ├── views/           # Templating engine views (HTML, etc.)
│   └── services/        # Business logic and service classes
├── bootstrap/
│   └── init.go          # Application initialization and bootstrap logic
├── config/
│   ├── app.yaml         # Application-level configuration (env, debug mode, etc.)
│   ├── database.yaml    # Database configuration (connections, pools, etc.)
│   └── routes.yaml      # Define application routes in a structured format
├── core/
│   ├── routing/         # Core routing logic (handling requests, parameters, etc.)
│   ├── orm/             # Object Relational Mapping (ORM) layer
│   ├── validation/      # Validation utilities
│   ├── middleware/      # Core middlewares (like auth, CORS, etc.)
│   ├── cache/           # Cache management (using in-memory, Redis, etc.)
│   ├── logging/         # Logging utilities
│   ├── response/        # HTTP response handling utilities
│   └── utils/           # Generic utility functions (helpers)
├── database/
│   ├── migrations/      # Database migrations
│   └── seeds/           # Database seeding
├── public/
│   └── assets/          # Static assets (CSS, JS, images)
├── storage/
│   ├── logs/            # Log files
│   ├── uploads/         # User uploads or other files
│   └── cache/           # Cached files
├── tests/
│   └── integration/     # Integration tests for the framework
│   └── unit/            # Unit tests for framework components
├── go.mod               # Go module file
├── go.sum               # Dependency management file
└── main.go              # Entry point of the framework
```

## Contributing

Feel free to contribute to Icepeak! If you have suggestions or find any bugs, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## What We Learn?

- How to set up a basic web framework in Go.
- How to handle routing, middleware, and dynamic requests.
- How to structure a scalable and maintainable web application.

---

*Icepeak* is inspired by Laravel but built on the powerful Go language, aiming to provide a clean and efficient development experience!
