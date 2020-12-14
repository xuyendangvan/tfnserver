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
	db "git_source_release/db"
	model "git_source_release/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var noti model.Notification
	err := decoder.Decode(&noti)
	log.Println(noti)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	e := createRecordNotification(noti)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createRecordNotification(noti model.Notification) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}
	Priority := noti.Priority
	Type := noti.Type_
	Title := noti.Title
	Content := noti.Content
	PosterID := noti.PosterID
	SeenCount := noti.SeenCount
	DateExpired := noti.DateExpired
	DateCreate := time.Now()
	DateUpdate := time.Now()

	insForm, err := db.SQLExec(tx, "INSERT INTO Notification(type, priority, title, content, poster_id, seen_count,expired_date, created_date,update_date, update_count) VALUES(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	if _, err := insForm.Exec(Type, Priority, Title, Content, PosterID, SeenCount, DateExpired, DateCreate, DateUpdate, 0); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteNotificationByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	e := deleteRecordNotification(ID)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteRecordNotification(ID string) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}
	insForm, err := database.Prepare("DELETE FROM Notification WHERE id= ?")
	if _, err := insForm.Exec(ID); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func GetNotificationByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	jsonResponse := getDataNotificationFromDB(ID)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func getDataNotificationFromDB(id string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data    model.Notification
		records []model.Notification
	)
	rows, err := database.Query("SELECT * FROM Notification WHERE id= ?", id)
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

func UpdateNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	decoder := json.NewDecoder(r.Body)
	var t model.Notification
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	log.Println(t)
	e := updateRecordNotification(ID, t)
	if e != nil {
		log.Printf(e.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateRecordNotification(ID string, noti model.Notification) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		log.Println(err)
		return err
	}
	sid, err := strconv.Atoi(ID)
	Priority := noti.Priority
	Type := noti.Type_
	Title := noti.Title
	Content := noti.Content
	PosterID := noti.PosterID
	SeenCount := noti.SeenCount
	DateExpired := noti.DateExpired
	updateDate := time.Now()
	insForm, err := db.SQLExec(tx, "Update Notification Set type = ?,priority= ?,title = ?,content= ?,poster_id = ?,seen_count = ?,expired_date=?,date_update= ?, update_count = update_count + 1 where id= ?")
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := insForm.Exec(Type, Priority, Title, Content, PosterID, SeenCount, DateExpired, updateDate, sid); err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	tx.Commit()
	return nil
}
