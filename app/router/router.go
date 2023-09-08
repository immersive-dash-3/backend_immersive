package router

import (
	"immersive_project/klp3/app/middleware"
	classData "immersive_project/klp3/features/classes/data"
	classHandler "immersive_project/klp3/features/classes/handler"
	classService "immersive_project/klp3/features/classes/service"
	usersData "immersive_project/klp3/features/users/data"
	usersHandler "immersive_project/klp3/features/users/handler"
	usersService "immersive_project/klp3/features/users/service"

	menteeLogData "immersive_project/klp3/features/menteelogs/data"
	menteeLogHandler "immersive_project/klp3/features/menteelogs/handler"
	menteeLogService "immersive_project/klp3/features/menteelogs/service"

	menteeData "immersive_project/klp3/features/mentee/data"
	menteeHandler "immersive_project/klp3/features/mentee/handler"
	menteeService "immersive_project/klp3/features/mentee/service"

	dashboardData "immersive_project/klp3/features/dashboard/data"
	dashboardHandler "immersive_project/klp3/features/dashboard/handler"
	dashboardService "immersive_project/klp3/features/dashboard/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	CData := classData.New(db)
	CService := classService.New(CData)
	CHandler := classHandler.New(CService)

	c.POST("/classes", CHandler.Add, middleware.JWTMiddleware())
	c.GET("/classes", CHandler.GetAll, middleware.JWTMiddleware())
	c.PUT("/classes/:class_id", CHandler.Edit, middleware.JWTMiddleware())
	c.GET("/classes/:class_id", CHandler.GetById, middleware.JWTMiddleware())
	c.DELETE("/classes/:class_id", CHandler.Delete, middleware.JWTMiddleware())

	MData := menteeData.New(db)
	MService := menteeService.New(MData)
	MHandler := menteeHandler.New(MService)

	c.POST("/mentees", MHandler.AddMentee, middleware.JWTMiddleware())
	c.GET("/mentees", MHandler.GetAllMentee, middleware.JWTMiddleware())
	c.GET("/mentees/:mentee_id", MHandler.GetMenteeById, middleware.JWTMiddleware())
	c.PUT("/mentees/:mentee_id", MHandler.UpdateMentee, middleware.JWTMiddleware())
	c.DELETE("/mentees/:mentee_id", MHandler.DeleteMentee, middleware.JWTMiddleware())

	MLData := menteeLogData.New(db)
	MLService := menteeLogService.New(MLData)
	MLHandler := menteeLogHandler.New(MLService)

	c.POST("/mentees/:mentee_id/logs", MLHandler.Add, middleware.JWTMiddleware())
	c.GET("/mentees/:mentee_id/logs", MLHandler.Get, middleware.JWTMiddleware())
	c.PUT("/logs/:log_id", MLHandler.Edit, middleware.JWTMiddleware())
	c.GET("/logs", MLHandler.GetAll, middleware.JWTMiddleware())
	c.DELETE("/logs/:log_id", MLHandler.Delete, middleware.JWTMiddleware())

	userData := usersData.New(db)
	userService := usersService.New(userData)
	userHandler := usersHandler.New(userService)

	c.POST("/login", userHandler.UserLogin)
	c.GET("/users", userHandler.FindAll, middleware.JWTMiddleware())
	c.GET("/users/:user_id", userHandler.FindById, middleware.JWTMiddleware())
	c.PUT("/users/:user_id", userHandler.Update, middleware.JWTMiddleware())
	c.DELETE("/users/:user_id", userHandler.Delete, middleware.JWTMiddleware())
	c.POST("/users", userHandler.Create, middleware.JWTMiddleware())

	dashData := dashboardData.New(db)
	dashService := dashboardService.New(dashData)
	dashHandler := dashboardHandler.New(dashService)

	c.GET("/dashboard", dashHandler.GetData)
}
