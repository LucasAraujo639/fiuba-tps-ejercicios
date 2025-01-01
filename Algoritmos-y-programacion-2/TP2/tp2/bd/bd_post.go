package bd

import (
	TDAPost "algogram/posts"
)

type PostsBD interface {

	// ObtenerPost obtiene el post pasado por id
	ObtenerPost(int) TDAPost.Post

	// ObtenerTexto obtiene el texto del post pasado por id
	ObtenerTexto(int) string

	// GuardarPost guarda el post pasado por parametro en la base de datos posts
	GuardarPost(TDAPost.Post)

	// Cantidad muestra la cantidad de posts creados por todos los usuarios hasta el momento
	Cantidad() int

	// VerLikes muestra todos los que usuarios que dieron like a mi post en orden alfabetico
	VerLikes(int) ([]string, error)

	// LikearPost El usuario logueado le da like al post pasado por id
	LikearPost(int, string) error
}
