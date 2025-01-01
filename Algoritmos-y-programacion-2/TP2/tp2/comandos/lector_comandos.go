package comandos

import (
	TDABD "algogram/bd"
	"algogram/errores"
	"algogram/salida"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Comandos posibles:
const (
	COMANDO_LOGIN              = "login"
	COMANDO_LOGOUT             = "logout"
	COMANDO_PUBLICAR           = "publicar"
	COMANDO_VER_SIGUIENTE_FEED = "ver_siguiente_feed"
	COMANDO_LIKEAR_POST        = "likear_post"
	COMANDO_MOSTRAR_LIKES      = "mostrar_likes"
	COMANDO_A_EJECUTAR_POS_CMD = 0
)

type Comando int

const (
	LOGIN Comando = iota
	LOGOUT
	PUBLICAR
	VER_SIGUIENTE_FEED
	LIKEAR_POST
	MOSTRAR_LIKES
	ERROR
)

func LectorComandos(usuariosBD TDABD.UsuariosBD, postsBD TDABD.PostsBD) {
	var err error
	linea := ""
	reader := bufio.NewReader(os.Stdin)

	for {
		linea, err = reader.ReadString('\n')
		if err != nil {
			// 	Verifica si se llego al final del archivo
			if err.Error() == "EOF" {
				break
			}

		}

		entradaSeparada := separarLineaEntrada(linea)
		comandoAEjecutar := obtenerComando(entradaSeparada)
		parametros := entradaSeparada[COMANDO_A_EJECUTAR_POS_CMD+1:]

		switch comandoAEjecutar {
		case LOGIN:
			nombreUsuario := strings.Join(parametros, " ")

			err := usuariosBD.Login(nombreUsuario)
			if err != nil {
				fmt.Println(err)
				break
			}

			usuarioLogueado, err := usuariosBD.ObtenerUsuario(nombreUsuario)
			if err != nil {
				fmt.Println(err)
				break
			}
			salida.ImprimirLogin(usuarioLogueado.VerNombre())

		case LOGOUT:
			err := verificarSinParametros(parametros)
			if err != nil {
				fmt.Println(err)
				break
			}

			err = usuariosBD.Logout()
			if err != nil {
				fmt.Println(err)
				break
			}
			salida.ImprimirLogout()

		case PUBLICAR:
			usuarioLogueado, err := usuariosBD.ObtenerLogueado()
			if err != nil {
				fmt.Println(err)
				break
			}
			textoPublicacion := strings.Join(parametros, " ")
			postNuevo := usuarioLogueado.CrearPost(postsBD.Cantidad(), textoPublicacion)

			// Actualizo la base de datos de posts con el nuevo post
			postsBD.GuardarPost(postNuevo)
			//Actualizo la base de datos usuarios con el nuevo post
			usuariosBD.ActualizarFeed(postNuevo)

			salida.ImprimirPostPublicado()

		case VER_SIGUIENTE_FEED:
			err := verificarSinParametros(parametros)
			if err != nil {
				fmt.Println(err)
				break
			}
			usuarioLogueado, err := usuariosBD.ObtenerLogueado()
			if err != nil {
				fmt.Println(errores.ErrorUsuarioNoLoggeadoONoMasPosts{})
				break
			}
			post, err := usuarioLogueado.VerSiguientePost()
			if err != nil {
				fmt.Println(err)
				break
			}
			salida.ImprimirSiguienteFeed(post)

		case LIKEAR_POST:
			// Verificar parametros de comando para likear post
			id, err := verificarUnParametro(parametros)
			if err != nil {
				fmt.Println(err)
				break
			}
			// Se transforma el ID de texto a número.
			idEntero, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err)
				break
			}
			// Se obtiene el usuario logueado.
			usuarioLogueado, err := usuariosBD.ObtenerLogueado()
			if err != nil {
				fmt.Println(errores.ErrorPostInexistenteOUsuarioNoLogueado{})
				break
			}
			nombreUsuarioLogueado := usuarioLogueado.VerNombre()
			err = postsBD.LikearPost(idEntero, nombreUsuarioLogueado)
			if err != nil {
				fmt.Println(err)
				break
			}
			salida.ImprimirLikearPost()

		case MOSTRAR_LIKES:
			// Verificar parametros de comando para likear post
			id, err := verificarUnParametro(parametros)
			if err != nil {
				fmt.Println(err)
				break
			}
			idEntero, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err)
				break
			}

			cadenaLikes, err := postsBD.VerLikes(idEntero)
			if err != nil {
				fmt.Println(err)
				break
			}
			salida.ImprimirMostrarLikes(cadenaLikes)

		default:
			fmt.Println(errores.ErrorComandoDesconocido{})
		}
	}

}

// obtenerComando convierte la cadena de caracteres del comando ingresado y lo transforma en un tipo de dato más cómodo de utilizar.
func obtenerComando(campos []string) Comando {
	switch campos[COMANDO_A_EJECUTAR_POS_CMD] {
	case COMANDO_LOGIN:
		return LOGIN
	case COMANDO_LOGOUT:
		return LOGOUT
	case COMANDO_PUBLICAR:
		return PUBLICAR
	case COMANDO_VER_SIGUIENTE_FEED:
		return VER_SIGUIENTE_FEED
	case COMANDO_LIKEAR_POST:
		return LIKEAR_POST
	case COMANDO_MOSTRAR_LIKES:
		return MOSTRAR_LIKES
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
