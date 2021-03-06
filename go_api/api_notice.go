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

func AddNotice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var notice model.Notice
	err := decoder.Decode(&notice)
	log.Println(notice)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	e := createRecordNotice(notice)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createRecordNotice(notice model.Notice) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}

	// check data
	Type := notice.Type_
	id := int32(0)
	query := ""

	if Type == 1 {
		query = "INSERT INTO Notice(severity,type,class_id,student_id,parent_id,date_occur,date_expire,title,content,date_create, date_update,update_count) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)"
		id = notice.ParentID
	} else {
		query = "INSERT INTO Notice(severity,type,class_id,student_id,teacher_id,date_occur,date_expire,title,content, date_create, date_update,update_count) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)"
		id = notice.TeacherID
	}
	for _, contentValue := range notice.Content {

		insForm, err := db.SQLExec(tx, query)
		if err != nil {
			return err
		}
		_, er := insForm.Exec(1, Type, notice.ClassID, notice.StudentID, id, time.Now(), time.Now(), "title", contentValue, time.Now(), time.Now(), 0)
		if er != nil {
			tx.Rollback()
			log.Println(er.Error())
			return er
		}
	}

	tx.Commit()
	return nil
}

func FindNoticeListByIndex(w http.ResponseWriter, r *http.Request) {
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

	offsets, value := r.URL.Query()["offset"]

	if !value || len(offsets[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	offset := offsets[0]

	jsonResponse := getDataNoticeFromDBWithIndex(startIndex, offset)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func getDataNoticeFromDBWithIndex(startIndex string, offset string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data    model.Notice
		records []model.Notice
	)
	limitValue, err := strconv.Atoi(offset)
	startValue, err := strconv.Atoi(startIndex)
	rows, err := database.Query("SELECT * FROM Notice limit ? offset ?", limitValue, startValue)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		data.Content = nil
		var date, datecreate time.Time
		var severity, title, file1, caption1, file2, caption2, file3, caption3 string
		var count int
		var content string
		rows.Scan(&data.Id, &severity, &data.Type_, &data.StudentID, &data.ParentID, &data.DateOccur, &data.DateExpired, &data.TeacherID, &title, &content, &data.ConfirmMessage, &file1, &caption1, &file2, &caption2, &file3, &caption3, &datecreate, &date, &count)
		data.Content = append(data.Content, content)
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

func UpdateNotice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	decoder := json.NewDecoder(r.Body)
	var t model.Notice
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	log.Println(t)
	e := updateRecordNotice(ID, t)
	if e != nil {
		log.Printf(e.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateRecordNotice(ID string, notice model.Notice) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		log.Println(err)
		return err
	}
	sid, err := strconv.Atoi(ID)
	Type := notice.Type_
	DateOccur := notice.DateOccur
	Content := notice.Content
	ConfirmMessage := notice.ConfirmMessage
	StudentID := notice.StudentID
	ParentID := notice.ParentID
	TeacherID := notice.TeacherID
	DateExpired := notice.DateExpired
	updateDate := time.Now()
	for _, contentValue := range Content {
		insForm, err := db.SQLExec(tx, "Update Notice Set type= ?,date_occur = ?,date_expired=?,content= ?,confirm_message = ?,class_id=?,parent_id=?,poster_id=?,date_update= ?, update_count = update_count + 1 where id= ?")
		if err != nil {
			log.Println(err)
			return err
		}
		if _, err := insForm.Exec(Type, DateOccur, DateExpired, contentValue, ConfirmMessage, StudentID, ParentID, TeacherID, updateDate, sid); err != nil {
			tx.Rollback()
			log.Println(err)
			return err
		}
	}
	tx.Commit()
	return nil
}
