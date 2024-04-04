package main

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"

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
	go func() {
		err := e.Start(":8080")
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for server to start
	for {
		_, err := http.Get("http://localhost:8080/")
		if err == nil {
			break
		}
	}

	testRequests()
}

func testRequests() {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	baseURL := "http://localhost:8080"

	// Set a cookie
	log.Println("Setting cookie...")
	makeRequest(baseURL+"/set", client)

	// Get the cookie
	log.Println("Getting cookie value...")
	makeRequest(baseURL+"/get", client)

	// Clear the cookie
	log.Println("Clearing cookie...")
	makeRequest(baseURL+"/clear", client)

	// Get the cookie
	log.Println("Getting cookie value...")
	makeRequest(baseURL+"/get", client)
}

func makeRequest(url string, client *http.Client) {
	response, err := client.Get(url)
	if err != nil {
		log.Printf("Failed to make request: %v\n", err)

		return
	}
	defer response.Body.Close()

	log.Printf("Headers from %s:", url)
	for name, values := range response.Header {
		// Loop over all headers and print them if they are related to cookies
		if name == "Set-Cookie" || name == "Cookie" {
			log.Printf("%v: %v\n", name, values)
		}
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	log.Printf("Response: %s\n", string(body))
	log.Println("----------------")
}
