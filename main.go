package main

import (
	"CS445FinalProject/entities"
	"encoding/json"
	"fmt"
	"time"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


var accountid, messageid, rideid, reportid, ratingid, joinrideid int
var accounts = make(map[int]entities.Account)

type shortenedAccount struct{
		Aid int `json:"aid"`
		Name string `json:"name"`
		DateCreated string `json:"date_created"`
		IsActive bool `json:"is_active"`
}

func cutAccount(aid int, a entities.Account) shortenedAccount{
	var b = shortenedAccount{
		aid,
		a.FirstName + " " + a.LastName,
		a.DateCreated,
		a.IsActive,
	}
	return b
}

func stripVars(vars map[string]string) int {
	i, err := strconv.Atoi(vars["accountid"])
	if err != nil {
        // handle error
        fmt.Println(err)
	}
	return i
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	var a entities.Account
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	accountid++
	counter := struct {
		Aid int `json:"aid"`
	}{
		Aid: accountid,
	}
	dt := time.Now()
	a.DateCreated = dt.Format("01-02-2006, 15:04:05")
	accounts[accountid] = a
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(counter)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	i := stripVars(mux.Vars(r))
	account := accounts[i]
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(account)

}
func activateAccount(w http.ResponseWriter, r *http.Request){
	i :=  stripVars(mux.Vars(r))
	account := accounts[i]
	var a entities.Account
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if a.IsActive == true {
		fmt.Println("in here")
		account.IsActive = true
		accounts[i] = account
	}
	fmt.Println(accounts)
}
func updateAccount(w http.ResponseWriter, r *http.Request){
	i :=  stripVars(mux.Vars(r))
	var a entities.Account
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	account := accounts[i]
	a, account = account, a
	accounts[i] = account
}
func deleteAccount(w http.ResponseWriter, r *http.Request){
	i :=  stripVars(mux.Vars(r))
	delete(accounts, i)
}
func getAccounts(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	if len(query) != 0 {

	} else {
		cutted := make([]shortenedAccount, 0)
		for key, value := range accounts {
			a := cutAccount(key, value)
			cutted = append(cutted, a)
		}
		w.Header().Set("Content-Type", "application-json")
		json.NewEncoder(w).Encode(cutted)
	}
	
}
func searchAccounts(w http.ResponseWriter, r *http.Request){
	fmt.Println(mux.Vars(r))
}
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/sar/accounts", createAccount).Methods("POST")
	r.HandleFunc("/sar/accounts/{accountid}", getAccount).Methods("GET")
	r.HandleFunc("/sar/accounts/{accountid}/status", activateAccount).Methods("PUT")
	r.HandleFunc("/sar/accounts/{accountid}", updateAccount).Methods("PUT")
	r.HandleFunc("/sar/accounts/{accountid}", deleteAccount).Methods("DELETE")
	r.HandleFunc("/sar/accounts", getAccounts).Methods("GET") // this route handles both searching and getting all accounts
	log.Fatal(http.ListenAndServe(":8080", r))
}

