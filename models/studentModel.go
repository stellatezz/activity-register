package models

import (
	"fmt"
	"github.com/ivancc666/tkp-register-go/database"
)

type Student struct {
	Sid         	int
	StudentID		string
	Email   		string
	Class   		string
	Classno 		string
}

func QueryUserWithCon(con string) int {
	sql := fmt.Sprintf("select sid from students %s", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

func QueryUserWithEmail(email string) int {
	sql := fmt.Sprintf("where email='%s'", email)
	return QueryUserWithCon(sql)
}

func QueryUserWithId(sid int) Student {
	sql := fmt.Sprintf("select * from students where sid=%d", sid)
	row := database.QueryRowDB(sql)
	var student Student
	row.Scan(&student.Sid, &student.StudentID, &student.Email, &student.Class, &student.Classno)
	return student
}