package main

import (
	"context"
	"fmt"

	"github.com/senzing/g2-sdk-go-grpc/g2configclient"
	"github.com/senzing/g2-sdk-go-grpc/g2diagnosticclient"
	"github.com/senzing/g2-sdk-go/g2config"
	"github.com/senzing/g2-sdk-go/g2diagnostic"
	pbg2config "github.com/senzing/g2-sdk-proto/go/g2config"
	pbg2diagnostic "github.com/senzing/g2-sdk-proto/go/g2diagnostic"
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

// ----------------------------------------------------------------------------
// Senzing object lifecycle functions
// ----------------------------------------------------------------------------

func getG2config(ctx context.Context, local bool) (g2config.G2config, error) {
	var result g2config.G2config

	// Determine which instantiation of the G2Configinterface to create.

	if local {
		result = &g2config.G2configImpl{}

	} else {
		grpcConnection := getGrpcConnection()
		result = &g2configclient.G2configClient{
			GrpcClient: pbg2config.NewG2ConfigClient(grpcConnection),
		}
	}

	// Initialize the object.

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

func getG2diagnostic(ctx context.Context, local bool) (g2diagnostic.G2diagnostic, error) {
	var result g2diagnostic.G2diagnostic

	// Determine which instantiation of the G2Diagnostic interface to create.

	if local {
		result = &g2diagnostic.G2diagnosticImpl{}

	} else {
		grpcConnection := getGrpcConnection()
		result = &g2diagnosticclient.G2diagnosticClient{
			GrpcClient: pbg2diagnostic.NewG2DiagnosticClient(grpcConnection),
		}
	}

	// Initialize the object.

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

func getSenzingObjects(ctx context.Context, isLocal bool) (g2config.G2config, g2diagnostic.G2diagnostic, error) {
	g2Config, err := getG2config(ctx, isLocal)
	if err != nil {
		fmt.Printf("getG2config: %v", err)
	}

	g2Diagnostic, err := getG2diagnostic(ctx, isLocal)
	if err != nil {
		fmt.Printf("getG2diagnostic: %v", err)
	}

	return g2Config, g2Diagnostic, err

}

func destroySenzingObjects(ctx context.Context, g2Config g2config.G2config, g2Diagnostic g2diagnostic.G2diagnostic) error {
	var err error = nil

	err = g2Config.Destroy(ctx)
	if err != nil {
		fmt.Printf("g2Config.Destroy: %v", err)
		return err
	}

	err = g2Diagnostic.Destroy(ctx)
	if err != nil {
		fmt.Printf("g2Diagnostic.Destroy: %v", err)
		return err
	}

	return err
}

// ----------------------------------------------------------------------------
// demonstrateXxxx
// ----------------------------------------------------------------------------

func demonstrateG2config(ctx context.Context, g2Config g2config.G2config) {
	configHandle, err := g2Config.Create(ctx)
	if err != nil {
		fmt.Printf("g2Config.Create: %v\n", err)
	}

	result, err := g2Config.ListDataSources(ctx, configHandle)
	if err != nil {
		fmt.Printf("g2Config.ListDataSources: %v\n", err)
	}

	fmt.Printf("Data Sources: %s\n", result)
}

func demonstrateG2diagnostic(ctx context.Context, g2Diagnostic g2diagnostic.G2diagnostic) {
	result, err := g2Diagnostic.GetTotalSystemMemory(ctx)
	if err != nil {
		fmt.Printf("g2Diagnostic.GetTotalSystemMemory: %v\n", err)
	}
	fmt.Printf("Memory: %d\n", result)
}

func demonstrateSenzingObjects(ctx context.Context, g2Config g2config.G2config, g2Diagnostic g2diagnostic.G2diagnostic) error {
	var err error = nil
	demonstrateG2config(ctx, g2Config)
	demonstrateG2diagnostic(ctx, g2Diagnostic)
	return err
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()
	booleans := []bool{true, false}
	for _, isLocal := range booleans {

		// Get Senzing objects.

		g2Config, g2Diagnostic, err := getSenzingObjects(ctx, isLocal)
		if err != nil {
			fmt.Printf("Error in getSenzingObjects: %v\n", err)
		}

		// Demonstrate Senzing objects.

		err = demonstrateSenzingObjects(ctx, g2Config, g2Diagnostic)
		if err != nil {
			fmt.Printf("Error in demonstrateSenzingObjects: %v\n", err)
		}

		// Destroy Senzing objects.

		err = destroySenzingObjects(ctx, g2Config, g2Diagnostic)
		if err != nil {
			fmt.Printf("Error in destroySenzingObjects: %v\n", err)
		}

	} // for _, isLocal := range booleans

}
