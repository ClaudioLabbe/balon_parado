package models

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type PorPuntos []Club

//se implementaron los métodos Len, Less y Swap para permitir el ordenamiento de la tabla de posiciones por puntos
func (p PorPuntos) Len() int           { return len(p) }
func (p PorPuntos) Less(i, j int) bool { return p[i].Points > p[j].Points }
func (p PorPuntos) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Match struct {
	Fecha   int
	Results [2]int
	Club    [2]Club
}

func (m Match) ExampleMatch(nameLiga string) {

	league := League{}
	// Se busca la liga
	league = league.MatchTournament(nameLiga)

	equipos := league.Club

	if len(equipos) == 16 {

		numEquipos := len(equipos)

		numFechas := (numEquipos - 1) * 2

		for fecha := 1; fecha <= numFechas; fecha++ {
			match := &Match{}
			fmt.Printf("Fecha %d:\n", fecha)

			//barajar el orden de los equipos en cada fecha
			rand.Shuffle(numEquipos, func(i, j int) {
				equipos[i], equipos[j] = equipos[j], equipos[i]
			})

			// Se utiliza una espera grupal (sync.WaitGroup) para asegurarse de que todas las gorroutines finalicen antes de continuar con la siguiente fecha.
			var wg sync.WaitGroup
			wg.Add(numEquipos / 2)

			for i := 0; i < numEquipos-1; i += 2 {
				go func(local, visitante *Club) {
					defer wg.Done()
					simularPartido(local, visitante, match, fecha)
				}(&equipos[i], &equipos[i+1])
			}
			wg.Wait()
			fmt.Println()
			time.Sleep(2 * time.Second)
		}

		//Ordenamiento
		sort.Sort(PorPuntos(equipos))
		time.Sleep(2 * time.Second)

		// Se lista la tabla de posiciones
		fmt.Println("Tabla de Posiciones Final:")
		for i, equipo := range equipos {
			fmt.Printf("%d. %s - Puntos: %d\n", i+1, equipo.Name, equipo.Points)
		}
	} else {
		fmt.Println("La liga no tiene el minimo (16) de equipos")
	}
}

// Esta función simula el resultado del partido y actualiza los puntos de los equipos correspondientes.
func simularPartido(local, visitante *Club, match *Match, fecha int) {
	randTime := rand.NewSource(time.Now().UnixNano())
	randomNew := rand.New(randTime)
	resultado := [2]int{randomNew.Intn(5), randomNew.Intn(5)}

	// Se agregan los puntos dependiendo del resultado
	if resultado[0] > resultado[1] {
		local.Points += 3
	} else if resultado[0] < resultado[1] {
		visitante.Points += 3
	} else if resultado[0] == resultado[1] {
		local.Points += 1
		visitante.Points += 1
	}

	match.Fecha = fecha
	match.Club[0] = *local
	match.Club[1] = *visitante
	match.Results[0] = resultado[0]
	match.Results[1] = resultado[1]

	fmt.Printf("%s vs %s - Resultado: %d-%d\n", local.Name, visitante.Name, resultado[0], resultado[1])
}
