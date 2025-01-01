package edu.fiuba.algo3.modelo.Penalidad;

import edu.fiuba.algo3.modelo.Bonificador.*;
import edu.fiuba.algo3.modelo.Jugador;

import edu.fiuba.algo3.modelo.RespuestaPuntuada;

import java.util.ArrayList;


public class SinPenalidad implements Penalidad {
    public ArrayList<BonificadorDecorador> obtenerBonificadoresDisponibles(Jugador jugador){
        ArrayList<BonificadorDecorador> exclusividades = jugador.getBonificadores().stream().filter(b -> b.getNombreBonificador().equals("Exclusividad")).collect(java.util.stream.Collectors.toCollection(ArrayList::new));
        ArrayList<BonificadorDecorador> anuladores = jugador.getBonificadores().stream().filter(b -> b.getNombreBonificador().equals("Anulador")).collect(java.util.stream.Collectors.toCollection(ArrayList::new));
        ArrayList<BonificadorDecorador> bonificadoresDisponibles = new ArrayList<>(exclusividades);
        bonificadoresDisponibles.addAll(anuladores);
        return bonificadoresDisponibles;
    }

    @Override
    public RespuestaPuntuada asignarPuntajeRespuesta(Jugador jugador, int cantidadCorrectas, int cantidadIncorrectas) {
        int puntajeAAsignar;
        if(cantidadIncorrectas > 0) {
            puntajeAAsignar = 0;
        }
        else {
            puntajeAAsignar = cantidadCorrectas;
        }
        return new RespuestaPuntuada(jugador, puntajeAAsignar);
    }
}
