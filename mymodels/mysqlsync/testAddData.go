package main

import (
	"os"
	"godev/mymodels/mysqlsync/bussiness"
	"strings"
	"time"
	"fmt"
	"database/sql"
	"math/rand"
	_ "github.com/go-sql-driver/mysql"
	)

var (
	NeedCol1 string

)
type Col1 struct {
	id                          int
	begin_time                  string
	status, comment, alarm_type string
}

func Insert1(db *sql.DB, data []*Col1, sql string) {

	//同一个insert语句插入多条语句，同一个链接 //10s  todo 网速不好的，这种性能最好
	var sqldata string
	//sql = "INSERT INTO ja_alarm_alarminstance(begin_time,status,comment,alarm_type) VALUES"
	l := len(data)
	for i, d := range data {

		if i == l-1 { // 最后 不能是，结尾 要是;
			tmpData := fmt.Sprintf("( '%s','%s','%s','%s');", d.begin_time, d.status, d.comment, d.alarm_type)
			sqldata = sqldata + tmpData
			break
		}
		tmpData := fmt.Sprintf("( '%s','%s','%s','%s'),", d.begin_time, d.status, d.comment, d.alarm_type)
		sqldata = sqldata + tmpData
	}
	fmt.Println(sql+sqldata)
	res, err := db.Exec(sql + sqldata)
	if err != nil {
		fmt.Println(err)
		return
	}
	lastId, err := res.LastInsertId() // 首先插入的ID
	//if err != nil {
	//	log.Fatal(err)
	//}
	rowCnt, err := res.RowsAffected() //本次影响的行数
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Printf("起始插入的ID:%d, 本次影响的行数:%d\n", lastId, rowCnt)
}

var alaremTypesl []string = []string{"mem", "cpu"}
var statussl []string = []string{"skipped", "converged"}


func createDb1(user, passwd, ip, dbname string) *sql.DB {
	db := &sql.DB{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, passwd, ip, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Can't connect to database")
		return db
	} else {
		err = db.Ping()
		if err != nil {
			fmt.Println("db.Ping %s failed:", ip)
			return db
		}
	}
	//db.SetMaxIdleConns(5)
	//db.SetMaxOpenConns(5)
	return db
}

func closeDb1(db *sql.DB) {
	db.Close()
}

func main() {
	//不能添加 -d false，不后台 就不加-d
	var cfgfile string

	if len(os.Args) == 1 { //daemon没有传入任何参数， 通过执行文件直接运行dd.exe
		cfgfile = "C:\\work\\go-dev\\src\\godev\\mymodels\\mysqlsync\\cfg.json"
	} else {

		if os.Args[1] == "-c" { //dd.exe -c  cfgfile
			cfgfile = os.Args[2]
		} else { // dd.exe -d true or -d true -c cfgfile
			cfgfile = "cfg.json"
		}
	}

	//cfgfile := flag.String("c", "cfg.json", "configuration file")
	bussiness.Cfg = cfgfile

	// 读取配置文件
	bussiness.InitConfig()
	cfg := bussiness.Config
	NeedCol1 = strings.Join(cfg.SrcDb.NeedCol, ",")
	// init log
	//bussiness.Initlog()
	//Log = bussiness.Returnlog()
	srcdb := createDb1("root","W3b5Ev!c3","192.168.130.17:3306","test")
	for {

		//dstdb := createDb(cfg.DstDb.DbUser, cfg.DstDb.DbPass, cfg.DstDb.Ip, cfg.DstDb.DbName)
		if srcdb.Stats().OpenConnections == 0 { //目标库与源库 如果某一个连接不上，进入下一次循环
			srcdb.Close()

			time.Sleep(time.Second * time.Duration(cfg.Interval))
			continue
		}

		var sql string

		sql = fmt.Sprintf("INSERT INTO ja_alarm_alarminstance(begin_time,status,comment,alarm_type) VALUES")
		t1 := time.Now()
		tmp := &Col1{rand.Intn(10), t1.Format("2006-01-02 15:04:05"),
		statussl[rand.Intn(len(statussl))], "", alaremTypesl[rand.Intn(len(alaremTypesl))]}
		insertData:=[]*Col1{tmp}
		Insert1(srcdb, insertData, sql)



		time.Sleep(time.Second * 2)
	}
	closeDb1(srcdb)
}
