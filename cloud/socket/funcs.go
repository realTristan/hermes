package ws

import (
	Hermes "github.com/realTristan/Hermes"
	"github.com/realTristan/Hermes/cloud/socket/handlers"
	Utils "github.com/realTristan/Hermes/cloud/socket/utils"
)

// Map of functions that can be called from the client
var Functions = map[string]func(*Utils.Params, *Hermes.Cache) []byte{
	"cache.length":        handlers.Length,
	"cache.clean":         handlers.Clean,
	"cache.set":           handlers.Set,
	"cache.delete":        handlers.Delete,
	"cache.get":           handlers.Get,
	"cache.get.all":       handlers.GetAll,
	"cache.keys":          handlers.Keys,
	"cache.info":          handlers.Info,
	"cache.info.testing":  handlers.InfoForTesting,
	"cache.exists":        handlers.Exists,
	"ft.init":             handlers.FTInit,
	"ft.init.json":        handlers.FTInitJson,
	"ft.clean":            handlers.FTClean,
	"ft.search":           handlers.Search,
	"ft.search.oneword":   handlers.SearchOneWord,
	"ft.search.values":    handlers.SearchValues,
	"ft.search.withkey":   handlers.SearchWithKey,
	"ft.maxbytes.set":     handlers.FTSetMaxBytes,
	"ft.maxlength.set":    handlers.FTSetMaxLength,
	"ft.storage":          handlers.FTStorage,
	"ft.storage.size":     handlers.FTStorageSize,
	"ft.storage.length":   handlers.FTStorageLength,
	"ft.isinitialized":    handlers.FTIsInitialized,
	"ft.indices.sequence": handlers.FTSequenceIndices,
}
