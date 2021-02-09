package rest

import (
	"encoding/json"
	"log"
	"ms-asset/pkg/asset/application"
	"ms-asset/pkg/asset/infrastructure/request"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	assetApp application.AssetUseCaseInterface
)

type AssetRoute interface {
	AddRoutes(router *mux.Router)
}

type assetRoute struct {
	app application.AssetUseCaseInterface
}

// New ...
func New(
	app application.AssetUseCaseInterface,
) AssetRoute {
	assetApp = app
	return &assetRoute{app: assetApp}
}

func (pRoute *assetRoute) AddRoutes(router *mux.Router) {
	router.HandleFunc("/assets", add).Methods(http.MethodPost)
	router.HandleFunc("/assets/client/{clientId:[0-9-\\d]+}", searchByClient).Methods(http.MethodGet)
	router.HandleFunc("/assets/{sku:[0-9-\\d]+}", search).Methods(http.MethodGet)

}

func add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var cl request.AssetRequest
	err := decoder.Decode(&cl)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Error unmarshalling request body")
		return
	}
	//	AddProduct(ctx context.Context, requestData request.ProductRequest) error

	if err := assetApp.Add(r.Context(), cl); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Can't create the record")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func searchByClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if result, error := assetApp.GetByClient(r.Context(), vars["businessId"]); error != nil {
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Business not found")
		return
	} else {
		_ = json.NewEncoder(w).Encode(result)
	}

}

func search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if result, error := assetApp.Get(r.Context(), vars["sku"]); error != nil {
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Record not found")
		return
	} else {
		_ = json.NewEncoder(w).Encode(result)
	}

}
