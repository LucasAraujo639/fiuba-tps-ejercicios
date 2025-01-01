package edu.fiuba.algo3.modelo.Penalidad;

import edu.fiuba.algo3.modelo.Bonificador.BonificadorDecorador;
import edu.fiuba.algo3.modelo.Jugador;
import edu.fiuba.algo3.modelo.RespuestaPuntuada;

import java.util.ArrayList;

public interface Penalidad {

    RespuestaPuntuada asignarPuntajeRespuesta(Jugador jugador, int cantidadCorrectas, int cantidadIncorrectas);
    ArrayList<BonificadorDecorador> obtenerBonificadoresDisponibles(Jugador jugador);

}
