package main

import "teachEcho/router"

func main() {
	// Echo instance
	/*e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8021"))*/
	router.Run()
}
