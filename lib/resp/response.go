package resp

import (
	"fmt"
	"gin_example/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSONResponse struct {
	Errcode  int         `json:"errcode"`
	Errmsg string      `json:"errmsg"`
	Data  interface{} `json:"data,omitempty"`
}

func JSON(c * gin.Context, obj interface{})  {
	c.JSON(http.StatusOK, JSONResponse{
		Errcode: 0,
		Errmsg: "OK",
		Data: obj,
	})
}

func serverError(c * gin.Context, code int, msg interface{})  {
	c.JSON(http.StatusInternalServerError, JSONResponse{
		Errcode: code,
		Errmsg: fmt.Sprintln(msg),
	})
}

func badRequest(c * gin.Context, code int, msg interface{})  {
	c.JSON(http.StatusBadRequest, JSONResponse{
		Errcode: code,
		Errmsg: fmt.Sprintln(msg),
	})
}

func unauthorized(c * gin.Context, code int, msg interface{})  {
	c.JSON(http.StatusUnauthorized, JSONResponse{
		Errcode: code,
		Errmsg: fmt.Sprintln(msg),
	})
}

func forbidden(c * gin.Context, code int, msg interface{})  {
	c.JSON(http.StatusForbidden, JSONResponse{
		Errcode: code,
		Errmsg: fmt.Sprintln(msg),
	})
}

func notFound(c * gin.Context, code int, msg interface{})  {
	c.JSON(http.StatusNotFound, JSONResponse{
		Errcode: code,
		Errmsg: fmt.Sprintln(msg),
	})
}

func Unauthorized(c * gin.Context, msg interface{})  {
	unauthorized(c, 40100, msg)
}

func ParamsInvalid(c * gin.Context, msg interface{})  {
	badRequest(c, 40000, msg)
}

func RecoreNotFound(c * gin.Context, msg interface{})  {
	if app.Env == "production" {
		badRequest(c, 40400, "Record Not Found")
		return
	}
	notFound(c, 40400, msg)
}

func SaveFailed(c * gin.Context, msg interface{})  {
	badRequest(c, 40001, msg)
}

func AuthorizedRequired(c * gin.Context, msg interface{})  {
	unauthorized(c, 40101, msg)
}

func JwtUnauthorized(c * gin.Context, msg interface{})  {
	unauthorized(c, 40102, msg)
}

func LimitedLogin(c * gin.Context, msg interface{})  {
	forbidden(c, 40301, msg)
}

func QueryFailed(c * gin.Context, msg interface{})  {
	serverError(c, 500001, msg)
}
