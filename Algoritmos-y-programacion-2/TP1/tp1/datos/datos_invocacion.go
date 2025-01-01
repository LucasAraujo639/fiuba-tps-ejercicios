package datos

import (
	//"rerepolez/votos"

	"os"
	"rerepolez/errores"
)

// Posición del nombre del ejecutable en la línea de comandos al ejecutar el programa.
const NOMBRE_ARCHIVO_POS_CMD = 0

// Posiciones de cada elemento parámetro en la línea de comando sin contar el nombre del ejecutable.
const (
	ARCHIVO_LISTA_POS_CMD  = 0
	ARCHIVO_PADRON_POS_CMD = 1
	CANT_PARAM_CMD         = 2
)

func ObtenerParametrosEjecucion() []string {
	//Obtener parametros
	params := os.Args[NOMBRE_ARCHIVO_POS_CMD+1:]

	return params

}

func ObtenerNombreArchivoPadron(params []string) string {
	archivoPadron := params[ARCHIVO_PADRON_POS_CMD]
	return archivoPadron
}
func ObtenerNombreArchivoLista(params []string) string {
	archivoLista := params[ARCHIVO_LISTA_POS_CMD]
	return archivoLista
}

// VerificarParametrosEjecucion verifica que los parámetros de ejecución del programa son correctos.
func VerificarParametrosEjecucion(params []string) error {
	if len(params) < CANT_PARAM_CMD {
		return errores.ErrorParametros{}
	}

	return nil
}
