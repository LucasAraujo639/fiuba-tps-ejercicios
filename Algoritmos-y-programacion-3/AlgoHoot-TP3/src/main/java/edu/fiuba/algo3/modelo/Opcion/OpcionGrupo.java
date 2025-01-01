package edu.fiuba.algo3.modelo.Opcion;

import java.util.HashSet;

public class OpcionGrupo implements Opcion {
    private String nombre;
    private HashSet<OpcionSimple> opciones;

    public OpcionGrupo(String nombre){
        this.nombre = nombre;
        this.opciones = new HashSet<>();
    }

    public OpcionGrupo(String nombre, HashSet<OpcionSimple> valores){
        this.nombre = nombre;
        this.opciones = valores;
    }

    public String getNombre() {
        return nombre;
    }

    public void agregarOpcion(OpcionSimple valor){
        this.opciones.add(valor);
    }

    public HashSet<OpcionSimple> getValores() {
        return opciones;
    }

    @Override
    public boolean esIgualA(Opcion opcionGrupo) {
        OpcionGrupo opcionComparada = (OpcionGrupo) opcionGrupo;
        HashSet<String> valoresComparada = new HashSet<>();
        HashSet<String> valoresGrupo = new HashSet<>();
        for(OpcionSimple opcion : opciones){
            valoresGrupo.add(opcion.getValor());
        }
        for(OpcionSimple opcion : opcionComparada.getValores()){
            valoresComparada.add(opcion.getValor());
        }
        return valoresComparada.equals(valoresGrupo);
    }

    @Override
    public String getTexto() {
        StringBuilder texto = new StringBuilder(nombre + ": ");
        for (OpcionSimple opcion : opciones) {
            texto.append(opcion.getTexto()).append(", ");
        }
        // Eliminar la última coma y espacio
        if (texto.length() > 0) {
            texto.setLength(texto.length() - 2);
        }
        return texto.toString();
    }

    @Override
    public String getName() {
        return nombre;
    }
}
