package edu.fiuba.algo3.modelo.TipoDePregunta;

import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Respuesta;
import edu.fiuba.algo3.modelo.RespuestaCorregida;

import java.util.ArrayList;

public interface TipoDePregunta {
    ArrayList<OpcionSimple> obtenerOpciones();
    RespuestaCorregida corregirRespuesta(Respuesta respuestaJugador);
}
