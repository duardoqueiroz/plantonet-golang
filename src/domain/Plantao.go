package main

type Plantao struct{
	DataPlantao string	`db: "data_plantao"`
	tipo string			`db:"tipo"`
	CpfMedico string	`db: "cpf_medico"`
}