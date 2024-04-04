package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// DefaultCookieExpiration determines how long a cookie is valid by default.
const DefaultCookieExpiration = 24 * time.Hour * 365 * 10

// Set sets a browser cookie.
func Set(c echo.Context, name string, value string, expiration *time.Time) {
	var exp time.Time
	if expiration == nil || expiration.IsZero() {
		exp = time.Now().Add(DefaultCookieExpiration)
	} else {
		exp = *expiration
	}

	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Path:    "/",
		Expires: exp,
	}

	// Set the cookie in the provided context
	c.SetCookie(cookie)
}

// Get gets a browser cookie.
func Get(c echo.Context, name string) string {
	cookie, err := c.Cookie(name)
	if err != nil {
		return ""
	}

	return cookie.Value
}

// Clear unsets a browser cookie.
func Clear(c echo.Context, name string) {
	cookie := &http.Cookie{
		Name:    name,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}

	c.SetCookie(cookie)
}
