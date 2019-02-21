package service

import (
	"context"
	"strconv"

	"gitlab.com/z547743799/irissearch/elastics"
	"gitlab.com/z547743799/irissearch/models"

	"github.com/xormplus/xorm"
	"gitlab.com/z547743799/iriscommon/utils"
	"gitlab.com/z547743799/irissearch/db"
)

type TbSearchItemService interface {
	ImportAllItems() *utils.E3Result
}

type tbSearchItemService struct {
	engine *xorm.Engine
}

func NewTbSearchItemService() TbSearchItemService {
	return &tbSearchItemService{
		engine: db.X,
	}
}

func (d *tbSearchItemService) ImportAllItems() *utils.E3Result {
	datalist := make([]models.SearchItem, 0)
	sql := `SELECT 
            a.id,
			a.title,
			a.sell_point,
			a.price,
			a.image,
			b. NAME category_name
		FROM
			tb_item a
		LEFT JOIN tb_item_cat b ON a.cid = b.id
		WHERE
			a.status = 1`

	err := d.engine.SQL(sql).Find(&datalist)
	if err != nil {
		panic(err)
	}

	//	datalist := make([]models.SearchItem, 0)
	//	err := d.engine.Desc("id").Find(&datalist)
	//	if err != nil {
	//		return nil
	//	}
	for _, v := range datalist {
		id := strconv.FormatInt(v.Id, 10)
		_, err := elastics.Client.Index().
			Index("mega").
			Type("TbItem").
			Id(id).
			BodyJson(v).
			Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
	return utils.Ok(nil)

}
