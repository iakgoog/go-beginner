package webserver
import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


// RunWebServer function
func RunWebServer() {
	e := routers.Router()
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
