package edu.fiuba.algo3.modelo;

public class RespuestaPuntuada {
    private final Jugador jugador;
    private int cantidadRespuestasCorrectas;
    private int puntos;

    public RespuestaPuntuada(Jugador jugador, int puntos) {
        this.jugador = jugador;
        this.puntos = puntos;
        this.cantidadRespuestasCorrectas = 0;
    }

    public RespuestaPuntuada(Jugador jugador, int puntos, int cantidadRespuestasCorrectas) {
        this.jugador = jugador;
        this.puntos = puntos;
        this.cantidadRespuestasCorrectas = cantidadRespuestasCorrectas;
    }

    public void setCantidadRespuestasCorrectas(int cantidadRespuestasCorrectas) {
        this.cantidadRespuestasCorrectas = cantidadRespuestasCorrectas;
    }

    public Jugador getJugador() {
        return jugador;
    }

    public void multiplicarPuntaje(int factor){
        this.puntos = this.puntos * factor;
    }

    public int getCantidadRespuestasCorrectas(){
        return cantidadRespuestasCorrectas;
    }

    public boolean esCorrecta(){
        return puntos > 0;
    }

    public void asignarPuntos() {
        jugador.sumarPuntos(puntos);
    }
}
