package models

import (
	"fmt"
	"github.com/ivancc666/tkp-register-go/database"
	//"github.com/ivancc666/tkp-register-go/models"
)

type Enrollment struct {
	Cid 		int
	Sid 		int
	Enrolltime 	string
	Status 		int
}

type EnrollDetail struct {
	Name		string
	Class 		string
	Classno  	string
	Enrolltime 	string
}

func AddEnroll(enroll Enrollment) (int64, error) {
	i, err := insertEnroll(enroll)
	return i, err
} 

func insertEnroll(enroll Enrollment) (int64, error) {
	return database.ModifyDB("insert into enrollment(cid, sid, enrolltime, status) values($1, $2, $3, $4)", enroll.Cid, enroll.Sid, enroll.Enrolltime, enroll.Status)
}

func ShowEnrollmentWithCid(classid string) []EnrollDetail {

	sql := fmt.Sprintf("select * from enrollment where cid =%s", classid)
	rows, _ := database.QueryDB(sql)
	enrolldetails := make([]EnrollDetail, 0)
	for rows.Next() {
		enroll := Enrollment{}
		rows.Scan(&enroll.Cid, &enroll.Sid, &enroll.Enrolltime, &enroll.Status)

		enrolldetail := EnrollDetail{}
		enrolldetail.Enrolltime = enroll.Enrolltime
		fmt.Println(enroll.Enrolltime)
		student := QueryUserWithId(enroll.Sid)
		enrolldetail.Name = student.StudentID
		enrolldetail.Class = student.Class
		enrolldetail.Classno = student.Classno
		

		enrolldetails = append(enrolldetails, enrolldetail)
	}
	return enrolldetails
}
