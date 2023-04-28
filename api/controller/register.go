package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/programming02/osg/api/models"
	"github.com/programming02/osg/jwt"
	"github.com/programming02/osg/storage"
	"net/http"
	"strings"
)

type RegisterService struct {
	storage storage.IStorage
}

func NewRegisterService(db *sqlx.DB) *RegisterService {
	return &RegisterService{
		storage: storage.New(db),
	}
}

func (r *RegisterService) Login(req models.RegisterRequest) (models.RegisterResponse, error) {
	res, err := r.storage.Register().Login(req)
	if err != nil {
		return models.RegisterResponse{}, err
	}
	return res, nil
}

func (a Api) Login(c *gin.Context) {
	req := models.RegisterRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	if !strings.Contains(req.PhoneNumber, "+") || len(req.PhoneNumber) != 13 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errors.New("phone number not found"),
		})
		return
	}

	res, err := a.RegisterService.Login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	token, err := jwt.GenerateNewTokens(res.UserID, map[string]string{"role": res.Role})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.LoginResponse{
		UserID:    res.UserID,
		Authorize: token.Access,
	})
}
