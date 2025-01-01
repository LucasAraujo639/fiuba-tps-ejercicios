package comandos

import (
	"algogram/errores"
)

// Cantidad de elementos de cada comando al invocar:
const (
	CANT_PARAMETROS_COMANDO_SIN_PARAMETROS   = 0
	CANT_PARAMETROS_COMANDO_CON_UN_PARAMETRO = 1
	COMANDO_POS_CMD                          = 0
)

// verificarSinParametros comprueba que no tenga parametros
func verificarSinParametros(parametros []string) error {
	if len(parametros) != CANT_PARAMETROS_COMANDO_SIN_PARAMETROS {
		return errores.ErrorMalaInvocacionComando{}
	}

	return nil
}

// verificarUnParametros comprueba que tenga solo un parametro
func verificarUnParametro(parametros []string) (string, error) {

	if len(parametros) != CANT_PARAMETROS_COMANDO_CON_UN_PARAMETRO {
		return "", errores.ErrorMalaInvocacionComando{}
	}
	parametro := parametros[COMANDO_POS_CMD]
	return parametro, nil
}
