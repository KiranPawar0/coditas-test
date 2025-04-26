package user

import (
	"net/http"

	"github.com/KiranPawar0/coditas-test/pkg/helper/structvalidator"
	"github.com/KiranPawar0/coditas-test/pkg/user/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateUser(c *gin.Context) {
	var user config.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validator := structvalidator.StructValidator{Struct: user}
	if err := validator.Validate(); err != nil {
		logrus.WithError(err).Error("Validation failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
