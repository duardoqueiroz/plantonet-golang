CREATE DATABASE plantonetgo;

\connect plantonetgo

CREATE TABLE manager(
   name VARCHAR(50) NOT NULL,
   cell_number VARCHAR(15) NOT NULL,
   cpf VARCHAR(11) PRIMARY KEY NOT NULL,
   email VARCHAR(30) NOT NULL,
   password VARCHAR(30) NOT NULL
);

CREATE TABLE doctor(
   name VARCHAR(50) NOT NULL,
   cell_number VARCHAR(15) NOT NULL,
   cpf VARCHAR(11) PRIMARY KEY NOT NULL,
   email VARCHAR(30) NOT NULL,
   password VARCHAR(30) NOT NULL,
   specialty VARCHAR(30) NOT NULL,
   crm VARCHAR(15) NOT NULL
);

CREATE TABLE medical_duty(
   date_duty DATE,
   type VARCHAR(20),
   cpf_doctor VARCHAR(11),
   CONSTRAINT cpf_doctor_fk FOREIGN KEY (cpf_doctor) REFERENCES doctor(cpf)
);