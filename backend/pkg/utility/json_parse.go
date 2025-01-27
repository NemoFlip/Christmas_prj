package utility

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ParseJson(ctx *gin.Context, requestStruct any) error {
	if err := ctx.BindJSON(requestStruct); err != nil {
		return fmt.Errorf("unable to bind user: %s", err)
	}
	return nil
}
