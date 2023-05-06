package utilsHeader

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetHeaderAuthorizationValue(c *gin.Context) (value string, err error) {
	bearertoken := c.Request.Header["Authorization"]
	if bearertoken == nil || len(strings.Split(bearertoken[0], " ")) != 2 {
		return "", errors.New("Authorization Header Not Found")
	}

	return strings.Split(bearertoken[0], " ")[1], nil
	
}