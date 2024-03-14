package main

import (
	"net/http"

	"github.com/raphaelmb/quick-db/internal/database"
	"github.com/raphaelmb/quick-db/internal/sdk"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sdk.Setup(database.NewMySQL("mysql", "", "", "", "", "", "", false))
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":9000", nil)
}
