package funcs

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"godev/mymodels/falconApi/models"
	"log"
)

type Page struct {
	PageNo     int                  // 当前页码
	PageSize   int                  // 每页多少行
	TotalPage  int                  // 共多少页
	TotalCount int                  // 共多少行
	FirstPage  bool                 // 是否第一页
	LastPage   bool                 // 是否最后一页
	List       []*models.EventCases // 数据
	Error      string
}

type CountTop struct {
	Error  string
	Events []*models.EventCases
}

func GetEventCase(count int) (c CountTop) {
	o1 := orm.NewOrm()
	o1.Using("default")
	//var events []*models.EventCases
	// todo  前count条，  时间倒序
	var events []*models.EventCases
	num, err := o1.QueryTable("event_cases").Filter("id__gt", 0).OrderBy("-timestamp").Limit(count).All(&events)
	if err != nil {
		c.Error = err.Error()
		log.Println("Returned Rows Num: %d, %s\n", num, err)
	}
	c.Events = events
	return
}

func GetSingleEvent(ep string, pageNo int, pageSize int) (rp Page) {
	var all []*models.EventCases //这是 params是orm模块自带的，不是模型

	o1 := orm.NewOrm()
	o1.Using("default")
	qs := o1.QueryTable("event_cases").Filter("endpoint", ep) // 数据结果集

	var off int
	if pageNo == 1 { // 如果第一页 偏移0个
		off = 0
	} else {
		off = (pageNo - 1) * pageSize // 不是第一页 页码-1   在 乘以 每页个数
	}

	//var rp Page
	rp.PageNo = pageNo     // 当前页码
	rp.PageSize = pageSize // 每页多少行
	rc, err := qs.Count()
	if err != nil {
		rp.Error = err.Error()
		return
	}
	rp.TotalCount = int(rc)             // 共多少行
	if rp.TotalCount%rp.PageSize == 0 { // 如果总行数 能被 每页行数 除尽
		rp.TotalPage = rp.TotalCount / rp.PageSize // 共有多少页
	} else {
		rp.TotalPage = rp.TotalCount/rp.PageSize + 1
	}
	rp.FirstPage = false
	if rp.PageNo == 1 { // 如果当前页码是1  ，是首页
		rp.FirstPage = true
	}
	rp.LastPage = false
	if rp.TotalPage == rp.PageNo { // 如果当前页码和 总共页数相等，是最后页
		rp.LastPage = true
	}

	_, err = qs.OrderBy("-timestamp").Limit(pageSize).Offset(off).All(&all)
	if err != nil {
		rp.Error = err.Error()
		return
	}
	fmt.Printf("Returned Rows: %+v,rows:%d\n", all, len(all))
	//for _, m := range all {
	//	fmt.Println(m["Name"], m["Update"]) //todo 这里要用大写
	//}
	rp.List = all // 当前页所有数据
	return
	//fmt.Printf("%+v", rp)
	//c.Data["json"] = &rp
}
