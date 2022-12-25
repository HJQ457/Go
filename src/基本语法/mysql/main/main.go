package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Tel  int    `json:"tel"`
}

func main() {

	db_conn, db_conn_err := sqlx.Open("mysql", "oper_dc:Yonyou*123@tcp(10.16.1.64:3306)/hjq")
	if db_conn_err != nil {
		fmt.Println(db_conn_err)
		return
	}

	//insert操作
	insert, insert_err := db_conn.Exec("insert into test(id, name, tel)values(?, ?, ?)", "5", "sam", "110")
	if insert_err != nil {
		fmt.Println("insert失败：", insert_err)
		return
	}

	insert_id, insert_id_err := insert.LastInsertId()
	if insert_id_err != nil {
		fmt.Println("insert_id_err失败：", insert_id_err)
		return
	}

	insert_row, insert_row_err := insert.RowsAffected()
	if insert_row_err != nil {
		fmt.Println("insert_row_err失败：", insert_id_err)
	}

	fmt.Println("插入成功，id为：", insert_id)
	fmt.Println("受影响的行数为：", insert_row)

	//select操作
	var select_res []Test
	select_err := db_conn.Select(&select_res, "select * from test where id = ?", 1)
	if select_err != nil {
		fmt.Println("select 失败：", select_err)
		return
	}
	fmt.Println(select_res)

	//update操作
	update, update_err := db_conn.Exec("update test set tel = ? where id = ?", 120, 5)
	if update_err != nil {
		fmt.Println("update 失败:", update_err)
		return
	}
	update_row, update_row_err := update.RowsAffected()
	if update_row_err != nil {
		fmt.Println("update 失败：", update_row_err)
		return
	}
	fmt.Println(update_row)

	//delete操作
	delete, delete_err := db_conn.Exec("delete from test where id = ?", 5)
	if delete_err != nil {
		fmt.Println("delete 失败", delete_err)
		return
	}
	delete_row, delete_row_err := delete.RowsAffected()
	if delete_row_err != nil {
		fmt.Println("delete_row失败", delete_row)
		return
	}
	fmt.Println(delete_row)
}
