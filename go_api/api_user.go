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
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	db "tfnserver/db"
	helper "tfnserver/helper"
	model "tfnserver/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	log.Println(user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	e := createRecordUser(user)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createRecordUser(user model.User) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}
	UserName := user.Username
	Name := user.Name
	Email := user.Email
	Password, err := helper.HashPassword(user.Password)
	//Phone := user.Phone
	Role := user.Role

	UserStatus := user.UserStatus
	DateCreate := time.Now()
	DateUpdate := time.Now()

	insForm, err := db.SQLExec(tx, "INSERT INTO User(type, name, login_name, password, email, status, date_create,date_update, update_count) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	if _, err := insForm.Exec(Role, Name, UserName, Password, Email, UserStatus, DateCreate, DateUpdate, 0); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func CreateUsersWithListInput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var user []model.User
	err := decoder.Decode(&user)
	log.Println(user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	e := createRecordUsers(user)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createRecordUsers(list []model.User) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}
	for _, user := range list {
		ID := user.Id
		UserName := user.Username
		Name := user.Name
		Email := user.Email
		Password, err := helper.HashPassword(user.Password)
		//Phone := user.Phone
		Role := user.Role

		UserStatus := user.UserStatus
		DateCreate := time.Now()
		DateUpdate := time.Now()

		insForm, err := db.SQLExec(tx, "INSERT INTO User(id, type, name, login_name, password, email, status, date_create,date_update, update_count) VALUES(?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			return err
		}
		if _, err := insForm.Exec(ID, Role, Name, UserName, Password, Email, UserStatus, DateCreate, DateUpdate, 0); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	e := deleteRecordUser(ID)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteRecordUser(ID string) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}
	insForm, err := database.Prepare("DELETE FROM User WHERE id= ?")
	if _, err := insForm.Exec(ID); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	jsonResponse := getDataUserFromDB(ID)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}
func getDataUserFromDB(id string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		data model.User
	)
	rows, err := database.Query("SELECT * FROM User WHERE id= ?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count int
		rows.Scan(&data.Id, &data.Role, &data.Name, &data.Username, &data.Password, &data.Email, &data.UserStatus, &datecreate, &date, &count)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(data)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	usernames, value := r.URL.Query()["username"]

	if !value || len(usernames[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	username := usernames[0]
	fmt.Printf(username)
	log.Println("Url Param 'key' is: " + string(username))
	passwords, values := r.URL.Query()["password"]

	if !values || len(passwords[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	password := passwords[0]

	log.Println("Url Param 'key' is: " + string(password))
	jsonResponse := login(username, password)
	if jsonResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

//login user
func login(username string, password string) []byte {
	database := db.DBConn()
	defer database.Close()
	var (
		user model.User
		data model.LoginResponse
		hash string
	)
	rows, err := database.Query("SELECT * FROM User WHERE login_name= ?", username)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//add token in model user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	for rows.Next() {
		var date, datecreate time.Time
		var count, id int32
		rows.Scan(&user.Id, &user.Role, &user.Name, &user.Username, &user.Password, &user.Email, &user.UserStatus, &datecreate, &date, &count)
		user.Token = tokenString
		if user.Role == 1 {
			id = getDataIntQuery("SELECT id FROM Parent WHERE user_id= ?", database, int(user.Id))
		} else if user.Role == 2 {
			id = getDataIntQuery("SELECT id FROM Teacher WHERE user_id= ?", database, int(user.Id))
		} else {
			id = int32(user.Id)
		}
		data.Id = id
		data.Role = user.Role
		data.Token = user.Token
		data.Message = user.Name
		hash = user.Password
	}
	defer rows.Close()

	if !helper.CheckPasswordHash(password, hash) {
		return nil
	}

	jsonResponse, jsonError := json.Marshal(data)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")

	defer r.Body.Close()

	ids, _ := r.URL.Query()["id"]

	ID := ids[0]

	types, _ := r.URL.Query()["type"]
	Type := types[0]

	log.Printf(ID)
	e := logUserOut(ID, Type)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ID is parent/teacherID, Type is the type of user
func logUserOut(ID string, Type string) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		return err
	}

	temp, _ := strconv.Atoi(Type)
	query := ""

	if temp == 1 {
		query = "DELETE FROM Device_token WHERE parent_id= ?"
	} else {
		query = "DELETE FROM Device_token WHERE teacher_id= ?"
	}

	insForm, err := database.Prepare(query)
	if _, err := insForm.Exec(ID); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	decoder := json.NewDecoder(r.Body)
	var t model.User
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	log.Println(t)
	e := updateRecordUser(ID, t)
	if e != nil {
		log.Printf(e.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateRecordUser(ID string, t model.User) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		log.Println(err)
		return err
	}
	sid, err := strconv.Atoi(ID)
	UserName := t.Username
	Name := t.Name
	Email := t.Email
	Password, err := helper.HashPassword(t.Password)
	//Phone := user.Phone
	Role := t.Role

	UserStatus := t.UserStatus
	updateDate := time.Now()
	insForm, err := db.SQLExec(tx, "Update User Set name= ?,type = ?,login_name= ?,password = ?, email= ? ,status= ?,date_update= ?, update_count = update_count + 1 where id= ?")
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := insForm.Exec(Name, Role, UserName, Password, Email, UserStatus, updateDate, sid); err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	tx.Commit()
	return nil
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Connection", "close")
	r.Header.Set("Connection", "close")
	defer r.Body.Close()
	ID := mux.Vars(r)["id"]
	log.Printf(ID)
	decoder := json.NewDecoder(r.Body)
	var data model.PasswordParam
	err := decoder.Decode(&data)
	log.Println(data)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	e := updatePasswordRecordUser(ID, data)
	if e != nil {
		log.Printf(e.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updatePasswordRecordUser(ID string, data model.PasswordParam) (err error) {
	database := db.DBConn()
	defer database.Close()
	tx, err := db.SQLBegin(database)
	if err != nil {
		log.Println(err)
		return err
	}
	var userID int
	var password string
	sid, err := strconv.Atoi(ID)
	dataValue, err := helper.HashPassword(data.NewPassword)
	if data.Type == 1 {
		userID, err = strconv.Atoi(getDataQuery("SELECT user_id FROM Parent WHERE id= ?", database, sid))
	} else {
		userID, err = strconv.Atoi(getDataQuery("SELECT user_id FROM Teacher WHERE id= ?", database, sid))
	}
	password = getDataQuery("SELECT password FROM User WHERE id= ?", database, userID)
	if !helper.CheckPasswordHash(data.OldPassword, password) {
		return nil
	}
	insForm, err := db.SQLExec(tx, "Update User Set password= ? where id= ?")
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := insForm.Exec(dataValue, userID); err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	tx.Commit()
	return nil
}

func getDataQuery(query string, database *sql.DB, sid int) string {
	var result string
	rows, err := database.Query(query, sid)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for rows.Next() {
		rows.Scan(&result)
	}
	defer rows.Close()
	return result
}
