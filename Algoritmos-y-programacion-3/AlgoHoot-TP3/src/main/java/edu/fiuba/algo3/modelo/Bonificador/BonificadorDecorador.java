package edu.fiuba.algo3.modelo.Bonificador;

import edu.fiuba.algo3.modelo.RespuestaPuntuada;

public class BonificadorDecorador implements Bonificador{
    private Bonificador wrapped;

    BonificadorDecorador(Bonificador wrapped){
        this.wrapped = wrapped;
    }

    public Bonificador envolverBonificador(Bonificador bonificador){
        this.wrapped = bonificador;
        return this;
    }

    public RespuestaPuntuada modificarPuntaje(RespuestaPuntuada respuestaPuntuada){
        return wrapped.modificarPuntaje(respuestaPuntuada);
    }

    @Override
    public String getNombreBonificador() {
        return "Base";
    }
}