package votos_test

import (
	"rerepolez/votos"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestVotanteNuevo(t *testing.T) {
	votante := votos.CrearVotante(100)

	require.NotEqual(t, nil, votante)
	require.Equal(t, 100, votante.LeerDNI())

	err := votante.Deshacer()
	require.Equal(t, "ERROR: Sin voto a deshacer", err.Error())

	err = votante.Votar(5000, 3)
	require.Equal(t, "ERROR: Tipo de voto inválido", err.Error())

	err = votante.Votar(-5, 777)
	require.Equal(t, "ERROR: Tipo de voto inválido", err.Error())

	votoRealizado := votante.FinVoto()
	require.Nil(t, err)
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.False(t, votoRealizado.Impugnado)
	require.Equal(t, 100, votante.LeerDNI())
}

func TestDNINoValidos(t *testing.T) {
	votante := votos.CrearVotante(0)

	require.Nil(t, votante)

	votante = votos.CrearVotante(-57)

	require.Nil(t, votante)
}

func TestVotoPartidoUnico(t *testing.T) {
	votante := votos.CrearVotante(47)

	require.NotNil(t, votante)
	require.Equal(t, 47, votante.LeerDNI())

	err := votante.Votar(votos.PRESIDENTE, 1)
	require.Nil(t, err)

	err = votante.Votar(votos.GOBERNADOR, 1)
	require.Nil(t, err)

	err = votante.Votar(votos.INTENDENTE, 1)
	require.Nil(t, err)

	votoRealizado := votante.FinVoto()
	require.Nil(t, err)
	require.Equal(t, 1, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, 1, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, 1, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.False(t, votoRealizado.Impugnado)
	require.Equal(t, 47, votante.LeerDNI())
}

func TestVotoVariosPartidos(t *testing.T) {
	votante := votos.CrearVotante(137)

	require.NotNil(t, votante)
	require.Equal(t, 137, votante.LeerDNI())

	err := votante.Votar(votos.PRESIDENTE, 5)
	require.Nil(t, err)

	err = votante.Votar(votos.GOBERNADOR, 4)
	require.Nil(t, err)

	err = votante.Votar(votos.INTENDENTE, 6)
	require.Nil(t, err)

	votoRealizado := votante.FinVoto()
	require.Nil(t, err)
	require.Equal(t, 5, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, 4, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, 6, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.False(t, votoRealizado.Impugnado)
	require.Equal(t, 137, votante.LeerDNI())
}

func TestVotoImpugnado(t *testing.T) {
	// Voto impugnado en voto a presidente.
	votante1 := votos.CrearVotante(137)

	require.NotNil(t, votante1)
	require.Equal(t, 137, votante1.LeerDNI())

	err := votante1.Votar(votos.PRESIDENTE, votos.VOTO_EN_BLANCO)
	require.Nil(t, err)

	err = votante1.Votar(votos.GOBERNADOR, 4)
	require.Nil(t, err)

	err = votante1.Votar(votos.INTENDENTE, 6)
	require.Nil(t, err)

	votoRealizado := votante1.FinVoto()
	require.Nil(t, err)
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, 4, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, 6, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.True(t, votoRealizado.Impugnado)
	require.Equal(t, 137, votante1.LeerDNI())

	// Voto impugnado en voto a gobernador.
	votante2 := votos.CrearVotante(555)

	require.NotNil(t, votante2)
	require.Equal(t, 555, votante2.LeerDNI())

	err = votante2.Votar(votos.PRESIDENTE, 5)
	require.Nil(t, err)

	err = votante2.Votar(votos.GOBERNADOR, votos.VOTO_EN_BLANCO)
	require.Nil(t, err)

	err = votante2.Votar(votos.INTENDENTE, 6)
	require.Nil(t, err)

	votoRealizado = votante2.FinVoto()
	require.Nil(t, err)
	require.Equal(t, 5, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, 6, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.True(t, votoRealizado.Impugnado)
	require.Equal(t, 555, votante2.LeerDNI())

	// Voto impugnado en voto a intendente.
	votante3 := votos.CrearVotante(321)

	require.NotNil(t, votante3)
	require.Equal(t, 321, votante3.LeerDNI())

	err = votante3.Votar(votos.PRESIDENTE, 5)
	require.Nil(t, err)

	err = votante3.Votar(votos.GOBERNADOR, 4)
	require.Nil(t, err)

	err = votante3.Votar(votos.INTENDENTE, votos.VOTO_EN_BLANCO)
	require.Nil(t, err)

	votoRealizado = votante3.FinVoto()
	require.Nil(t, err)
	require.Equal(t, 5, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, 4, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.True(t, votoRealizado.Impugnado)
	require.Equal(t, 321, votante3.LeerDNI())
}

func TestAlternativaInvalida(t *testing.T) {
	votante := votos.CrearVotante(42)

	require.NotNil(t, votante)
	require.Equal(t, 42, votante.LeerDNI())

	err := votante.Votar(votos.CANT_VOTACION, 5)
	require.Equal(t, "ERROR: Tipo de voto inválido", err.Error())

	err = votante.Votar(votos.CANT_VOTACION+6, 5)
	require.Equal(t, "ERROR: Tipo de voto inválido", err.Error())

	err = votante.Votar(-5, 5)
	require.Equal(t, "ERROR: Tipo de voto inválido", err.Error())
}

func TestDeshacerVotoAVotoBlanco(t *testing.T) {
	votante := votos.CrearVotante(93)

	require.NotNil(t, votante)
	require.Equal(t, 93, votante.LeerDNI())

	err := votante.Deshacer()
	require.Equal(t, "ERROR: Sin voto a deshacer", err.Error())

	err = votante.Votar(votos.GOBERNADOR, 15)
	require.Nil(t, err)

	err = votante.Deshacer()
	require.Nil(t, err)

	votoRealizado := votante.FinVoto()
	require.Nil(t, err)
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.False(t, votoRealizado.Impugnado)
	require.Equal(t, 93, votante.LeerDNI())
}

func TestDeshacerVotoCompletoAVotoBlanco(t *testing.T) {
	votante := votos.CrearVotante(105)

	require.NotNil(t, votante)
	require.Equal(t, 105, votante.LeerDNI())

	err := votante.Deshacer()
	require.Equal(t, "ERROR: Sin voto a deshacer", err.Error())

	err = votante.Votar(votos.GOBERNADOR, 15)
	require.Nil(t, err)
	err = votante.Votar(votos.INTENDENTE, 5)
	require.Nil(t, err)
	err = votante.Votar(votos.PRESIDENTE, 1)
	require.Nil(t, err)

	err = votante.Deshacer()
	require.Nil(t, err)
	err = votante.Deshacer()
	require.Nil(t, err)
	err = votante.Deshacer()
	require.Nil(t, err)

	votoRealizado := votante.FinVoto()
	require.Nil(t, err)
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, votos.VOTO_EN_BLANCO, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.False(t, votoRealizado.Impugnado)
	require.Equal(t, 105, votante.LeerDNI())
}

func TestVariosVotosYDeshacer(t *testing.T) {
	votante := votos.CrearVotante(108)

	require.NotNil(t, votante)
	require.Equal(t, 108, votante.LeerDNI())

	err := votante.Deshacer()
	require.Equal(t, "ERROR: Sin voto a deshacer", err.Error())

	err = votante.Votar(votos.GOBERNADOR, 15)
	require.Nil(t, err)
	err = votante.Votar(votos.INTENDENTE, 5)
	require.Nil(t, err)
	err = votante.Votar(votos.PRESIDENTE, 1)
	require.Nil(t, err)

	err = votante.Votar(votos.GOBERNADOR, 7)
	require.Nil(t, err)
	err = votante.Votar(votos.GOBERNADOR, 23)
	require.Nil(t, err)
	err = votante.Votar(votos.INTENDENTE, 8)
	require.Nil(t, err)
	err = votante.Votar(votos.PRESIDENTE, 6)
	require.Nil(t, err)

	err = votante.Deshacer()
	require.Nil(t, err)

	votoRealizado := votante.FinVoto()
	require.Nil(t, err)
	require.Equal(t, 1, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, 23, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, 8, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.False(t, votoRealizado.Impugnado)
	require.Equal(t, 108, votante.LeerDNI())
}

func TestVolverAVotoCompletoAnterior(t *testing.T) {
	votante := votos.CrearVotante(117)

	require.NotNil(t, votante)
	require.Equal(t, 117, votante.LeerDNI())

	err := votante.Votar(votos.GOBERNADOR, 2)
	require.Nil(t, err)
	err = votante.Votar(votos.INTENDENTE, 8)
	require.Nil(t, err)
	err = votante.Votar(votos.PRESIDENTE, 6)
	require.Nil(t, err)

	err = votante.Votar(votos.INTENDENTE, 87)
	require.Nil(t, err)
	err = votante.Votar(votos.PRESIDENTE, 65)
	require.Nil(t, err)
	err = votante.Votar(votos.GOBERNADOR, 28)
	require.Nil(t, err)

	err = votante.Deshacer()
	require.Nil(t, err)
	err = votante.Deshacer()
	require.Nil(t, err)
	err = votante.Deshacer()
	require.Nil(t, err)

	votoRealizado := votante.FinVoto()
	require.Nil(t, err)
	require.Equal(t, 6, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, 2, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, 8, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.False(t, votoRealizado.Impugnado)
	require.Equal(t, 117, votante.LeerDNI())
}

func TestVotoFraudulento(t *testing.T) {
	// Crear un votante.
	votante := votos.CrearVotante(222)

	require.NotNil(t, votante)
	require.Equal(t, 222, votante.LeerDNI())

	err := votante.Votar(votos.PRESIDENTE, 5)
	require.Nil(t, err)

	err = votante.Votar(votos.GOBERNADOR, 4)
	require.Nil(t, err)

	err = votante.Votar(votos.INTENDENTE, 6)
	require.Nil(t, err)

	votoRealizado := votante.FinVoto()
	require.Nil(t, err)
	require.Equal(t, 5, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, 4, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, 6, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.False(t, votoRealizado.Impugnado)
	require.Equal(t, 222, votante.LeerDNI())

	// Votar luego de finalizado el voto.
	err = votante.Votar(votos.PRESIDENTE, 7)
	require.EqualErrorf(t, err, "ERROR: Votante FRAUDULENTO: 222", err.Error())

	err = votante.Votar(votos.GOBERNADOR, 7)
	require.EqualErrorf(t, err, "ERROR: Votante FRAUDULENTO: 222", err.Error())

	err = votante.Votar(votos.INTENDENTE, 7)
	require.EqualErrorf(t, err, "ERROR: Votante FRAUDULENTO: 222", err.Error())

	// Deshacer luego de finalizado el voto.
	err = votante.Deshacer()
	require.EqualErrorf(t, err, "ERROR: Votante FRAUDULENTO: 222", err.Error())

	// Volver a emitir voto luego de finalizado el voto.
	votoRealizado = votante.FinVoto()
	require.EqualErrorf(t, err, "ERROR: Votante FRAUDULENTO: 222", err.Error())
	require.Equal(t, votos.Voto{}, votoRealizado)
	require.Equal(t, 222, votante.LeerDNI())
}

func TestCambioDeVoto(t *testing.T) {
	votante := votos.CrearVotante(1000)

	require.NotNil(t, votante)
	require.Equal(t, 1000, votante.LeerDNI())

	err := votante.Votar(votos.PRESIDENTE, 5)
	require.Nil(t, err)

	err = votante.Votar(votos.GOBERNADOR, 4)
	require.Nil(t, err)

	err = votante.Votar(votos.INTENDENTE, 6)
	require.Nil(t, err)

	// Cambiar voto de presidente.
	err = votante.Votar(votos.PRESIDENTE, 1)
	require.Nil(t, err)

	// Imprimir resultado.
	votoRealizado := votante.FinVoto()
	require.Equal(t, 1, votoRealizado.VotoPorTipo[votos.PRESIDENTE])
	require.Equal(t, 4, votoRealizado.VotoPorTipo[votos.GOBERNADOR])
	require.Equal(t, 6, votoRealizado.VotoPorTipo[votos.INTENDENTE])
	require.False(t, votoRealizado.Impugnado)
	require.Equal(t, 1000, votante.LeerDNI())
}
