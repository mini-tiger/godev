package g

import "tjtools/nmap"

//var FieldsMap map[string]string=map[string]string{"客户端":"ReportClient","主机名":"HOSTNAME","总作业数":"TOTALJOB",
//													"已完成":"COMPLETED","完成但有错误":"COMPLETIONERROR","完成但有警告":"COMPLETIONWARN",
//													"已终止":"TERMINATION","不成功":}
var AllChEngSummaryMap *nmap.SafeMap= nmap.NewSafeMap()

func init() {

	AllSummary := map[string]interface{}{"客户端":"REPORTCLIENT","主机名":"HOSTNAME","总作业数":"TOTALJOB",
														"已完成":"COMPLETED","完成但有错误":"COMPLETIONERROR","完成但有警告":"COMPLETIONWARN",
														"已终止":"KILLE","不成功":"UNSUCCESSFUL",""}
	AllChEngSummaryMap.M=AllSummary
	//var a *nmap.SafeMap=&nmap.SafeMap{M:AllChEngSummaryMap}

	//var aa *nmap.SafeMap = &nmap.NewSafeMap(a)
}

// 摘要表
var SummaryFieldsMap map[int]string = map[int]string{0: "REPORTCLIENT", 1: "HOSTNAME", 2: "TOTALJOB",
	3: "COMPLETED", 4: "COMPLETEDWITHERRORS", 5: "COMPLETEDWITHWARNINGS", 6: "KILLED", 7: "UNSUCCESSFUL", 8: "RUNNING", 9: "DELAYED",
	10: "NORUN", 11: "NOSCHEDULE", 12: "COMMITTED", 13: "SIZEOFAPPLICATION", 14: "DATAWRITTEN", 15: "STARTTIME", 16: "ENDTIME", 17: "PROTECTEDOBJECTS",
	18: "FAILEDOBJECTS", 19: "FAILEDFOLDERS"}

// 摘要表cv10    缺少 1，12
var SummaryFieldsMapCv10 map[int]string = map[int]string{0: "REPORTCLIENT", 1: "TOTALJOB",
	2: "COMPLETED", 3: "COMPLETEDWITHERRORS", 4: "COMPLETEDWITHWARNINGS", 5: "KILLED", 6: "UNSUCCESSFUL", 7: "RUNNING", 8: "DELAYED",
	9: "NORUN", 10: "NOSCHEDULE", 11: "SIZEOFAPPLICATION", 12: "DATAWRITTEN", 13: "STARTTIME", 14: "ENDTIME", 15: "PROTECTEDOBJECTS",
	16: "FAILEDOBJECTS", 17: "FAILEDFOLDERS"}

// 摘要表cv8   缺少 1，5，9，12
var SummaryFieldsMapCv8 map[int]string = map[int]string{0: "REPORTCLIENT", 1: "TOTALJOB",
	2: "COMPLETED", 3: "COMPLETEDWITHERRORS", 4: "KILLED", 5: "UNSUCCESSFUL", 6: "RUNNING",
	7: "NORUN", 8: "NOSCHEDULE", 9: "SIZEOFAPPLICATION", 10: "DATAWRITTEN", 11: "STARTTIME", 12: "ENDTIME", 13: "PROTECTEDOBJECTS",
	14: "FAILEDOBJECTS", 15: "FAILEDFOLDERS"}

var SummaryFieldsMapPlus map[int]string = map[int]string{20: "COMMCELL", 21: "REPORTTIME"} // 这两个字段从页面上开头获取,唯一联合字段，也是和详细表 关系的字段

// 详细数据表
var DetailFieldsMap map[int]string = map[int]string{0: "DATACLIENT", 1: "AgentInstance", 2: "BackupSetSubclient", 3: "Job ID (CommCell)(Status)",
	4: "Type", 5: "Scan Type", 6: "Start Time(Write Start Time)", 7: "End Time or Current Phase", 8: "Size of Application", 9: "Data Transferred",
	10: "Data Written", 11: "Data Size Change", 12: "Transfer Time", 13: "Throughput (GB/Hour)", 14: "Protected Objects", 15: "Failed Objects",
	16: "Failed Folders"}

// 详细数据表 cv8
var DetailFieldsMapCv8 map[int]string = map[int]string{0: "DATACLIENT", 1: "AgentInstance", 2: "BackupSetSubclient", 3: "Job ID (CommCell)(Status)",
	4: "Type", 5: "Scan Type", 6: "Start Time(Write Start Time)", 7: "End Time or Current Phase", 8: "Size of Application",
	10: "Data Size Change", 11: "Transfer Time", 12: "Throughput (GB/Hour)", 13: "Protected Objects", 14: "Failed Objects",
	15: "Failed Folders"}

var DetailFieldsMapPlus map[int]string = map[int]string{
	17: "COMMCELL", 18: "REPORTTIME", // 与摘要表一样
	19: "JOBTYPE", 20: "START TIME", 21: "DATASUBCLIENT", //作业状态  ,通过开始时间 和 子客户端，格式化出来的字段
	22: "REASONFORFAILURE", 23: "SOLVETYPE", //失败原因(失败，完成..),，解决状态
	24: "SOLVETIME", 25: "ENGINEER"} // 解决时间，工程师， 这几个字段不用插入数据，

//todo 有问题的行 解决状态默认是未解决

type StatusColor struct {
	ColorStatus string
	Status      int
}

// 0 不检查直接录入， 1 检查下一行失败原因  2.不录入
var StatusColors map[string]StatusColor = map[string]StatusColor{"#66ABDD": StatusColor{"运行中", 0},
	"#CC99FF": StatusColor{"已延迟", 1},
	"#CCFFCC": StatusColor{"已完成", 0},
	"#FFCC99": StatusColor{"完成但有错误", 1},
	"#00FFCC": StatusColor{"完成但有警告", 1},
	"#FF99CC": StatusColor{"已终止", 1},
	"#FF3366": StatusColor{"失败", 1},
	"#CC9999": StatusColor{"过时", 0},
	"#FFFFFF": StatusColor{"无计划", 2},
	"#FF9999": StatusColor{"未运行", 1},
	"93C54B":  StatusColor{"提交", 0},
	"#CCFFFF": StatusColor{"数据大小按 10% 或更多增加/减少", 0}}

// 0 不检查直接录入， 1 检查下一行失败原因  2.不录入
var StatusColorsCv8 map[string]StatusColor = map[string]StatusColor{"#CCFFFF": StatusColor{"运行中", 0},
	"#CC99FF": StatusColor{"已延迟", 1},
	"#CCFFCC": StatusColor{"已完成", 0},
	"#FFCC99": StatusColor{"完成但有错误", 1},
	"#FF99CC": StatusColor{"已终止", 1},
	"#FF3366": StatusColor{"失败", 1},
	"#CC9999": StatusColor{"过时", 0},
	"#FFFFFF": StatusColor{"无保护", 0},
	"66ABDD":  StatusColor{"数据大小按 10% 或更多增加/减少", 0}}
