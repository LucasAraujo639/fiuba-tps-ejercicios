package bd

import (
	"algogram/errores"
	TDAPost "algogram/posts"
)

const _TAMANIO_INICIAL = 0

type postsBD struct {
	posts    []TDAPost.Post
	cantidad int
}

func CrearPostsBD() PostsBD {
	bd := new(postsBD)
	bd.posts = crearSlicePosts()
	return bd
}
func crearSlicePosts() []TDAPost.Post {
	slice := make([]TDAPost.Post, _TAMANIO_INICIAL)
	return slice
}
func (p *postsBD) ObtenerPost(id int) TDAPost.Post {
	return p.posts[id]
}
func (p *postsBD) VerLikes(id int) ([]string, error) {
	if id < 0 || id >= len(p.posts) {
		return nil, errores.ErrorPostInexistenteOSinLikes{}
	}
	// Una vez que verifique que el id existe en mi slice de posts traigo los likes del post indicado por el id
	cadenaLikes := p.posts[id].VerLikes()

	// Verifico que el post tenga likes
	if len(cadenaLikes) == 0 {
		return nil, errores.ErrorPostInexistenteOSinLikes{}
	}
	return cadenaLikes, nil
}
func (p *postsBD) LikearPost(id int, usuarioLogueado string) error {

	if id < 0 || id >= len(p.posts) {
		return errores.ErrorPostInexistenteOUsuarioNoLogueado{}
	}
	// Likea el post de mi bd post
	p.posts[id].Likear(usuarioLogueado)
	return nil
}
func (p postsBD) ObtenerTexto(id int) string {
	return p.posts[id].VerPost()
}
func (p *postsBD) GuardarPost(post TDAPost.Post) {
	p.posts = append(p.posts, post)
	p.cantidad++
}
func (p postsBD) Cantidad() int {
	return p.cantidad
}
