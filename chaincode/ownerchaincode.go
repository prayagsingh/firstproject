package main

/*
import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Car struct which store the car details
type Car struct {
	modelName    string
	color        string
	serialNo     string
	manufacturer string
	owner        Owner //composition
}

// Owner struct which stores car owner's details
type Owner struct {
	ID      string
	name    string
	address string
}

// CarChaincode struct of type Chaincode which further implements two methods Init and Invoke
type CarChaincode struct {
}

// Init method implemented by CarChaincode
func (t *CarChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	//Declare owners from Owner struct
	alpha := Owner{name: "Alpha", address: "1, Tumbbad", ID: "A1"}
	beta := Owner{name: "Beta", address: "2, Tumbbad", ID: "A2"}

	//Declaree carfrom Car struct
	car := Car{modelName: "Land Rover", color: "white", serialNo: "334712531234", manufacturer: "Tata Motors", owner: alpha}

	// Convert alpha Owner to []byte
	alphaJSON, err := json.Marshal(alpha)
	if err != nil {
		return shim.Error("Error in marshelling alpha Json")
	}

	// Add alpha to ledger
	error := stub.PutState(alpha.ID, alphaJSON)
	if error != nil {
		return shim.Error("Failed to create asset " + alpha.name)
	}

	// Add beta to ledger
	betaJSON, err := json.Marshal(beta)
	if err != nil {
		return shim.Error("Error in Marshelling betajson")
	}
	// Add beta to ledger
	err = stub.PutState(beta.ID, betaJSON)
	if err != nil {
		return shim.Error("Failed to create asset " + beta.name)
	}

	//Add car to ledger
	carAsJSONBytes, err := json.Marshal(car)
	if err != nil {
		return shim.Error("Error in Marshelling carJson")
	}
	err = stub.PutState(car.serialNo, carAsJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + car.serialNo)
	}
	return shim.Success([]byte("Assets created successfully."))
}

// Invoke implemented by CarChaincode
func (t *CarChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	// Read args from the transaction proposal.
	// fc=> method to invoke
	function, args := stub.GetFunctionAndParameters()
	if function != "TransferOwnership" {
		return shim.Error("Unkmown function call")
	}
	return t.TransferOwnership(stub, args)

}

// TransferOwnership method to change the owner
func (t *CarChaincode) TransferOwnership(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// args[0]=> car serial no
	// args[1]==> new owner national identity
	// Read existing car asset
	carAsBytes, _ := stub.GetState(args[0])
	if carAsBytes == nil {
		return shim.Error("car asset not found")
	}

	// Construct the struct Car
	car := Car{}
	err := json.Unmarshal(carAsBytes, &car)
	if err != nil {
		shim.Error("unable to unmarshal carBytes")
	}
	//Read newOwnerDetails
	ownerAsBytes, _ := stub.GetState(args[1])
	if ownerAsBytes == nil {
		return shim.Error("owner asset not found")
	}
	// Construct the struct Owner
	newOwner := Owner{}
	err = json.Unmarshal(ownerAsBytes, &newOwner)
	if err != nil {
		shim.Error("unable to unmarshal ownerBytes")
	}

	// Update owner
	car.ChangeOwner(newOwner)
	carAsJSONBytes, _ := json.Marshal(car)

	// Update car ownership in the
	err = stub.PutState(car.serialNo, carAsJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + car.serialNo)
	}
	return shim.Success([]byte("Asset modified."))

}

// ChangeOwner method to change the car owner
func (c *Car) ChangeOwner(newOwner Owner) {
	c.owner = newOwner
}

func main() {
	logger.SetLevel(shim.LogInfo)
	// Start the chaincode process
	err := shim.Start(new(CarChaincode))
	if err != nil {
	logger.Error("Error starting PhantomChaincode - ", err.Error()
	}
}

*/
