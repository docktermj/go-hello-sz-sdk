package main

import (
	"context"
	"fmt"

	"github.com/senzing/g2-sdk-go-grpc/g2diagnosticclient"
	"github.com/senzing/g2-sdk-go/g2diagnostic"
	pb "github.com/senzing/g2-sdk-proto/go/g2diagnostic"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcAddress = "localhost:8258"
)

// ----------------------------------------------------------------------------
// Internal methods - names begin with lower case
// ----------------------------------------------------------------------------

func getGrpcConnection() *grpc.ClientConn {
	result, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Did not connect: %v\n", err)
	}
	return result
}

func getG2diagnostic(ctx context.Context, local bool) (g2diagnostic.G2diagnostic, error) {
	var result g2diagnostic.G2diagnostic

	// Determine which instantiation of the G2Diagnostic interface to create.

	if local {
		result = &g2diagnostic.G2diagnosticImpl{}

	} else {
		grpcConnection := getGrpcConnection()
		result = &g2diagnosticclient.G2diagnosticClient{
			G2DiagnosticGrpcClient: pb.NewG2DiagnosticClient(grpcConnection),
		}
	}

	// Initialize the G2diagnostic object.

	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		fmt.Printf("Cannot construct system configuration. Error: %v\n", err)
	}
	err = result.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		fmt.Printf("Cannot Init. Error: %v\n", err)
	}
	return result, err

}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()

	booleans := []bool{true, false}
	for _, isLocal := range booleans {

		// Get Senzing objects.

		g2Diagnostic, err := getG2diagnostic(ctx, isLocal)
		if err != nil {
			fmt.Printf("getG2diagnostic: %v", err)
		}

		// Demonstrate tests.

		result, err := g2Diagnostic.GetTotalSystemMemory(ctx)
		if err != nil {
			fmt.Printf("g2Diagnostic.GetTotalSystemMemory: %v\n", err)
		}
		fmt.Printf("Memory: %d\n", result)

		// Destroy Senzing objects

		err = g2Diagnostic.Destroy(ctx)
		if err != nil {
			fmt.Printf("g2Diagnostic.Destroy: %v", err)
		}
	} // for _, isLocal := range booleans

}
