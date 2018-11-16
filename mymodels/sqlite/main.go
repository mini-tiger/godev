package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "/home/go/src/godev/mymodels/sqlite/111.sqlite")

	//插入数据
	stmt, err := db.Prepare("INSERT INTO agent(version ,uuid ,time) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("v1", "123-33dfd-vc-dfdf", "123435355")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
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
	rows, err := db.Query("SELECT uuid FROM agent limit 1")
	checkErr(err)

	for rows.Next() {
		var uuid string
		err = rows.Scan(&uuid)
		checkErr(err)
		fmt.Println(uuid)

	}

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

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
