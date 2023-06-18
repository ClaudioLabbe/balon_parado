package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

const pathClub string = "json/Club.json"

type Club struct {
	Points                        int
	Name, Stadium, Team_formation string
	Players                       []Player
}

func newFileClub(data Club) {

	file, err := ioutil.ReadFile(pathClub)
	var c, cTrue []Club
	if err != nil {
		cTrue = []Club{data}
	} else {

		err = json.Unmarshal(file, &c)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range c {
			if v.Name != data.Name {
				cTrue = append(cTrue, v)
			}
		}
		cTrue = append(cTrue, data)

	}

	newFile, _ := json.MarshalIndent(cTrue, "", " ")

	ioutil.WriteFile(pathClub, newFile, 0644)

}

func readFileClub() []Club {
	file, err := ioutil.ReadFile(pathClub)

	var clubList []Club

	if err != nil {
		return clubList
	} else {

		err = json.Unmarshal(file, &clubList)
		if err != nil {
			log.Fatal(err)
		}
	}

	return clubList
}

func (club Club) Add(name, stadium, team_formation string) Club {

	fmt.Println("* Creando registros de Club")

	var clubList []Club = readFileClub()

	for _, v := range clubList {
		if v.Name == name {
			fmt.Println("Club ya existe")
			club = v
			return club
		}
	}

	club = Club{
		Name:           name,
		Stadium:        stadium,
		Team_formation: team_formation,
	}

	newFileClub(club)

	fmt.Println("Club guardado correctamente!")

	return club
}

func (club Club) AddPlayerToClub(id_psn, nameClub string) Club {
	fmt.Println("Agregando jugador al club")
	var player Player

	for _, v := range filePlayer() {
		if v.ID_PSN == id_psn {
			fmt.Println("Player encontrado")
			player = v
		}
	}

	if player == (Player{}) {
		fmt.Println("player no encontrado")
		return club
	}

	for _, v := range readFileClub() {
		if v.Name == nameClub {
			club = v
		}
	}

	for _, v := range club.Players {
		if v.ID_PSN == id_psn {
			fmt.Println("Jugador ya existe en el plantel de " + club.Name)
			return club
		}
	}

	fmt.Println(club)
	if club.Name == "" {
		fmt.Println("Club no encontrado")
		return club
	}

	club.Players = append(club.Players, player)
	newFileClub(club)

	return club

}

func (club Club) List() {

	for _, v := range readFileClub() {

		fmt.Println("--------------- CLUB --------------")
		fmt.Println("Nombre: " + v.Name)
		fmt.Println("Estadio: " + v.Stadium)
		fmt.Println("Formacion: " + v.Team_formation)

		for i, p := range v.Players {
			fmt.Printf("-------- Jugador: %s --------\n", strconv.Itoa(i+1))
			fmt.Println("Nombre: " + p.User.Name)
			fmt.Println("Apellido: " + p.User.LastName)
			fmt.Println("ID_PSN: " + p.ID_PSN)
			fmt.Println("Posicion: " + p.Position)
			fmt.Println("Dorsal: " + strconv.Itoa(p.Dorsal))
			fmt.Println("----------------")
		}

		fmt.Println("-----------------------------------")
	}
}
