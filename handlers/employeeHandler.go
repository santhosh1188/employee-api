package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santhosh1188/employee-api/database"
	"github.com/santhosh1188/employee-api/models"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employee
	database.DB.Find(&employees)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var employee models.Employee
	result := database.DB.First(&employee, params["id"])
	if result.Error != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(employee)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Create(&emp)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var emp models.Employee
	if err := database.DB.First(&emp, params["id"]).Error; err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	var updated models.Employee
	json.NewDecoder(r.Body).Decode(&updated)

	emp.Name = updated.Name
	emp.Age = updated.Age
	emp.Department = updated.Department

	database.DB.Save(&emp)
	json.NewEncoder(w).Encode(emp)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var emp models.Employee
	if err := database.DB.First(&emp, params["id"]).Error; err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&emp)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee deleted"})
}
