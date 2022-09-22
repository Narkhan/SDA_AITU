package main

func main() {
	hhkz := JobSite{}
	Bob := NewPerson("Bob")
	hhkz.addVacancy("Senior Golang Developer")
	hhkz.subscribe(Bob)
	hhkz.addVacancy("Junior Golang Developer")
	hhkz.removeVacancy("Senior Golang Developer")
}
