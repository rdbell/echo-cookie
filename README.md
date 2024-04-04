# Echo Cookie

A simple package designed for managing browser cookies within the Echo web framework. It offers functionalities to set, get, and clear cookies, providing a streamlined approach to handling user sessions and preferences in Go web applications.

## Features

- **Set Cookies**: Easily set cookies with customizable names, values, and expiration times.
- **Get Cookies**: Retrieve the value of a specified cookie from a browser request.
- **Clear Cookies**: Invalidate a cookie by setting its expiration to a past date.

## Usage Example

## Build a server
```go
package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	cookie "github.com/rdbell/echo-cookie"
)

func setHandler(c echo.Context) error {
	// Set a cookie
	cookie.Set(c, "exampleCookie", "Hello, Echo!", nil)

	return c.String(http.StatusOK, "Cookie set")
}

func getHandler(c echo.Context) error {
	// Get the cookie
	value := cookie.Get(c, "exampleCookie")
	if value == "" {
		return c.String(http.StatusNotFound, "Cookie not found")
	}

	return c.String(http.StatusOK, "Cookie value: "+value)
}

func clearHandler(c echo.Context) error {
	// Clear the cookie
	cookie.Clear(c, "exampleCookie")

	return c.String(http.StatusOK, "Cookie cleared")
}

func main() {
	e := echo.New()

	// Set routes
	e.GET("/set", setHandler)
	e.GET("/get", getHandler)
	e.GET("/clear", clearHandler)

	// Start server
    err := e.Start(":8080")
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
```

## Run the server
```bash
go run main.go
```

## Test the server
```bash
curl -i http://localhost:8080/set
```


## Demo
Navigate to the [`example`](/example) folder within this repository and run `go run main.go` in your terminal. This command will launch an Echo server demonstrating cookie setting, getting, and clearing functionalities.

## Contribution

Contributions are welcome! If you'd like to improve or suggest new features, feel free to fork the repository, make your changes, and submit a pull request.

# License

This project is licensed under the MIT License. See the [LICENSE](/LICENSE) file for more details.
