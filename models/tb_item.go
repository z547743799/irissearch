package models

type SearchItem struct {
	Id            int64  `json:"id" xorm:"pk comment('商品id，同时也是商品编号') BIGINT(20)"`
	Title         string `json:"title" xorm:"not null comment('商品标题') VARCHAR(100)"`
	SellPoint     string `json:"sell_point" xorm:"comment('商品卖点') VARCHAR(500)"`
	Price         int64  `json:"price" xorm:"not null comment('商品价格，单位为：分') BIGINT(20)"`
	Image         string `json:"image" xorm:"comment('商品图片') VARCHAR(500)"`
	Category_name string `json:"category_name" xorm:"comment('类目名称') VARCHAR(50)"`
}
