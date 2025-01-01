package edu.fiuba.algo3.entrega_1;

import edu.fiuba.algo3.modelo.*;
import edu.fiuba.algo3.modelo.Jugador;

import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Penalidad.ConPenalidad;
import edu.fiuba.algo3.modelo.Penalidad.Penalidad;
import edu.fiuba.algo3.modelo.Penalidad.SinPenalidad;
import edu.fiuba.algo3.modelo.TipoDePregunta.TipoDePregunta;
import edu.fiuba.algo3.modelo.TipoDePregunta.VerdaderoFalso;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;

public class TestPreguntasVF {
    @Test
    public void Test01PreguntaVFClasicoBienRespondida() {
       Jugador jugador1 = new Jugador("Bob");
       ArrayList<OpcionSimple> opciones = new ArrayList<>();
       opciones.add(new OpcionSimple("Verdadero",1));
       opciones.add(new OpcionSimple("Falso",2));

        TipoDePregunta consigna = new VerdaderoFalso(opciones, new OpcionSimple("Verdadero", 1));
        Penalidad penalidad = new SinPenalidad();
        String enunciado = "Messi es el mejor jugador de la historia?";
        String tema = "Deportes";

        Pregunta pregunta = new Pregunta(consigna, penalidad, enunciado,tema);
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Verdadero", 1));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        Assertions.assertEquals(1, jugador1.getPuntajeTotal());
    }

    @Test
    public void Test02PreguntaVFClasicoMalRespondida() {
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Verdadero",1));
        opciones.add(new OpcionSimple("Falso",2));

        TipoDePregunta consigna = new VerdaderoFalso(opciones, new OpcionSimple("Verdadero", 1));
        Penalidad penalidad = new SinPenalidad();
        String enunciado = "Messi es el mejor jugador de la historia?";
        String tema = "Deportes";

        Pregunta pregunta = new Pregunta(consigna, penalidad, enunciado,tema);
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Falso", 2));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        Assertions.assertEquals(0, jugador1.getPuntajeTotal());
    }

    @Test
    public void Test03PreguntaVFConPenalidadBienRespondida() {
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Verdadero",1));
        opciones.add(new OpcionSimple("Falso",2));

        TipoDePregunta consigna = new VerdaderoFalso(opciones, new OpcionSimple("Verdadero", 1));
        Penalidad penalidad = new ConPenalidad();
        String enunciado = "Messi es el mejor jugador de la historia?";
        String tema = "Deportes";

        Pregunta pregunta = new Pregunta(consigna, penalidad, enunciado,tema);
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Verdadero", 1));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        // Verificar puntajes
        Assertions.assertEquals(1, jugador1.getPuntajeTotal());
    }

    @Test
    public void Test04PreguntaVFConPenalidadMalRespondida() {
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Verdadero",1));
        opciones.add(new OpcionSimple("Falso",2));

        TipoDePregunta consigna = new VerdaderoFalso(opciones, new OpcionSimple("Verdadero", 1));
        Penalidad penalidad = new ConPenalidad();
        String enunciado = "Messi es el mejor jugador de la historia?";
        String tema = "Deportes";

        Pregunta pregunta = new Pregunta(consigna, penalidad, enunciado,tema);
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Falso", 2));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        // Verificar puntajes
        Assertions.assertEquals(-1, jugador1.getPuntajeTotal());
    }
}
