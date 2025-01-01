package datos

import (
	"algogram/errores"
	"os"
)

const (
	NOMBRE_ARCHIVO_POS_CMD      = 0
	ARCHIVO_USUARIOS_POS_PARAMS = 0
	CANT_PARAM_CMD              = 1
)

func ObtenerParametrosEjecucion() []string {
	//Obtener parametros
	params := os.Args[NOMBRE_ARCHIVO_POS_CMD+1:]
	return params

}

func ObtenerNombreArchivoUsuarios(params []string) string {
	archivoUsuarios := params[ARCHIVO_USUARIOS_POS_PARAMS]
	return archivoUsuarios
}

// VerificarParametrosEjecucion verifica que los parámetros de ejecución del programa son correctos.
func VerificarParametrosEjecucion(params []string) error {
	if len(params) < CANT_PARAM_CMD {
		return errores.ErrorParametros{}
	}

	return nil
}
