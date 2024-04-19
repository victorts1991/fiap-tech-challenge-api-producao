package error

import (
	"fiap-tech-challenge-api/internal/core/commons"
	"github.com/joomcode/errorx"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const layout = "2006-01-02T15:04:05.999999Z07:00"

type Response struct {
	StatusCode int    `json:"-"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Timestamp  string `json:"timestamp"`
}

func NewErrorResponse(error *errorx.Error) Response {
	status := getHttpCode(error)

	return Response{
		StatusCode: status,
		Status:     http.StatusText(status),
		Message:    error.Message(),
		Timestamp:  time.Now().Format(layout),
	}
}

func getHttpCode(err *errorx.Error) int {
	switch {
	case err.IsOfType(errorx.IllegalFormat) || err.IsOfType(errorx.IllegalArgument) || err.IsOfType(commons.BadRequest):
		return 400
	case err.IsOfType(commons.NotFound):
		return 404
	case err.IsOfType(commons.Unauthorized):
		return 401
	default:
		return 500
	}
}

func HandleError(ctx echo.Context, err *errorx.Error) error {
	errResponse := NewErrorResponse(err)
	return ctx.JSON(errResponse.StatusCode, errResponse)
}

func ResponseJson(ctx echo.Context, o interface{}) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	ctx.Response().WriteHeader(http.StatusOK)
	return jsoniter.NewEncoder(ctx.Response()).Encode(o)
}
