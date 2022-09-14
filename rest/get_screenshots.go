package rest

import (
	"net/http"

	"github.com/arelate/gaugin/gaugin_middleware"
	"github.com/arelate/gaugin/view_models"
	"github.com/arelate/vangogh_local_data"
	"github.com/boggydigital/nod"
)

func GetScreenshots(w http.ResponseWriter, r *http.Request) {

	// GET /screenshots?id

	id := r.URL.Query().Get("id")

	gaugin_middleware.DefaultHeaders(nil, w)

	idRedux, _, err := getRedux(
		http.DefaultClient,
		id,
		false,
		vangogh_local_data.ScreenshotsProperty)

	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	svm := view_models.NewScreenshots(idRedux[id])

	if err := tmpl.ExecuteTemplate(w, "screenshots-page", svm); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}