CREATE DATABASE plantonet;

use plantonet;

CREATE TABLE gerente(
   nome VARCHAR(50) NOT NULL,
   telefone VARCHAR(15) NOT NULL,
   cpf VARCHAR(11) PRIMARY KEY NOT NULL,
   email VARCHAR(30) NOT NULL,
   senha VARCHAR(30) NOT NULL
);

CREATE TABLE medico(
   nome VARCHAR(50) NOT NULL,
   telefone VARCHAR(15) NOT NULL,
   cpf VARCHAR(11) PRIMARY KEY NOT NULL,
   email VARCHAR(30) NOT NULL,
   senha VARCHAR(30) NOT NULL,
   especialidade VARCHAR(30) NOT NULL,
   crm VARCHAR(15) NOT NULL,
);

CREATE TABLE plantao(
   data_plantao DATE,
   tipo VARCHAR(20),
   cpf_medico VARCHAR(11),
   CONSTRAINT cpf_medico_fk FOREIGN KEY (cpf_medico) REFERENCES medico(cpf),
);