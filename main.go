package main

import (
	"fmt"
	"net"
	"os"

	"grpc-service-template/db"
	"grpc-service-template/pb"
	"grpc-service-template/services/users"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := readConfig()
	if err != nil {
		log.WithError(err).Fatal("failed to read config")
		os.Exit(1)
	}

	dbc, err := db.New(config.DBConnString)
	if err != nil {
		log.WithError(err).Error("failed connect to database")
		os.Exit(1)
	}
	fmt.Println(config.DBConnString)
	if err := db.Migrate(config.DBConnString); err != nil {
		log.WithError(err).Error("failed to run migrations")
		os.Exit(1)
	}

	usersServiceHandlers := users.New(dbc)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.WithError(err).Fatal("failed to start listen")
		os.Exit(1)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	pb.RegisterUsersServer(server, usersServiceHandlers)

	if err := server.Serve(lis); err != nil {
		log.WithError(err).Fatal("failed to start serve")
		os.Exit(1)
	}
}
