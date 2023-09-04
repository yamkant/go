package main

import (
	"html/template"
	"os"
)

type User struct {
	Name	string
	Email	string
	Age		int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

// NOTE: string template
func mainV1() {
	user := User{Name:"tester", Email: "tester@example.com", Age: 23}
	user2 := User{Name:"tester2", Email: "tester2@example.com", Age: 18} 
	tmpl, err := template.New("Tmpl1").Parse("Name: {{.Name}}\nEmail: {{.Email}}\nAge: {{.Age}}")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, user)
	tmpl.Execute(os.Stdout, user2)
}

// NOTE: file template
func mainV2() {
	user := User{Name:"tester", Email: "tester@example.com", Age: 23}
	user2 := User{Name:"tester2", Email: "tester2@example.com", Age: 40} 
	fileDir := "templates/"
	fileName := "tmpl1.tmpl"
	tmpl, err := template.New("Tmpl1").ParseFiles(fileDir + fileName)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, fileName, user)
	tmpl.ExecuteTemplate(os.Stdout, fileName, user2)
}

// NOTE: nested file template
func mainV3() {
	user := User{Name:"tester", Email: "tester@example.com", Age: 23}

	fileDir := "templates/"
	fileName := "tmpl1.tmpl"
	fileName2 := "tmpl2.tmpl"

	tmpl, err := template.New("Tmpl1").ParseFiles(fileDir + fileName, fileDir + fileName2)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, fileName2, user)
}

// NOTE: nested file template with list data
func main() {
	user1 := User{Name:"tester", Email: "tester@example.com", Age: 23}
	user2 := User{Name:"tester2", Email: "tester2@example.com", Age: 40} 
	users := []User{user1, user2}

	fileDir := "templates/"
	fileName := "tmpl1.tmpl"
	fileName3 := "tmpl3.tmpl"

	tmpl, err := template.New("Tmpl1").ParseFiles(fileDir + fileName, fileDir + fileName3)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, fileName3, users)
}