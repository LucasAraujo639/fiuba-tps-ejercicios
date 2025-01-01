package votos

import (
	"rerepolez/errores"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni              int
	votoActual       Voto
	historialCambios TDAPila.Pila[Voto]
}

func CrearVotante(dni int) Votante {
	if dni <= 0 {
		return nil
	}
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.historialCambios = TDAPila.CrearPilaDinamica[Voto]()
	return votante
}

// LeerDNI devuelve el valor del DNI del votante.
func (votante *votanteImplementacion) LeerDNI() int {
	return votante.dni
}

// actualizarVoto cambia el valor de un voto del votante.
func (votante *votanteImplementacion) actualizarVoto(tipo TipoVoto, alternativa int) {
	votante.votoActual.VotoPorTipo[tipo] = alternativa
}

// Votar cambia el valor del voto del votante de forma que luego pueda deshacerse la operación si es necesario.
// Devuelve un error si el tipo de voto es inválido con mensaje "ERROR: Tipo de voto inválido".
func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {

	if tipo >= CANT_VOTACION || tipo < 0 {
		return errores.ErrorTipoVoto{}
	}

	votoAnterior := votante.votoActual
	votante.historialCambios.Apilar(votoAnterior)
	votante.actualizarVoto(tipo, alternativa)

	if alternativa == LISTA_IMPUGNA {
		votante.impugnarVoto()
	}
	return nil
}

// impugnarVoto marca el voto como impugnado si encuentra algún voto con el valor LISTA_IMPUGNA.
func (votante *votanteImplementacion) impugnarVoto() {
	votante.votoActual.Impugnado = true
}

// Deshacer cambia el último voto realizado por su valor anterior.
// Devuelve un error si no se realizaron votos anteriores con mensaje "ERROR: Sin voto a deshacer".
func (votante *votanteImplementacion) Deshacer() error {

	if votante.historialCambios.EstaVacia() {
		return errores.ErrorNoHayVotosAnteriores{}
	}
	votoAnterior := votante.historialCambios.Desapilar()
	votante.votoActual = votoAnterior
	return nil
}

// FinVoto establece que el votante ya emitió su voto.
// Devuelve un error si el votante ya emitió su voto con mensaje "ERROR: Votante FRAUDULENTO: <NRO DNI>".
func (votante *votanteImplementacion) FinVoto() Voto {

	for !votante.historialCambios.EstaVacia() {
		votante.historialCambios.Desapilar()
	}
	return votante.votoActual
}
