package handlers

import (
	"Guardian/models"
	"github.com/alexedwards/argon2id"
	"github.com/labstack/echo"
	"net/http"
)

type AuthHandler struct {
	Handler
}

func (h AuthHandler) Login(c echo.Context) (err error) {

	cred := models.Credentials{}
	user := models.UserWithPassword{}

	if err = c.Bind(cred); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if err := h.Database.Where(&models.User{Email: cred.Email}).First(&user).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	_, err = argon2id.ComparePasswordAndHash(user.Password, cred.Password)
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.NoContent(http.StatusOK)
}

func (h AuthHandler) Logour(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
