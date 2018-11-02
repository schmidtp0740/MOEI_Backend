CREATE TABLE rxlist (
  patientID VARCHAR(5) NOT NULL,
  rxid VARCHAR(5) NOT NULL,
  timestamp BIGINT NOT NULL,
  status VARCHAR(10) NOT NULL,
  CONSTRAINT UC_rx UNIQUE (patientID, rxid, timestamp)
);