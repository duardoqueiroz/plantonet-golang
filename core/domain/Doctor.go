package main

type Doctor struct {
	name       string `db:"name"`
	cellNumber string `db:"cell_number"`
	cpf        string `db:"cpf"`
	email      string `db:"email"`
	password   string `db:"password"`
	specialty  string `db:"specialty"`
	crm        string `db:"crm"`
}

func (d Doctor) getCpf() string {
	return d.cpf
}
func (d Doctor) getName() string {
	return d.name
}
func (d Doctor) getCellNumber() string {
	return d.cellNumber
}
func (d Doctor) getEmail() string {
	return d.email
}
func (d Doctor) getCrm() string {
	return d.crm
}
func (d Doctor) getSpecialty() string {
	return d.specialty
}
func (d Doctor) getPassword() string {
	return d.password
}
func (d Doctor) setName(name string) {
	d.name = name
}
func (d Doctor) setCpf(cpf string) {
	d.cpf = cpf
}
func (d Doctor) setEmail(email string) {
	d.email = email
}
func (d Doctor) setPassword(password string) {
	d.password = password
}
func (d Doctor) setSpecialty(specialty string) {
	d.specialty = specialty
}
func (d Doctor) setCrm(crm string) {
	d.crm = crm
}
