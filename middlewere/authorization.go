package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/config"
	_ "github.com/programming02/osg/config"
	"github.com/programming02/osg/jwt"
	"log"
	"net/http"
)

// JWTRoleAuthorizer is a structure for a Role Authorizer type
type JWTRoleAuthorizer struct {
	enforcer   *casbin.Enforcer
	SigningKey []byte
	//	logger     logger.Logger
}

// NewJWTRoleAuthorizer creates and returns new Role Authorizer
func NewJWTRoleAuthorizer(cfg *config.Config) (*JWTRoleAuthorizer, error) {
	enforcer, err := casbin.NewEnforcer(cfg.CasbinConfigPath, cfg.MiddlewareRolesPath)
	if err != nil {
		log.Fatal("could not initialize new enforcer:", err.Error())
		return nil, err
	}

	return &JWTRoleAuthorizer{
		enforcer:   enforcer,
		SigningKey: []byte(cfg.JWTSecretKey),
	}, nil
}

func Authorizer(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		enforcer, err := casbin.NewEnforcer(cfg.CasbinConfigPath, cfg.MiddlewareRolesPath)
		if err != nil {
			log.Fatal("enforcer not initialized, ", err)
			return
		}

		claims, err := jwt.ExtractClaims(accessToken, []byte(cfg.SigningKey))
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
