package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const pathLeague = "json/League.json"

type League struct {
	Name string
	Club []Club
}

func newFileTournament(data League) {

	file, err := ioutil.ReadFile(pathLeague)
	var t []League
	var tOficial []League
	if err != nil {
		t = []League{data}

	} else {

		err = json.Unmarshal(file, &t)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range t {
			if v.Name != data.Name {
				tOficial = append(tOficial, v)
			}
		}

		t = append(tOficial, data)

	}

	newFile, _ := json.MarshalIndent(t, "", " ")

	ioutil.WriteFile(pathLeague, newFile, 0644)

}

func validateTournament() []League {

	file, err := ioutil.ReadFile(pathLeague)

	var leagueList []League

	if err != nil {
		fmt.Println("No se encontro registro")
		return leagueList
	} else {

		err = json.Unmarshal(file, &leagueList)
		if err != nil {
			log.Fatal(err)
		}
	}

	return leagueList
}

func (league League) Add(name string) League {

	for _, v := range validateTournament() {
		if v.Name == name {
			fmt.Println("Torneo ya existe")
			return league
		}
	}

	league = League{
		Name: name,
	}

	newFileTournament(league)

	return league
}

func (league League) AddClubToTournament(name, nameClub string) League {

	for _, v := range validateTournament() {
		if v.Name == name {

			league = v
		}
	}

	var club Club

	for _, c := range readFileClub() {
		if c.Name == nameClub {
			club = c
			fmt.Println("club encontrado")
		}
	}

	if club.Name == "" {
		fmt.Println("club no encontrado")
		return league
	}

	for _, v := range league.Club {
		if v.Name == club.Name {
			fmt.Println("club ya se encuentra registrado")
			return league
		}
	}

	league.Club = append(league.Club, club)

	newFileTournament(league)

	return league
}

func (league League) List() {
	file, err := ioutil.ReadFile(pathLeague)

	var t []League

	if err != nil {
		log.Fatal(err)
	} else {

		err = json.Unmarshal(file, &t)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("------ Listado de Torneos ------")
		for _, v := range t {
			fmt.Println("Nombre: " + v.Name)
			fmt.Println("----------------------------")
		}

		fmt.Println("------ Fin Listado de Torneos ------")
	}
}

func (league League) MatchTournament(nameLiga string) League {
	for _, t := range validateTournament() {
		if t.Name == nameLiga {
			league = t
		}
	}

	return league

}
