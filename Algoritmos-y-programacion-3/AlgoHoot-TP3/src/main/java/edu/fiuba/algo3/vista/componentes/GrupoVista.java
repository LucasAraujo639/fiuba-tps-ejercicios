package edu.fiuba.algo3.vista.componentes;

import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import javafx.collections.ObservableList;
import javafx.scene.control.ListCell;
import javafx.scene.control.ListView;

import java.util.ArrayList;

public class GrupoVista extends ListView<OpcionSimple> {
    ArrayList<OpcionSimple> opciones;

    public GrupoVista(ObservableList<OpcionSimple> opciones){
        super(opciones);
        this.setCellFactory(param -> new ListCell<OpcionSimple>() {
            @Override
            protected void updateItem(OpcionSimple item, boolean empty) {
                super.updateItem(item, empty);
                if (empty || item == null) {
                    setText(null);
                } else {
                    setText(item.getTexto());
                }
            }
        });
        this.setPrefHeight(200);
    }

}
