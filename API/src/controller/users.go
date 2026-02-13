package controller

import (
	"api/src/answers"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repository.NewUserRepository(db)

	lastID, err := repositorio.Create(user)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = uint64(lastID)
	answers.JSON(w, http.StatusCreated, user)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := db.Connect()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	user, err := repository.Search(nameOrNick)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	answers.JSON(w, http.StatusOK, users)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["UserId"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.Connect()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	user, err := repository.SearchById(userId)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	answers.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user."))
}
