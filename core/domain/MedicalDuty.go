package main

type MedicalDuty struct {
	dateDuty  string `db:"date_duty"`
	typeDuty  string `db:"type"`
	cpfDoctor string `db:"cpf_doctor"`
}

func (md MedicalDuty) getDate() string {
	return md.dateDuty
}
func (md MedicalDuty) getType() string {
	return md.typeDuty
}
func (md MedicalDuty) getCpfDoctor() string {
	return md.cpfDoctor
}

func (md MedicalDuty) setDate(date string) {
	md.dateDuty = date
}
func (md MedicalDuty) setType(typ string) {
	md.typeDuty = typ
}
func (md MedicalDuty) setCpfDoctor(cpf string) {
	md.cpfDoctor = cpf
}
