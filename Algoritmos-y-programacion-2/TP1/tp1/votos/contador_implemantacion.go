package votos

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"rerepolez/errores"
)

type contadorVotosImplementacion struct {
	votosImpugnados int
	listaPartidos   *[]Partido
}

const (
	NOMBRE_PARTIDO_POS    = 0
	NOMBRE_PRESIDENTE_POS = 1
	NOMBRE_GOBERNADOR_POS = 2
	NOMBRE_INTENDENTE_POS = 3
)

func CrearContadorVotos() ContadorVotos {
	return new(contadorVotosImplementacion)
}

func (contador *contadorVotosImplementacion) CargarListaAlternativas(archivo *os.File) error {
	if archivo == nil {
		return errores.ErrorLeerArchivo{}
	}
	contador.listaPartidos = obtenerListaPartidos(archivo)
	return nil
}

func (contador *contadorVotosImplementacion) CantidadAlternativas() int {
	return len(*contador.listaPartidos) - 1
}

func (contador *contadorVotosImplementacion) sumarVotoCandidato(tipo TipoVoto, alternativa int) {
	(*contador.listaPartidos)[alternativa].VotadoPara(tipo)
}

func (contador *contadorVotosImplementacion) SumarVotos(votoRealizado Voto) {
	if votoRealizado.Impugnado {
		contador.votosImpugnados++
		return
	}

	for tipo := PRESIDENTE; tipo < CANT_VOTACION; tipo++ {
		alternativa := votoRealizado.VotoPorTipo[tipo]
		contador.sumarVotoCandidato(tipo, alternativa)
	}
}

func (contador *contadorVotosImplementacion) ImprimirResultados(archivoSalida *os.File) {
	fmt.Fprintf(archivoSalida, "%s\n", MENSAJE_VOTOS_PRESIDENTE)
	contador.imprimirResultadosTipo(archivoSalida, PRESIDENTE)
	fmt.Fprint(archivoSalida, "\n")

	fmt.Fprintf(archivoSalida, "%s\n", MENSAJE_VOTOS_GOBERNADOR)
	contador.imprimirResultadosTipo(archivoSalida, GOBERNADOR)
	fmt.Fprint(archivoSalida, "\n")

	fmt.Fprintf(archivoSalida, "%s\n", MENSAJE_VOTOS_INTENDENTE)
	contador.imprimirResultadosTipo(archivoSalida, INTENDENTE)
	fmt.Fprint(archivoSalida, "\n")

	contador.imprimirVotosImpugnados(archivoSalida)
}

func (contador *contadorVotosImplementacion) imprimirResultadosTipo(archivoSalida *os.File, tipo TipoVoto) {
	for i := 0; i < len(*contador.listaPartidos); i++ {
		fmt.Fprintf(archivoSalida, "%s\n", (*contador.listaPartidos)[i].ObtenerResultado(tipo))
	}
}

func (contador *contadorVotosImplementacion) imprimirVotosImpugnados(archivoSalida *os.File) {
	var mensajeVoto string
	if contador.votosImpugnados == 1 {
		mensajeVoto = MENSAJE_VOTO_SINGULAR
	} else {
		mensajeVoto = MENSAJE_VOTO_PLURAL
	}
	fmt.Fprintf(archivoSalida, "Votos Impugnados: %d %s\n", contador.votosImpugnados, mensajeVoto)
}

// ObtenerListaPartidos devuelve un puntero a un slice de Partidos
func obtenerListaPartidos(archivoLista *os.File) *[]Partido {
	defer archivoLista.Close()

	listaPartidos := make([]Partido, 1)

	// Agregar partido de votos en blanco.
	listaPartidos[0] = CrearVotosEnBlanco()

	// Agregar partidos de la lista electoral.
	escaner := bufio.NewScanner(archivoLista)

	candidatosPartidos := [3]string{}

	for escaner.Scan() {
		linea := escaner.Text()

		// Separa los datos en campos.
		lineaSeparada := strings.Split(linea, ",")

		// Guardan los nombres de los candidatos.
		candidatosPartidos[NOMBRE_PRESIDENTE_POS-1] = lineaSeparada[NOMBRE_PRESIDENTE_POS]
		candidatosPartidos[NOMBRE_GOBERNADOR_POS-1] = lineaSeparada[NOMBRE_GOBERNADOR_POS]
		candidatosPartidos[NOMBRE_INTENDENTE_POS-1] = lineaSeparada[NOMBRE_INTENDENTE_POS]
		partidoNuevo := CrearPartido(lineaSeparada[NOMBRE_PARTIDO_POS], candidatosPartidos)
		listaPartidos = append(listaPartidos, partidoNuevo)
	}

	return &listaPartidos
}
