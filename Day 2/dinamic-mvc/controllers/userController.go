package controllers

import (
	"mvc/lib/database"
	"mvc/models"

	"github.com/labstack/echo/v4"
)

func GetUser(e echo.Context) error {
	users, err := database.GetUser()
	if err != nil {
		return e.JSON(500, err)
	}
	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "Users retrieved",
		"data":    users,
	})
}

func GetUserByID(e echo.Context) error {
	id := e.Param("id")

	user, err := database.GetUserByID(id)
	if err != nil {
		return e.JSON(500, err)
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User retrieved",
		"data":    user,
	})
}

func CreateUser(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)

	newUser, err := database.CreateUser(user)
	if err != nil {
		return e.JSON(500, err)
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User created",
		"data":    newUser,
	})
}

func DeleteUser(e echo.Context) error {
	id := e.Param("id")

	err := database.DeleteUser(id)
	if err != nil {
		return e.JSON(500, err)
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User deleted",
	})
}

func UpdateUser(e echo.Context) error {
	id := e.Param("id")
	user := models.User{}
	e.Bind(&user)

	updatedUser, err := database.UpdateUser(id, user)
	if err != nil {
		return e.JSON(500, err)
	}

	return e.JSON(200, map[string]interface{}{
		"succes":  true,
		"message": "User updated",
		"data":    updatedUser,
	})
}
