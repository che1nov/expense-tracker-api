package handlers

import (
	"encoding/json"
	database "expense-tracker-api/internal/db"
	"expense-tracker-api/internal/models"

	"net/http"
)

func AddExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uint)
	expense.UserID = userID

	if err := database.DB.Create(&expense).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Expense added"})
}

func ListExpenses(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	var expenses []models.Expense
	database.DB.Where("user_id = ?", userID).Find(&expenses)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}
