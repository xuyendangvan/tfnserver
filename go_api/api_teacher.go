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
	"encoding/json"
	"fmt"
	db "tfnserver/db"
	model "tfnserver/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func AddTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var teacher []model.Teacher
	err := decoder.Decode(&teacher)
	log.Println(teacher)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	e := createRecordTeachers(teacher)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createRecordTeachers(list []model.Teacher) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}
	for _, item := range list {
		UserID := item.UserID
		Name := item.Name
		Email := item.Email
		LoginName := item.LoginName
		Address := item.Address
		AddressCity := item.AddressCity
		AddressDistrict := item.AddressDistrict
		AddressStreet := item.AddressStreet
		Phone := item.Phone
		Status := item.Status
		DateCreated := time.Now()
		DateUpdate := time.Now()

		insForm, err := db.SQLExec(tx, "INSERT INTO Teacher(user_id, name, login_name, password, email, tel, address_city,address_district, address_ward,address_street,address,status,date_create,date_update, update_count) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			return err
		}
		if _, err := insForm.Exec(UserID, Name, LoginName, "", Email, Phone, AddressCity, AddressDistrict, "", AddressStreet, Address, Status, DateCreated, DateUpdate, 0); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func DeleteTeacherByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	e := deleteRecordTeacher(ID)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteRecordTeacher(ID string) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}
	insForm, err := database.Prepare("DELETE FROM Teacher WHERE id= ?")
	if _, err := insForm.Exec(ID); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func FindClassByTeacherID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	jsonResponse := getDataClassByTeacherFromDB(ID)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func getDataClassByTeacherFromDB(id string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data    model.Class
		records []model.Class
	)
	rows, err := database.Query("SELECT c.id,c.school_year_id,c.level,c.teacher_id,c.name,c.date_create,c.date_update,c.update_count FROM Teacher t inner join Class c ON t.id = c.teacher_id WHERE t.id= ?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count int
		rows.Scan(&data.Id, &data.YearID, &data.Level, &data.TeacherID, &data.Name, &datecreate, &date, &count)
		records = append(records, data)
	}
	defer rows.Close()
	if records == nil {
		return nil
	}
	jsonResponse, jsonError := json.Marshal(records)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func FindNotificationByTeacherID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	jsonResponse := getDataNotificationByTeacherFromDB(ID)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func getDataNotificationByTeacherFromDB(id string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data    model.Notification
		records []model.Notification
	)
	rows, err := database.Query("SELECT * FROM Notification n WHERE n.type = 2 and DATEDIFF(n.expired_date, CURDATE()) >= 0")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count int
		rows.Scan(&data.Id, &data.Type_, &data.Priority, &data.Title, &data.Content, &data.PosterID, &data.SeenCount, &datecreate, &date, &count)
		records = append(records, data)
	}
	defer rows.Close()
	if records == nil {
		return nil
	}
	jsonResponse, jsonError := json.Marshal(records)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}
func FindTeacherByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	jsonResponse := getDataTeacherFromDB(ID)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func getDataTeacherFromDB(id string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data    model.Teacher
		records []model.Teacher
	)
	rows, err := database.Query("SELECT * FROM Teacher WHERE id= ?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count int
		var password, addressward string
		rows.Scan(&data.Id, &data.UserID, &data.Name, &data.LoginName, &password, &data.Email, &data.Phone, &data.AddressCity, &data.AddressDistrict, &addressward, &data.AddressStreet, &data.Address, &data.Status, &datecreate, &date, &count)
		//user.UpdateDate = h.ConvertDateToString(date, time.RFC3339)
		data.DateCreated = datecreate.Format(time.RFC3339)
		records = append(records, data)
	}
	defer rows.Close()
	if records == nil {
		return nil
	}
	jsonResponse, jsonError := json.Marshal(records)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func FindTeacherListByIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	startIndexs, value := r.URL.Query()["startIndex"]

	if !value || len(startIndexs[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	startIndex := startIndexs[0]

	offsets, value := r.URL.Query()["keyword"]

	if !value || len(offsets[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	offset := offsets[0]

	jsonResponse := getDataTeacherFromDBWithIndex(startIndex, offset)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}
func getDataTeacherFromDBWithIndex(startIndex string, offset string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data    model.Teacher
		records []model.Teacher
	)
	limitValue, err := strconv.Atoi(offset)
	startValue, err := strconv.Atoi(startIndex)
	rows, err := database.Query("SELECT * FROM Teacher limit ? offset ?", limitValue, startValue)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count int
		var password, addressward string
		rows.Scan(&data.Id, &data.UserID, &data.Name, &data.LoginName, &password, &data.Email, &data.Phone, &data.AddressCity, &data.AddressDistrict, &addressward, &data.AddressStreet, &data.Address, &data.Status, &datecreate, &date, &count)
		//user.UpdateDate = h.ConvertDateToString(date, time.RFC3339)
		data.DateCreated = datecreate.Format(time.RFC3339)
		records = append(records, data)
	}
	defer rows.Close()
	if records == nil {
		return nil
	}
	jsonResponse, jsonError := json.Marshal(records)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	decoder := json.NewDecoder(r.Body)
	var t model.Teacher
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	log.Println(t)
	e := updateRecordTeacher(ID, t)
	if e != nil {
		log.Printf(e.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateRecordTeacher(ID string, item model.Teacher) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		log.Println(err)
		return err
	}
	sid, err := strconv.Atoi(ID)
	Name := item.Name
	Email := item.Email
	LoginName := item.LoginName
	Address := item.Address
	AddressCity := item.AddressCity
	AddressDistrict := item.AddressDistrict
	AddressStreet := item.AddressStreet
	Phone := item.Phone
	Status := item.Status
	DateUpdate := time.Now()
	insForm, err := db.SQLExec(tx, "Update Teacher Set name= ?,login_name= ?, email= ? ,tel= ?,address_city= ?,address_district = ?, address_street = ?,address_ward=?,address=?,status = ?, date_update= ?, update_count = update_count + 1 where id= ?")
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := insForm.Exec(Name, LoginName, Email, Phone, AddressCity, AddressDistrict, AddressStreet, "ward", Address, Status, DateUpdate, sid); err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	tx.Commit()
	return nil
}

func FindActivityByTeacherID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	jsonResponse := getDataActivityByTeacherFromDB(ID)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func getDataActivityByTeacherFromDB(id string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data    model.Activity
		records []model.Activity
	)
	rows, err := database.Query("SELECT a.id,a.class_id,a.date_occur,a.date_expire,a.poster_id,a.title,a.content,a.photo1,a.caption1,a.photo2,a.caption2,a.photo3,a.caption3,a.photo4,a.caption4,a.photo5,a.caption5,a.date_create,a.date_update,a.update_count FROM Activity a inner join Teacher t ON a.poster_id = t.id WHERE t.id= ?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count int
		rows.Scan(&data.Id, &data.ClassID, &data.DateOccur, &data.DateExpired, &data.Title, &data.Content, &data.Photo1, &data.Caption1, &data.Photo2, &data.Caption2, &data.Photo3, &data.Caption3, &data.Photo4, &data.Caption4, &data.Photo5, &data.Caption5, &datecreate, &date, &count)
		records = append(records, data)
	}
	defer rows.Close()
	if records == nil {
		return nil
	}
	jsonResponse, jsonError := json.Marshal(records)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}
