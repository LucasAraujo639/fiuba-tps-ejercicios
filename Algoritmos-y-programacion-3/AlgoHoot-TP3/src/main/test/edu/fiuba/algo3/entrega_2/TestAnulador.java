package edu.fiuba.algo3.entrega_2;

import edu.fiuba.algo3.modelo.*;
import edu.fiuba.algo3.modelo.Bonificador.BonificadorDecorador;
import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Penalidad.ConPenalidad;
import edu.fiuba.algo3.modelo.Penalidad.Penalidad;
import edu.fiuba.algo3.modelo.TipoDePregunta.TipoDePregunta;
import edu.fiuba.algo3.modelo.TipoDePregunta.VerdaderoFalso;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class TestAnulador {

    @Test
    public void test01jugador1AplicaAnuladorySeAnulaElPuntajeDelRival() {
        Jugador jugador1 = new Jugador("Bob");
        Jugador jugador2 = new Jugador("Alice");

        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Verdadero",1));
        opciones.add(new OpcionSimple("Falso",2));

        TipoDePregunta consigna = new VerdaderoFalso(opciones, new OpcionSimple("Verdadero", 1));
        Penalidad penalidad = new ConPenalidad();

        Pregunta pregunta = new Pregunta(consigna, penalidad, "Messi es el mejor jugador de la historia?","deporte");

        Respuesta respuestaJugador1 = new Respuesta(jugador1);
        respuestaJugador1.agregarOpcion(new OpcionSimple("Verdadero", 1));

        Respuesta respuestaJugador2 = new Respuesta(jugador2);
        respuestaJugador2.agregarOpcion(new OpcionSimple("Verdadero", 1));

        jugador1.responder(pregunta, respuestaJugador1);
        jugador2.responder(pregunta, respuestaJugador2);

        ArrayList<BonificadorDecorador> bonificadores = pregunta.obtenerBonificadoresDisponibles(jugador1);
        BonificadorDecorador bonificadorDecorador = bonificadores.get(2);
        jugador1.usarBonificador(bonificadorDecorador,pregunta);

        pregunta.validarRespuestas();

        assertEquals(1, jugador1.getPuntajeTotal());
        assertEquals(0, jugador2.getPuntajeTotal());
    }

    @Test
    public void test01jugador1y2AplicaAnuladorySeAnulaElPuntajeDeTodos() {
        Jugador jugador1 = new Jugador("Bob");
        Jugador jugador2 = new Jugador("Alice");

        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Verdadero",1));
        opciones.add(new OpcionSimple("Falso",2));

        TipoDePregunta consigna = new VerdaderoFalso(opciones, new OpcionSimple("Verdadero", 1));
        Penalidad penalidad = new ConPenalidad();

        Pregunta pregunta = new Pregunta(consigna, penalidad, "Messi es el mejor jugador de la historia?","deporte");

        Respuesta respuestaJugador1 = new Respuesta(jugador1);
        respuestaJugador1.agregarOpcion(new OpcionSimple("Verdadero", 1));

        Respuesta respuestaJugador2 = new Respuesta(jugador2);
        respuestaJugador2.agregarOpcion(new OpcionSimple("Verdadero", 1));

        jugador1.responder(pregunta, respuestaJugador1);
        jugador2.responder(pregunta, respuestaJugador2);

        ArrayList<BonificadorDecorador> bonificadores = pregunta.obtenerBonificadoresDisponibles(jugador1);
        BonificadorDecorador bonificadorDecorador = bonificadores.get(2);
        ArrayList<BonificadorDecorador> bonificadores2 = pregunta.obtenerBonificadoresDisponibles(jugador2);
        BonificadorDecorador bonificadorDecorador2 = bonificadores2.get(2);
        jugador1.usarBonificador(bonificadorDecorador,pregunta);
        jugador2.usarBonificador(bonificadorDecorador2,pregunta);

        pregunta.validarRespuestas();

        assertEquals(0, jugador1.getPuntajeTotal());
        assertEquals(0, jugador2.getPuntajeTotal());
    }

}
