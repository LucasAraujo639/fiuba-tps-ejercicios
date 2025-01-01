package edu.fiuba.algo3.modelo;

import edu.fiuba.algo3.modelo.Opcion.Opcion;

import java.util.ArrayList;

public class Respuesta {
    Jugador jugador;
    ArrayList<Opcion> opciones;

    public Respuesta(Jugador jugador) {
        this.jugador = jugador;
        this.opciones = new ArrayList<>();
    }

    public void agregarOpcion(Opcion opcion) {
        this.opciones.add(opcion);
    }

    public void removerOpcion(ArrayList<Opcion> opciones) {
        this.opciones.removeAll(opciones);
    }

    public Jugador getJugador() {
        return jugador;
    }

    public ArrayList<Opcion> getOpciones() {
        return opciones;
    }
}
