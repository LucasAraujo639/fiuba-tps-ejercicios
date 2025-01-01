package usuarios

import (
	TDAPost "algogram/posts"
)

type Usuario interface {

	// VerNombre devuelve el nombre del usuario.
	VerNombre() string

	// CrearPost crea un post con el texto y el ID pasados por argumento.
	CrearPost(int, string) TDAPost.Post

	// VerSiguientePost devuelve el siguiente post a ver en el feed.
	VerSiguientePost() (TDAPost.Post, error)

	// GuardarPostFeed guarda el post guardado por argumento en el feed del usuario con la afinidad también pasada por argumento.
	GuardarPostFeed(TDAPost.Post, int)

	// VerPosicionLista devuelve la posición del usuario en el archivo donde se almacenan.
	VerPosicionLista() int
}
