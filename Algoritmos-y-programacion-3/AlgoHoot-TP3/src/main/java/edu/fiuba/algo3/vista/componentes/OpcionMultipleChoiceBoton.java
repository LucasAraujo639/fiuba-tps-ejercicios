package edu.fiuba.algo3.vista.componentes;

import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import javafx.scene.control.CheckBox;

public class OpcionMultipleChoiceBoton extends CheckBox {
    private OpcionSimple opcion;

    public OpcionMultipleChoiceBoton(OpcionSimple opcion){
        super(opcion.getTexto());
        this.opcion = opcion;
        this.setStyle("-fx-font-size: 18px; -fx-padding: 10px;");
    }

    public OpcionSimple getOpcion() {
        return opcion;
    }
}
