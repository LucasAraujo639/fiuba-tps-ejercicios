package edu.fiuba.algo3.modelo.Penalidad;

import edu.fiuba.algo3.modelo.Bonificador.BonificadorDecorador;
import edu.fiuba.algo3.modelo.Jugador;

import edu.fiuba.algo3.modelo.RespuestaPuntuada;

import java.util.ArrayList;
import java.util.stream.Collectors;


public class PenalidadParcial implements Penalidad {

    public ArrayList<BonificadorDecorador> obtenerBonificadoresDisponibles(Jugador jugador){
        return jugador.getBonificadores().stream().filter(b -> b.getNombreBonificador().equals("Anulador")).collect(Collectors.toCollection(ArrayList::new));
    }

    @Override
    public RespuestaPuntuada asignarPuntajeRespuesta(Jugador jugador, int cantidadCorrectas, int cantidadIncorrectas) {
        if(cantidadIncorrectas == 0){
            return new RespuestaPuntuada(jugador, cantidadCorrectas);
        }
        return new RespuestaPuntuada(jugador, 0);
    }
}
