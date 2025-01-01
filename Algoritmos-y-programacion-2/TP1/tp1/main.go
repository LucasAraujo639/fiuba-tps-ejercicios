package main

import (
	"fmt"
	"os"

	"rerepolez/comandos"
	"rerepolez/datos"
	"rerepolez/salida"
	"rerepolez/votos"
)

func main() {

	// Obtener datos de invocación.
	params := datos.ObtenerParametrosEjecucion()
	err := datos.VerificarParametrosEjecucion(params)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	// Carga de datos.
	archivoPadronElectoral, err := datos.AbrirArchivo(datos.ObtenerNombreArchivoPadron(params))
	if err != nil {
		salida.ImprimirError(err)
		return
	}

	archivoListaPartidos, err := datos.AbrirArchivo(datos.ObtenerNombreArchivoLista(params))
	if err != nil {
		salida.ImprimirError(err)
		return
	}

	// Crear padrón electoral y contador de votos.
	padronElectoral := votos.CrearPadronElectoral()
	contadorVotos := votos.CrearContadorVotos()

	err = padronElectoral.CargarPadron(archivoPadronElectoral)
	if err != nil {
		salida.ImprimirError(err)
		return
	}

	err = contadorVotos.CargarListaAlternativas(archivoListaPartidos)
	if err != nil {
		salida.ImprimirError(err)
		return
	}

	// Leer comandos de la entrada estándar.
	comandos.LectorComandos(padronElectoral, contadorVotos)

	// Imprimir resultados.
	if err == nil {
		contadorVotos.ImprimirResultados(os.Stdout)
	} else {
		salida.ImprimirError(err)
	}
}
