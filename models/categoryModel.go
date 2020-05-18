package models

import (
	"fmt"
	"github.com/ivancc666/tkp-register-go/database"
)

type Category struct {
	Id         int32
	Name		string
}

func ShowAllCategory() []Category {
	sql := "select * from category"
	fmt.Println(sql)
	rows, _ := database.QueryDB(sql)
	cats := make([]Category, 0)
	for rows.Next() {
		cat := Category{}
		rows.Scan(&cat.Id, &cat.Name)
    	cats = append(cats, cat)
	}
	return cats
}