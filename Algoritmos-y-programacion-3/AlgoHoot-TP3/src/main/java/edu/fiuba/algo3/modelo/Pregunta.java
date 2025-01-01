package edu.fiuba.algo3.modelo;

import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Penalidad.Penalidad;
import edu.fiuba.algo3.modelo.TipoDePregunta.TipoDePregunta;
import edu.fiuba.algo3.modelo.Bonificador.*;

import java.util.ArrayList;

public class Pregunta {
    private String enunciado;
    private TipoDePregunta tipo;
    private Penalidad penalidad;
    private ArrayList<Respuesta> respuestasJugador;
    private Bonificador bonificador;
    private String tema;
    private int cantidadRespuestasCorrectas;


    public Pregunta(TipoDePregunta tipo, Penalidad penalidad, String enunciado, String tema){
        this.tipo = tipo;
        this.penalidad = penalidad;
        this.enunciado = enunciado;
        this.respuestasJugador = new ArrayList<>();
        this.bonificador = new BonificadorConcreto();
        this.tema = tema;
        this.cantidadRespuestasCorrectas = 0;
    }

    public ArrayList<OpcionSimple> obtenerOpciones(){
        return tipo.obtenerOpciones();
    }

    public void agregarRespuesta(Respuesta respuestaJugador){
        respuestasJugador.add(respuestaJugador);
    }

    private RespuestaCorregida corregirRespuesta(Respuesta respuesta){
        RespuestaCorregida respuestaCorregida = tipo.corregirRespuesta(respuesta);
        if (respuestaCorregida.esCorrecta()) { this.cantidadRespuestasCorrectas += 1; }
        return respuestaCorregida;
    }

    private RespuestaPuntuada puntuarRespuesta(RespuestaCorregida respuesta){
        RespuestaPuntuada respuestaPuntuada = respuesta.asignarPuntaje(penalidad);
        respuestaPuntuada.setCantidadRespuestasCorrectas(cantidadRespuestasCorrectas);
        return respuestaPuntuada;
    }

    private RespuestaPuntuada aplicarBonificadores(RespuestaPuntuada respuesta){
        return bonificador.modificarPuntaje(respuesta);
    }

    private void puntuarJugadores(ArrayList<RespuestaCorregida> respuestasCorregidas){
        for(RespuestaCorregida respuestaCorregida : respuestasCorregidas){
            RespuestaPuntuada respuestaPuntuada = puntuarRespuesta(respuestaCorregida);
            respuestaPuntuada = aplicarBonificadores(respuestaPuntuada);
            respuestaPuntuada.asignarPuntos();
        }
    }

    public void validarRespuestas(){
        ArrayList<RespuestaCorregida> respuestasCorregidas = new ArrayList<>();
        for(Respuesta respuesta : respuestasJugador){
            RespuestaCorregida respuestaCorregida = corregirRespuesta(respuesta);
            respuestasCorregidas.add(respuestaCorregida);
        }
        puntuarJugadores(respuestasCorregidas);
    }

    public void agregarBonificador(BonificadorDecorador bonificador) {
        this.bonificador = bonificador.envolverBonificador(this.bonificador);
    }

    public String getEnunciado(){
        return this.enunciado;
    }

    public String getTema() {
        return tema;
    }

    public Penalidad getPenalidad(){ return penalidad;}

    public ArrayList<BonificadorDecorador> obtenerBonificadoresDisponibles(Jugador jugador){
        return penalidad.obtenerBonificadoresDisponibles(jugador);
    }

    public TipoDePregunta getTipo() {
        return tipo;
    }

    public boolean respondieronTodos(int cantidadJugadores){
        return respuestasJugador.size() == cantidadJugadores;
    }
}