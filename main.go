package main

import (
	"embed"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"goro/model"
	"goro/service"
	"goro/views"
	"goro/worker"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed static
var static embed.FS

func Static() http.Handler {
	dist, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(dist))
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cant load env")
	}

	e := echo.New()
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", Static())))

	// bind views to the server
	views.Routes(e)
	e.Use(setupContext)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	v1 := e.Group("/api/v1")
	//v1.Use(echojwt.JWT(os.Getenv("JWT_SECRET")))
	v1.GET("/currencies", func(c echo.Context) error {
		//cc := c.(*model.EchoCustomContext)
		cr, err := service.GetAssets()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, cr)
	})

	go worker.SyncAssets()

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(addr))
}

func setupContext(next echo.HandlerFunc) echo.HandlerFunc {
	conf := service.GetContext()

	return func(c echo.Context) error {
		cc := &model.EchoCustomContext{Context: c, CustomContext: conf}

		return next(cc)
	}
}
