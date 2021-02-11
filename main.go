package main

// Main or entry point for our application
import (
	"flag"
	"fmt"
	"log"

	"ohs-user/pkg/user/infrastructure/usermanagerimpl"

	"ohs-user/shared/server"

	"net/http"
	"os"
	"strconv"

	"github.com/apex/gateway"
)

// ClientHandler set up all dependencies
func ClientHandler() {
	var (
		defaultHost    = os.Getenv("CLIENTAPI_SERVER_HOST")
		defaultPort, _ = strconv.Atoi(os.Getenv("CLIENTAPI_SERVER_PORT"))
		dbDriver       = os.Getenv("DATABASE_DRIVER")
	)
	host := flag.String("host", defaultHost, "define host of the server")
	port := flag.Int("port", defaultPort, "define port of the server")
	database := flag.String("database", dbDriver, "initialize the api using the given db engine")

	// Injecting services and repos to Application Layer
	userUseCase := userUseCase.New(userService.New(usermanagerimpl.New()))

	httpAddr := fmt.Sprintf("%s:%d", *host, *port)

	// Injecting server configuration
	userRoute := userRoute.New(userUseCase)
	server := server.New(userRoute)

	// Next two lines are for AWS Conf
	http.Handle("/", server.Router())
	log.Fatal(gateway.ListenAndServe(httpAddr, nil))

	// Next line is for Local conf
	//log.Fatal(http.ListenAndServe(httpAddr, server.Router()))
	fmt.Println("The client server is running", httpAddr)

}

func main() {
	fmt.Println("V1.0.0")
	ClientHandler()
}
