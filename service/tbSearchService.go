package service

import (
	"context"
	"fmt"
	"reflect"

	"github.com/olivere/elastic"
	"github.com/xormplus/xorm"
	"gitlab.com/z547743799/iriscommon/pojo"
	"gitlab.com/z547743799/irissearch/db"
	"gitlab.com/z547743799/irissearch/elastics"
	"gitlab.com/z547743799/irissearch/models"
)

type TbSearchService interface {
	Search(keyword string, page, rows int) *pojo.SearchResult
}

type tbSearchService struct {
	engine *xorm.Engine
}

func NewTbSearchService() TbSearchService {
	return &tbSearchService{
		engine: db.X,
	}
}

func (d *tbSearchService) Search(keyword string, page, rows int) *pojo.SearchResult {


	q := elastic.NewQueryStringQuery(keyword)
	res, err := elastics.Client.Search("mega").Size(9999).Type("TbItem").Query(q).Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	var yp models.SearchItem
	searc := make([]models.SearchItem, 0)

	for _, item := range res.Each(reflect.TypeOf(yp)) { //从搜索结果中取数据的方法
		t := item.(models.SearchItem)
		fmt.Printf("%#v\n", t)
		searc = append(searc, t)

		//fmt.Println(item)
	}

	//数量
	numFound := len(searc)
	//页数
	totalPage := numFound / rows
	if totalPage%rows != 0 {
		totalPage++
	}
	//分页
	if rows > numFound {
		rows = numFound
	}
	searcc := searc[(page-1)*rows : page*rows]

	searchResult := &pojo.SearchResult{RecordCount: int64(numFound), TotalPages: totalPage, ItemList: searcc}
	return searchResult

}
