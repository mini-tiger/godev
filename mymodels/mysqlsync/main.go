package main

import (
	"database/sql"
	"fmt"
	_ "gitee.com/taojun319/godaemon"
	log "github.com/ccpaging/nxlog4go"
	slog "log"
	_ "github.com/go-sql-driver/mysql"
	"godev/mymodels/mysqlsync/bussiness"
	"strconv"
	"strings"
	"time"
	"flag"
	"github.com/toolkits/file"
	"os"
)

//const (
//	INTERVAL = 10
//)

var (
	NeedCol string
	Log     *log.Logger
)

type Col struct {
	id                          int
	begin_time                  string
	status, comment, alarm_type string
}

func mulitGet(db *sql.DB, sql string) []*Col {
	//fmt.Println(sql)
	tmp := make([]*Col, 0)
	rows, err := db.Query(sql)
	if err != nil {
		Log.Error("sql err:%s", sql)
	}
	defer rows.Close() // 必要

	for rows.Next() { //如果在rows 没有循环完之前就退出，rows不会关闭，链接还是打开状态，要rows.Close()
		tmpRow := &Col{}
		err := rows.Scan(&tmpRow.begin_time, &tmpRow.status, &tmpRow.comment, &tmpRow.alarm_type) //这里顺序要与上面 SQL中获取字段顺序一样
		if err != nil {
			Log.Error("sql scan err:%s", sql)
		}
		//log.Println(tmpRow)
		tmp = append(tmp, tmpRow)
	}
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
	return tmp

}

// 获取表数据
func oneGet(db *sql.DB, sql string) int64 {
	//直接执行语句

	rows := db.QueryRow(sql)

	//defer rows.Close()
	//cloumns, err := rows.Columns()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(cloumns)
	//fmt.Println(rows.Columns())
	var tmp string

	err := rows.Scan(&tmp)
	if err != nil {
		Log.Error("sql err:%s sql:%s", err, sql)
		return int64(0) //空值
	}
	t, err := strconv.ParseInt(tmp, 10, 64)
	return t
}

func subInsert(tx *sql.Tx, data []*Col, sql string) {
	for _, d := range data {
		_, err := tx.Exec(sql, d.id, d.begin_time, d.status, d.comment, d.alarm_type)
		if err != nil {
			Log.Error("sql db.Exec %s err %s", sql, err)
			tx.Rollback()
		}
	}
	err := tx.Commit()

	if err != nil {
		Log.Error("sql db.commit err %s", sql)
	}
}

