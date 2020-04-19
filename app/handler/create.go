package handler

import (
    "net/http"
    "github.com/labstack/echo"
    "app/model"
)

func CreateUser(c echo.Context) (err error) {
    u := &model.User{
        ID : model.Seq,
    }
    if err := c.Bind(u); err != nil {
        return err
    }
    model.Users[u.ID] = u
    model.Seq++
    return c.JSON(http.StatusCreated, u)
}