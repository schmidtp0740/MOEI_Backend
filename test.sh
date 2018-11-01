#!/bin/bash
clear

# printf "GetPeople\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/pd 
# printf "\n"

# printf "GetPerson\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/pd/p01
# printf "\n"

printf "insertRx\n"
curl -H "Content-type:application/json" -X POST http://localhost:8080/rx -d '{"patientID":"p01","rxid":"rx02","timestamp":1541022553032,"doctor":"rick","docLicense":"doc001","prescription":"mitodel","refills":1,"quantity":2.5,"expDate": 1541022553034,"status":"prescribed","approved":"false"}'
printf "\n"

printf "fillRx\n"
curl -H "Content-type:application/json" -X PATCH http://localhost:8080/rx -d '{"patientID":"p01","rxid":"rx02","timestamp":15123124,"pharmacist": "phar","phLicense": "phar001", "prescription":"pre","refills":1,"expDate": 12345,"status":"filled","approved":"true"}'
printf "\n"

printf "approveRx\n"
curl -H "Content-type:application/json" -X PUT http://localhost:8080/rx -d '{"patientID":"p01","rxid":"rx01","timestamp":15123124,"approved":"true"}'
printf "\n"

printf "getRxForPatient\n"
curl -H "Content-type:application/json" -X GET http://localhost:8080/rx/p01 -d '{"patientID":"p01"}'
printf "\n"

# printf "insertHeartRate\n"
# curl -H "Content-type:application/json" -X POST http://localhost:8080/hr -d '{"patientID":"p01","heartRate":80,"timestamp":123457}'
# printf "\n"

# printf "getHeartRateDataForPatient\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/hr/p01 
# printf "\n"

printf "getAllRxHistory\n"
curl -H "Content-type:application/json" -X GET http://localhost:8080/rxledger 
printf "\n"

# printf "hack\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/hack 
# printf "\n"

# printf "isHacked\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/bcs 
# printf "\n"

# printf "insertInsurance\n"
# curl -H "Content-type:application/json" -X POST http://localhost:8080/insurance -d '{"patientID": "p01", "insuranceName": "gieco", "expDate": 12345, "policyID": "p0000123"}'
# printf "\n"

# printf "getInsurance\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/insurance/p01
# printf "\n"

# printf "insertBloodPressure\n"
# curl -H "Content-type:application/json" -X POST http://localhost:8080/bp -d '{"patientID": "p01", "low": 10, "high": 81, "timestamp": 15123124}'
# printf "\n"

# printf "getBloodPressure\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/bp/p01
# printf "\n"




