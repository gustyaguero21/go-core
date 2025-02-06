package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}

func (er *Error) Error() string {
	return fmt.Sprintf("Status: %d, Error: %s", er.Status, er.Err)
}

func NewError(ctx *gin.Context, status int, err string) {
	errorStruct := &Error{
		Status: status,
		Err:    err,
	}

	ctx.JSON(status, errorStruct)
}
