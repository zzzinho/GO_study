package errors

import "github.com/gin-gonic/gin"

var ValidationErrors = []string{}

func HandleErr(c *gin.Context, err error) error {
	if err != nil {
		c.Error(err)
	}
	return err
}
