package edu.fiuba.algo3.modelo.Bonificador;

import edu.fiuba.algo3.modelo.Jugador;
import edu.fiuba.algo3.modelo.RespuestaPuntuada;

public class AnuladorDecorador extends BonificadorDecorador{
    private Jugador jugador;
    private String nombreBonificador;

    public AnuladorDecorador(Bonificador wrapped, Jugador jugador){
        super(wrapped);
        this.jugador = jugador;
        nombreBonificador = "Anulador";
    }

    public String getNombreBonificador() {
        return nombreBonificador;
    }

    public RespuestaPuntuada modificarPuntaje(RespuestaPuntuada respuestaPuntuada) {
        RespuestaPuntuada respuesta = super.modificarPuntaje(respuestaPuntuada);
        if(!jugador.equals(respuesta.getJugador()) && respuesta.esCorrecta()){
            respuesta.multiplicarPuntaje(0);
        }
        return respuesta;
    }
}