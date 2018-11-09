package models

type AddhostResp struct {
	Result string
	Bk_error_code int64
	Bk_error_msg string
	Data map[string]interface{}`json:"data"`
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
