# IwIValidAPI

Official API from US Fachschaft IWI at HKA (Hochschule Karlsruhe) that provides valid user information for the RaumZeit tool.

## Overview

IwIValidAPI is a lightweight REST API built with Go that serves as a data provider for the RaumZeit application. It provides a secure endpoint to retrieve a list of valid users maintained by the student council (Fachschaft IWI) at HKA.

## Features

- 🔒 **Token-based Authentication** - Secure API access via Authorization header
- 📋 **User List Management** - Simple file-based storage of valid users
- 🚀 **Lightweight & Fast** - Built with Go for optimal performance
- 📝 **JSON Response Format** - Standard RESTful API responses
- 🔍 **Request Logging** - Middleware for tracking API usage

## Prerequisites

- Go 1.25.5 or higher
- A `.env` file with the required environment variables (see Configuration section)
- A `users` file containing the list of valid users (one username per line)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/fsiwi-hka/IwIValidAPI.git
cd IwIValidAPI
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the project root (see Configuration section below)

4. Create a `users` file in the project root with valid usernames (one per line):
```
user1
user2
user3
```

## Configuration

Create a `.env` file in the project root with the following variables:

```env
PORT=8080
API_TOKEN=your-secret-token-here
```

### Environment Variables

- `PORT` - The port on which the API server will run (e.g., 8080)
- `API_TOKEN` - Secret token required in the Authorization header for API authentication

**Note**: Make sure to use a strong, randomly generated token for `API_TOKEN` in production environments.

## Usage

### Starting the Server

Run the API server:
```bash
go run main.go
```

The server will start on the port specified in your `.env` file.

### API Endpoints

#### GET /GetValidUsers

Retrieves the list of valid users.

**Request Headers:**
- `Authorization`: Your API token (must match `API_TOKEN` in .env)

**Example Request:**
```bash
curl -X GET http://localhost:8080/GetValidUsers \
  -H "Authorization: your-secret-token-here"
```

**Success Response (200 OK):**
```json
{
  "users": [
    "user1",
    "user2",
    "user3"
  ]
}
```

**Error Response (400 Bad Request):**
```json
{
  "error": "No valid Token in Authorization header"
}
```

**Error Response (400 Bad Request):**
```json
{
  "error": "Error when Reading Current Valid Users"
}
```

## Project Structure

```
IwIValidAPI/
├── main.go           # Entry point, server setup, and routing
├── src/
│   └── routes.go     # HTTP handlers and middleware
├── go.mod            # Go module dependencies
├── go.sum            # Dependency checksums
├── .env              # Environment variables (not in version control)
├── .gitignore        # Git ignore rules
└── users             # List of valid users (one per line, not in version control)
```

## Security Considerations

- The `API_TOKEN` should be kept secret and never committed to version control
- The `.env` file is excluded from git via `.gitignore`
- The `users` file is excluded from git and should be managed separately
- All requests are validated via the `LoggingMiddleware` which checks for a valid Authorization header

## Integration with RaumZeit

<!-- TODO: Add information about how this API integrates with the RaumZeit tool -->
<!-- Please provide details about:
     - What RaumZeit uses the user data for
     - How often RaumZeit polls this API
     - Any specific requirements or expectations from RaumZeit -->

This API is designed to be consumed by the RaumZeit application. RaumZeit uses the valid user list to verify authorized users.

## Development

### Running in Development
```bash
go run main.go
```

### Building for Production
```bash
go build -o iwivalid-api
./iwivalid-api
```

## Dependencies

- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router and URL matcher
- [joho/godotenv](https://github.com/joho/godotenv) - Environment variable loader

## Contributing

<!-- TODO: Add contribution guidelines if this is an open-source project -->
<!-- Please specify:
     - How to report bugs
     - How to suggest features
     - Pull request process
     - Code style guidelines -->

## License

<!-- TODO: Add license information -->
<!-- Please specify the license under which this project is released -->

## Maintainer

**US Fachschaft IWI, HKA (Hochschule Karlsruhe)**

## Support

<!-- TODO: Add support contact information -->
<!-- Please provide:
     - Contact email or issue tracker
     - Documentation links
     - FAQ or troubleshooting guide -->

For questions or issues, please contact the Fachschaft IWI at HKA.
