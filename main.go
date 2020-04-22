package main

import (
	"CS445FinalProject/entities"
	"encoding/json"
	"fmt"
	"time"
	"reflect"
	"strings"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


var accountid, messageid, rideid, reportid, ratingid, joinrideid int
var accounts = make(map[int]entities.Account)
var rides = make(map[int]entities.Ride)

type shortenedAccount struct{
		Aid int `json:"aid"`
		Name string `json:"name"`
		DateCreated string `json:"date_created"`
		IsActive bool `json:"is_active"`
}
type shortenedRide struct {
	Rid int `json:"rid"`
	LocationInfo struct {
		FromCity string `json:"from_city"`
		FromZip  string `json:"from_zip"`
		ToCity   string `json:"to_city"`
		ToZip    string `json:"to_zip"`
	} `json:"location_info"`
	DateTime struct {
		Date string `json:"date"`
		Time string `json:"time"`
	} `json:"date_time"`
}
func cutRide(rid int, a entities.Ride) shortenedRide{
	var b = shortenedRide{
		rid,
		a.LocationInfo,
		a.DateTime,
	}
	return b
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

func stripVars(vars map[string]string, entity string) int {
	i, err := strconv.Atoi(vars[entity +"id"])
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
	i := stripVars(mux.Vars(r), "account")
	account := accounts[i]
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(account)

}
func activateAccount(w http.ResponseWriter, r *http.Request){
	i :=  stripVars(mux.Vars(r), "account")
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
	i :=  stripVars(mux.Vars(r), "account")
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
	i :=  stripVars(mux.Vars(r), "account")
	delete(accounts, i)
}
func getAccounts(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	if (len(query) != 0) {
	q := searchAccounts(query["key"][0])
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(q)
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
func searchAccounts(key string) []shortenedAccount{
	
	matching := make([]shortenedAccount, 0)
	for idx, v := range accounts {
		flag := false
		s := reflect.ValueOf(&v).Elem()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			if f.Type().String() == "string" {
				k := strings.ToLower(key)
				val := strings.ToLower(fmt.Sprintf("%s", f.Interface()))
				if(strings.Contains(val, k) == true){
					flag = true
				}
			}
		}
		if(flag){
			a := cutAccount(idx, v)
			matching = append(matching, a)
		}
		
	}
	return matching
	
}
func rateAccount(w http.ResponseWriter, r *http.Request){}
func getDriverRatings(w http.ResponseWriter, r *http.Request){}
func getRiderRatings(w http.ResponseWriter, r *http.Request) {}
func postRide(w http.ResponseWriter, r *http.Request) {
	var a entities.Ride
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rideid++
	counter := struct {
		Rid int `json:"rid"`
	}{
		Rid: rideid,
	}
	rides[rideid] = a
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(counter)

}
func updateRide(w http.ResponseWriter, r *http.Request){
	i :=  stripVars(mux.Vars(r), "ride")
	var a entities.Ride
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ride := rides[i]
	a, ride = ride, a
	rides[i] = ride
}
func deleteRide(w http.ResponseWriter, r *http.Request) {
	i :=  stripVars(mux.Vars(r), "ride")
	delete(rides, i)
}
func getRides(w http.ResponseWriter, r *http.Request) {

		cutted := make([]shortenedRide, 0)
		for key, value := range rides {
			a := cutRide(key, value)
			cutted = append(cutted, a)
		}
		w.Header().Set("Content-Type", "application-json")
		json.NewEncoder(w).Encode(cutted)
}
func getRide(w http.ResponseWriter, r *http.Request) {
	i :=  stripVars(mux.Vars(r), "ride")
	ride := rides[i]
	a := entities.CreateRideDetail(ride, accounts[ride.Aid])

}
func searchRides(keys map[string]string){}
func joinRide(w http.ResponseWriter, r *http.Request) {}
func confirmJoinRide(w http.ResponseWriter, r *http.Request) {}
func addMessage(w http.ResponseWriter, r *http.Request) {}
func getMessages(w http.ResponseWriter, r *http.Request) {}
func main() {

	r := mux.NewRouter()

	// Account related endpoints
	r.HandleFunc("/sar/accounts", createAccount).Methods("POST")
	r.HandleFunc("/sar/accounts/{accountid}", getAccount).Methods("GET")
	r.HandleFunc("/sar/accounts/{accountid}/status", activateAccount).Methods("PUT")
	r.HandleFunc("/sar/accounts/{accountid}", updateAccount).Methods("PUT")
	r.HandleFunc("/sar/accounts/{accountid}", deleteAccount).Methods("DELETE")
	r.HandleFunc("/sar/accounts", getAccounts).Methods("GET") // this route handles both searching and getting all accounts
	r.HandleFunc("/sar/accounts/{accountid}/ratings", rateAccount).Methods("POST") // rate account 
	r.HandleFunc("/sar/accounts/{accountid}/driver", getDriverRatings).Methods("GET") // view driver account ratings 
	r.HandleFunc("/sar/accounts/{accountid}/rider", getRiderRatings).Methods("GET") // view rider account ratings 

	// end Account related Endpoints

	// ***************************** //

	// Rides related endpoints
	r.HandleFunc("/sar/rides", postRide).Methods("POST") // Post a ride
	r.HandleFunc("/sar/rides/{rideid}", updateRide).Methods("PUT") // Update s ride
	r.HandleFunc("/sar/rides/{rideid}", deleteRide).Methods("DELETE") // Delete a ride
	r.HandleFunc("/sar/rides", getRides).Methods("GET") // Get all rides && search rides with searchRides()
	r.HandleFunc("/sar/rides/{rideid}", getRide).Methods("GET") // Get a ride
	r.HandleFunc("/sar/rides/{rideid}/join_requests", joinRide).Methods("POST") // Request to join a ride
	r.HandleFunc("/sar/rides/{rideid}/join_requests/{joinid}", confirmJoinRide).Methods("PATCH") // Confirm / Deny join request && confirm passenger pickup
	r.HandleFunc("/sar/rides/{rideid}/messages", addMessage).Methods("POST") // Add a message to a ride
	r.HandleFunc("/sar/rides/{rideid}/messages", getMessages).Methods("GET") // Get all rides && search rides with searchRides()
	
	// end Ride related Endpoints

	// ***************************** //

	// Reports related endpoints


	log.Fatal(http.ListenAndServe(":8080", r))
}

