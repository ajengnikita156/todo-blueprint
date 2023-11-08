package routes

import (
	"fmt"
	"membuattasktodo/controller"
	"membuattasktodo/db"
	middleware "membuattasktodo/middleware"
	"membuattasktodo/repository"
	"membuattasktodo/service"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Init() error {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		return err
	}
	defer db.Close()

	// Routes

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	task := e.Group("/task")
	task.Use(middleware.ValidateToken)

	e.GET("", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "Application is Running",
		})
	})

	e.GET("/get", controller.UserController)

	//inith buat gambar
	e.Static("/uploads", "/uploads")

	task.GET("", controller.GetAlltaskController)
	task.GET("/:id", controller.GetTaskById)
	task.POST("/add", controller.CreateTasksController)
	task.PUT("/:id", controller.UpdateTaskController)
	task.DELETE("/:id", controller.DeleteTasksController)
	task.DELETE("", controller.BulkDeleteTask)
	task.POST("", controller.SearchTasksFormController)

	//AUTH
	e.POST("/login", controller.Login)
	e.POST("/logout", controller.Logout)
	e.POST("/register", controller.RegisterController)

	//kategori
	task.GET("/kategori", controller.GetKategoriController)
	task.POST("/kategori/add", controller.AddKategoriController)
	task.DELETE("/delete/:id", controller.DeleteKategoriController)
	task.PUT("/edit/:id", controller.EditKategoriController)

	return e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
