package main

type Gerente struct{
	Nome string		`db: "nome"`
	Telefone string `db: "telefone"`
	Cpf 	 string `db: "cpf"`
	Email string	`db: "email"`
	Senha string	`db: "senha"`
}
