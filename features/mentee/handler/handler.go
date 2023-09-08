package handler

import (
	"immersive_project/klp3/features/mentee"
	"immersive_project/klp3/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type MenteeHandler struct {
	MenteeService mentee.MenteeServiceInterface
}

func New(service mentee.MenteeServiceInterface) *MenteeHandler {
	return &MenteeHandler{
		MenteeService: service,
	}
}

func (handler *MenteeHandler) GetAllMentee(c echo.Context) error {
	var pageConv, itemConv, statusConv, classConv int
	var errPageConv, errItemConv, errStatusConv, errClassConv error

	page := c.QueryParam("page")
	if page != "" {
		pageConv, errPageConv = strconv.Atoi(page)
		if errPageConv != nil {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
	}

	itemPerPage := c.QueryParam("itemPerPage")
	if itemPerPage != "" {
		itemConv, errItemConv = strconv.Atoi(itemPerPage)
		if errItemConv != nil {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
	}

	searchName := c.QueryParam("searchName")

	status_id := c.QueryParam("status_id")
	if status_id != "" {
		statusConv, errStatusConv = strconv.Atoi(status_id)
		if errStatusConv != nil {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
	} else {
		statusConv = 0
	}

	class_id := c.QueryParam("class_id")
	if class_id != "" {
		classConv, errClassConv = strconv.Atoi(class_id)
		if errClassConv != nil {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
	} else {
		classConv = 0
	}

	education_type := c.QueryParam("education_type")
	result, next, err := handler.MenteeService.GetAll(uint(pageConv), uint(itemConv), uint(statusConv), uint(classConv), education_type, searchName)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helper.WebResponse(http.StatusNotFound, "operation failed, resource not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error"+err.Error()+"err", nil))
	}
	var responseData []MenteeResponseAll
	for _, v := range result {
		var data = EntityToResponseAll(v)
		responseData = append(responseData, data)
	}
	return c.JSON(http.StatusOK, helper.FindAllWebResponse(http.StatusOK, "success", responseData, next))
}

func (handler *MenteeHandler) AddMentee(c echo.Context) error {
	var input MenteeRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}
	var inputEntity = RequestToEntity(input)
	err := handler.MenteeService.Insert(inputEntity)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error"+err.Error(), nil))
	}
	return c.JSON(http.StatusCreated, helper.WebResponse(http.StatusCreated, "created", nil))
}

func (handler *MenteeHandler) GetMenteeById(c echo.Context) error {
	id := c.Param("mentee_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}
	result, err := handler.MenteeService.Get(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helper.WebResponse(http.StatusNotFound, "operation failed, resource not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	var menteeResp = EntityToResponse(result)
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "get data success", menteeResp))
}

func (handler *MenteeHandler) UpdateMentee(c echo.Context) error {
	id := c.Param("mentee_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}
	var input MenteeRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}
	var inputEntity = RequestToEntity(input)
	err := handler.MenteeService.Update(uint(idConv), inputEntity)
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helper.WebResponse(http.StatusNotFound, "operation failed, resource not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", nil))
}

func (handler *MenteeHandler) DeleteMentee(c echo.Context) error {
	id := c.Param("mentee_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}
	err := handler.MenteeService.Delete(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helper.WebResponse(http.StatusNotFound, "operation failed, resource not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", nil))
}
