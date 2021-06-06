package api

import (
	"gin_example/app"
	"gin_example/lib/resp"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/gin-gonic/gin"
)

func DelaysCreate(c * gin.Context)  {

	asyncResult, err := app.Machinery.SendTask(&tasks.Signature{
		Name: "CountUsers",
	})

	if err != nil {
		resp.SaveFailed(c, err)
		return
	}

	state := asyncResult.GetState()

	resp.JSON(c, state)
}

func DelaysShow(c * gin.Context)  {

	id := c.Param("id")

	state, err := app.Machinery.GetBackend().GetState(id)

	if err != nil {
		resp.QueryFailed(c, err)
		return
	}

	resp.JSON(c, state)
}