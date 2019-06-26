package main

import (
	"fmt"

	"github.com/hyperledger/firstproject/blockchain"
	"github.com/hyperledger/firstproject/web"
	"github.com/hyperledger/firstproject/web/controllers"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.firstproject.com",

		// Channel parameters
		ChannelID:     "mychannel",
		ChannelConfig: "/c/Projects/Go/src/github.com/hyperledger/firstproject/firstproject-network/artifacts/channel.tx",

		// Chaincode parameters
		ChainCodeID:     "firstproject",
		ChaincodeGoPath: "/c/Projects/Go",
		ChaincodePath:   "github.com/hyperledger/firstproject/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}
	/*
		// Using mux to expose to REST Api's
		router := mux.NewRouter()
		router.HandleFunc("/api/{func}", fSetup.QueryHello).Methods("GET")
		router.HandleFunc("/api/{func}", fSetup.InvokeHello).Methods("POST")

		log.Fatal(http.ListenAndServe("localhost:8000", router))

	*/
	/*
		// Query the chaincode
		response, err := fSetup.QueryHello()
		if err != nil {
			fmt.Printf("Unable to query hello on the chaincode: %v\n", err)
		} else {
			fmt.Printf("Response from the query hello: %s\n", response)
		}

		// Invoke the chaincode
		txID, err := fSetup.InvokeHello("FirstProject")
		if err != nil {
			fmt.Printf("Unable to invoke hello on the chaincode: %v\n", err)
		} else {
			fmt.Printf("Successfully invoke hello, transaction ID: %s\n", txID)
		}

		// Query again the chaincode
		response, err = fSetup.QueryHello()
		if err != nil {
			fmt.Printf("Unable to query hello on the chaincode: %v\n", err)
		} else {
			fmt.Printf("Response from the query hello: %s\n", response)
		}
	*/
	// Launch the web application listening
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)

}
