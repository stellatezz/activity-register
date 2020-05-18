package models

import (
	"fmt"
	"github.com/ivancc666/tkp-register-go/database"
)

type Class struct {
	Cid         int
	Name		string
	Date   		string
	Time   		string
	Venue 		string
	Tutor		string
	Quota 		int
	Fee 		string
	Description string
	Category 	int
}

func ShowAllClass() []Class {
	sql := "select * from classes"
	fmt.Println(sql)
	rows, _ := database.QueryDB(sql)
	classes := make([]Class, 0)
	for rows.Next() {
		class := Class{}
		rows.Scan(&class.Cid, &class.Name, &class.Date, &class.Time, &class.Venue, &class.Tutor, &class.Quota, &class.Fee, &class.Description, &class.Category)
    	classes = append(classes, class)
	}
	return classes
}

func QueryClassWithID(classid string) Class {
	sql := fmt.Sprintf("select * from classes where cid=%s", classid)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	var class Class
	row.Scan(&class.Cid, &class.Name, &class.Date, &class.Time, &class.Venue, &class.Tutor, &class.Quota, &class.Fee, &class.Description, &class.Category)
	fmt.Println("classid = ", class.Cid)
	return class
}

func AddClass(class Class) (int64, error) {
	i, err := insertClass(class)
	return i, err
} 

func insertClass(class Class) (int64, error) {
	return database.ModifyDB("insert into classes(name,date,time,venue,tutor,quota,fee,description,category) values($1,$2,$3,$4,$5,$6,$7,$8,$9)",
	 class.Name, class.Date, class.Time, class.Venue, class.Tutor, class.Quota, class.Fee, class.Description, class.Category)
}