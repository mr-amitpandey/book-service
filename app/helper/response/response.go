package response

import (
	"fmt"
	"net/http"

	"github.com/book-service/api/app/utils/common/constants"
	common "github.com/book-service/api/app/utils/common/src"
	"github.com/gin-gonic/gin"
)

func SendSuccessResponse[T any](c *gin.Context, data T) {
	common.CreateJSONResponse(c, constants.ResponseOK, http.StatusOK, "", data)
}

func SendBadRequestResponse[T comparable](c *gin.Context, message T) {
	common.CreateJSONResponse(c, constants.ResponseBad, http.StatusBadRequest, fmt.Sprintf("%v", message), nil)
}

func SendErrorResponse[T comparable](c *gin.Context, statusCode int, message T) {
	common.CreateJSONResponse(c, constants.ResponseError, statusCode, fmt.Sprintf("%v", message), nil)
}

func SendConflictResponse(c *gin.Context, message string) {
	common.CreateJSONResponse(c, constants.ResponseConflict, http.StatusConflict, message, nil)
}

func SendForbiddenResponse(c *gin.Context, message string) {
	common.CreateJSONResponse(c, constants.ResponseForbidden, http.StatusForbidden, message, nil)
}

func SendUnauthorizedResponse(c *gin.Context, message string) {
	common.CreateJSONResponse(c, constants.ResponseUnauthorized, http.StatusUnauthorized, message, nil)
}

func SendNotFoundResponse(c *gin.Context, message string) {
	common.CreateJSONResponse(c, constants.ResponseNotFound, http.StatusNotFound, message, nil)
}

func SendInternalServerErrorResponse(c *gin.Context, message string) {
	common.CreateJSONResponse(c, constants.ResponseError, http.StatusInternalServerError, message, nil)
}
