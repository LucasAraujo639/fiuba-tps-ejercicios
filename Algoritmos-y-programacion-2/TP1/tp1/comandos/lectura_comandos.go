package comandos

import (
	"bufio"
	"os"
	"rerepolez/errores"
	"rerepolez/salida"
	"rerepolez/votos"
	"strings"
	TDACola "tdas/cola"
)

// Comandos posibles:
const (
	COMANDO_INGRESAR_VOTANTE     = "ingresar"
	COMANDO_REALIZAR_VOTO        = "votar"
	COMANDO_DESHACER_ULTIMO_VOTO = "deshacer"
	COMANDO_FINALIZAR_VOTO       = "fin-votar"
)

// Tipo enumerativo de comando:
type Comando int

const (
	INGRESAR Comando = iota
	VOTAR
	DESHACER
	FIN_VOTO
	ERROR
)

// LectorEntradaEstandar realiza el ingreso de los comandos por la entrada estándar del programa.
func LectorComandos(padronElectoral votos.PadronElectoral, contadorVotos votos.ContadorVotos) {
	var err error
	linea := ""
	var votanteActual votos.Votante = nil
	colaVotantes := TDACola.CrearColaEnlazada[int]()

	reader := bufio.NewReader(os.Stdin)

	for {
		// posible error
		linea, err = reader.ReadString('\n')
		if err != nil {
			// Verifica si se llego al final del archivo
			if err.Error() == "EOF" {
				break
			}

		}

		entradaSeparada := separarLineaEntrada(linea)
		comandoAEjecutar := obtenerComando(entradaSeparada)
		parametros := entradaSeparada[COMANDO_A_EJECUTAR_POS_CMD+1:]

		switch comandoAEjecutar {
		case INGRESAR:
			dni := 0

			// Verificar parámetros de comando para ingresar votante.
			dni, err = verificarParamsIngresarVotante(parametros)
			if err != nil {
				break
			}

			// Verificar que el DNI ingresado esté en el padrón.
			if !padronElectoral.PerteneceAlPadron(dni) {
				err = errores.DNIFueraPadron{}
				break
			}

			// Verificar que no haya votado.
			/*if padronElectoral.YaVoto(dni) {
				err = errores.ErrorVotanteFraudulento{Dni: dni}
				break
			}*/

			// Encolar al DNI si ya hay un votante.
			if votanteActual != nil {
				colaVotantes.Encolar(dni)
				break
			}

			// Si no hay un votante, desencolar a algún votante de la fila.
			// Si no hay nadie, vota quien haya ingresado.
			if !colaVotantes.EstaVacia() {
				dni = colaVotantes.Desencolar()
			}

			votanteActual = votos.CrearVotante(dni)

			err = nil

		case VOTAR:
			var tipo votos.TipoVoto
			alternativa := -1

			// Verificar si hay algún votante.
			if votanteActual == nil {
				err = errores.FilaVacia{}
				break
			}

			// Verificar que no haya votado.
			if padronElectoral.YaVoto(votanteActual.LeerDNI()) {
				err = errores.ErrorVotanteFraudulento{Dni: votanteActual.LeerDNI()}
				// Hace que vote la siguiente persona en la fila, si la hay.
				votanteActual = generarSiguienteVotante(colaVotantes)
				break
			}

			// Verificar formato de parámetros para realizar un voto.
			tipo, alternativa, err = verificarParamsRealizarVoto(parametros)
			if err != nil {
				break
			}

			// Verificar que exista la alternativa deseada.
			if alternativa > contadorVotos.CantidadAlternativas() {
				err = errores.ErrorAlternativaInvalida{}
				break
			}

			// Realizar voto.
			err = EjecutarComandoRealizarVoto(votanteActual, tipo, alternativa)

		case DESHACER:
			if votanteActual == nil {
				err = errores.FilaVacia{}
				break
			}

			// Verificar que no haya votado.
			if padronElectoral.YaVoto(votanteActual.LeerDNI()) {
				err = errores.ErrorVotanteFraudulento{Dni: votanteActual.LeerDNI()}
				// Hace que vote la siguiente persona en la fila, si la hay.
				votanteActual = generarSiguienteVotante(colaVotantes)
				break
			}

			// Verificar parámetros.
			err = verificarParamsDeshacerVoto(parametros)

			if err != nil {
				break
			}

			// Verificar que no haya votado.
			if padronElectoral.YaVoto(votanteActual.LeerDNI()) {
				err = errores.ErrorVotanteFraudulento{Dni: votanteActual.LeerDNI()}
				break
			}
			err = EjecutarComandoDeshacerVoto(votanteActual)

		case FIN_VOTO:
			var votoRealizado votos.Voto

			if votanteActual == nil {
				err = errores.FilaVacia{}
				break
			}

			// Verificar que no haya votado.
			if padronElectoral.YaVoto(votanteActual.LeerDNI()) {
				err = errores.ErrorVotanteFraudulento{Dni: votanteActual.LeerDNI()}
				break
			}

			// Verificar parámetros.
			err = verificarParamsFinalizarVoto(parametros)
			if err != nil {
				break
			}

			votoRealizado = EjecutarComandoFinalizarVoto(votanteActual)

			// Sumar votos del votante.
			contadorVotos.SumarVotos(votoRealizado)

			// Establecer que ya voto.
			padronElectoral.FinVoto(votanteActual.LeerDNI())

			// Hace que vote la siguiente persona en la fila, si la hay.
			votanteActual = generarSiguienteVotante(colaVotantes)

		default:
			err = errores.ErrorComandoDesconocido{}
		}

		// Imprimir resultado.
		if err != nil {
			salida.ImprimirError(err)
		} else {
			salida.ImprimirComandoOK()
		}
	}

	if votanteActual != nil {
		salida.ImprimirError(errores.ErrorCiudadanosSinVotar{})
	}
}

func obtenerComando(campos []string) Comando {
	switch campos[COMANDO_A_EJECUTAR_POS_CMD] {
	case COMANDO_INGRESAR_VOTANTE:
		return INGRESAR
	case COMANDO_REALIZAR_VOTO:
		return VOTAR
	case COMANDO_DESHACER_ULTIMO_VOTO:
		return DESHACER
	case COMANDO_FINALIZAR_VOTO:
		return FIN_VOTO
	}
	return ERROR
}

// SepararLineaEntrada realiza todas las acciones necesarias para separar la línea de comando en sus respectivos campos
// según el protocolo definido.
func separarLineaEntrada(linea string) []string {
	// Quitar espacios en el inicio y el fin si los hubiera.
	lineaSinSaltoDeLinea := strings.TrimSpace(linea)

	//Dividir la linea por cada espacio que contenga.
	partes := strings.Fields(lineaSinSaltoDeLinea)

	return partes
}

// generarSiguienteVotante devuelve el siguiente votante en la fila, si lo hay.
func generarSiguienteVotante(colaVotantes TDACola.Cola[int]) votos.Votante {
	if !colaVotantes.EstaVacia() {
		return votos.CrearVotante(colaVotantes.Desencolar())
	}
	return nil
}
