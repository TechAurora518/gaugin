package api

import (
	"github.com/arelate/vangogh_local_data"
	"log"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	idRedux, err := getRedux(http.DefaultClient, id, vangogh_local_data.ReduxProperties()...)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Add("Content-Type", "text/html")

	pvm, err := productViewModelFromRedux(idRedux)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.ExecuteTemplate(w, "product", pvm); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
	}
}
