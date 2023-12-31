package handler

import (
	"fmt"
	"immersive_project/klp3/app/middleware"
	"immersive_project/klp3/features/menteelogs"
	"immersive_project/klp3/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenteeLogHandler struct {
	menteeLogHandler menteelogs.MenteeLogServiceInterface
}
func (handler *MenteeLogHandler)Add(c echo.Context)error{
	idUser:=middleware.ExtractTokenUserId(c)
	id:=c.Param("mentee_id")
	idMentee,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,helper.WebResponse(400, "id not valid", nil))
	}
	fmt.Println(id)
	var input MenteeLogRequest
	errBind:=c.Bind(&input)
	if errBind != nil{
		return c.JSON(http.StatusBadRequest,helper.WebResponse(400,"input tidak sesuai",nil))
	}
	inputEntity:=RequestToEntity(input)
	inputEntity.MenteeID=uint(idMentee)
	inputEntity.UserID=uint(idUser)
	fmt.Println(inputEntity)
	err:=handler.menteeLogHandler.Add(inputEntity)
	if err != nil{
		return c.JSON(http.StatusInternalServerError,helper.WebResponse(500,err.Error(),nil))
	}
	return c.JSON(http.StatusCreated,helper.WebResponse(201,"success create feedback",nil))
}

func (handler *MenteeLogHandler)Get(c echo.Context)error{
	id:=c.Param("mentee_id")
	idMentee,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,helper.WebResponse(400, "id not valid", nil))
	}
	data,err :=handler.menteeLogHandler.Get(uint(idMentee))	
	if err != nil{
		return c.JSON(http.StatusInternalServerError,helper.WebResponse(500,err.Error(),nil))
	}
	response:=EntityResponseById(data)
	return c.JSON(http.StatusOK,helper.WebResponse(200,"success get feedback",response))
}

func (handler *MenteeLogHandler)Edit(c echo.Context)error{
	id:=c.Param("log_id")
	idLog,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,helper.WebResponse(400, "id not valid", nil))
	}
	var input MenteeLogRequest
	errBind:=c.Bind(&input)
	if errBind != nil{
		return c.JSON(http.StatusBadRequest,helper.WebResponse(400,"input tidak sesuai",nil))
	}
	inputEntity:=RequestToEntity(input)
	err:=handler.menteeLogHandler.Edit(uint(idLog),inputEntity)	
	if err != nil{
		return c.JSON(http.StatusInternalServerError,helper.WebResponse(500,err.Error(),nil))
	}
	return c.JSON(http.StatusOK,helper.WebResponse(200,"success edit feedback",nil))
}

func (handler *MenteeLogHandler)Delete(c echo.Context)error{
	id:=c.Param("log_id")
	idLog,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,helper.WebResponse(400, "id not valid", nil))
	}
	err:=handler.menteeLogHandler.Delete(uint(idLog))	
	if err != nil{
		return c.JSON(http.StatusInternalServerError,helper.WebResponse(500,err.Error(),nil))
	}
	return c.JSON(http.StatusOK,helper.WebResponse(200,"success delete feedback",nil))
}

func (handler *MenteeLogHandler)GetAll(c echo.Context)error{
	data,err:=handler.menteeLogHandler.GetAll()
	if err != nil{
		return c.JSON(http.StatusInternalServerError,helper.WebResponse(500,err.Error(),nil))
	}
	output:=EntityToResponseAll(data)
	return c.JSON(http.StatusOK,helper.WebResponse(200,"success delete feedback",output))	
}
func New(handler menteelogs.MenteeLogServiceInterface)*MenteeLogHandler{
	return &MenteeLogHandler{
		menteeLogHandler: handler,
	}
}