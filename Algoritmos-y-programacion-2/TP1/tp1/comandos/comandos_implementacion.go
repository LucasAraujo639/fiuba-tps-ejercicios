package comandos

import (
	//"fmt"
	//"os"
	"strconv"

	"rerepolez/errores"
	"rerepolez/votos"

	TDACola "tdas/cola"
)

// Cantidad de elementos de cada comando al invocar:
const (
	CANT_PARAMETROS_COMANDO_INGRESAR = 1
	CANT_PARAMETROS_COMANDO_VOTAR    = 2
	CANT_PARAMETROS_COMANDO_DESHACER = 0
	CANT_PARAMETROS_COMANDO_FIN_VOTO = 0
)

// Posición de comando a ejecutar en entrada:
const COMANDO_A_EJECUTAR_POS_CMD int = 0

// -------------------------------------------------------------------------------------------------------------
//                                 Comando de ingreso de votante nuevo.
// -------------------------------------------------------------------------------------------------------------

// Posiciones de parámetros del comando para ingresar votante sin contar el nombre del comando:
const (
	DNI_INGRESADO_POS_CMD = 0
)

// EjecutarComandoIngresar
func EjecutarComandoIngresar(dni string, padrones []votos.Votante, miCola TDACola.Cola[votos.Votante]) {

}

// verificarParamsIngresarVotante comprueba que los parámetros pasados por argumento sean válidos.
func verificarParamsIngresarVotante(parametros []string) (int, error) {
	if len(parametros) != CANT_PARAMETROS_COMANDO_INGRESAR {
		return 0, errores.ErrorMalaInvocacionComando{}
	}

	dni, err := strconv.Atoi(parametros[DNI_INGRESADO_POS_CMD])

	if err != nil || dni <= 0 {
		return 0, errores.DNIError{}
	}
	return dni, nil
}

// -------------------------------------------------------------------------------------------------------------
//                                    Comando de realización de voto.
// -------------------------------------------------------------------------------------------------------------

// Posiciones de parámetros del comando para realizar voto sin contar el nombre del comando:
const (
	TIPO_VOTO_POS_CMD          = 0
	ALTERNATIVA_VOTADA_POS_CMD = 1
)

const (
	PARAMETRO_TIPO_VOTO_PRESIDENTE = "Presidente"
	PARAMETRO_TIPO_VOTO_GOBERNADOR = "Gobernador"
	PARAMETRO_TIPO_VOTO_INTENDENTE = "Intendente"
)

// EjecutarComandoRealizarVoto
func EjecutarComandoRealizarVoto(votante votos.Votante, tipo votos.TipoVoto, alternativa int) error {
	return votante.Votar(tipo, alternativa)
}

func verificarParamsRealizarVoto(parametros []string) (votos.TipoVoto, int, error) {
	var tipo votos.TipoVoto

	// Verificar cantidad de parámetros de entrada.
	if len(parametros) != CANT_PARAMETROS_COMANDO_VOTAR {
		return votos.PRESIDENTE, 0, errores.ErrorMalaInvocacionComando{}
	}

	// Verificar que la alternativa sea válida.
	alternativa, err := strconv.Atoi(parametros[ALTERNATIVA_VOTADA_POS_CMD])
	if err != nil || alternativa < 0 {
		return tipo, alternativa, errores.ErrorAlternativaInvalida{}
	}

	// Verificar que el tipo de dato sea el correcto.
	switch parametros[TIPO_VOTO_POS_CMD] {
	case PARAMETRO_TIPO_VOTO_PRESIDENTE:
		tipo = votos.PRESIDENTE
	case PARAMETRO_TIPO_VOTO_GOBERNADOR:
		tipo = votos.GOBERNADOR
	case PARAMETRO_TIPO_VOTO_INTENDENTE:
		tipo = votos.INTENDENTE
	default:
		err = errores.ErrorTipoVoto{}
	}

	return tipo, alternativa, err
}

// -------------------------------------------------------------------------------------------------------------
//                              Comando para deshacer el último voto realizado.
// -------------------------------------------------------------------------------------------------------------

// EjecutarComandoDeshacerVoto cambia el último voto realizado por el votante a su estado anterior.
func EjecutarComandoDeshacerVoto(votante votos.Votante) error {
	if votante == nil {
		return errores.ErrorNoHayVotosAnteriores{}
	}
	return votante.Deshacer()
}

func verificarParamsDeshacerVoto(parametros []string) error {
	// Verificar cantidad de parámetros de entrada.
	if len(parametros) != CANT_PARAMETROS_COMANDO_DESHACER {
		return errores.ErrorMalaInvocacionComando{}
	}
	return nil
}

// -------------------------------------------------------------------------------------------------------------
//                                    Comando de finalización de voto.
// -------------------------------------------------------------------------------------------------------------

// EjecutarComandoFinalizarVoto marca que el votante ha finalizado su voto.
func EjecutarComandoFinalizarVoto(votante votos.Votante) votos.Voto {
	return votante.FinVoto()
}

func verificarParamsFinalizarVoto(parametros []string) error {
	// Verificar cantidad de parámetros de entrada.
	if len(parametros) != CANT_PARAMETROS_COMANDO_FIN_VOTO {
		return errores.ErrorMalaInvocacionComando{}
	}
	return nil
}
