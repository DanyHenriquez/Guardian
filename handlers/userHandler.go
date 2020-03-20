package handlers

import (
	"crypto/rand"
	"fmt"
	"github.com/DanyHenriquez/Guardian/models"
	"github.com/labstack/echo"
	"github.com/pquerna/otp/totp"
	"net/http"
)

type UserHandler struct {
	Handler
}

func (h UserHandler) Index(c echo.Context) (err error) {
	users := models.Users{}

	if err := h.Database.Find(&users).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}

func (h UserHandler) Get(c echo.Context) (err error) {
	id := c.Param("id")
	user := models.User{}

	if err := h.Database.First(&user, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

func (h UserHandler) Post(c echo.Context) (err error) {
	user := models.User{}

	if err = c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	b := make([]byte, 4)
	rand.Read(b)
	user.Username = fmt.Sprintf("%x", b)

	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "kutklanten.nl",
		AccountName: user.Username,
	})

	user.Secret = key.Secret()

	if err := h.Database.Create(&user).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, user)
}

func (h UserHandler) Put(c echo.Context) (err error) {
	id := c.Param("id")
	user := models.User{}

	if err := h.Database.First(&user, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err = c.Bind(&user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.Database.Save(&user).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (h UserHandler) Delete(c echo.Context) (err error) {
	id := c.Param("id")
	user := models.User{}

	if err := h.Database.First(&user, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := h.Database.Delete(&user).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
