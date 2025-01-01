package edu.fiuba.algo3.modelo;

import edu.fiuba.algo3.modelo.Bonificador.*;

import java.util.ArrayList;

public class Jugador {
    private String nombre;
    private int puntajeTotal;
    private ArrayList<BonificadorDecorador> bonificadores = new ArrayList<>();
    private BonificadorDecorador ultimoBonificador;

    public Jugador(String nombreJugador) {
        nombre = nombreJugador;
        inicializarBonificadores();
    }

    private void inicializarBonificadores() {
        bonificadores.add(new MultiplicadorDecorador(new BonificadorConcreto(), this,2));
        bonificadores.add(new MultiplicadorDecorador(new BonificadorConcreto(), this, 3));
        bonificadores.add(new ExclusividadDecorador(new BonificadorConcreto(), this));
        bonificadores.add(new ExclusividadDecorador(new BonificadorConcreto(),this));
        bonificadores.add(new AnuladorDecorador(new BonificadorConcreto(),this));
    }

    public ArrayList<BonificadorDecorador> getBonificadores(){ return bonificadores; }

    public String getNombre(){ return this.nombre; }

    public int getPuntajeTotal(){
        return puntajeTotal;
    }

    public void usarBonificador(BonificadorDecorador bonificador, Pregunta pregunta) {
        pregunta.agregarBonificador(bonificador);
        bonificadores.remove(bonificador);
        ultimoBonificador = bonificador;
    }

    public void responder(Pregunta pregunta, Respuesta respuestaJugador) {
        pregunta.agregarRespuesta(respuestaJugador);
    }

    public void sumarPuntos(int puntaje) {
        puntajeTotal += puntaje;
    }

    public BonificadorDecorador getUltimoBonificador() {
        return ultimoBonificador;
    }

    public void setUltimoBonificador(BonificadorDecorador ultimoBonificador) {
        this.ultimoBonificador = ultimoBonificador;
    }
}