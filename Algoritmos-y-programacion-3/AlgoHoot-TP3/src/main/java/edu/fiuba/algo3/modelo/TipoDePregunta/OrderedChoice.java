package edu.fiuba.algo3.modelo.TipoDePregunta;

import edu.fiuba.algo3.modelo.Opcion.Opcion;
import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Respuesta;
import edu.fiuba.algo3.modelo.RespuestaCorregida;

import java.util.ArrayList;

public class OrderedChoice implements TipoDePregunta {
    ArrayList<OpcionSimple> opciones;
    ArrayList<OpcionSimple> opcionCorrecta;

    public OrderedChoice(ArrayList<OpcionSimple> opciones, ArrayList<OpcionSimple> opcionCorrecta){
        this.opciones = opciones;
        this.opcionCorrecta = opcionCorrecta;
    }

    private boolean esIgualA(ArrayList<Opcion> opcionesJugador){
        boolean esIgual = true;
        int i = 0;
        while(esIgual && i < opcionesJugador.size()){
            OpcionSimple opcionJugador = (OpcionSimple) opcionesJugador.get(i);
            OpcionSimple opcionCorrecta = this.opcionCorrecta.get(i);
            if(!opcionJugador.esIgualA(opcionCorrecta)){
                esIgual = false;
            }
            i++;
        }
        return esIgual;
    }

    @Override
    public ArrayList<OpcionSimple> obtenerOpciones() {
        return opciones;
    }

    @Override
    public RespuestaCorregida corregirRespuesta(Respuesta respuestaJugador) {
        int cantidadCorrectas = 0;
        int cantidadIncorrectas = 0;
        ArrayList<Opcion> opcionesJugador = respuestaJugador.getOpciones();
        if(esIgualA(opcionesJugador)) {
            cantidadCorrectas += 1;
        } else {
            cantidadIncorrectas += 1;
        }
        return new RespuestaCorregida(respuestaJugador.getJugador(), cantidadCorrectas, cantidadIncorrectas);
    }
}
