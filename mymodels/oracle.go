package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-oci8"
	"os"
)

// todo 运行环境linux
func main() {
	os.Setenv("NLS_LANG", "")
	//if len(os.Args) != 2 {
	//	log.Fatalln(os.Args[0] + " user/password@host:port/sid")
	//}

	db, err := sql.Open("oci8", "test/test@192.168.43.22:1521/orcl")
	//fmt.Printf("%+v\n",db)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	rows, err := db.Query("select \"Client\" from BCD ")
	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	for rows.Next() {
		var data string
		rows.Scan(&data)
		fmt.Println(data)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 5000; i++ {
		db.Exec(fmt.Sprintf("insert into BCD(\"Client\",\"AgentInstance\",\"BackupSetSubclient\") values('a','b',%d)", i))
	}
}
