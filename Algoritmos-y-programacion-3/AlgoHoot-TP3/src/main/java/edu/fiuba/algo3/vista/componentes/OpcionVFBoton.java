package edu.fiuba.algo3.vista.componentes;

import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import javafx.scene.control.RadioButton;
import javafx.scene.control.ToggleGroup;

public class OpcionVFBoton extends RadioButton {
    OpcionSimple opcion;

    public OpcionVFBoton(ToggleGroup grupo, OpcionSimple opcion){
        super(opcion.getTexto());
        this.opcion = opcion;
        this.setToggleGroup(grupo);
        this.setStyle("-fx-font-size: 18px; -fx-padding: 10px;");
    }

    public OpcionSimple getOpcion() {
        return opcion;
    }
}
