package handler

import (
	"errors"
	"immersive_project/klp3/exception"
	"immersive_project/klp3/features/users"
	"immersive_project/klp3/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandlerImplementation struct {
	service users.UserServiceInterface
}

func New(service users.UserServiceInterface) users.UserHandlerInterface {
	return &UserHandlerImplementation{service: service}
}

// Create implements users.UserHandlerInterface
func (handler *UserHandlerImplementation) Create(c echo.Context) error {
	var userRequest UserRequest

	err := c.Bind(&userRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation falied, request resource not valid", nil))
	}

	userEntity := UserRequestToEntity(userRequest)
	err = handler.service.Create(userEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}

	return c.JSON(http.StatusCreated, helper.WebResponse(http.StatusCreated, "created", nil))

}

// Delete implements users.UserHandlerInterface
func (handler *UserHandlerImplementation) Delete(c echo.Context) error {
	id := c.Param("user_id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "bad request", nil))
	}

	err = handler.service.Delete(uint(intId))
	if err != nil {
		if errors.Is(err, exception.ErrIdIsNotFound) {
			return c.JSON(http.StatusNotFound, helper.WebResponse(http.StatusNotFound, "operation failed, resource not found", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
		}
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", nil))
}

// FindAll implements users.UserHandlerInterface
func (handler *UserHandlerImplementation) FindAll(c echo.Context) error {
	page := c.QueryParam("page")
	itemsPerPage := c.QueryParam("itemsPerPage")
	searchName := c.QueryParam("searchName")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "bad request", nil))
	}

	intItemsPerPage, err := strconv.Atoi(itemsPerPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "bad request", nil))
	}

	res, prevPage, nextPage, err := handler.service.FindAll(intPage, intItemsPerPage, searchName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	resResponse := UserEntityToResponseAll(res)

	return c.JSON(http.StatusOK, helper.FindAllWebResponse(http.StatusOK, "success", resResponse, prevPage, nextPage))

}

// FindById implements users.UserHandlerInterface
func (handler *UserHandlerImplementation) FindById(c echo.Context) error {
	id := c.Param("user_id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "bad request", nil))
	}
	res, err := handler.service.FindById(intId)
	if err != nil {
		if errors.Is(err, exception.ErrIdIsNotFound) {
			return c.JSON(http.StatusNotFound, helper.WebResponse(http.StatusNotFound, "operation failed, resource not found", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
		}
	}
	userResponse := UserEntityToResponse(res)
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", userResponse))
}

// Update implements users.UserHandlerInterface
func (handler *UserHandlerImplementation) Update(c echo.Context) error {
	var userRequest UserRequest
	id := c.Param("user_id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "bad request", nil))
	}

	err = c.Bind(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "bad request", nil))
	}

	userEntity := UserRequestToEntity(userRequest)
	userEntity.Id = uint(intId)

	err = handler.service.Update(userEntity)
	if err != nil {
		if errors.Is(err, exception.ErrIdIsNotFound) {
			return c.JSON(http.StatusNotFound, helper.WebResponse(http.StatusNotFound, "operation failed, resource not found", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
		}
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", nil))
}

// UserLogin implements users.UserHandlerInterface
func (handler *UserHandlerImplementation) UserLogin(c echo.Context) error {
	var userRequest UserRequest

	err := c.Bind(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "bad request", nil))
	}

	res, token, err := handler.service.Login(userRequest.Email, userRequest.Password)
	if err != nil {
		if errors.Is(err, exception.ErrIdIsNotFound) {
			return c.JSON(http.StatusNotFound, helper.WebResponse(http.StatusNotFound, "operation failed, resource not found", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
		}
	}

	dataLogin := LoginResponse{
		Role:  res.Role,
		Token: token,
	}

	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", dataLogin))

}
