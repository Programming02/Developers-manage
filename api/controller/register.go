package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/programming02/osg/api/models"
	"github.com/programming02/osg/jwt"
	"net/http"
)

func (a Api) Register(c *gin.Context) {
	user := models.Users{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "can't bind body " + err.Error()})
		return
	}
	id := uuid.New().String()
	token, err := jwt.GenerateNewTokens(id, map[string]string{"role": user.Positions})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "can't generate token" + err.Error()})
		return
	}
	user.Id = id
	if err := a.storage.Register().RegisterUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "can't register user " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.RegisterResponseModel{
		ID:           id,
		AccessToken:  token.Access,
		ExpiresIn:    token.AccExpire,
		RefreshToken: token.Refresh,
	})
}

func (a Api) Login(c *gin.Context) {
	user := models.LoginRequestModel{}
	meta, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "cannot parse token, " + err.Error()})
		return
	}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "can't bind body " + err.Error()})
		return
	}
	if err := a.storage.Register().Login(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "can't login " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.LoginResponseModel{
		ID: meta.UserID.String(),
	})
}
