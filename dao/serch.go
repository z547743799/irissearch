package dao

import (
	"fmt"
	"reflect"

	"github.com/olivere/elastic"
)

func printEmployee(res *elastic.SearchResult, err error, str interface{}) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ str
	//em := make([]Employee, 0)
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(str)
		fmt.Printf("%#v\n", t)
		//em = append(em, t)

		//fmt.Println(item)
	}
	//sb := sdf{Ss: em, Ssi: "---"}
	//js, err := json.Marshal(sb)
	//fmt.Println(err)
	//fmt.Println(string(js))
}
