package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"godev/mymodels/beego/models"
)

type Study struct {
	beego.Controller
}

type User struct {
	Name interface{} `form:"username",json:"name"`
	Age  int         `form:"age",json:"age"`
}

func (c *Study) Get() {
	fmt.Println(c.GetString("abcpost")) // todo 获取 post, get参数
	fmt.Println(c.Input().Get("abcpost"))
	c.Data["data1"] = "taojun"
	c.Data["data2"] = "this is test data"

	o1 := orm.NewOrm()
	o1.Using("default")
	fmt.Println(o1.Driver().Name())                // 数据库别名，default  切换不同数据库用
	fmt.Println(o1.Driver().Type() == orm.DRMySQL) // 数据库类型

	user := models.User{Id: 0} // todo 空数据库 先查询
	err := o1.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name)
	}

	// todo 读 若不存在则创建
	user = models.User{Name: "11", Email: "@111"}
	// 默认必须传入一个参数作为条件字段，同时也支持多个参数多个条件字段
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	if created, id, err := o1.ReadOrCreate(&user, "Name"); err == nil { // todo 通过Name作为判断是否有数据标准
		if created {
			fmt.Println("新建立的数据. Id:", id)
		} else {
			fmt.Println("找到一条数据. Id:", id)
		}
	}

	// todo 根据 人 添加对应的多个 课程
	c1 := models.Class{Name: "shuxue", User: &user}
	id, err := o1.Insert(&c1)
	if err == nil {
		fmt.Println(id)
	}

	c2 := models.Class{Name: "yuwen", User: &user}
	id, err = o1.Insert(&c2)
	if err == nil {
		fmt.Println(id)
	}
	//todo 查询所有user id大于0
	fmt.Println("===========查询所有user id大于0===================")
	var users []*models.User
	num, err := o1.QueryTable("auth_user").Filter("id__gt", 0).All(&users)
	fmt.Printf("Returned Rows Num: %d, %s\n", num, err)

	//todo 查询所有user  返回MAP
	fmt.Println("===========查询所有user  返回MAP===================")
	//uu:=map[string]interface{}
	var users2 []orm.Params                                                //这是 params是orm模块自带的，不是模型
	num, err = o1.QueryTable("auth_user").Values(&users2, "name", "email") // todo 可以分字段或者全部
	fmt.Printf("Returned Rows Num: %d, %s\n", num, err)
	for _, m := range users2 {
		fmt.Println(m["Name"], m["Email"]) //todo 这里要用大写
	}

	//todo 查询所有user  返回slice
	//  每个子元素都是一个切片，排列与 Model 中定义的 Field 顺序一致
	fmt.Println("===========查询所有user  返回slice===================")
	var lists []orm.ParamsList
	num, err = o1.QueryTable("auth_user").ValuesList(&lists, "name", "class__name")
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		fmt.Println(lists)
	}

	//todo 查询所有user  ,每个user 单个列 存入slice
	fmt.Println("===========查询所有user  ,每个user 单个列 存入slice===================")
	var lists1 orm.ParamsList
	num, err = o1.QueryTable("class").ValuesFlat(&lists1, "name")
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		fmt.Println(lists1)
	}

	// todo 通过人 查找多个课程
	fmt.Println("==============todo 通过人 查找多个课程===========")
	var clss []*models.Class
	num, err = o1.QueryTable("Class").Filter("User", user.Id).RelatedSel().All(&clss)
	if err == nil {
		fmt.Printf("%d posts read\n", num)
		for _, post := range clss {
			fmt.Printf("Id: %d, Name: %s User:%+v,\n", post.Id, post.Name, post.User)
		}
	}

	// todo 通过课程是数学的 属于哪个人
	// 方法一
	fmt.Println("==============课程是数学的 属于哪个人===========")
	var user11 []orm.Params
	num, err = o1.QueryTable("auth_user").Filter("class__name", "shuxue").Values(&user11)
	if err == nil {
		//fmt.Println(user11)
		fmt.Printf("数学user:%+v 拥有\n", user11[0]["Name"])
	}
	// 方法二
	var cc1 models.Class
	var user2 models.User
	err = o1.QueryTable("class").Filter("name__icontains", "shuxue").Limit(1).One(&cc1)
	if err == nil {
		fmt.Println(cc1)
		o1.QueryTable("auth_user").Filter("id", cc1.User).Limit(1).One(&user2)
		fmt.Printf("数学user:%+v 拥有\n", user2.Name)
	}

	//todo 查询所有 分页
	p := 3  // 当前是第几页
	pp := 2 // 每页多少个

	fmt.Println("===========查询所有 分页===================")

	var all []orm.Params             //这是 params是orm模块自带的，不是模型
	qs := o1.QueryTable("auth_user") // 数据结果集
	var off int
	if p == 1 { // 如果第一页 偏移0个
		off = 0
	} else {
		off = (p - 1) * pp // 不是第一页 页码-1   在 乘以 每页个数
	}

	var rp Page
	rp.PageNo = p    // 当前页码
	rp.PageSize = pp // 每页多少行
	rc, err := qs.Count()
	rp.TotalCount = int(rc) // 共多少行
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

	qs.OrderBy("-update").Limit(pp).Offset(off).Values(&all)
	fmt.Printf("Returned Rows Num: %+v\n", all)
	//for _, m := range all {
	//	fmt.Println(m["Name"], m["Update"]) //todo 这里要用大写
	//}
	rp.List = all // 当前页所有数据
	fmt.Printf("%+v", rp)
	c.Data["json"] = &rp
	c.ServeJSON()		//	 返回JSON

	// todo 级联删除 用户下的课程
	//o1.QueryTable("auth_user").Filter("id",1).Delete()

	// todo 返回HTML
	//c.TplName = "study/index.tpl" //
}

type Page struct {
	PageNo     int         // 当前页码
	PageSize   int         // 每页多少行
	TotalPage  int         // 共多少页
	TotalCount int         // 共多少行
	FirstPage  bool        // 是否第一页
	LastPage   bool        // 是否最后一页
	List       interface{} // 数据
}

func (c *Study) Post() {
	//u := User{}
	//if err := c.ParseForm(&u); err != nil { // todo formdata
	//	fmt.Println(err)
	//}
	//fmt.Println(u)
	fmt.Println(string(c.Ctx.Input.RequestBody))
	u := User{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &u)
	fmt.Println(u)

	//c.Data["data1"] = "taojun"
	//c.Data["data2"] = "this is test data"
	//c.TplName = "study/index.tpl" //
	c.Data["json"] = u
	c.ServeJSON()

}
