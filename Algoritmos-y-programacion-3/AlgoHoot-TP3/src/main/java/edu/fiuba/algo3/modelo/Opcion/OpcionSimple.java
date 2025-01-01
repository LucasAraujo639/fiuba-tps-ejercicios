package edu.fiuba.algo3.modelo.Opcion;

public class OpcionSimple implements Opcion {
    private final String valor;
    private int id;

    public OpcionSimple(String valor, int id) {
        this.id = id;
        this.valor = valor;
    }

    public String getValor() {
        return valor;
    }
    public int getId() {
        return id;
    }

    @Override
    public boolean esIgualA(Opcion opcion) {
        OpcionSimple opcionSimple = (OpcionSimple) opcion;
        return this.id == opcionSimple.getId();
    }

    @Override
    public String getTexto() {
        return valor;
    }

    @Override
    public String getName() {
        return valor;
    }

}
