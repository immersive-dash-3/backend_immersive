package handler

import (
	"immersive_project/klp3/features/classes"
	"immersive_project/klp3/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	classHandler classes.ClassServiceInterface
}
func (handler *ClassHandler)Add(c echo.Context)error{
	var input ClassesRequest
	errBind:=c.Bind(&input)
	if errBind != nil{
		return c.JSON(http.StatusBadGateway,helper.WebResponse(400, "error bind data", nil))
	}
	inputEntity:= RequestToEntity(input)
	errAdd:=handler.classHandler.Add(inputEntity)
	if errAdd != nil{
		return c.JSON(http.StatusInternalServerError,helper.WebResponse(500,errAdd.Error(),nil))
	}
	return c.JSON(http.StatusCreated,helper.WebResponse(201,"success insert class",nil))
}

func (handler *ClassHandler)GetAll(c echo.Context)error{
	data,err:=handler.classHandler.GetAll()
	if err != nil{
		return c.JSON(http.StatusInternalServerError,helper.WebResponse(500,err.Error(),nil))
	}
	response:=ResponseToEntityAll(data)
	return c.JSON(http.StatusOK,helper.WebResponse(200,"success get all class",response))
}

func (handler *ClassHandler)Edit(c echo.Context)error{
	id:=c.Param("class_id")
	idConv,errConv:= strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusNotFound,helper.WebResponse(404,"id not valid",nil))
	}
	var input ClassesRequest
	errBind:=c.Bind(&input)
	if errBind != nil{
		return c.JSON(http.StatusBadGateway,helper.WebResponse(400, "error bind data", nil))
	}
	inputEntity:= RequestToEntity(input)
	err:=handler.classHandler.Edit(uint(idConv),inputEntity)
	if err!= nil{
		return c.JSON(http.StatusInternalServerError,helper.WebResponse(500,err.Error(),nil))
	}
	return c.JSON(http.StatusOK,helper.WebResponse(200,"success update class",nil))
}

func (handler *ClassHandler)GetById(c echo.Context)error{
	id:=c.Param("class_id")
	idConv,errConv:= strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusNotFound,helper.WebResponse(404,"id not valid",nil))
	}
	data,err:=handler.classHandler.GetById(uint(idConv))
	if err!= nil{
		return c.JSON(http.StatusInternalServerError,helper.WebResponse(500,err.Error(),nil))
	}
	output:=ResponseToEntity(data)
	return c.JSON(http.StatusOK,helper.WebResponse(200,"success get by id class",output))	
}

func (handler *ClassHandler)Delete(c echo.Context)error{
	id:=c.Param("class_id")
	idConv,errConv:= strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusNotFound,helper.WebResponse(404,"id not valid",nil))
	}
	err:=handler.classHandler.Delete(uint(idConv))	
	if err!= nil{
		return c.JSON(http.StatusInternalServerError,helper.WebResponse(500,err.Error(),nil))
	}
	return c.JSON(http.StatusOK,helper.WebResponse(200,"success deleted",nil))	
}
func New(handler classes.ClassServiceInterface)*ClassHandler{
	return &ClassHandler{
		classHandler: handler,
	}
}