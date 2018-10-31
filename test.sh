#!/bin/bash
clear

# curl -H "Content-type:application/json" -X GET http://localhost:8080/pd 
# printf "\n"
# curl -H "Content-type:application/json" -X POST http://localhost:4001/bcsgw/rest/v1/transaction/query -d '{"channel":"mychannel","chaincode":"emrcc","method":"getPeople","args":[],"chaincodeVer": "v1" }'
# printf "\n"

# curl -H "Content-type:application/json" -X GET http://localhost:8080/pd/p01
# printf "\n"
# curl -H "Content-type:application/json" -X POST http://localhost:4001/bcsgw/rest/v1/transaction/query -d '{"channel":"mychannel","chaincode":"emrcc","method":"getPerson","args":["p01"],"chaincodeVer": "v1" }'
# printf "\n"

curl -H "Content-type:application/json" -X POST http://localhost:8080/rx -d '{"patientID":"p01","rxid":"rx01","timestamp":15123123,"doctor":"doctor","prescription":"pre","refills":1,"status":"prescribed","approved":"true"}'
printf "\n"

curl -H "Content-type:application/json" -X PATCH http://localhost:8080/rx -d '{"patientID":"p01","rxid":"rx01","timestamp":15123124,"doctor":"doctor","pharmacist": "phar", "prescription":"pre","refills":1,"status":"filled","approved":"true"}'
printf "\n"

curl -H "Content-type:application/json" -X GET http://localhost:8080/rx/p01 -d '{"patientID":"p01"}'
printf "\n"
