package main

import (
	"awesome-go/internal/app/cardholder"
	"awesome-go/internal/app/transaction"
	"awesome-go/internal/boot"
	"awesome-go/internal/configs"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"net"
	"net/http"

	gw "awesome-go/rpc/protos" // Update
)

var (
	grpcEndpoint = net.JoinHostPort(configs.Config.Host, configs.Config.GRPCPort)
	host         = net.JoinHostPort(configs.Config.Host, configs.Config.Port)
)

func runHTTPServer() (*http.Server, error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := gw.RegisterCardHolderServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		return nil, err
	}

	err = gw.RegisterTransactionServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		return nil, err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	boot.Lgr.Infof("starting api server.")
	server := http.Server{
		Addr:    host,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		return nil, err
	}
	return &server, nil
}

func runGRPCServer() (*grpc.Server, error) {
	server := grpc.NewServer()

	cardholderFactory := cardholder.NewFactory()
	cardholderRepository := cardholder.NewRepository()
	cardHolderService := cardholder.NewService(cardholderRepository, cardholderFactory)
	cardholderServer := cardholder.NewCardHolderServer(cardHolderService)

	transactionFactory := transaction.NewFactory()
	transactionRepository := transaction.NewRepository()
	transactionService := transaction.NewService(transactionRepository, transactionFactory, cardHolderService)
	transactionServer := transaction.NewTransactionServer(transactionService)

	gw.RegisterCardHolderServiceServer(server, cardholderServer)
	gw.RegisterTransactionServiceServer(server, transactionServer)
	// start listening on server port for a tcp connection
	if l, err := net.Listen("tcp", grpcEndpoint); err != nil {
		boot.Lgr.Fatalf("error in listening on port :%d, err:%s", configs.Config.GRPCPort, err.Error())
	} else {
		boot.Lgr.Info("starting grpc server.")
		if err := server.Serve(l); err != nil {
			boot.Lgr.Fatal("unable to start server", err)
		}
	}
	return server, nil
}
func main() {
	boot.Setup()
	go func() {
		_, err := runGRPCServer()
		if err != nil {
			panic(err)
		}
	}()
	_, err := runHTTPServer()
	if err != nil {
		grpclog.Fatal(err)
	}
}
