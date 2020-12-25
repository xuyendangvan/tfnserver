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
	"log"
	"net/http"
	"strconv"
	db "tfnserver/db"
	model "tfnserver/model"
	"time"

	"github.com/gorilla/mux"
)

func AddParent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var parent model.Parent
	err := decoder.Decode(&parent)
	log.Println(parent)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	e := createRecordParent(parent)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createRecordParent(parent model.Parent) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}
	UserID := parent.UserID
	Name := parent.Name
	Email := parent.Email
	LoginName := parent.LoginName
	Address := parent.Address
	AddressCity := parent.AddressCity
	AddressDistrict := parent.AddressDistrict
	AddressStreet := parent.AddressStreet
	Phone := parent.Phone
	//Photo := parent.Photo
	Status := parent.Status
	DateCreated := time.Now()
	DateUpdate := time.Now()
	//Phone := students.Phone
	//StudentStatus := students.StudentStatus

	insForm, err := db.SQLExec(tx, "INSERT INTO Parent(user_id, name, login_name, password, email, tel, address_city,address_district, address_ward,address_street,address,status,date_create,date_update, update_count) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	if _, err := insForm.Exec(UserID, Name, LoginName, "", Email, Phone, AddressCity, AddressDistrict, "", AddressStreet, Address, Status, DateCreated, DateUpdate, 0); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func UpdateParent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	decoder := json.NewDecoder(r.Body)
	var t model.Parent
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	log.Println(t)
	e := updateRecordParent(ID, t)
	if e != nil {
		log.Printf(e.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateRecordParent(ID string, item model.Parent) (err error) {
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
	//Photo := parent.Photo
	Status := item.Status
	DateUpdate := time.Now()
	insForm, err := db.SQLExec(tx, "Update Parent Set name= ?,login_name= ?, email= ? ,tel= ?,address_city= ?,address_district = ?, address_street = ?,address_ward=?,address=?,status = ?, date_update= ?, update_count = update_count + 1 where id= ?")
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

func DeleteParent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	e := deleteRecordParent(ID)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteRecordParent(ID string) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}
	insForm, err := database.Prepare("DELETE FROM Parent WHERE id= ?")
	if _, err := insForm.Exec(ID); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func FindStudentByParentID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	jsonResponse := getDataStudentByParentFromDB(ID)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func getDataStudentByParentFromDB(id string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data    model.Student
		records []model.Student
	)
	rows, err := database.Query("SELECT * FROM Student s WHERE s.parent_id= ?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count int
		var face_photo string
		rows.Scan(&data.Id, &data.ParentID, &data.ClassID, &data.Name, &data.Birthday, &face_photo, &datecreate, &date, &count)
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

func GetParentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	jsonResponse := getDataParentFromDB(ID)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func getDataParentFromDB(id string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data model.Parent
	)
	rows, err := database.Query("SELECT * FROM Parent WHERE id= ?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count int
		var password, addressward string
		rows.Scan(&data.Id, &data.UserID, &data.Name, &data.LoginName, &password, &data.Email, &data.Phone, &data.AddressCity, &data.AddressDistrict, &addressward, &data.AddressStreet, &data.Address, &data.Status, &datecreate, &date, &count)
		data.DateCreated = datecreate.Format(time.RFC3339)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(data)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func GetParentNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	jsonResponse := getDataNotificationByParentFromDB(ID)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func getDataNotificationByParentFromDB(id string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data    model.NotificationData
		records []model.NotificationData
	)
	rows, err := database.Query("SELECT * FROM Notification n WHERE n.type = 1 and DATEDIFF(n.expired_date, CURDATE()) >= 0")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count int
		rows.Scan(&data.Id, &data.Type_, &data.Priority, &data.ClassID, &data.Category, &data.Title, &data.Content, &data.PosterID, &data.SeenCount, &data.DateExpired, &datecreate, &date, &count)
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

func GetFormsByParentID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	jsonResponse := getDataFormByParentFromDB(ID)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func getDataFormByParentFromDB(id string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data    model.Form
		records []model.Form
	)
	rows, err := database.Query("SELECT a.id,a.repeat_id,a.student_id,a.application_date,a.application_time,a.type,a.note,a.meal_absent,a.late_meal,a.picker_name,a.picker_face_photo,a.direction,a.approved,a.approver,a.date_create,a.date_update,a.update_count FROM Student s inner join Application a ON s.id = a.student_id WHERE s.parent_id= ?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count int
		var picker_name, picker_face_photo, direction, approved, approver string
		rows.Scan(&data.Id, &data.PosterID, &data.StudentID, &data.DateRequest, &data.TimeRequest, &data.Type_, &data.Content, &data.CancelMeal, &data.LateMeal, &picker_name, &picker_face_photo, &direction, &approved, &approver, &datecreate, &date, &count)
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

func UpdateParenNotificationStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	parentIDs, value := r.URL.Query()["parentID"]

	if !value || len(parentIDs[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	parentID := parentIDs[0]

	notificationIDs, value := r.URL.Query()["notificationID"]

	if !value || len(notificationIDs[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	notificationID := notificationIDs[0]

	e := updateNotificationStatus(parentID, notificationID)
	if e != nil {
		log.Printf(e.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateNotificationStatus(parentID string, notificationID string) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		log.Println(err)
		return err
	}
	sid, err := strconv.Atoi(parentID)
	nid, err := strconv.Atoi(notificationID)
	DateUpdate := time.Now()
	insForm, err := db.SQLExec(tx, "INSERT INTO Parent_Notification(parent_id, notification_id, class_id, update_date, update_count) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	if _, err := insForm.Exec(sid, nid, nid, DateUpdate, 0); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
