package edu.fiuba.algo3.modelo;

import edu.fiuba.algo3.modelo.Penalidad.Penalidad;

public class RespuestaCorregida {
    private Jugador jugador;
    private int cantidadCorrectas;
    private int cantidadIncorrectas;

    public RespuestaCorregida(Jugador jugador, int cantidadCorrectas, int cantidadIncorrectas){
        this.jugador = jugador;
        this.cantidadCorrectas = cantidadCorrectas;
        this.cantidadIncorrectas = cantidadIncorrectas;
    }

    public boolean esCorrecta(){
        return cantidadIncorrectas == 0;
    }

    public RespuestaPuntuada asignarPuntaje(Penalidad penalidad) {
        return penalidad.asignarPuntajeRespuesta(jugador, cantidadCorrectas, cantidadIncorrectas);
    }

}
