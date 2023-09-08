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

func (handler *ClassHandler) Add(c echo.Context) error {
	var input ClassesRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadGateway, helper.WebResponse(400, "error bind data", nil))
	}
	inputEntity := RequestToEntity(input)
	errAdd := handler.classHandler.Add(inputEntity)
	if errAdd != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(500, errAdd.Error(), nil))
	}
	return c.JSON(http.StatusCreated, helper.WebResponse(201, "success insert class", nil))
}

func (handler *ClassHandler) GetAll(c echo.Context) error {

	var qparams classes.QueryParams

	itemsPerPage := c.QueryParam("itemsPerPage")
	page := c.QueryParam("page")

	if itemsPerPage == "" {
		qparams.IsClassDashboard = false
	} else {
		qparams.IsClassDashboard = true
		intItemsPerPage, err := strconv.Atoi(itemsPerPage)
		if err != nil {
			c.JSON(http.StatusBadRequest, helper.WebResponse(400, "bad request", nil))
		}
		qparams.ItemsPerPage = intItemsPerPage

	}

	if page == "" {
		qparams.Page = 1
	} else {
		intPage, err := strconv.Atoi(page)
		if err != nil {
			c.JSON(http.StatusBadRequest, helper.WebResponse(400, "bad request", nil))
		}
		qparams.Page = intPage
	}

	searchName := c.QueryParam("searchNameClass")
	qparams.SearchName = searchName
	bol, data, err := handler.classHandler.GetAll(qparams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(500, err.Error(), nil))
	}
	response := EntityToResponseAll(data)
	return c.JSON(http.StatusOK, helper.FindAllWebResponse(200, "success get all class", response, bol))
}

func (handler *ClassHandler) Edit(c echo.Context) error {
	id := c.Param("class_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusNotFound, helper.WebResponse(404, "id not valid", nil))
	}
	var input ClassesRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadGateway, helper.WebResponse(400, "error bind data", nil))
	}
	inputEntity := RequestToEntity(input)
	err := handler.classHandler.Edit(uint(idConv), inputEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(500, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(200, "success update class", nil))
}

func (handler *ClassHandler) GetById(c echo.Context) error {
	id := c.Param("class_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusNotFound, helper.WebResponse(404, "id not valid", nil))
	}
	data, err := handler.classHandler.GetById(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(500, err.Error(), nil))
	}
	output := EntityToResponse(data)
	return c.JSON(http.StatusOK, helper.WebResponse(200, "success get by id class", output))
}

func (handler *ClassHandler) Delete(c echo.Context) error {
	id := c.Param("class_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusNotFound, helper.WebResponse(404, "id not valid", nil))
	}
	err := handler.classHandler.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(500, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(200, "success deleted", nil))
}
func New(handler classes.ClassServiceInterface) *ClassHandler {
	return &ClassHandler{
		classHandler: handler,
	}
}
