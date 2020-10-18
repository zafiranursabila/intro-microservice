package handler

import (
	"net/http"

	"github.com/zafiranursabila/intro-microservice/utils"
)

func ValidateAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	if authToken != "asdfghjk" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	utils.WrapAPISuccess(w, r, "success", 200)
}

func (db *AuthDB) SignUp (w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		utils.WrapAPIError(w,r,http.StatusText(http.StatusMethodNotAllowed),http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	deper r.Body.Close()
	if err != nill {
		utils.WrapAPIError(w, r, "can't read body", http.StatusBadRequest)
		return
	}

	var signup database.Auth

	err = json.Unmarshal(body, &signup)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal : "+err.Error(), http.StatusBadRequest)
		return
	}
	
	signup.Token = utils.IdGenerator()
	err = signup.SignUp(db.Db); if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	signup.Token = utils.IdGenerator()

	err = signup.SignUp(db.DB); if err != nill {
		utils.WrapAPISuccess(w,r,Error(),http.StatusBadRequest)
		return
	}

	utils.WrapAPISuccess(w,r,"success",http.StatusOK)
	return
}

func (db *AuthDB) Login (w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		utils.WrapAPIError(w,r,http.StatusText(http.StatusMethodNotAllowed),http.StatusMethodNotAllowed)
		return
	}
	///1000 buat Login
}