// 插入数据
func Insert(db *sql.DB, data []*Col, sql string) {
	//使用同一个事务，不用重复发起链接，使用同一个链接循环插入数据  2m39s
	//tx, err := db.Begin()     //
	//if err != nil {
	//	Log.Error("sql db.Begin err")
	//}
	//subInsert(tx, data, sql)

	//使用同一个stmt链接，不用重复发起链接,使用同一个链接循环插入数据  1m22s
	//sql="INSERT INTO ja_alarm_alarminstance(begin_time,status,comment,alarm_type) VALUES(?,?,?,?);"   //
	////var sqldata string
	//stm,_:=db.Prepare(sql)
	//for _,d:=range data {
	//	stm.Exec(d.begin_time,d.status,d.comment,d.alarm_type)
	//}
	//stm.Close()

	//同一个insert语句插入多条语句，同一个链接 //10s  todo 网速不好的，这种性能最好
	var sqldata string
	sql = "INSERT INTO ja_alarm_alarminstance(begin_time,status,comment,alarm_type) VALUES"
	l := len(data)
	for i, d := range data {
		tmpData := fmt.Sprintf("( '%s','%s','%s','%s'),", d.begin_time, d.status, d.comment, d.alarm_type)
		sqldata = sqldata + tmpData
		if i == l-1 { // 最后 不能是，结尾 要是;
			tmpData := fmt.Sprintf("( '%s','%s','%s','%s');", d.begin_time, d.status, d.comment, d.alarm_type)
			sqldata = sqldata + tmpData
		}
	}
	res, err := db.Exec(sql + sqldata)
	if err != nil {
		Log.Info("sql : %s err", sql)
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
	Log.Info("起始插入的ID:%d, 本次影响的行数:%d", lastId, rowCnt)
}

// 删除数据
//func Delete(db *sql.DB) {
//	stmt, err := db.Prepare("DELETE FROM user WHERE name='python'")
//	if err != nil {
//		log.Fatal(err)
//	}
//	res, err := stmt.Exec()
//	lastId, err := res.LastInsertId()
//	if err != nil {
//		log.Fatal(err)
//	}
//	rowCnt, err := res.RowsAffected()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
//}

//// 更新数据
//func Update(db *sql.DB) {
//	stmt, err := db.Prepare("UPDATE user SET age=27 WHERE name='python'")
//	if err != nil {
//		log.Fatal(err)
//	}
//	res, err := stmt.Exec()
//	lastId, err := res.LastInsertId()
//	if err != nil {
//		log.Fatal(err)
//	}
//	rowCnt, err := res.RowsAffected()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
//}

func createDb(user, passwd, ip, dbname string) *sql.DB {
	db := &sql.DB{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, passwd, ip, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		Log.Info("Can't connect to database")
		return db
	} else {
		err = db.Ping()
		if err != nil {
			Log.Info("db.Ping %s failed:", ip)
			return db
		}
	}
	//db.SetMaxIdleConns(5)
	//db.SetMaxOpenConns(5)
	return db
}

func closeDb(db *sql.DB) {
	db.Close()
}

//10.240.grant all on *.* to 'sync'@'%' identified by 'sync@!123';
func main() {
	cfgfile := flag.String("c", "cfg.json", "configuration file")
	flag.Parse()
	f, _ := os.OpenFile("mysqlsync.log", os.O_RDWR|os.O_CREATE, 0755) //加载日志前用
	logger := slog.New(f, "",slog.Llongfile)

	if !file.IsExist(*cfgfile) {
		logger.Printf("arg -c config file: %s,is not existent. use default cfg.json`",*cfgfile)
		//*cfgfile = "C:\\work\\go-dev\\src\\godev\\mymodels\\mysqlsync\\cfg.json"
	}
	f.Close()
	// 读取配置文件
	bussiness.InitConfig(*cfgfile)
	cfg := bussiness.Config
	NeedCol = strings.Join(cfg.SrcDb.NeedCol, ",")
	// init log
	bussiness.Initlog()
	Log=bussiness.Returnlog()

	for {
		srcdb := createDb(cfg.SrcDb.DbUser, cfg.SrcDb.DbPass, cfg.SrcDb.Ip, cfg.SrcDb.DbName)

		dstdb := createDb(cfg.DstDb.DbUser, cfg.DstDb.DbPass, cfg.DstDb.Ip, cfg.DstDb.DbName)
		if dstdb.Stats().OpenConnections == 0 || srcdb.Stats().OpenConnections == 0 { //目标库与源库 如果某一个连接不上，进入下一次循环
			srcdb.Close()
			dstdb.Close()

			time.Sleep(time.Second * time.Duration(cfg.Interval))
			continue
		}

		var sql string
		//查找 源库与目标库最新 UNIX时间

		sql = fmt.Sprintf("SELECT UNIX_TIMESTAMP(e.%s) from ja_alarm_alarminstance as e ORDER BY begin_time desc limit 1;", cfg.SrcDb.SortCol)
		srcDate := oneGet(srcdb, sql) //取 源库 目标库 最新的时间戳
		sql = fmt.Sprintf("SELECT UNIX_TIMESTAMP(e.%s) from ja_alarm_alarminstance as e ORDER BY begin_time desc limit 1;", cfg.DstDb.SortCol)
		dstDate := oneGet(dstdb, sql)
		Log.Info("srcDb New Datadate: %s, dstDb New Datadate: %s",
			time.Unix(srcDate, 0).Format("2006-01-02 15:04:05"),
			time.Unix(dstDate, 0).Format("2006-01-02 15:04:05"))

		if srcDate == int64(0) { //源库最新是0 ，代表被清空
			Log.Info("srcDb no data")
			continue //结束本次循环
		}
		//INSERT INTO ja_alarm_alarminstance(id,begin_time,status,comment,alarm_type) VALUES(1,'2017-03-02 15:22:22',3,4,5);
		if srcDate > dstDate { //源数据库中数据 有比 目标库中新的数据
			sql = fmt.Sprintf("select %s from ja_alarm_alarminstance as e where UNIX_TIMESTAMP(e.begin_time) > %d", NeedCol, dstDate)
			insertData := mulitGet(srcdb, sql) //查找大于目标库时间戳的数据
			Log.Info("Need insertData Len:%d", len(insertData))
			sql = fmt.Sprintf("INSERT INTO ja_alarm_alarminstance(id,begin_time,status,comment,alarm_type) VALUES(?,?,?,?,?);")
			Insert(dstdb, insertData, sql)
		}
		closeDb(srcdb)
		closeDb(dstdb)

		time.Sleep(time.Second * time.Duration(cfg.Interval))
	}
}