package controllers

import (
	"groot/middleware"
	"groot/models"
	. "groot/services"
)

func Tags(ctx *middleware.Context) {
	tags, err := TagService.Find()

	if err != nil {
		ctx.Go(1, "获取标签失败")
		return
	}

	ctx.Go(tags)
}

func Tag(ctx *middleware.Context) {
	id, err := ctx.Params().GetInt("id")

	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	tag, err := TagService.FindByID(uint(id))

	if err != nil {
		ctx.Go(500, "查询失败")
		return
	}

	ctx.Go(tag)
}

func SearchTag(ctx *middleware.Context) {
	name := ctx.Params().Get("name")

	if name == "all" {
		Tags(ctx)
		return
	}

	tags, _ := TagService.FindByName(name)

	ctx.Go(tags)
}

func CreateTag(ctx *middleware.Context) {
	var tag models.Tag
	
	err := ctx.ReadJSON(&tag)
	if err != nil {
		ctx.Go(406, "参数有误")
		return
	}

	tag.AuthorID = 10000
	err = tag.Save()
	if err != nil {
		ctx.Go(500, "保存失败")
		return
	}

	ctx.Go(tag)
}
