package rest

import (
	"github.com/arelate/vangogh_local_data"
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
)

var (
	Auth     = middleware.BasicHttpAuth
	GetOnly  = middleware.GetMethodOnly
	PostOnly = middleware.PostMethodOnly
	Gzip     = middleware.Gzip
	Static   = middleware.Static
	Log      = nod.RequestLog
	Redirect = http.RedirectHandler
)

var port int

func HandleFuncs(p int) {

	port = p

	patternHandlers := map[string]http.Handler{
		// unauth data endpoints
		"/updates":       Gzip(GetOnly(Static(Log(http.HandlerFunc(GetUpdates))))),
		"/product":       Gzip(GetOnly(Static(Log(http.HandlerFunc(GetProduct))))),
		"/search":        Gzip(GetOnly(Static(Log(http.HandlerFunc(GetSearch))))),
		"/digest":        Gzip(GetOnly(Log(http.HandlerFunc(GetDigest)))),
		"/description":   Gzip(GetOnly(Log(http.HandlerFunc(GetDescription)))),
		"/downloads":     Gzip(GetOnly(Log(http.HandlerFunc(GetDownloads)))),
		"/changelog":     Gzip(GetOnly(Log(http.HandlerFunc(GetChangelog)))),
		"/screenshots":   Gzip(GetOnly(Log(http.HandlerFunc(GetScreenshots)))),
		"/videos":        Gzip(GetOnly(Log(http.HandlerFunc(GetVideos)))),
		"/steam-news":    Gzip(GetOnly(Log(http.HandlerFunc(GetSteamNews)))),
		"/steam-reviews": Gzip(GetOnly(Log(http.HandlerFunc(GetSteamReviews)))),
		// unauth media endpoints
		"/image":      GetOnly(Log(http.HandlerFunc(GetImage))),
		"/video":      GetOnly(Log(http.HandlerFunc(GetVideo))),
		"/thumbnails": GetOnly(Log(http.HandlerFunc(GetThumbnails))),
		"/items/":     GetOnly(Log(http.HandlerFunc(GetItems))),
		// auth data endpoints
		"/wishlist/add":     Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetWishlistAdd)))), AdminRole),
		"/wishlist/remove":  Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetWishlistRemove)))), AdminRole),
		"/tags/edit":        Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetTagsEdit)))), AdminRole),
		"/local-tags/edit":  Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetLocalTagsEdit)))), AdminRole),
		"/tags/apply":       Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetTagsApply)))), AdminRole),
		"/local-tags/apply": Auth(Gzip(GetOnly(Log(http.HandlerFunc(GetLocalTagsApply)))), AdminRole),
		// auth media endpoints
		"/files":       Auth(GetOnly(Log(http.HandlerFunc(GetFiles))), AdminRole, SharedRole),
		"/local-file/": Auth(GetOnly(Log(http.HandlerFunc(GetLocalFile))), AdminRole, SharedRole),
		// prerender
		"/prerender": PostOnly(Log(http.HandlerFunc(PostPrerender))),
		// products redirects
		"/products": Redirect("/search", http.StatusPermanentRedirect),
		// start at the updates
		"/": Redirect("/updates", http.StatusPermanentRedirect),
		// robots.txt
		"/robots.txt": Gzip(GetOnly(Log(http.HandlerFunc(GetRobotsTxt)))),
	}

	for route, path := range searchRoutes() {
		patternHandlers["/products/"+route] = Redirect(path, http.StatusPermanentRedirect)
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}

func searchRoutes() map[string]string {
	routes := make(map[string]string)

	searchPath := "/search"

	routes["filter"] = searchPath

	q := make(url.Values)
	q.Set(vangogh_local_data.TypesProperty, vangogh_local_data.AccountProducts.String())
	q.Set(vangogh_local_data.SortProperty, vangogh_local_data.GOGOrderDateProperty)
	q.Set(vangogh_local_data.DescendingProperty, vangogh_local_data.TrueValue)
	routes["owned"] = searchPath + "?" + q.Encode()

	q = make(url.Values)
	q.Set(vangogh_local_data.WishlistedProperty, vangogh_local_data.TrueValue)
	q.Set(vangogh_local_data.SortProperty, vangogh_local_data.GOGReleaseDateProperty)
	q.Set(vangogh_local_data.DescendingProperty, vangogh_local_data.TrueValue)
	routes["wishlist"] = searchPath + "?" + q.Encode()

	q = make(url.Values)
	q.Set(vangogh_local_data.TypesProperty, vangogh_local_data.CatalogProducts.String())
	q.Set(vangogh_local_data.OwnedProperty, vangogh_local_data.FalseValue)
	q.Set(vangogh_local_data.IsDiscountedProperty, vangogh_local_data.TrueValue)
	q.Set(vangogh_local_data.SortProperty, vangogh_local_data.DiscountPercentageProperty)
	q.Set(vangogh_local_data.DescendingProperty, vangogh_local_data.TrueValue)
	routes["sale"] = searchPath + "?" + q.Encode()

	q = make(url.Values)
	q.Set(vangogh_local_data.TypesProperty, vangogh_local_data.CatalogProducts.String())
	q.Set(vangogh_local_data.SortProperty, vangogh_local_data.GOGReleaseDateProperty)
	q.Set(vangogh_local_data.DescendingProperty, vangogh_local_data.TrueValue)
	routes["all"] = searchPath + "?" + q.Encode()

	return routes
}
