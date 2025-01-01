package edu.fiuba.algo3.modelo.Bonificador;

import edu.fiuba.algo3.modelo.*;

public class ExclusividadDecorador extends BonificadorDecorador {
    Jugador jugador;
    String nombreBonificador;

    public ExclusividadDecorador(Bonificador wrapped,Jugador jugador){
        super(wrapped);
        this.jugador = jugador;
        this.nombreBonificador = "Exclusividad";
    }

    public String getNombreBonificador() {
        return nombreBonificador;
    }

    @Override
    public RespuestaPuntuada modificarPuntaje(RespuestaPuntuada respuestaPuntuada){
        // Llamamos al método del decorador envuelto primero
        RespuestaPuntuada respuesta = super.modificarPuntaje(respuestaPuntuada);
        if(respuesta.getCantidadRespuestasCorrectas() == 1 && respuesta.esCorrecta()){
            respuesta.multiplicarPuntaje(2);
        }

        return respuesta;
    }

}
