package usuarios_test

// import (
// 	TDAPost "tp2/posts"
// 	TDAUsuario "tp2/usuarios"

// 	TDADicccionario "tdas/diccionario"

// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestUsuarioNuevo(t *testing.T) {
// 	basePosts := TDADicccionario.CrearHash[int, TDAPost.Post]()
// 	usr := TDAUsuario.CrearUsuario("", basePosts)

// 	require.Equal(t, "", usr.VerNombre())

// 	_, err := usr.VerSiguientePost()

// 	require.NotEqual(t, nil, err)
// 	require.Equal(t, "Error: no hay mas posts para ver", err.Error())

// 	err = usr.LikearPost(0)

// 	require.NotEqual(t, nil, err)
// 	require.Equal(t, "Error: Post inexistente", err.Error())

// }

// func TestUsuarioPersonalizado(t *testing.T) {
// 	basePosts := TDADicccionario.CrearHash[int, TDAPost.Post]()
// 	usr := TDAUsuario.CrearUsuario("MiguelMerentiel", basePosts)

// 	require.Equal(t, "MiguelMerentiel", usr.VerNombre())

// 	_, err := usr.VerSiguientePost()

// 	require.NotEqual(t, nil, err)
// 	require.Equal(t, "Error: no hay mas posts para ver", err.Error())

// 	err = usr.LikearPost(0)

// 	require.NotEqual(t, nil, err)
// 	require.Equal(t, "Error: Post inexistente", err.Error())
// }
