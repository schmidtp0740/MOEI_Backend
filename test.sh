#!/bin/bash
clear

# printf "GetPeople\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/pd 
# printf "\n"

# printf "GetPerson\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/pd/p01
# printf "\n"

# printf "insertRx\n"
# curl -H "Content-type:application/json" -X POST http://localhost:8080/rx -d '{"patientID":"p01","rxid":"rx01","timestamp":1541022553032,"doctor":"rick","prescription":"mitodel","refills":1,"expDate": 1541022553034,"status":"prescribed","approved":"false"}'
# printf "\n"

# printf "modifyRx\n"
# curl -H "Content-type:application/json" -X PATCH http://localhost:8080/rx -d '{"patientID":"p01","rxid":"rx03","timestamp":15123124,"doctor":"doctor","pharmacist": "phar", "prescription":"pre","refills":1,"expDate": 12345,"status":"filled","approved":"true"}'
# printf "\n"

# printf "getRxForPatient\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/rx/p01 -d '{"patientID":"p01"}'
# printf "\n"

# printf "insertHeartRate\n"
# curl -H "Content-type:application/json" -X POST http://localhost:8080/hr -d '{"patientID":"p01","heartRate":80,"timestamp":123457}'
# printf "\n"

# printf "getHeartRateDataForPatient\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/hr/p01 
# printf "\n"

# printf "getAllRxHistory\n"
# curl -H "Content-type:application/json" -X GET http://localhost:8080/rx 
# printf "\n"

printf "hack\n"
curl -H "Content-type:application/json" -X GET http://localhost:8080/hack 
printf "\n"

printf "isHacked\n"
curl -H "Content-type:application/json" -X GET http://localhost:8080/bcs 
printf "\n"




