package posts

// La interfaz Post viene a representar las acciones que se esperan de un post en una red social.
type Post interface {

	// VerCreador devuelve el nombre de usuario del creador del post.
	VerCreador() string

	// VerID devuelve el ID del post.
	VerID() int

	// VerPost devuelve el texto de la publicación.
	VerPost() string

	// Likear le da un "Me gusta" al post. Se pasa por parámetro el nombre del usuario que le da "Me gusta" al post.
	// No se verifica la existencia del usuario pasado por argumento.
	Likear(string)

	// VerLikes devuelve los nombres de los usuarios que dieron "Me gusta" al post ordenados en orden alfabético.
	VerLikes() []string
}
