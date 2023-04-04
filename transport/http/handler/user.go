package handler

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/onelab-hw2/model"
	"github.com/labstack/echo"
)

func (h Manager) GetUser(c echo.Context) error {
	idString, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return err
	}
	id := uint(idString)

	resp, err := h.srv.User.Get(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) CreateUser(c echo.Context) error {
	var req model.User
	if err := c.Bind(&req); err != nil {
		return err
	}
	resp, err := h.srv.User.Create(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) UpdateUser(c echo.Context) error {
	idString, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return err
	}
	id := uint(idString)

	var req model.User
	if err := c.Bind(&req); err != nil {
		return err
	}
	resp, err := h.srv.User.Update(id, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) ListUsers(c echo.Context) error {
	resp, err := h.srv.User.List()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) DeleteUser(c echo.Context) error {
	idString, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return err
	}
	id := uint(idString)

	resp, err := h.srv.User.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}
