package errores

type ErrorUsuarioNoLoggeado struct{}

func (e ErrorUsuarioNoLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorUsuarioNoExiste struct{}

func (e ErrorUsuarioNoExiste) Error() string {
	return "Error: usuario no existente"
}

type ErrorUsuarioNoLoggeadoONoMasPosts struct{}

func (e ErrorUsuarioNoLoggeadoONoMasPosts) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorUsuarioYaLoggeado struct{}

func (e ErrorUsuarioYaLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorPostInexistenteOUsuarioNoLogueado struct{}

func (e ErrorPostInexistenteOUsuarioNoLogueado) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorPostInexistenteOSinLikes struct{}

func (e ErrorPostInexistenteOSinLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}

type ErrorNoMasPosts struct{}

type ErrorLecturaArchivo struct{}

func (e ErrorLecturaArchivo) Error() string {
	return "Error: no se pudo leer el archivo"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "Error: Faltan Parametros"
}

type ErrorMalaInvocacionComando struct{}

func (e ErrorMalaInvocacionComando) Error() string {
	return "ERROR: Comando mal invocado"
}

type ErrorComandoDesconocido struct{}

func (e ErrorComandoDesconocido) Error() string {
	return "ERROR: Comando desconocido"
}
