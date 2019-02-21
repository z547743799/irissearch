package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gitlab.com/z547743799/irissearch/service"
)

type PageController struct {
	Ctx     iris.Context
	Service service.TbSearchService
}

func (c *PageController) GetSearch() mvc.Result {
	keyword := c.Ctx.URLParamDefault("keyword", "手机")
	page := c.Ctx.URLParamIntDefault("page", 1)

	list := c.Service.Search(keyword, page, 30)
	return mvc.View{
		Name: "search.html",
		Data: iris.Map{
			"itemList":    list.ItemList,
			"query":       keyword,
			"recordCount": list.RecordCount,
			"totalPages":  list.TotalPages,
		},
	}
}
