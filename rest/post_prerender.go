package rest

import (
	"github.com/arelate/gaugin/paths"
	"github.com/arelate/vangogh_local_data"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/stencil/stencil_rest"
	"net/http"
)

func PostPrerender(w http.ResponseWriter, _ *http.Request) {

	// POST /prerender

	// the following pages will be pre-rendered:
	// - default path (/updates)
	// - every top-level search route (/search, owned, wishlist, sale, all)
	// - every product updated at the last sync

	if err := setPrerender(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}

func setPrerender() error {
	ps := createListsPaths()
	var err error

	ps, err = appendUpdatedItemsPaths(ps)
	if err != nil {
		return err
	}

	return stencil_rest.Prerender(ps, true, port)
}

func updatePrerender(ids ...string) error {
	ps := createListsPaths()

	for _, id := range ids {
		ps = append(ps, paths.ProductId(id))
	}

	return stencil_rest.Prerender(ps, false, port)
}

func createListsPaths() []string {

	ps := make([]string, 0)
	ps = append(ps, "/updates")

	for _, p := range searchRoutes() {
		ps = append(ps, p)
	}

	return ps
}

func appendUpdatedItemsPaths(p []string) ([]string, error) {

	updRdx, _, err := getRedux(
		http.DefaultClient,
		"",
		true,
		vangogh_local_data.LastSyncUpdatesProperty)

	if err != nil {
		return nil, err
	}

	keys := make(map[string]interface{})
	for _, rdx := range updRdx {
		for _, id := range rdx[vangogh_local_data.LastSyncUpdatesProperty] {
			keys[id] = nil
		}
	}

	for id := range keys {
		p = append(p, paths.ProductId(id))
	}

	return p, nil
}
