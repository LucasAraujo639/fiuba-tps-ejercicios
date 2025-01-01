package edu.fiuba.algo3;

import edu.fiuba.algo3.modelo.Juego;
import edu.fiuba.algo3.modelo.Jugador;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class TestJuego {

    private Juego juego;

    @BeforeEach
    public void setUp() {
        juego = new Juego();
        juego.agregarJugador(new Jugador("Jugador 1"));
        juego.agregarJugador(new Jugador("Jugador 2"));
    }

    @Test
    public void testAgregarJugador() {
        assertEquals(2, juego.devolverJugadores().size(), "Debe haber dos jugadores");
    }

    @Test
    public void testTurnos() {
        Jugador jugadorActual = juego.getJugadorActual();
        assertEquals("Jugador 1", jugadorActual.getNombre());

        juego.siguienteTurno();
        jugadorActual = juego.getJugadorActual();
        assertEquals("Jugador 2", jugadorActual.getNombre());

        juego.siguienteTurno();
        jugadorActual = juego.getJugadorActual();
        assertEquals("Jugador 1", jugadorActual.getNombre());
    }

    /*@Test
    void queTipoDePreguntaToca() {
        TipoDePregunta tipo;

        tipo = juego.queTipoDePreguntaToca();
        System.out.println(tipo);
        juego.siguienteTurno();

        tipo = juego.queTipoDePreguntaToca();
        System.out.println(tipo);
        juego.siguienteTurno();

        tipo = juego.queTipoDePreguntaToca();
        System.out.println(tipo);
        juego.siguienteTurno();

        tipo = juego.queTipoDePreguntaToca();
        System.out.println(tipo);
        juego.siguienteTurno();
    }*/
}
