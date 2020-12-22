/*
 * Demo App
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	db "tfnserver/db"
	model "tfnserver/model"
	"net/http"
	"time"
)

func CalculateTuitionFee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	fee := getDataStatistic()
	e := createRecordTuitionFee(fee)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getDataStatistic() []model.TuitionFee {
	database := db.DBConn()
	defer database.Close()
	countValue := 0
	var (
		data         model.StudentStatistic
		dataResult   model.TuitionFee
		recordResult []model.TuitionFee
	)
	rows, err := database.Query("SELECT id FROM Student")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		countValue++
		var id int
		rows.Scan(&id)
		var month int32 = int32(time.Now().Month())
		var year int = time.Now().Year()
		data.Quater = (month / 3) + 1
		data.Total = 6 * 4 * 3
		queryAttendance := getQueryValue(data.Quater, year, "attentdance_date")
		queryApplication := getQueryValue(data.Quater, year, "application_date")
		data.RequestedAbsences = getDataIntQuery("SELECT Count(*) FROM Attendance a WHERE a.student_id = ? and a.status = 2 and"+queryAttendance, database, id)
		data.Absences = getDataIntQuery("SELECT Count(*) FROM Attendance a WHERE a.student_id = ? and a.status = 3 and"+queryAttendance, database, id)
		data.UsedMeals = getDataIntQuery("SELECT Count(*) FROM Application a WHERE a.student_id = ? and"+queryApplication, database, id)
		data.Id = int32(countValue)
		data.StudentID = int32(id)
		data.DateCreated = time.Now().Format("2006-01-02 15:04:05")
		dataResult.Id = data.Id
		dataResult.StudentID = data.StudentID
		dataResult.Other = 1500000
		dataResult.Bus = 1000000
		dataResult.Tuition = (data.Total - data.RequestedAbsences) * 100000
		dataResult.Food = data.UsedMeals * 15000
		data.Total = dataResult.Other + dataResult.Bus + dataResult.Tuition + dataResult.Food
		dataResult.Paid = 0
		dataResult.Refund = 0
		recordResult = append(recordResult, dataResult)
	}
	defer rows.Close()
	return recordResult
}

func createRecordTuitionFee(fees []model.TuitionFee) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}
	for _, fee := range fees {
		Quater := fee.Quater
		Paid := fee.Paid
		//ClassID := fee.ClassID
		StudentID := fee.StudentID
		Total := fee.Total
		Tuition := fee.Tuition
		Food := fee.Food
		Bus := fee.Bus
		Refund := fee.Refund
		Other := fee.Other
		DateCreate := time.Now()
		DateUpdate := time.Now()
		//Phone := students.Phone
		//StudentStatus := students.StudentStatus

		insForm, err := db.SQLExec(tx, "INSERT INTO Tuition_fee(student_id,quater_id, tuition_fee, cleaning_fee, bus_fee, meal_fee, refund,total,status,date_create, date_update, update_count) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			return err
		}
		if _, err := insForm.Exec(StudentID, Quater, Tuition, Other, Bus, Food, Refund, Total, Paid, DateCreate, DateUpdate, 0); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func GetFeedbacks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
