package votos

import "os"

type PadronElectoral interface {

	// CargarPadron carga los datos almacenados en un archivo pasado por argumento.
	// Ocurre un error cuando el archivo no es válido.
	CargarPadron(*os.File) error

	// PerteneceAPadron indica si un DNI pertenece o no al padrón electoral.
	// Se asume que el DNI es correcto de antemano. No se garantiza el correcto funcionamiento para un DNI incorrecto.
	PerteneceAlPadron(int) bool

	// YaVoto verifica si el DNI indicado ya emitió su voto.
	// Se asume que el DNI es correcto de antemano. No se garantiza el correcto funcionamiento para un DNI incorrecto.
	YaVoto(int) bool

	// FinVoto almacena que el DNI ingresado ya emitió su voto.
	FinVoto(int) error
}
