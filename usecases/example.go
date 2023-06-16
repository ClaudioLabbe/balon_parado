package usecases

import (
	"time"

	"github.com/ClaudioLabbe/balon_parado/interfaces"
	"github.com/ClaudioLabbe/balon_parado/models"
)

var metod interfaces.Metod

func Create() {
	// Creacion de Usuario
	var user models.User

	// Para crear usuario solo se necesita nombre y apellido
	user.AddUser("Felipe", "Calvo")
	time.Sleep(2 * time.Second)
	user.AddUser("Max", "Calvo")
	time.Sleep(2 * time.Second)
	metod = user.AddUser("Jorge", "Carrasco")
	time.Sleep(2 * time.Second)

	metod.List()

	//Creacion de Player
	player := models.Player{}
	// Para la creacion de un jugador se necesita nombre del usario creado anteriormente, Id de PS, posicion y dorsal
	metod = player.Add("Jorge", "Carrasco", "Jorge_1234", "MCD", 34)
	time.Sleep(2 * time.Second)

	metod.List()

	//Creacion de Club
	club := models.Club{}
	// Para la creacion de un club se necesita nombre del club, estadio, formacion
	club = club.Add("Sudacraks FC", "Metropolitano", "3-4-1-2")

	// Se agrega un jugador al club
	club.AddPlayerToClub("Jorge_1234", "Sudacraks FC")

	club.List()

	//Creacion de DT
	dt := models.Dt{}
	// Para la creacion de un DT se necesita nombre del usuario y apellido del usuario creado anteriormente y club
	dt.Add("Jorge", "Carrasco", models.Club{})
	time.Sleep(2 * time.Second)
	metod = dt.Update("Max", "Calvo", club)
	time.Sleep(2 * time.Second)

	dt.List()

	// Simular liga
	match := models.Match{}
	// Para simular liga se necesita el nombre de una liga ya creada, que se encuentre en el archivo League.json
	match.ExampleMatch("Liga chilena")

	//Creacion de liga
	league := models.League{}
	// solo se necesita el nombre de la liga
	league.Add("Liga EFA")

	// Se agrega el club al torneo
	league.AddClubToTournament("Liga EFA", "Sudacraks FC")
	league.List()
}
