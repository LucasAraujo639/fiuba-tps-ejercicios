package bd

import (
	"algogram/errores"
	TDAPost "algogram/posts"
	TDAUsuario "algogram/usuarios"
	"math"
	TDAHash "tdas/diccionario"
)

type usuariosBD struct {
	usuarios        TDAHash.Diccionario[string, TDAUsuario.Usuario]
	usuarioLogueado TDAUsuario.Usuario
}

// CrearUsuariosBD crea una base de datos de usuarios.
func CrearUsuariosBD(dicUsuarios TDAHash.Diccionario[string, TDAUsuario.Usuario]) UsuariosBD {
	bd := new(usuariosBD)
	bd.usuarios = dicUsuarios
	return bd
}

// Login loguea al usuario con el nombre pasado por argumento si existe y si no devuelve un error.
func (bd *usuariosBD) Login(usuario string) error {
	if !bd.usuarios.Pertenece(usuario) {
		return errores.ErrorUsuarioNoExiste{}
	}
	if bd.HayConectado() {
		return errores.ErrorUsuarioYaLoggeado{}
	}
	usuarioConectado := bd.usuarios.Obtener(usuario)
	bd.usuarioLogueado = usuarioConectado
	return nil
}

// Logout desloguea el usuario que se encuentre logueado. Si no hay ninún usuario logueado, devuelve un error.
func (bd *usuariosBD) Logout() error {
	if !(bd.HayConectado()) {
		return errores.ErrorUsuarioNoLoggeado{}
	}
	bd.usuarioLogueado = nil
	return nil
}

// Existe informa si el usuario con el nombre pasado por argumento existe.
func (bd *usuariosBD) Existe(nombreUsuario string) bool {
	return bd.usuarios.Pertenece(nombreUsuario)
}

// Si el usuario pertenece obtiene el usuario
func (bd *usuariosBD) ObtenerUsuario(nombreUsuario string) (TDAUsuario.Usuario, error) {
	if !bd.Existe(nombreUsuario) {
		return nil, errores.ErrorUsuarioNoExiste{}
	}
	return bd.usuarios.Obtener(nombreUsuario), nil
}

// ObtenerLogueado devuelve al usuario logueado. Si no hay ninguno devuelve un error.
func (bd *usuariosBD) ObtenerLogueado() (TDAUsuario.Usuario, error) {
	if !bd.HayConectado() {
		return nil, errores.ErrorUsuarioNoLoggeado{}
	}
	return bd.usuarioLogueado, nil
}

// HayConectado informa si hay un usuario logueado.
func (bd *usuariosBD) HayConectado() bool {
	return bd.usuarioLogueado != nil
}

// ActualizarFeed actualiza los feeds de todos los usuarios de la red con el post pasado por argumento.
func (bd *usuariosBD) ActualizarFeed(post TDAPost.Post) {
	creadorNombre := post.VerCreador() // usario logueado
	creadorUsuario := bd.usuarios.Obtener(creadorNombre)
	//Itero todo el diccionario uno por uno para ir agregando el post a cada usuario, excepto al creador
	iter := bd.usuarios.Iterador()
	for iter.HaySiguiente() {
		clave, usuario := iter.VerActual()

		// verifica que el usuario no sea el creador del post
		if clave != creadorNombre {
			// Calcula la afinidad del post con el usuario.
			afinidad := bd.definirAfinidad(usuario, creadorUsuario)

			// Guarda el post en el feed del usuario con la afinidad calculada.
			usuario.GuardarPostFeed(post, afinidad)
		}

		// Avanza al siguiente usuario en el diccionario
		iter.Siguiente()
	}
}

// Se calcula la afinidad entre el usuario recien que publica un post y cada usuario en el feed durante la actualizacion
func (bd *usuariosBD) definirAfinidad(usuario TDAUsuario.Usuario, creadorUsuario TDAUsuario.Usuario) int {
	posUsuario := usuario.VerPosicionLista()
	posCreador := creadorUsuario.VerPosicionLista()
	return int(math.Abs(float64(posUsuario - posCreador)))

}
