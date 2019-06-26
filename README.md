# firstproject
Project based on HLF1.4 with WEB and REST Api's. This poroject is tested on Windows OS. To run it on Ubuntu, please make the appropiate changes like changing the Path in main.go file(ChannelConfig and ChaincodeGoPath)

### Start the project
1. go to $GOPATH/src/github.com/hyperledger/firstproject
2. run "make" command to start the docker

### REST API's 
1. To get the value: curl -X GET http://localhost:3000/Index
2. To change the value: curl -v -X POST -d '{"Value":"Shiv"}' http://localhost:3000/changevalue
