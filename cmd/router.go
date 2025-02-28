package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/naveeshkumar24/internal/handlers"
	"github.com/naveeshkumar24/internal/middleware"
	"github.com/naveeshkumar24/repository"
)

func registerRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.CorsMiddleware)

	billingporepo := repository.NewBillingPoRepository(db)
	BillingPoHandler := handlers.NewBillingPoHandler(billingporepo)
	router.HandleFunc("/dropdown", BillingPoHandler.FetchDropDown).Methods("GET")
	router.HandleFunc("/submit", BillingPoHandler.SubmitFormBillingPoData).Methods("POST")
	router.HandleFunc("/fetch", BillingPoHandler.FetchBillingPoData).Methods("GET")
	router.HandleFunc("/update", BillingPoHandler.UpdateBillingPoData).Methods("POST")
	router.HandleFunc("/delete/{id}", BillingPoHandler.DeleteBillingPoHandler).Methods("POST")
	excelDownloadBillingHandler := handlers.NewExcelDownloadBPOHandler(billingporepo)
	router.HandleFunc("/download", excelDownloadBillingHandler.DownloadBPO).Methods("GET")

	return router
}
