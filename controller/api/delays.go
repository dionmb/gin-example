package api

import (
	"gin_example/app"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DelaysCreate(c * gin.Context)  {

	asyncResult, err := app.Machinery.SendTask(&tasks.Signature{
		Name: "CountUsers",
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"error": err,
		})
		return
	}

	state := asyncResult.GetState()

	c.JSON(http.StatusOK, state)
}

func DelaysShow(c * gin.Context)  {

	id := c.Param("id")

	state, err := app.Machinery.GetBackend().GetState(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, state)
}