package usuarios

import (
	"algogram/errores"
	TDAPost "algogram/posts"
	TDAColaPrioridad "tdas/cola_prioridad"
)

// Implementación de un usuario.
type usuarioImplementacion struct {
	posLista int
	nombre   string
	feed     TDAColaPrioridad.ColaPrioridad[*nodoHeap]
}

// Implementación del nodo de la cola de prioridad.
type nodoHeap struct {
	post     TDAPost.Post
	afinidad int
}

// FUNCIONES DE CREACIÓN:

// CrearUsuario crea un usuario que posea el nombre pasado por argumento.
func CrearUsuario(posLista int, nombre string) Usuario {
	nuevoUsuario := new(usuarioImplementacion)
	nuevoUsuario.posLista = posLista
	nuevoUsuario.nombre = nombre
	nuevoUsuario.feed = TDAColaPrioridad.CrearHeap[*nodoHeap](funcCompFeed)
	return nuevoUsuario
}

// Función de comparacion que considera afinidad y fecha de creacion
func funcCompFeed(nodoHeap1, nodoHeap2 *nodoHeap) int {

	// Comparar por afinidad primero.
	if nodoHeap1.afinidad == nodoHeap2.afinidad {
		return nodoHeap2.post.VerID() - nodoHeap1.post.VerID()
	}
	return nodoHeap2.afinidad - nodoHeap1.afinidad
}

// PRIMITIVAS:

// VerNombre devuelve el nombre del usuario.
func (usr *usuarioImplementacion) VerNombre() string {
	return usr.nombre
}

// CrearPost crea un post con el texto y el ID pasados por argumento.
func (usr *usuarioImplementacion) CrearPost(id int, texto string) TDAPost.Post {
	postNuevo := TDAPost.CrearPost(id, usr.VerNombre(), texto)
	return postNuevo
}

// VerSiguientePost devuelve el siguiente post a ver en el feed.
func (usr *usuarioImplementacion) VerSiguientePost() (TDAPost.Post, error) {
	if usr.feed.EstaVacia() {
		return nil, errores.ErrorUsuarioNoLoggeadoONoMasPosts{}
	}

	return usr.feed.Desencolar().post, nil
}

// VerNombre devuelve la posicion de lista del usuario
func (usr *usuarioImplementacion) VerPosicionLista() int {
	return usr.posLista
}

// GuardarPostFeed guarda el post guardado por argumento en el feed del usuario con la afinidad también pasada por argumento.
func (usr *usuarioImplementacion) GuardarPostFeed(postNuevo TDAPost.Post, afinidad int) {
	nuevoNodoHeap := new(nodoHeap)
	nuevoNodoHeap.post = postNuevo
	nuevoNodoHeap.afinidad = afinidad
	usr.feed.Encolar(nuevoNodoHeap)
}
