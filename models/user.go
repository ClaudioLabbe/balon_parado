package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"
)

const (
	pathUser = "json/User.json"
)

type User struct {
	Name, LastName string
}

func newFileUser(data User) {

	file, err := ioutil.ReadFile(pathUser)
	var o []User
	if err != nil {
		o = []User{data}
		newFile, _ := json.MarshalIndent(o, "", " ")

		ioutil.WriteFile(pathUser, newFile, 0644)
	} else {

		err = json.Unmarshal(file, &o)
		if err != nil {
			log.Fatal(err)
		}

		o = append(o, data)

		newFile, _ := json.MarshalIndent(o, "", " ")

		ioutil.WriteFile(pathUser, newFile, 0644)
	}

}

func (user User) AddUser(name, lastName string) User {

	var userList []User = user.ReadFile()

	for _, v := range userList {
		if v.Name == name && v.LastName == lastName {
			fmt.Println("Usuario ya existe")
			return user
		}
	}

	user = User{
		Name:     name,
		LastName: lastName,
	}

	newFileUser(user)

	return user
}

func (user User) ReadFile() []User {
	file, err := ioutil.ReadFile(pathUser)

	var userList []User

	if err != nil {
		fmt.Println("Creando registro Usuario")
		return userList
	} else {

		err = json.Unmarshal(file, &userList)
		if err != nil {
			log.Fatal(err)
		}
	}

	return userList
}

func (user User) List() {

	var userList []User = user.ReadFile()

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Printf(lineInicio, "Listado de usuarios")
	fmt.Fprintln(w, "Nombres\tApellidos\t")
	for _, v := range userList {
		fmt.Fprintln(w, "\t"+v.Name+"\t"+v.LastName+"\t")
	}
	w.Flush()

	fmt.Println(lineFin)

}
