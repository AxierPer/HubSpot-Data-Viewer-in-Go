# HubSpot API Integration in Go

This project is a simple Go-based API that integrates with HubSpot to fetch and expose its data. Using the Gin framework, it provides a lightweight and RESTful interface for consuming HubSpot services and delivering JSON-based responses.

## Features

- **HubSpot API Integration**: Connects to HubSpot to fetch data such as contacts, deals, or custom objects.
- **Go + Gin**: Uses the Gin framework for efficient routing and handling of HTTP requests.
- **JSON-based Responses**: The API exclusively returns JSON data, ideal for integration with other applications or frontend frameworks.
- **Environment Variables**: Securely manages sensitive information like HubSpot tokens using environment variables.

## Project Structure
  ```plaintext
  api_hubspot_go/ 
  ├─ src/ 
  │ ├─ config/ 
  │ │ └─ config.go 
  │ └─ routes/ 
  │ └─ routes.go 
  ├─ .env 
  ├─ .gitignore 
  ├─ README.md 
  ├─ go.mod 
  ├─ go.sum 
  └─ main.go
  ```

- **`src/config`**: Handles configuration, including loading environment variables (e.g., HubSpot API tokens).
- **`src/routes`**: Contains API routes that interact with HubSpot's endpoints and return JSON responses.
- **`main.go`**: Initializes the Gin server and sets up routing for the API.

## Prerequisites

1. **HubSpot Account**: You must have a HubSpot account and an API key or access token to use their services.
2. **Go Environment**: Ensure that Go is installed on your system.

## Getting Started

1. **Install dependencies**:
   
   ```bash
   go mod tidy
   ```

2. **Set environment variables**:

- Create a .env file in the root of your project with the following content:
    
    ```makefile
    TOKEN_HUBSPOT=your_hubspot_api_toke
    ```

- Replace your_hubspot_api_token with your HubSpot API token.

3. **Run the application**:

    ```bash
    go run main.go
    ```

4. **Access the application**:

- Open your browser and navigate to:
`http://localhost:8080`

5. **Test the API**:

- Use tools like curl, Postman, or any HTTP client to send requests to the API.
- Example endpoint:

  ```bash
  GET http://localhost:8080/contacts
  ```
  This might fetch a list of contacts from your HubSpot account.

## Example Endpoints

- **GET /contacts**: Retrieves a list of contacts from HubSpot.
- **GET /contacts/**:id: Fetches details for a specific contact by ID.
- **POST /custom-object**: Adds a custom object to HubSpot (example functionality).
Endpoints can be customized in `routes/routes.go`.

## Customization
- Add or modify routes in `routes.go` to extend the API functionality.
- Use `config.go` to manage additional environment variables or configuration settings.
- Extend the API to consume other HubSpot endpoints like Deals, Companies, or Pipelines.

## Additional Notes
- This API acts as a middleware to interact with HubSpot services, abstracting the complexity of their API.
- Designed for scalability, you can build additional microservices or frontend applications on top of this API.

# License
This project is open source and available under the MIT License.

