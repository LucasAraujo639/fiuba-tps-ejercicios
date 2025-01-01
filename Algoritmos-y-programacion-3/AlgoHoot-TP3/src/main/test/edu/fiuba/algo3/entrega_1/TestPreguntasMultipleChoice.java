package edu.fiuba.algo3.entrega_1;

import edu.fiuba.algo3.modelo.Jugador;
import edu.fiuba.algo3.modelo.Opcion.Opcion;
import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Penalidad.ConPenalidad;
import edu.fiuba.algo3.modelo.Penalidad.Penalidad;
import edu.fiuba.algo3.modelo.Penalidad.PenalidadParcial;
import edu.fiuba.algo3.modelo.Penalidad.SinPenalidad;
import edu.fiuba.algo3.modelo.Pregunta;
import edu.fiuba.algo3.modelo.Respuesta;
import edu.fiuba.algo3.modelo.TipoDePregunta.MultipleChoice;
import edu.fiuba.algo3.modelo.TipoDePregunta.TipoDePregunta;
import org.junit.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class TestPreguntasMultipleChoice {

    @Test
    public void test01PreguntaMultipleChoiceClasicoBienRespondida() {
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Neymar",1));
        opciones.add(new OpcionSimple("Messi",2));
        opciones.add(new OpcionSimple("Cristiano Ronaldo",3));
        opciones.add(new OpcionSimple("Lewandowski",4));

        ArrayList<OpcionSimple> opcionesCorrectas = new ArrayList<>();
        opcionesCorrectas.add(new OpcionSimple("Neymar",1));
        opcionesCorrectas.add(new OpcionSimple("Lewandowski",4));
        String enunciado = "Cual/es de estos jugadores nunca ganaron un Balon de Oro?";
        String tema = "Deportes";

        TipoDePregunta consigna = new MultipleChoice(opciones, opcionesCorrectas);
        Penalidad penalidad = new SinPenalidad();

        Pregunta pregunta = new Pregunta(consigna, penalidad, enunciado,tema);
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Neymar",1));
        respuesta.agregarOpcion(new OpcionSimple("Lewandowski",4));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        assertEquals(2, jugador1.getPuntajeTotal());
    }

    @Test
    public void test02PreguntaMultipleChoiceClasicoMalRespondida(){
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Neymar",1));
        opciones.add(new OpcionSimple("Messi",2));
        opciones.add(new OpcionSimple("Cristiano Ronaldo",3));
        opciones.add(new OpcionSimple("Lewandowski",4));

        ArrayList<OpcionSimple> opcionesCorrectas = new ArrayList<>();
        opcionesCorrectas.add(new OpcionSimple("Neymar",1));
        opcionesCorrectas.add(new OpcionSimple("Lewandowski",4));
        String enunciado = "Cual/es de estos jugadores nunca ganaron un Balon de Oro?";
        String tema = "Deportes";

        TipoDePregunta consigna = new MultipleChoice(opciones, opcionesCorrectas);
        Penalidad penalidad = new SinPenalidad();

        Pregunta pregunta = new Pregunta(consigna, penalidad, enunciado,tema);
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Messi",2));
        respuesta.agregarOpcion(new OpcionSimple("Cristiano Ronaldo",3));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        assertEquals(0, jugador1.getPuntajeTotal());
    }

    @Test
    public void test03PreguntaMultipleChoiceConPenalidadBienRespondida(){
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Neymar",1));
        opciones.add(new OpcionSimple("Messi",2));
        opciones.add(new OpcionSimple("Cristiano Ronaldo",3));
        opciones.add(new OpcionSimple("Lewandowski",4));

        ArrayList<OpcionSimple> opcionesCorrectas = new ArrayList<>();
        opcionesCorrectas.add(new OpcionSimple("Neymar",1));
        opcionesCorrectas.add(new OpcionSimple("Lewandowski",4));
        String enunciado = "Cual/es de estos jugadores nunca ganaron un Balon de Oro?";
        String tema = "Deportes";

        TipoDePregunta consigna = new MultipleChoice(opciones, opcionesCorrectas);
        Penalidad penalidad = new ConPenalidad();

        Pregunta pregunta = new Pregunta(consigna, penalidad, enunciado,tema);
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Neymar",1));
        respuesta.agregarOpcion(new OpcionSimple("Lewandowski",4));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        assertEquals(2, jugador1.getPuntajeTotal());
    }

    @Test
    public void test04PreguntaMultipleChoiceConPenalidadMalRespondida(){
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Neymar",1));
        opciones.add(new OpcionSimple("Messi",2));
        opciones.add(new OpcionSimple("Cristiano Ronaldo",3));
        opciones.add(new OpcionSimple("Lewandowski",4));

        ArrayList<OpcionSimple> opcionesCorrectas = new ArrayList<>();
        opcionesCorrectas.add(new OpcionSimple("Neymar",1));
        opcionesCorrectas.add(new OpcionSimple("Lewandowski",4));
        String enunciado = "Cual/es de estos jugadores nunca ganaron un Balon de Oro?";
        String tema = "Deportes";

        TipoDePregunta consigna = new MultipleChoice(opciones, opcionesCorrectas);
        Penalidad penalidad = new ConPenalidad();

        Pregunta pregunta = new Pregunta(consigna, penalidad, enunciado,tema);
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Messi",2));
        respuesta.agregarOpcion(new OpcionSimple("Cristiano Ronaldo",3));
        respuesta.agregarOpcion(new OpcionSimple("Lewandowski",4));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        assertEquals(-1, jugador1.getPuntajeTotal());
    }

    @Test
    public void test05PreguntaMultipleChoiceParcialBienRespondida(){
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Neymar",1));
        opciones.add(new OpcionSimple("Messi",2));
        opciones.add(new OpcionSimple("Cristiano Ronaldo",3));
        opciones.add(new OpcionSimple("Lewandowski",4));

        ArrayList<OpcionSimple> opcionesCorrectas = new ArrayList<>();
        opcionesCorrectas.add(new OpcionSimple("Neymar",1));
        opcionesCorrectas.add(new OpcionSimple("Lewandowski",4));
        String enunciado = "Cual/es de estos jugadores nunca ganaron un Balon de Oro?";
        String tema = "Deportes";

        TipoDePregunta consigna = new MultipleChoice(opciones, opcionesCorrectas);
        Penalidad penalidad = new PenalidadParcial();

        Pregunta pregunta = new Pregunta(consigna, penalidad, enunciado,tema);
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Neymar",1));
        respuesta.agregarOpcion(new OpcionSimple("Lewandowski",4));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        assertEquals(2, jugador1.getPuntajeTotal());
    }
    @Test
    public void test06PreguntaMultipleChoiceParcialMalRespondida(){
        Jugador jugador1 = new Jugador("Bob");
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        opciones.add(new OpcionSimple("Neymar",1));
        opciones.add(new OpcionSimple("Messi",2));
        opciones.add(new OpcionSimple("Cristiano Ronaldo",3));
        opciones.add(new OpcionSimple("Lewandowski",4));

        ArrayList<OpcionSimple> opcionesCorrectas = new ArrayList<>();
        opcionesCorrectas.add(new OpcionSimple("Neymar",1));
        opcionesCorrectas.add(new OpcionSimple("Lewandowski",4));
        String enunciado = "Cual/es de estos jugadores nunca ganaron un Balon de Oro?";
        String tema = "Deportes";

        TipoDePregunta consigna = new MultipleChoice(opciones, opcionesCorrectas);
        Penalidad penalidad = new PenalidadParcial();

        Pregunta pregunta = new Pregunta(consigna, penalidad, enunciado,tema);
        Respuesta respuesta = new Respuesta(jugador1);

        respuesta.agregarOpcion(new OpcionSimple("Neymar",1));
        respuesta.agregarOpcion(new OpcionSimple("Lewandowski",4));
        respuesta.agregarOpcion(new OpcionSimple("Cristiano Ronaldo",3));

        jugador1.responder(pregunta, respuesta);

        pregunta.validarRespuestas();

        assertEquals(0, jugador1.getPuntajeTotal());
    }
}
