package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"ohs-user/pkg/user/application"
	"ohs-user/pkg/user/infrastructure/request"

	"github.com/gorilla/mux"
)

var (
	userApp application.UserUseCaseInterface
)

type UserRoute interface {
	AddRoutes(router *mux.Router)
}

type userRoute struct {
	app application.UserUseCaseInterface
}

// New ...
func New(
	app application.UserUseCaseInterface,
) UserRoute {
	userApp = app
	return &userRoute{app: userApp}
}

func (pRoute *userRoute) AddRoutes(router *mux.Router) {
	router.HandleFunc("/user", add).Methods(http.MethodPost)
	router.HandleFunc("/user/client/{clientId:[0-9-\\d]+}", searchByClient).Methods(http.MethodGet)
	router.HandleFunc("/user/{sku:[0-9-\\d]+}", search).Methods(http.MethodGet)

}

func add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var cl request.UserRequest
	err := decoder.Decode(&cl)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Error unmarshalling request body")
		return
	}
	//	AddProduct(ctx context.Context, requestData request.ProductRequest) error

	if err := userApp.SignUp(r.Context(), cl); err != nil {
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
	if result, error := userApp.SignIn(r.Context(), vars["businessId"]); error != nil {
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
	if result, error := userApp.ChangePassword(r.Context(), vars["sku"]); error != nil {
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Record not found")
		return
	} else {
		_ = json.NewEncoder(w).Encode(result)
	}

}
