package stencil_app

import (
	"github.com/arelate/gaugin/data"
	"github.com/arelate/vangogh_local_data"
)

var PropertyTitles = map[string]string{
	vangogh_local_data.TextProperty:                 "Any Text",
	vangogh_local_data.TitleProperty:                "Title",
	vangogh_local_data.TagIdProperty:                "Account Tags",
	vangogh_local_data.LocalTagsProperty:            "Local Tags",
	vangogh_local_data.SteamTagsProperty:            "Steam Tags",
	vangogh_local_data.OperatingSystemsProperty:     "OS",
	vangogh_local_data.DevelopersProperty:           "Developers",
	vangogh_local_data.PublishersProperty:           "Publishers",
	vangogh_local_data.SeriesProperty:               "Series",
	vangogh_local_data.GenresProperty:               "Genres",
	vangogh_local_data.StoreTagsProperty:            "Store Tags",
	vangogh_local_data.FeaturesProperty:             "Features",
	vangogh_local_data.LanguageCodeProperty:         "Language",
	vangogh_local_data.IncludesGamesProperty:        "Includes",
	vangogh_local_data.IsIncludedByGamesProperty:    "Included By",
	vangogh_local_data.RequiresGamesProperty:        "Requires",
	vangogh_local_data.IsRequiredByGamesProperty:    "Required By",
	vangogh_local_data.ProductTypeProperty:          "Product Type",
	vangogh_local_data.WishlistedProperty:           "Wishlisted",
	vangogh_local_data.OwnedProperty:                "Owned",
	vangogh_local_data.IsFreeProperty:               "Free",
	vangogh_local_data.IsDiscountedProperty:         "On Sale",
	vangogh_local_data.PreOrderProperty:             "Pre-order",
	vangogh_local_data.ComingSoonProperty:           "Coming Soon",
	vangogh_local_data.TBAProperty:                  "TBA",
	vangogh_local_data.InDevelopmentProperty:        "In Development",
	vangogh_local_data.IsUsingDOSBoxProperty:        "Using DOSBox",
	vangogh_local_data.IsUsingScummVMProperty:       "Using ScummVM",
	vangogh_local_data.TypesProperty:                "Data Type",
	vangogh_local_data.SteamReviewScoreDescProperty: "Steam Reviews",
	vangogh_local_data.SortProperty:                 "Sort",
	vangogh_local_data.DescendingProperty:           "Descending",
	vangogh_local_data.GlobalReleaseDateProperty:    "Global Release",
	vangogh_local_data.GOGReleaseDateProperty:       "GOG.com Release",
	vangogh_local_data.GOGOrderDateProperty:         "GOG.com Order",
	vangogh_local_data.ValidationResultProperty:     "Validation Result",
	vangogh_local_data.RatingProperty:               "Rating",

	data.GauginGOGLinksProperty:           "GOG.com Links",
	vangogh_local_data.ForumUrlProperty:   "Forum",
	vangogh_local_data.StoreUrlProperty:   "Store",
	vangogh_local_data.SupportUrlProperty: "Support",
	data.GauginSteamLinksProperty:         "Steam Links",
	data.GauginSteamCommunityUrlProperty:  "Community",
}
