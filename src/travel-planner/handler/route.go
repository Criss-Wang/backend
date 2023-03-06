package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"travel-planner/model"
	"travel-planner/service"
	"github.com/gorilla/mux"
)

func GetPlanHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a get sites request in the get all plans handler")
	va := mux.Vars(r)["vacation_id"]
	fmt.Println(va)
	vacationID, err := strconv.ParseUint(va, 10, 32)
	if err != nil {
		fmt.Println("cannot convert requestion id to vacationID")
	}

	fmt.Printf("vacationID: %v\n", vacationID)
	w.Header().Set("Content-Type", "application/json")

	//line 66 is hardcode for test, we cannot get info from http yet, we should use line65
	//vacationId := mux.Vars(r)["vacationid"]
	var plans *model.ListOfShowPlan
	plans, err = service.ShowRoute(uint32(vacationID))

	if err != nil || plans == nil {
		http.Error(w, "Failed to get sites from bd", http.StatusInternalServerError)
		return
	}

	// change to json
	js, err := json.Marshal(plans)
	if err != nil {
		http.Error(w, "Failed to parse sites to JSON format", http.StatusInternalServerError)
	}

	w.Write(js)
}
