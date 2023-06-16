package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"
)

var path string = "json/Dt.json"

type Dt struct {
	User User
	Club Club
}

func newFile(data Dt) {

	file, err := ioutil.ReadFile(path)
	var o, dts []Dt
	if err != nil {
		o = []Dt{data}
		newFile, _ := json.MarshalIndent(o, "", " ")

		ioutil.WriteFile(path, newFile, 0644)
	} else {

		err = json.Unmarshal(file, &o)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range o {
			if v.User.Name != data.User.Name && v.User.LastName != data.User.LastName {
				dts = append(dts, v)
			}
		}

		o = append(dts, data)

		newFile, _ := json.MarshalIndent(o, "", " ")

		ioutil.WriteFile(path, newFile, 0644)
	}

}

func (dt Dt) Add(name, lastName string, club Club) Dt {

	var userFile User

	for _, v := range user.ReadFile() {
		if v.Name == name && v.LastName == lastName {
			userFile = v
		}
	}

	if userFile.Name == "" {
		fmt.Println("Usuario no existe")
		return dt
	}

	dt = Dt{
		User: userFile,
		Club: club,
	}

	var dtList []Dt = ReadFileDt()

	for _, v := range dtList {
		if v.User.Name == name {
			fmt.Println("DT ya existe")
			return dt
		}
	}

	newFile(dt)

	return dt
}

func (dt Dt) Update(name, lastName string, club Club) Dt {

	dt = Dt{}

	var dtList []Dt = ReadFileDt()

	for _, v := range dtList {
		if v.User.Name == name && v.User.LastName == lastName {
			fmt.Println("DT ya existe")
			dt = v
		}
	}

	dt.Club = club

	newFile(dt)

	return dt
}

func ReadFileDt() []Dt {

	file, err := ioutil.ReadFile(path)

	var dtList []Dt

	if err != nil {
		fmt.Println("Creando registro de DT")
		return dtList
	} else {

		err = json.Unmarshal(file, &dtList)
		if err != nil {
			log.Fatal(err)
		}
	}

	return dtList
}

func (dt Dt) List() {

	var dts []Dt = ReadFileDt()

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Printf(lineInicio, "Listado de DTs")
	fmt.Fprintln(w, "Nombres\tApellidos\tClub\t")
	for _, v := range dts {
		fmt.Fprintln(w, v.User.Name+"\t"+v.User.LastName+"\t"+v.Club.Name+"\t")
	}

	w.Flush()
	fmt.Println(lineFin)

}
