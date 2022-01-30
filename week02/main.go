package main

import "fmt"

func main() {
	sql, err := initSql()
	defer sql.Close()
	if err != nil {
		fmt.Println("initSql error:", err)
		fmt.Printf("%+v", err, "\n")
	}

	err = query(sql, "select * from account where id=10001")
	if err != nil {
		fmt.Println("query sql error :", err)
		fmt.Printf("%+v", err, "\n")
	}
}
