package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"

	mid "github.com/coderbunker/heikenet-backend/middleware"
	"github.com/coderbunker/heikenet-backend/models"
)

func CreateUser(c echo.Context) error {
	new_user := new(models.NewUser)
	if err := c.Bind(new_user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "can't read json",
		})
	}

	db, err := mid.GetDB(c)
	if err != nil {
		log.Fatal(err)
	}

	user, err := models.CreateUser(db, new_user)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusCreated, map[string]string{
			"id": user.ID,
		})
	}
}

func GetUser(c echo.Context) error {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "can't read id",
		})
	}

	db, err := mid.GetDB(c)
	if err != nil {
		log.Fatal(err)
	}

	user, err := models.GetUser(db, id)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c echo.Context) error {
	new_user := new(models.NewUser)
	if err := c.Bind(new_user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "can't read json",
		})
	}

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "can't read id",
		})
	}

	db, err := mid.GetDB(c)
	if err != nil {
		log.Fatal(err)
	}

	user, err := models.UpdateUser(db, id, new_user)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, map[string]string{
			"id": user.ID,
		})
	}
}

func DeleteUser(c echo.Context) error {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "can't read id",
		})
	}

	db, err := mid.GetDB(c)
	if err != nil {
		log.Fatal(err)
	}

	err = models.DeleteUser(db, id)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})
	} else {
		return c.NoContent(http.StatusNoContent)
	}
}
