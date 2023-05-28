package main

import (
	"log"
	"os"

	"github.com/aman-singh7/go-hex-arch/internal/adapters/framework/driven/db"
	"github.com/aman-singh7/go-hex-arch/internal/application/api"
	"github.com/aman-singh7/go-hex-arch/internal/application/core/arithmetic"
	"github.com/aman-singh7/go-hex-arch/internal/ports"

	gRPC "github.com/aman-singh7/go-hex-arch/internal/adapters/framework/driver/grpc"
)

func main() {
	var err error

	var dbAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbAdapter, err = db.NewAdapter(dbDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()
}
