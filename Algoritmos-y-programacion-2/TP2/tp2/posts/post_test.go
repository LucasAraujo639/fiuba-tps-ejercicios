package posts_test

import (
	TDAPost "tp2/posts"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostVacio(t *testing.T) {
	post := TDAPost.CrearPost(0, "", "")

	require.Equal(t, 0, post.VerID())
	require.Equal(t, "", post.VerCreador())
	require.Equal(t, "", post.VerPost())
	require.Equal(t, []string{}, post.VerLikes())
}

func TestPostPersonalizado(t *testing.T) {
	post := TDAPost.CrearPost(10, "Yo", "Robot")

	require.Equal(t, 10, post.VerID())
	require.Equal(t, "Yo", post.VerCreador())
	require.Equal(t, "Robot", post.VerPost())
	require.Equal(t, []string{}, post.VerLikes())
}

func TestLikearPost(t *testing.T) {
	post := TDAPost.CrearPost(1492, "Cristobal", "Colon")

	require.Equal(t, 1492, post.VerID())
	require.Equal(t, "Cristobal", post.VerCreador())
	require.Equal(t, "Colon", post.VerPost())
	require.Equal(t, []string{}, post.VerLikes())

	post.Likear("Jesus")

	require.Equal(t, []string{"Jesus"}, post.VerLikes())
	require.Equal(t, 1492, post.VerID())
	require.Equal(t, "Cristobal", post.VerCreador())
	require.Equal(t, "Colon", post.VerPost())
}

func TestVariosLikes(t *testing.T) {
	post := TDAPost.CrearPost(2022, "Messi", "Anda pa alla")
	
	require.Equal(t, 2022, post.VerID())
	require.Equal(t, "Messi", post.VerCreador())
	require.Equal(t, "Anda pa alla", post.VerPost())
	require.Equal(t, []string{}, post.VerLikes())

	post.Likear("DiMaria")
	post.Likear("Argentina")
	post.Likear("Scaloni")
	post.Likear("DePaul")
	post.Likear("AntoRocuzzo")

	require.Equal(t, []string{"AntoRocuzzo", "Argentina", "DePaul", "DiMaria", "Scaloni"}, post.VerLikes())
	require.Equal(t, 2022, post.VerID())
	require.Equal(t, "Messi", post.VerCreador())
	require.Equal(t, "Anda pa alla", post.VerPost())
}