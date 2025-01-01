package edu.fiuba.algo3.entrega_1;


import edu.fiuba.algo3.modelo.Jugador;
import edu.fiuba.algo3.modelo.Opcion.OpcionGrupo;
import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Penalidad.Penalidad;
import edu.fiuba.algo3.modelo.Penalidad.SinPenalidad;
import edu.fiuba.algo3.modelo.Pregunta;
import edu.fiuba.algo3.modelo.Respuesta;
import edu.fiuba.algo3.modelo.TipoDePregunta.GroupChoice;
import edu.fiuba.algo3.modelo.TipoDePregunta.TipoDePregunta;
import org.junit.Test;

import java.util.ArrayList;
import java.util.HashSet;

import static org.junit.Assert.assertEquals;

public class TestPreguntaGroupChoice {

    @Test
    public void test01JugadorRespondeGroupChoiceCorrectamenteBienRespondida(){
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Tigre",1));
        opciones.add(new OpcionSimple("Pato", 2));
        opciones.add(new OpcionSimple("Gallina",3));
        opciones.add(new OpcionSimple("Perro",4));
        opciones.add(new OpcionSimple("Pajaro",5));
        opciones.add(new OpcionSimple("Gato",6));

        HashSet<OpcionSimple> mamiferosCorrectos = new HashSet<>();
        mamiferosCorrectos.add(new OpcionSimple("Tigre",1));
        mamiferosCorrectos.add(new OpcionSimple("Perro",4));
        mamiferosCorrectos.add(new OpcionSimple("Gato",6));
        HashSet<OpcionSimple> oviparosCorrectos = new HashSet<>();
        oviparosCorrectos.add(new OpcionSimple("Gallina",3));
        oviparosCorrectos.add(new OpcionSimple("Pajaro",5));
        oviparosCorrectos.add(new OpcionSimple("Pato",2));

        OpcionGrupo grupo1 = new OpcionGrupo("Mamiferos", mamiferosCorrectos);
        OpcionGrupo grupo2 = new OpcionGrupo("Oviparos", oviparosCorrectos);

        ArrayList<OpcionGrupo> gruposCorrectos = new ArrayList<>();
        gruposCorrectos.add(grupo1);
        gruposCorrectos.add(grupo2);

        TipoDePregunta consigna = new GroupChoice(opciones, gruposCorrectos);
        Penalidad penalidad = new SinPenalidad();
        String enunciado = "Colocar los animales en la categoria correspondiente";
        String tema = "Animales";

        Pregunta pregunta = new Pregunta(consigna, penalidad,enunciado,tema);

        OpcionGrupo mamiferosJugador = new OpcionGrupo("Mamiferos", new HashSet<>());
        OpcionGrupo oviparosJugador = new OpcionGrupo("Oviparos", new HashSet<>());

        Respuesta respuesta = new Respuesta(jugador1);

        oviparosJugador.agregarOpcion(new OpcionSimple("Gallina", 3));
        oviparosJugador.agregarOpcion(new OpcionSimple("Pato",2));
        oviparosJugador.agregarOpcion(new OpcionSimple("Pajaro",5));
        mamiferosJugador.agregarOpcion(new OpcionSimple("Tigre",1));
        mamiferosJugador.agregarOpcion(new OpcionSimple("Perro",4));
        mamiferosJugador.agregarOpcion(new OpcionSimple("Gato",6));

        respuesta.agregarOpcion(oviparosJugador);
        respuesta.agregarOpcion(mamiferosJugador);

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        assertEquals(1, jugador1.getPuntajeTotal());
    }
    @Test
    public void test02JugadorRespondeGroupChoiceMalRespondida(){
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Tigre",1));
        opciones.add(new OpcionSimple("Pato", 2));
        opciones.add(new OpcionSimple("Gallina",3));
        opciones.add(new OpcionSimple("Perro",4));
        opciones.add(new OpcionSimple("Pajaro",5));
        opciones.add(new OpcionSimple("Gato",6));

        HashSet<OpcionSimple> mamiferosCorrectos = new HashSet<>();
        mamiferosCorrectos.add(new OpcionSimple("Tigre",1));
        mamiferosCorrectos.add(new OpcionSimple("Perro",4));
        mamiferosCorrectos.add(new OpcionSimple("Gato",6));
        HashSet<OpcionSimple> oviparosCorrectos = new HashSet<>();
        oviparosCorrectos.add(new OpcionSimple("Gallina",3));
        oviparosCorrectos.add(new OpcionSimple("Pajaro",5));
        oviparosCorrectos.add(new OpcionSimple("Pato",2));

        OpcionGrupo grupo1 = new OpcionGrupo("Mamiferos", mamiferosCorrectos);
        OpcionGrupo grupo2 = new OpcionGrupo("Oviparos", oviparosCorrectos);

        ArrayList<OpcionGrupo> gruposCorrectos = new ArrayList<>();
        gruposCorrectos.add(grupo1);
        gruposCorrectos.add(grupo2);

        TipoDePregunta consigna = new GroupChoice(opciones, gruposCorrectos);
        Penalidad penalidad = new SinPenalidad();
        String enunciado = "Colocar los animales en la categoria correspondiente";
        String tema = "Animales";

        Pregunta pregunta = new Pregunta(consigna, penalidad,enunciado,tema);

        OpcionGrupo mamiferosJugador = new OpcionGrupo("Mamiferos", new HashSet<>());
        OpcionGrupo oviparosJugador = new OpcionGrupo("Oviparos", new HashSet<>());

        Respuesta respuesta = new Respuesta(jugador1);

        oviparosJugador.agregarOpcion(new OpcionSimple("Gato", 6));
        oviparosJugador.agregarOpcion(new OpcionSimple("Pato",2));
        oviparosJugador.agregarOpcion(new OpcionSimple("Pajaro",5));
        mamiferosJugador.agregarOpcion(new OpcionSimple("Tigre",1));
        mamiferosJugador.agregarOpcion(new OpcionSimple("Perro",4));
        mamiferosJugador.agregarOpcion(new OpcionSimple("Gallina",3));

        respuesta.agregarOpcion(oviparosJugador);
        respuesta.agregarOpcion(mamiferosJugador);

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        assertEquals(0, jugador1.getPuntajeTotal());
    }
}
