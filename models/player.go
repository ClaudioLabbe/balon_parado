package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

var user User

type Player struct {
	ID_PSN, Position string
	Dorsal           int
	User             User
}

func newFilePlayer(data Player) {

	path := "json/Player.json"

	file, err := ioutil.ReadFile(path)
	var p []Player
	if err != nil {
		p = []Player{data}
		newFile, _ := json.MarshalIndent(p, "", " ")

		ioutil.WriteFile(path, newFile, 0644)
	} else {

		err = json.Unmarshal(file, &p)
		if err != nil {
			log.Fatal(err)
		}

		p = append(p, data)

		newFile, _ := json.MarshalIndent(p, "", " ")

		ioutil.WriteFile(path, newFile, 0644)
	}

	fmt.Println("Player guardado")
}

func (player *Player) Add(name, lastName, id_psn, position string, dorsal int) Player {

	var userFile User

	for _, v := range user.ReadFile() {
		if v.Name == name && v.LastName == lastName {
			userFile = v
		}
	}

	if userFile.Name == "" {
		log.Fatal("Usuario no existe")
	}

	player = &Player{
		ID_PSN:   id_psn,
		Position: position,
		Dorsal:   dorsal,
		User:     userFile,
	}

	var playerList []Player = filePlayer()

	for _, v := range playerList {
		if v.ID_PSN == id_psn {
			fmt.Println("Player ya existe")
			return *player
		}
	}

	newFilePlayer(*player)

	return *player
}

func filePlayer() []Player {

	path := "json/Player.json"

	file, err := ioutil.ReadFile(path)

	var playersList []Player

	if err != nil {
		fmt.Println("Creando registro Player")
		return playersList
	} else {

		err = json.Unmarshal(file, &playersList)
		if err != nil {
			log.Fatal(err)
		}
	}

	return playersList

}

func (player Player) List() {
	var playerList []Player = filePlayer()

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Printf(lineInicio, "Listado de jugadores")
	fmt.Fprintln(w, "Nombre\tApellido\tID_PSN\tPosicion\tDorsal")
	for _, v := range playerList {
		fmt.Fprintln(w, v.User.Name+"\t"+v.User.LastName+"\t"+v.ID_PSN+"\t"+v.Position+"\t"+strconv.Itoa(v.Dorsal)+"\t")
	}
	w.Flush()

	fmt.Println(lineFin)
}
