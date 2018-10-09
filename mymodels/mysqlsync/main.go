package main

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

type Col struct {
	id                          int
	begin_time                  string
	status, comment, alarm_type string
}

var NeedColSlice = []string{"id", "begin_time", "status", "comment", "alarm_type"}
var NeedCol = strings.Join(NeedColSlice, ",")

func mulitGet(db *sql.DB, sql string) []*Col {
	//fmt.Println(sql)
	tmp := make([]*Col, 0)
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		tmpRow := &Col{}
		err := rows.Scan(&tmpRow.id, &tmpRow.begin_time, &tmpRow.status, &tmpRow.comment, &tmpRow.alarm_type) //这里顺序要与上面 SQL中获取字段顺序一样
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(tmpRow)
		tmp = append(tmp, tmpRow)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
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
		return int64(0) //空值
	}
	t, err := strconv.ParseInt(tmp, 10, 64)
	return t
	//fmt.Println(rows.NextResultSet())
	//for rows.Next() {
	//
	//	var name string
	//	rows.Scan(&name)
	//	fmt.Println(name)
	//}
	// for rows.Next() {
	//  err := rows.Scan(&cloumns[0], &cloumns[1], &cloumns[2])
	//  if err != nil {
	//      log.Fatal(err)
	//  }
	//  fmt.Println(cloumns[0], cloumns[1], cloumns[2])
	// }
	//values := make([]sql.RawBytes, len(cloumns))
	//scanArgs := make([]interface{}, len(values))
	//for i := range values {
	//	scanArgs[i] = &values[i]
	//}
	//for rows.Next() {
	//	err = rows.Scan(scanArgs...)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	var value string
	//	for i, col := range values {
	//		if col == nil {
	//			value = "NULL"
	//		} else {
	//			value = string(col)
	//		}
	//		fmt.Println(cloumns[i], ": ", value)
	//	}
	//	fmt.Println("------------------")
	//}
	//if err = rows.Err(); err != nil {
	//	log.Fatal(err)
	//}
}

// 插入数据
func Insert(db *sql.DB, data []*Col, sql string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // danger!
	for _, d := range data {
		_, err := tx.Stmt(stmt).Exec(d.id, d.begin_time, d.status, d.comment, d.alarm_type)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	//res, err := stmt.Exec("python", 19)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//lastId, err := res.LastInsertId() // 最后插入的ID
	//if err != nil {
	//	log.Fatal(err)
	//}
	//rowCnt, err := res.RowsAffected() //本次影响的行数
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
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

func main() {
	// username: root; password: 123456; database: test
	srcdb, err := sql.Open("mysql", "root:W3b5Ev!c3@tcp(192.168.130.17:3306)/bkdata_monitor_alert?charset=utf8")
	srcdb.SetMaxIdleConns(5)
	srcdb.SetMaxOpenConns(5)

	dstdb, err := sql.Open("mysql", "root:W3b5Ev!c3@tcp(192.168.130.17:3306)/test?charset=utf8")
	dstdb.SetMaxIdleConns(5)
	dstdb.SetMaxOpenConns(5)
	if err != nil {
		log.Println("Can't connect to database")
	} else {
		err = srcdb.Ping()
		if err != nil {
			fmt.Println("db.Ping failed:", err)
		}
	}
	defer srcdb.Close()
	// Insert(db)
	// Update(db)
	// Delete(db)
	//查找 源库与目标库最新 UNIX时间
	var sql string

	sql = "SELECT UNIX_TIMESTAMP(e.begin_time) from ja_alarm_alarminstance as e ORDER BY begin_time desc limit 1;"
	srcDate := oneGet(srcdb, sql)
	dstDate := oneGet(dstdb, sql)
	fmt.Println(srcDate)
	fmt.Println(dstDate)
	if srcDate == int64(0) { //源库最新是0 ，代表被清空
		return //结束本次循环
	}
	//INSERT INTO ja_alarm_alarminstance(id,begin_time,status,comment,alarm_type) VALUES(1,'2017-03-02 15:22:22',3,4,5);
	if srcDate > dstDate {
		sql = fmt.Sprintf("select %s from ja_alarm_alarminstance as e where UNIX_TIMESTAMP(e.begin_time) > %d", NeedCol, dstDate)
		insertData := mulitGet(srcdb, sql) //查找大于目标库时间戳的数据
		sql = fmt.Sprintf("//INSERT INTO ja_alarm_alarminstance(id,begin_time,status,comment,alarm_type) VALUES(?,?,?,?,?);")
		Insert(dstdb, insertData, sql)
	}
}
