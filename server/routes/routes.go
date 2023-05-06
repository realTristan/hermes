package routes

import (
	"github.com/gorilla/mux"
	Hermes "github.com/realTristan/Hermes"
	"github.com/realTristan/Hermes/server/handlers"
)

// Set the routes
func Set(cache *Hermes.Cache, router *mux.Router) {
	// Cache handlers
	router.HandleFunc("/values", handlers.Values(cache)).Methods("GET")
	router.HandleFunc("/length", handlers.Length(cache)).Methods("GET")
	router.HandleFunc("/clean", handlers.Clean(cache)).Methods("POST")
	router.HandleFunc("/set", handlers.Set(cache)).Methods("POST")
	router.HandleFunc("/delete", handlers.Delete(cache)).Methods("DELETE")
	router.HandleFunc("/get", handlers.Get(cache)).Methods("GET")
	router.HandleFunc("/get/all", handlers.GetAll(cache)).Methods("GET")
	router.HandleFunc("/keys", handlers.Keys(cache)).Methods("GET")
	router.HandleFunc("/info", handlers.Info(cache)).Methods("GET")
	router.HandleFunc("/exists", handlers.Exists(cache)).Methods("GET")

	// Full text cache handlers
	router.HandleFunc("/ft/init", handlers.FTInit(cache)).Methods("POST")
	router.HandleFunc("/ft/init/json", handlers.FTInitJson(cache)).Methods("POST")
	router.HandleFunc("/ft/clean", handlers.FTClean(cache)).Methods("POST")
	router.HandleFunc("/ft/search", handlers.Search(cache)).Methods("GET")
	router.HandleFunc("/ft/searchoneword", handlers.SearchOneWord(cache)).Methods("GET")
	router.HandleFunc("/ft/searchvalues", handlers.SearchValues(cache)).Methods("GET")
	router.HandleFunc("/ft/searchvalueswithkey", handlers.SearchValuesWithKey(cache)).Methods("GET")
	router.HandleFunc("/ft/maxbytes", handlers.FTSetMaxBytes(cache)).Methods("POST")
	router.HandleFunc("/ft/maxwords", handlers.FTSetMaxWords(cache)).Methods("POST")
	router.HandleFunc("/ft/wordcache", handlers.FTWordCache(cache)).Methods("GET")
	router.HandleFunc("/ft/wordcachesize", handlers.FTWordCacheSize(cache)).Methods("GET")
	router.HandleFunc("/ft/isinitialized", handlers.FTIsInitialized(cache)).Methods("GET")
	router.HandleFunc("/ft/add", handlers.FTAdd(cache)).Methods("POST")
}
