package datos

import (
	"os"

	"rerepolez/errores"
)

// AbrirArchivo intenta abrir el archivo con nombre pasado como argumento y devuelve un error según el criterio usitlizado.
func AbrirArchivo(nombreArchivo string) (*os.File, error) {
	archivoAbierto, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, errores.ErrorLeerArchivo{}
	}
	return archivoAbierto, nil
}
