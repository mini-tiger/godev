package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
	"github.com/jander/golog/logger"
	"time"
)

type SqliteDb struct {
	sync.RWMutex
	DB *sql.DB
}

var SqlConn SqliteDb

func ReturnSqlDB() *SqliteDb {
	return &SqlConn
}

func NewConn(db string) (err error) {
	db1, err := sql.Open("sqlite3", db)
	SqlConn.DB = db1
	return
}

func CreateSelectSql(table string, col, where string) (sql string) {
	sql = fmt.Sprintf("select %s from %s where %s", col, table, where)
	return
}
func (d *SqliteDb) GetAgentVer() (ver string, err error) {
	d.Lock()
	defer func() {
		d.Unlock()
	}()
	rows, err := d.DB.Query("select version from Agent limit 1")
	defer func() {
		rows.Close()
	}()
	for rows.Next() {
		err = rows.Scan(&ver)
		if err != nil {
			return
		}
	}
	fmt.Println(ver)
	return
}

func (d *SqliteDb) GetData(sql string, ver *string) error {
	d.Lock()
	defer func() {
		d.Unlock()
	}()
	rows, err := d.DB.Query(sql + " limit 1")
	defer func() {
		rows.Close()
	}()
	if err != nil {
		return err
	}
	for rows.Next() {
		err = rows.Scan( ver)
		if err != nil {

			return err
		}
	}

	return nil
}
func (d *SqliteDb) GetExist(sql string) (error, bool) {
	d.Lock()
	defer func() {
		d.Unlock()
	}()
	rows, err := d.DB.Query(sql + " limit 1")

	if err != nil {

		return err, false
	}
	defer func() {
		rows.Close()
	}()
	return err, rows.Next()
	//checkErr(err)
	//var uuid string
	//var ver string
	//for rows.Next() {
	//
	//	err = rows.Scan(&uuid, &ver)
	//	checkErr(err)
	//	fmt.Println(uuid, ver)
	//}
	//if uuid !=""{
	//	return true
	//}
	//return false
}

func (d *SqliteDb) UpdateAgentVer(where string, version, time string) (err error) {
	d.Lock()
	defer func() {
		d.Unlock()
	}()
	stmt, err := d.DB.Prepare("update agent set version= ?, `time`=? where uuid = ?")
	//if err!=nil{
	//	return err
	//}

	res, err := stmt.Exec(version, time, where)
	//if err!=nil{
	//	return err
	//}

	affect, err := res.RowsAffected()
	//if err!=nil{
	//	return err
	//}

	if affect == 0 {
		return errors.New("update affect = 0")
	}
	return
}

func (d *SqliteDb) InsertAgent(ver, uuid, time string) (err error) {
	d.Lock()
	defer func() {
		d.Unlock()
	}()
	stmt, err := d.DB.Prepare("INSERT INTO agent(version ,uuid ,time) values(?,?,?)")

	res, err := stmt.Exec(ver, uuid, time)

	id, err := res.LastInsertId()
	if id == 0 {
		err = errors.New("insertAgent table == 0")
	}
	return
}
func (d *SqliteDb) InsertUuid(uuid string,bizid int) (err error) {
	d.Lock()
	defer func() {
		d.Unlock()
	}()
	stmt, err := d.DB.Prepare("INSERT INTO updateinfo(uuid ,bizid) values(?,?)")

	res, err := stmt.Exec(uuid,bizid)

	id, err := res.LastInsertId()
	if id == 0 {
		err = errors.New("insertAgent table == 0")
	}
	return
}
func (d *SqliteDb) Close() {
	d.DB.Close()
}

type AppSinge struct {
	AppName string
	Ver string
}

func (d *SqliteDb) GetAppData() (err error,tmp []AppSinge) {
	d.Lock()
	defer func() {
		d.Unlock()
	}()
	rows, err := d.DB.Query("select name,version from app")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		rows.Close()
	}()
	for rows.Next() {
		var ts AppSinge
		err = rows.Scan(&ts.AppName, &ts.Ver)
		//if err!=nil{
		//
		//}
		tmp=append(tmp,ts)
		//fmt.Println(ts)
	}

	return
}
func (d *SqliteDb) AppUpdate() (err error) {
	d.Lock()
	defer func() {
		d.Unlock()
	}()
	curtime := time.Now().Unix()
	stmt, err := d.DB.Prepare("update app set monitor_cmd=?,monitor_return=?,start_cmd=?,stop_cmd=?,time=?,version=?, rundir=? where name=?")
	if err!=nil{
		logger.Printf("AppUpdate name:%s err:%s\n",rep.AppName,err)
		return err
	}

	res, err := stmt.Exec("ps aux|grep vsftp|grep -v grep|wc -l","1","start.sh","stop.sh","")
	if err!=nil{
		logger.Printf("AppUpdate name:%s err:%s\n",rep.AppName,err)
		return err
	}

	affect, err := res.RowsAffected()

	//if err!=nil{
	//	return err
	//}

	if affect == 0 {
		return errors.New("update affect = 0")
	}
	return
}
func main() {
	err := NewConn("C:\\Users\\Administrator\\Desktop\\111.sqlite")
	if err != nil {
		log.Printf("sqlite conn fail err:%s\n", err)
	}

	//sql1 := CreateSelectSql("agent", "version", "1=1")

	defer func() {
		SqlConn.Close()
	}()

	sqlconn1 := ReturnSqlDB() // 其它模块时
	//sqlconn1.InsertUuid("123",1)
	//err, b := sqlconn1.GetExist(sql1)
	sqlconn1.GetAppExists("vsftpd")
	//err,tmp:=sqlconn1.GetAppData()
	//fmt.Println(err,tmp[0])
	//fmt.Println(err,b)
	//if err != nil {
	//	log.Printf("run sql :%s faile err:%s\n", sql1, err)
	//}
	//cver := "v3" // 当前版本
	//// 没记录 直接插入，  有记录 判断是否 版本一样，一样则跳过，不一样更新
	//// ubs 发送过来 必须带有版本号
	//var uuid, ver, timestr string
	//
	//timestr = strconv.FormatInt(time.Now().Unix(), 10)
	//if b {
	//	err := SqlConn.GetData(sql1,  &ver)
	//	if err != nil {
	//		log.Println("getdata err", err)
	//	}
	//	fmt.Println(uuid, ver)
	//	if ver == cver {
	//		log.Println("version same skip")
	//		return
	//	}
	//
	//	err = SqlConn.UpdateAgentVer(uuid, ver, timestr)
	//	if err != nil {
	//		log.Printf("update err:%s\n", err)
	//	}
	//	SqlConn.Close()
	//} else {
	//	err := SqlConn.InsertAgent(ver, uuid, timestr)
	//	if err != nil {
	//		log.Printf("insert agent err:%s\n", err)
	//	}
	//}

	////更新数据
	//stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec("astaxieupdate", id)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)
	//
	////查询数据
	//rows, err := db.Query("SELECT uuid FROM agent limit 1")
	//checkErr(err)
	//
	//for rows.Next() {
	//	var uuid string
	//	err = rows.Scan(&uuid)
	//	checkErr(err)
	//	fmt.Println(uuid)
	//
	//}

	////删除数据
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec(id)
	//checkErr(err)
	//
	//affect, err = res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)
	//
	//db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
