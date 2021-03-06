package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"sync"
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
	err=db1.Ping()
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

func (d *SqliteDb) Close() {
	d.DB.Close()
}

func main() {
	err := NewConn("C:\\work\\go-dev\\src\\godev\\mymodels\\sqlite\\info.sqlite")
	if err != nil {
		log.Printf("sqlite conn fail err:%s\n", err)
	}

	sql1 := CreateSelectSql("agent", "version", "1=1")

	defer func() {
		SqlConn.Close()
	}()

	sqlconn1 := ReturnSqlDB() // ???????????????

	err, b := sqlconn1.GetExist(sql1)
	fmt.Println(err,b)
	if err != nil {
		log.Printf("run sql :%s faile err:%s\n", sql1, err)
	}
	cver := "v3" // ????????????
	// ????????? ???????????????  ????????? ???????????? ????????????????????????????????????????????????
	// ubs ???????????? ?????????????????????
	var uuid, ver, timestr string

	timestr = strconv.FormatInt(time.Now().Unix(), 10)
	if b {
		err := SqlConn.GetData(sql1,  &ver)
		if err != nil {
			log.Println("getdata err", err)
		}
		fmt.Println(uuid, ver)
		if ver == cver {
			log.Println("version same skip")
			return
		}

		err = SqlConn.UpdateAgentVer(uuid, ver, timestr)
		if err != nil {
			log.Printf("update err:%s\n", err)
		}
		SqlConn.Close()
	} else {
		err := SqlConn.InsertAgent(ver, uuid, timestr)
		if err != nil {
			log.Printf("insert agent err:%s\n", err)
		}
	}

	////????????????
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
	////????????????
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

	////????????????
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
