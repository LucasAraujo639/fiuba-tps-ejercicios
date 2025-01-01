package votos

import "os"

type ContadorVotos interface {

	// CargarListaAlternativas carga la lista presentada por los partidos para la votación.
	CargarListaAlternativas(*os.File) error

	// CantidadAlternativas dice la cantidad de partidos que se presentaron en las elecciones.
	CantidadAlternativas() int

	// SumarVotos suma los votos correspondientes a cada candidato según el voto realizado por el votante.
	SumarVotos(Voto)

	// MostrarResultados muestra los votos del candidato solicitado.
	// La alternativa 0 se reserva para los votos en blanco.
	ImprimirResultados(*os.File)
}
