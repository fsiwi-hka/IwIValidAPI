# IwIValidAPI

Official API from us Fachschaft IWI at HKA (Hochschule Karlsruhe) that provides valid user information for the RaumZeit tool.

## Overview

IwIValidAPI is a lightweight REST API built with Go that serves as a data provider for the RaumZeit application. It provides a endpoint to retrieve a list of valid users maintained by the student council (Fachschaft IWI) at HKA.

## Usage

### API Endpoints

#### GET /GetValidUsers

Retrieves the list of valid users.

**Request Headers:**
- `Authorization`: Your API token (must match `API_TOKEN` in .env)

**Example Request:**
```bash
curl -X GET https://api.users.iwi-hka.de/GetValidUsers \
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




## Development

## Prerequisites
- Go 1.25.5 or higher
- A `.env` file with the required environment variables (see Configuration section)
- A `users` file containing the list of valid users (one username per line)


## Dependencies
- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router and URL matcher
- [joho/godotenv](https://github.com/joho/godotenv) - Environment variable loader

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

4. Create a `users` file in the project root with valid usernames (8 char):

```
ipsm1023
lore2234
sole3245
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

### Starting the Server

Run the API server:
```bash
go run main.go
```

The server will start on the port specified in your `.env` file.

## Project Structure

## Security Considerations

- The `API_TOKEN` should be kept secret and never committed to version control
- The `.env` file is excluded from git via `.gitignore`
- The `users` file is excluded from git and should be managed separately
- All requests are validated via the `LoggingMiddleware` which checks for a valid Authorization header


### Running in Development
```bash
go run main.go
```

### Building for Production
```bash
go build -o iwivalid-api
./iwivalid-api
```


## Maintainer
**US Fachschaft IWI, HKA (Hochschule Karlsruhe)**

For questions or issues, please contact the Fachschaft IWI at HKA.
