package main

type Manager struct {
	name       string `db:"nome"`
	cellNumber string `db:"cell_number"`
	cpf        string `db:"cpf"`
	email      string `db:"email"`
	password   string `db:"password"`
}


func (m Manager) getCpf() string{
	return m.cpf
}
func (m Manager) getName() string{
	return m.name
}
func (m Manager) getCellNumber() string{
	return m.cellNumber
}
func (m Manager) getEmail() string{
	return m.email
}
func (m Manager) getPassword() string{
	return m.password
}
func (m Manager) setName(name string){
	m.name = name
}
func (m Manager) setCpf(cpf string){
	m.cpf = cpf
}
func (m Manager) setEmail(email string){
	m.email = email
}
func (m Manager) setPassword(password string) {
	m.password = password
}
func (m Manager) setCellNumber(cellNumber string) {
	m.cellNumber = cellNumber
}