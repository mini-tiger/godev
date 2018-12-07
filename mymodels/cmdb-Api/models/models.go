package models

type AddhostResp struct {
	Result bool
	Bk_error_code int64
	Bk_error_msg string
	Data map[string]interface{}`json:"data"`
}
type SearchResp struct {
	Result bool
	Bk_error_code int64
	Bk_error_msg string
	Data SearchRespSub`json:"data"`
}
type SearchRespSub struct {
	Count uint
	Info []SearchDataSub
}

type SearchDataSub struct {
	Set interface{}
	Biz interface{}
	Host map[string]interface{}
}

type UpdateHostResp struct {
	Result bool
	Bk_error_code uint
	Bk_error_msg string
	Data string
}




type AddhostRespDataField struct {
	Bk_comment string `json:"Bk_comment"`
	Error string	`json:"error"`
	Success string	`json:"success"`
	Update_error string	`json:"update_error"`

}

type AddHost struct {
	Host_info map[string]*AddHostInfo 	`json:"host_info"`
	Hk_supplier_id int	`json:"bk_supplier_id"`
	Hk_biz_id	int		`json:"bk_biz_id"`
}

type AddHostInfo struct {
	Bk_host_innerip string `json:"bk_host_innerip"`
	//Operator		string	`json:"operator"`
	//Bk_asset_id		string
	Bk_sn			string `json:"bk_sn"`
	Bk_comment		string	`json:"bk_comment"`
	//Bk_service_term	string
	//Bk_host_name	string
	//Bk_os_type		string
	//Bk_os_name		string
	//Bk_os_version	string
	//Bk_os_bit		string
	//Bk_cpu			string
	//Bk_cpu_mhz		string
	//Bk_cpu_module	string
	//Bk_mem			string
	//Bk_disk			string
	//Bk_mac			string
	//Bk_outer_mac	string
	//Import_from		string
}

type SearchHostIp struct {
	Data []interface{}
	Exact uint
	Flag string
}

type SearchHostConSub struct {
	Bk_obj_id string `json:"bk_obj_id"`
	Fields []interface{} `json:"fields"`
	Condition []*ConSub `json:"condition"`
}

type ConSub struct {
	Field string `json:"field"`
	Operator string `json:"operator"`
	Value string `json:"value"`
}

type SearchHost struct {
	Ip *SearchHostIp `json:"ip"`
	//Condition []*SearchHostConSub `json:"condition"`
	Condition []interface{} `json:"condition"`
	Page map[string]interface{} `json:"page"`
	//Pattern string `json:"pattern"`
}

type UpdateHost struct {
	Bk_host_id string `json:"bk_host_id"`
	Bk_sn string `json:"bk_sn"`


	Bk_host_innerip string `json:"bk_host_innerip"`
	//Bk_comment string `json:"bk_comment"`
	Uuid string	`json:"uuid"`
	Bk_host_name  string `json:"bk_host_name"`
	//Bk_os_type    int   `json:"bk_os_type,string"`
	Bk_os_name    string `json:"bk_os_name"`
	Bk_os_version string `json:"bk_os_version"`
	Bk_os_bit     string `json:"bk_os_bit"`
	//Bk_cpu        int    `json:"bk_cpu,string"`
	//Bk_cpu_mhz    int    `json:"bk_cpu_mhz"`
	Bk_cpu_module string `json:"bk_cpu_module"`
	Bk_mem        int64  `json:"bk_mem,int"`
	Bk_disk       int64  `json:"bk_disk,int"`
	Bk_mac        string `json:"bk_mac"`
	Bk_outer_mac  string `json:"bk_outer_mac"`
	//BK_cloud_id	  int	`json:"bk_cloud_id"`


}
