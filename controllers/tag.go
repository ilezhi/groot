package controllers

import (
	// "math/rand"
	// "time"
	// "fmt"
	"groot/models"
	"groot/middleware"
)

func CreateTag(ctx *middleware.Context) {
	var tag = new(models.Tag)
	ctx.ReadJSON(tag)
	user := ctx.Session().Get("user").(*models.User)
	tag.AuthorID = user.ID

	err := tag.FindAndSave()
	if err != nil {
		ctx.Error(500, "新增失败")
		return
	}

	ctx.Go(tag)
}
