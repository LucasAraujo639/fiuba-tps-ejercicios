package edu.fiuba.algo3.entrega_1;


import edu.fiuba.algo3.modelo.Jugador;
import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Penalidad.Penalidad;
import edu.fiuba.algo3.modelo.Penalidad.SinPenalidad;
import edu.fiuba.algo3.modelo.Pregunta;
import edu.fiuba.algo3.modelo.Respuesta;
import edu.fiuba.algo3.modelo.TipoDePregunta.OrderedChoice;
import edu.fiuba.algo3.modelo.TipoDePregunta.TipoDePregunta;
import org.junit.Test;

import java.util.ArrayList;

import static org.junit.Assert.assertEquals;

public class TestPreguntasOrderChoice {

    @Test
    public void Test01PreguntaOrderChoiceBienRespondida() {
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Macri",1));
        opciones.add(new OpcionSimple("Cristina",2));
        opciones.add(new OpcionSimple("Alberto",3));
        opciones.add(new OpcionSimple("Duhalde",4));
        opciones.add(new OpcionSimple("Nestor",5));


        ArrayList<OpcionSimple> opcionesCorrectas = new ArrayList<>();
        opcionesCorrectas.add(new OpcionSimple("Alberto",3));
        opcionesCorrectas.add(new OpcionSimple("Macri",1));
        opcionesCorrectas.add(new OpcionSimple("Cristina",2));
        opcionesCorrectas.add(new OpcionSimple("Nestor",5));
        opcionesCorrectas.add(new OpcionSimple("Duhalde",4));

        TipoDePregunta consigna = new OrderedChoice(opciones, opcionesCorrectas);
        Penalidad penalidad = new SinPenalidad();
        String enunciado = "Ordenar los presidentes por anio en orden descendiente";
        String tema = "Politica";

        Pregunta pregunta = new Pregunta(consigna, penalidad,enunciado,tema );
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Alberto",3));
        respuesta.agregarOpcion(new OpcionSimple("Macri",1));
        respuesta.agregarOpcion(new OpcionSimple("Cristina",2));
        respuesta.agregarOpcion(new OpcionSimple("Nestor",5));
        respuesta.agregarOpcion(new OpcionSimple("Duhalde",4));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        assertEquals(1, jugador1.getPuntajeTotal());
    }
    @Test
    public void Test02PreguntaOrderChoiceMalRespondida() {
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Macri",1));
        opciones.add(new OpcionSimple("Cristina",2));
        opciones.add(new OpcionSimple("Alberto",3));
        opciones.add(new OpcionSimple("Duhalde",4));
        opciones.add(new OpcionSimple("Nestor",5));


        ArrayList<OpcionSimple> opcionesCorrectas = new ArrayList<>();
        opcionesCorrectas.add(new OpcionSimple("Alberto",3));
        opcionesCorrectas.add(new OpcionSimple("Macri",1));
        opcionesCorrectas.add(new OpcionSimple("Cristina",2));
        opcionesCorrectas.add(new OpcionSimple("Nestor",5));
        opcionesCorrectas.add(new OpcionSimple("Duhalde",4));

        TipoDePregunta consigna = new OrderedChoice(opciones, opcionesCorrectas);
        Penalidad penalidad = new SinPenalidad();
        String enunciado = "Ordenar los presidentes por anio en orden descendiente";
        String tema = "Politica";

        Pregunta pregunta = new Pregunta(consigna, penalidad,enunciado,tema );
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Macri",1));
        respuesta.agregarOpcion(new OpcionSimple("Cristina",2));
        respuesta.agregarOpcion(new OpcionSimple("Alberto",3));
        respuesta.agregarOpcion(new OpcionSimple("Nestor",5));
        respuesta.agregarOpcion(new OpcionSimple("Duhalde",4));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        assertEquals(0, jugador1.getPuntajeTotal());
    }
}
