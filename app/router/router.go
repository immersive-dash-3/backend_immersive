package router

import (
	classData "immersive_project/klp3/features/classes/data"
	classHandler "immersive_project/klp3/features/classes/handler"
	classService "immersive_project/klp3/features/classes/service"

	menteeLogData "immersive_project/klp3/features/menteelogs/data"
	menteeLogHandler "immersive_project/klp3/features/menteelogs/handler"
	menteeLogService "immersive_project/klp3/features/menteelogs/service"

	menteeData "immersive_project/klp3/features/mentee/data"
	menteeHandler "immersive_project/klp3/features/mentee/handler"
	menteeService "immersive_project/klp3/features/mentee/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	CData := classData.New(db)
	CService := classService.New(CData)
	CHandler := classHandler.New(CService)

	c.POST("/classes", CHandler.Add)
	c.GET("/classes", CHandler.GetAll)
	c.PUT("/classes/:class_id", CHandler.Edit)
	c.GET("/classes/:class_id", CHandler.GetById)
	c.DELETE("/classes/:class_id", CHandler.Delete)

	MData := menteeData.New(db)
	MService := menteeService.New(MData)
	MHandler := menteeHandler.New(MService)

	c.POST("/mentees", MHandler.AddMentee)
	c.GET("/mentees", MHandler.GetAllMentee)
	c.GET("/mentees/:mentee_id", MHandler.GetMenteeById)
	c.PUT("/mentees/:mentee_id", MHandler.UpdateMentee)
	c.DELETE("/mentees/:mentee_id", MHandler.DeleteMentee)

	MLData := menteeLogData.New(db)
	MLService := menteeLogService.New(MLData)
	MLHandler := menteeLogHandler.New(MLService)

	c.POST("/mentees/:mentee_id/logs", MLHandler.Add)
	c.GET("/mentees/:mentee_id/logs", MLHandler.Get)
}
