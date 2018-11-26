package main

import (
	"godev/mymodels/cmdb-Api/models"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
	"log"
)

func HostAdd() {
	url := "http://192.168.43.202:8083/api/v3/hosts/add"
	addhost := models.AddHost{}
	addhost.Hk_biz_id = 3
	addhost.Hk_supplier_id = 0
	addhostinfo := models.AddHostInfo{}
	addhostinfo.Bk_host_innerip = "192.168.1.1"
	addhostinfo.Bk_sn = "VMware-42 3d bc 1d b1 76 04 9b-07 22 4d 40 a0 0c c3 f4"
	addhostinfo.Bk_comment = "备注new"
	addhost.Host_info = map[string]*models.AddHostInfo{"0": &addhostinfo}
	fmt.Println(addhost)
	// 方法一
	req := httplib.NewBeegoRequest(url, "POST")
	req.SetTimeout(time.Duration(10)*time.Second, time.Duration(10)*time.Second)

	req.Header("Content-Type", "application/json")
	req.Header("dataType", "json")
	req.Header("BK_USER", "admin")
	req, err := req.JSONBody(addhost)
	if err != nil {
		fmt.Printf("req fail err:%s\n", err)
	}
	var respdata models.AddhostResp
	err = req.ToJSON(&respdata)
	if err != nil {
		fmt.Printf("tojson err:%s\n", err)
	}
	fmt.Printf("%+v\n", respdata)

	//resp,err:=req.Response()
	//if err!=nil{
	//	fmt.Printf("response err:%s\n",err)
	//}
	//respBytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(string(respBytes))

	// 方法二
	//bytesData, err := json.Marshal(addhost)
	//if err != nil {
	//	fmt.Println(err.Error() )
	//	return
	//}
	//reader := bytes.NewReader(bytesData)
	//
	//
	//client := &http.Client{
	//	Timeout: 10 * time.Second,
	//}
	//req, err := http.NewRequest("POST", url, reader)
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("dataType", "json")
	//req.Header.Add("BK_USER", "admin")
	////req.Header.Add("HTTP_BLUEKING_SUPPLIER_ID", "0")
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//respBytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//var respdata models.AddhostResp
	//json.Unmarshal(respBytes,&respdata)
	//fmt.Printf("%+v\n",respdata)
	//
	//var cc map[string]interface{}
	//json.Unmarshal(respBytes, &cc) //解析JSON
	//fmt.Printf("%T,%+v\n", cc, cc)
}
func main() {

	//HostAdd()
	//SearchHost()
	updateHost()

}
func SearchHost() {
	url := "http://192.168.43.202:8083/api/v3/hosts/search"
	Searchhost := models.SearchHost{}
	tmpcon := map[string]interface{}{"bk_obj_id": "host", "fields": []interface{}{"bk_sn"}} // fields 返回的属性
	tmpcon["condition"] = []interface{}{
		map[string]interface{}{"field": "uuid", "operator": "$eq", "value": "ddddd"},
	}

	Searchhost.Condition = []interface{}{tmpcon}

	Searchhost.Ip = &models.SearchHostIp{[]interface{}{}, 1, "bk_host_innerip|bk_host_outerip"}
	//Searchhost.Pattern=""
	Searchhost.Page = map[string]interface{}{"start": 0, "limit": 1, "sort": "bk_sn"}

	fmt.Printf("%+v\n", Searchhost)

	req := httplib.NewBeegoRequest(url, "POST")
	req.SetTimeout(time.Duration(10)*time.Second, time.Duration(10)*time.Second)

	req.Header("Content-Type", "application/json")
	req.Header("dataType", "json")
	req.Header("BK_USER", "admin")
	req, err := req.JSONBody(Searchhost)
	if err != nil {
		fmt.Printf("req fail err:%s\n", err)
	}
	var respdata models.SearchResp
	err = req.ToJSON(&respdata)
	if err != nil {
		fmt.Printf("tojson err:%s\n", err)
	}
	fmt.Printf("%+v\n", respdata)

	if respdata.Result {
		fmt.Printf("%T,%+v\n", respdata.Data.Info, respdata.Data.Info)
		if len(respdata.Data.Info)>=1{
			s:=respdata.Data.Info[0]
			fmt.Printf("%T,%+v\n", s.Host, s.Host)
			fmt.Println(s.Host["bk_sn"])
		}else{
			log.Println("没有匹配到")
		}

	}else{
		log.Println("接口返回错误")
	}

}
func updateHost()  {
	url := "http://192.168.43.202:8083/api/v3/hosts/batch"
	updateHost := models.UpdateHost{}
	updateHost.Bk_host_id="17"
	updateHost.Bk_sn="ddddaaaacccc"

	fmt.Printf("%+v\n", updateHost)

	req := httplib.NewBeegoRequest(url, "PUT")
	req.SetTimeout(time.Duration(10)*time.Second, time.Duration(10)*time.Second)

	req.Header("Content-Type", "application/json")
	req.Header("dataType", "json")
	req.Header("BK_USER", "admin")
	req, err := req.JSONBody(updateHost)
	if err != nil {
		fmt.Printf("req fail err:%s\n", err)
	}
	var respdata models.UpdateHostResp
	err = req.ToJSON(&respdata)
	if err != nil {
		fmt.Printf("tojson err:%s\n", err)
	}
	fmt.Printf("%+v\n", respdata)

	if respdata.Result {
		fmt.Printf("%T,%+v\n", respdata, respdata)


	}else{
		log.Println("接口返回错误")
	}
}