package api

import (
	"github.com/boggydigital/nod"
	"net/http"
)

func HandleFuncs() {

	patternHandlers := map[string]func(http.ResponseWriter, *http.Request){
		"/":            http.RedirectHandler("/account", http.StatusPermanentRedirect).ServeHTTP,
		"/account":     nod.RequestLog(GetAccount),
		"/store":       nod.RequestLog(GetStore),
		"/product":     nod.RequestLog(GetProduct),
		"/search":      nod.RequestLog(GetSearch),
		"/images":      nod.RequestLog(GetImages),
		"/videos":      nod.RequestLog(GetVideos),
		"/files":       nod.RequestLog(GetFiles),
		"/local-file/": nod.RequestLog(GetLocalFile),
		"/css/":        http.FileServer(http.FS(cssFiles)).ServeHTTP,
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h)
	}
}
