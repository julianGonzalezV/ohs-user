package main

// Main or entry point for our application
import (
	"flag"
	"fmt"
	"log"

	userRepository "ohs-user/pkg/user/domain/repository"

	"ohs-user/shared/server"
	"ohs-user/shared/storageconn"

	"net/http"
	"os"
	"strconv"

	"github.com/apex/gateway"
)

/// initializeRepo returns a repository based on database type name
func initializeRepo(database *string) userRepository.UserRepository {
	switch *database {
	case "mongo":
		return newMongoRepository()
	default:
		return nil // we can have several implementation like in memory, postgress etc
	}
}

/// newMongoRepository returns the mongoDB implementation
func newMongoRepository() userManagerDependy.UserManagement {
	mongoAddr := os.Getenv("DATABASE_CONN")
	client := storageconn.Connect(mongoAddr)
	return userRepositoryImpl.New(client)
}

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
	userR := initializeRepo(database)
	userUseCase := userUseCase.New(userService.New(userR))

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
