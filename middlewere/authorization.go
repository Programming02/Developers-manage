package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/config"
	_ "github.com/programming02/osg/config"
	"log"
	"net/http"
	"strings"
)

func Authorizer(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		enforcer, err := casbin.NewEnforcer(cfg.CasbinConfigPath, cfg.MiddlewareRolesPath)
		if err != nil {
			log.Fatal("enforcer not initialized, ", err)
			return
		}

		claims, err := ExtractClaims(accessToken, []byte(cfg.SigningKey))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		role := claims["role"]

		ok, err := enforcer.Enforce(role, c.Request.URL.String(), c.Request.Method)
		if err != nil {
			log.Println("could not enforce:", err)
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusForbidden, map[string]string{
				"error": "user not allowed, there is a problem with authorization",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ExtractClaims(t string, signKey []byte) (jwtgo.MapClaims, error) {
	claims := jwtgo.MapClaims{}
	if t == "" {
		claims["role"] = "unauthorized"
		return claims, nil
	}
	if strings.Contains(t, "Basic") {
		claims["role"] = "unauthorized"
		return claims, nil
	}
	token, err := jwtgo.ParseWithClaims(t, claims, func(token *jwtgo.Token) (interface{}, error) {
		return signKey, nil
	})

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !(ok && token.Valid) {
		err = fmt.Errorf("JWT token bad")
		return nil, err
	}
	return claims, nil
}
