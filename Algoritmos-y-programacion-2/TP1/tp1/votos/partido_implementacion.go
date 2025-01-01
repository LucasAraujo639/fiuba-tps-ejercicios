package votos

import (
	"strconv"
)

type partidoImplementacion struct {
	nombre          string
	candidatos      [CANT_VOTACION]string
	votosCandidatos [CANT_VOTACION]uint
}

type partidoEnBlanco struct {
	votosBlanco [CANT_VOTACION]uint
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partidoNuevo := new(partidoImplementacion)
	partidoNuevo.nombre = nombre
	partidoNuevo.candidatos = candidatos
	return partidoNuevo
}

func CrearVotosEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votosCandidatos[tipo]++
}

func (partido *partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	var resultado string
	var mensajeVoto string

	if partido.votosCandidatos[tipo] == 1 {
		mensajeVoto = MENSAJE_VOTO_SINGULAR
	} else {
		mensajeVoto = MENSAJE_VOTO_PLURAL
	}
	resultado = partido.nombre + " - " + partido.candidatos[tipo] + ": " + strconv.Itoa(int(partido.votosCandidatos[tipo])) + " " + mensajeVoto

	return resultado

}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votosBlanco[tipo]++
}

func (blanco *partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	var resultado string
	var mensajeVoto string
	if blanco.votosBlanco[tipo] == 1 {
		mensajeVoto = MENSAJE_VOTO_SINGULAR
	} else {
		mensajeVoto = MENSAJE_VOTO_PLURAL
	}
	resultado += "Votos en Blanco: " + strconv.Itoa(int(blanco.votosBlanco[tipo])) + " " + mensajeVoto
	return resultado
}
