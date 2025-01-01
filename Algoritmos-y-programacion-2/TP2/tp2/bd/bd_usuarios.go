package bd

import (
	TDAPost "algogram/posts"
	TDAUsuario "algogram/usuarios"
)

type UsuariosBD interface {
	// Login perimte al usuario loggearse
	Login(string) error

	//Logout permite al usuario desloguearse
	Logout() error

	//Actualizar Feed actualiza el feed de los usuarios en mi base de datos
	ActualizarFeed(TDAPost.Post)

	// Existe me dice si existe el usuario en la base de datos
	Existe(string) bool

	// ObtenerUsuario obtiene el usuario del nombre pasado por parametro
	ObtenerUsuario(string) (TDAUsuario.Usuario, error)

	// HayConectado me dice si hay un usuario conectado a la base de datos
	HayConectado() bool

	//ObtenerLogueado obtiene el usuario logueado
	ObtenerLogueado() (TDAUsuario.Usuario, error)
}
