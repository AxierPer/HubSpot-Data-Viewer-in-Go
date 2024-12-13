# HubSpot Data Viewer in Go

This project is a Go-based application that retrieves and displays data from HubSpot, rendering it through HTML templates. It leverages the Gin framework for routing and serving pages, and integrates with the HubSpot API to fetch and present dynamic data.

## Features

- **Go + Gin**: Uses the lightweight and efficient Gin framework for HTTP handling.
- **HTML Templates**: Employs Go's templating engine to load and render HTML templates dynamically.
- **Data from HubSpot**: Fetches data from the HubSpot API and displays it in user-friendly HTML views.
- **Modular Structure**: Organized file structure, with separate directories for configuration, routing, and static assets.
- **Static Assets**: CSS styles are separated into the `static/styles` directory, and templates reside in `static/templates`.

## Project Structure
```plaintext
api_hubspot_go/ 
├─ src/ 
│ ├─ config/ 
│ │ └─ config.go 
│ └─ routes/ 
│ └─ routes.go 
├─ static/ 
│ ├─ styles/ │ 
│ ├─ general.css 
│ │ └─ getId.css 
│ └─ templates/ 
│ ├─ general.html 
│ └─ getId.html 
├─ .env 
├─ .gitignore 
├─ README.md 
├─ go.mod 
├─ go.sum 
└─ main.go
```

- **`src/config`**: Configuration logic (e.g., loading environment variables, setting up credentials).
- **`src/routes`**: Defines routes and handlers that map URLs to specific logic and templates.
- **`static/styles`**: Contains CSS files for styling the rendered pages.
- **`static/templates`**: Houses HTML templates that present data to the user.
- **`main.go`**: The main entry point of the application, initializing the server and loading routes.

## Getting Started

1. **Install dependencies**:
   ```bash
   go mod tidy
   ```

2. **Set environment variables**:

- Edit the .env file to include your HubSpot token and other necessary environment variables.

3. **Run the application**:

    ```bash
    go run main.go
    ```

4. Access the application:

- Open your browser and navigate to:
`http://localhost:8080`

## Customization

- Modify `config.go` to adjust how environment variables are loaded.
- Update `routes.go` to create or alter endpoints according to your data fetching needs.
- Edit HTML templates in `static/templates/` to change the page layout or data presentation.
- Adjust CSS in `static/styles/` to achieve the desired look and feel.

## Additional Notes
- This project is intended as a starting point for integrating Go, Gin, and HubSpot data into a basic web interface.
- Further enhancements can include adding authentication, pagination, error handling, and improved UI/UX.
