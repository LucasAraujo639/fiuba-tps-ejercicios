package edu.fiuba.algo3.modelo.Bonificador;


import edu.fiuba.algo3.modelo.RespuestaPuntuada;

public interface Bonificador {
    RespuestaPuntuada modificarPuntaje(RespuestaPuntuada respuestaPuntuada);

    String getNombreBonificador();
}