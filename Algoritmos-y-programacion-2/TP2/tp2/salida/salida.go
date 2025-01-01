package salida

import (
	TDAPost "algogram/posts"
	"fmt"
)

// ImprimirLogin muestra un mensaje luego de que se loguee un usuario
func ImprimirLogin(usuario string) {
	fmt.Println("Hola " + usuario)
}

// ImprimirLogout muestra un mensaje luego de que se desloguee un usuario
func ImprimirLogout() {
	fmt.Println("Adios")
}

// ImprimirLikearPost muestra un mensaje luego de que se likee un post
func ImprimirLikearPost() {
	fmt.Println("Post likeado")
}

// ImprimirMostrarLikes muestra un mensaje luego de que se solicite ver los likes de un post
func ImprimirMostrarLikes(likes []string) {
	fmt.Printf("El post tiene %d likes:\n", len(likes))

	// Enumerar los nombres de usuarios
	for _, usuario := range likes {
		fmt.Printf("\t%s\n", usuario)
	}
}

// ImprimirPostPublicado muestra un mensaje luego de que se publique un post
func ImprimirPostPublicado() {
	fmt.Println("Post publicado")
}

// ImprimirSiguienteFeed muestra toda la información del siguiente post del feed
func ImprimirSiguienteFeed(post TDAPost.Post) {
	fmt.Printf("Post ID %d\n", post.VerID())
	fmt.Printf("%s dijo: %s\n", post.VerCreador(), post.VerPost())
	fmt.Printf("Likes: %d\n", len(post.VerLikes()))
}
