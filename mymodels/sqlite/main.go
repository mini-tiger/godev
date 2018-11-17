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
	return
}

func CreateSelectSql(table string, col, where string) (sql string) {
	sql = fmt.Sprintf("select %s from %s where %s", col, table, where)
	return
}

func (d *SqliteDb) GetData(sql string, uuid, ver *string) error {
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
		err = rows.Scan(uuid, ver)
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
	defer func() {
		rows.Close()
	}()
	if err != nil {
		rows.Close()
		return err, false
	}

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
	err := NewConn("D:\\work\\project-dev\\src\\godev\\mymodels\\sqlite\\111.sqlite")
	if err != nil {
		log.Printf("sqlite conn fail err:%s\n", err)
	}

	sql1 := CreateSelectSql("agent", "uuid,version", "1=1")

	defer func() {
		SqlConn.Close()
	}()

	sqlconn1 := ReturnSqlDB() // 其它模块时
	err, b := sqlconn1.GetExist(sql1)
	if err != nil {
		log.Printf("run sql :%s faile err:%s\n", sql1, err)
	}
	var uuid, ver, timestr string
	uuid = "11111111"
	ver = "22222222222"
	timestr = strconv.FormatInt(time.Now().Unix(), 10)
	if b {
		err := SqlConn.GetData(sql1, &uuid, &ver)
		if err != nil {
			log.Println("getdata err", err)
		}
		fmt.Println(uuid, ver)

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
