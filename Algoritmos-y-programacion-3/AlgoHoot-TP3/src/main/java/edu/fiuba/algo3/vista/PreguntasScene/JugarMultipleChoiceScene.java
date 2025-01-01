package edu.fiuba.algo3.vista.PreguntasScene;

import edu.fiuba.algo3.controlador.SceneController;
import edu.fiuba.algo3.vista.componentes.OpcionMultipleChoiceBoton;
import edu.fiuba.algo3.modelo.Juego;
import edu.fiuba.algo3.modelo.Jugador;
import edu.fiuba.algo3.modelo.Penalidad.Penalidad;
import edu.fiuba.algo3.modelo.Pregunta;
import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Respuesta;
import javafx.geometry.Insets;
import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.layout.BorderPane;
import javafx.scene.layout.HBox;
import javafx.scene.layout.VBox;

import java.util.ArrayList;

public class JugarMultipleChoiceScene {

    private SceneController sceneController;
    private Juego juego;

    public JugarMultipleChoiceScene(SceneController sceneController, Juego juego) {
        this.sceneController = sceneController;
        this.juego = juego;
    }

    public Scene getScene() {
        Jugador jugador = juego.getJugadorActual();
        Pregunta pregunta = juego.getPreguntaActual();
        Penalidad penalidadDeLaPregunta = pregunta.getPenalidad();
        Respuesta respuesta = new Respuesta(jugador);
        ArrayList<OpcionSimple> opciones = pregunta.obtenerOpciones();

        // Crear la etiqueta de la pregunta
        Label preguntaLabel = new Label(pregunta.getEnunciado());
        preguntaLabel.setStyle("-fx-font-size: 24px; -fx-font-weight: bold; -fx-text-fill: #333;");
        preguntaLabel.setPadding(new Insets(20));
        preguntaLabel.setWrapText(true);
        VBox preguntaLabelVBox = new VBox(preguntaLabel);
        preguntaLabelVBox.setAlignment(Pos.CENTER);


        VBox botonesPenalidad = sceneController.MostrarBonificadores(pregunta,jugador);
        botonesPenalidad.setAlignment(Pos.CENTER_RIGHT);
        botonesPenalidad.setPadding(new Insets(20));

        VBox opcionesBox = new VBox(10);
        for (OpcionSimple opcion : opciones) {
            OpcionMultipleChoiceBoton checkBox = new OpcionMultipleChoiceBoton(opcion);
            checkBox.setStyle("-fx-font-size: 18px; -fx-padding: 10px;");
            opcionesBox.getChildren().add(checkBox);
        }
        opcionesBox.setPadding(new Insets(20));
        opcionesBox.setAlignment(Pos.CENTER_LEFT);

        // Botón para enviar la respuesta
        Button enviarButton = new Button("Enviar");
        enviarButton.setOnAction(e -> {
            ArrayList<OpcionSimple> opcionesSeleccionadas = new ArrayList<>();
            for (javafx.scene.Node node : opcionesBox.getChildren()) {
                if (node instanceof OpcionMultipleChoiceBoton) {
                    OpcionMultipleChoiceBoton checkBox = (OpcionMultipleChoiceBoton) node;
                    if (checkBox.isSelected()) {
                        respuesta.agregarOpcion(checkBox.getOpcion());
                    }
                }
            }
            jugador.responder(pregunta, respuesta);
            juego.siguienteTurno();
            sceneController.siguienteTurno();
        });
        enviarButton.setStyle("-fx-font-size: 18px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 10px;");
        // Boton para volver al menú
        Button backButton = new Button("Volver al Menú");
        backButton.setOnAction(e -> sceneController.switchToTheStartScene());
        backButton.setStyle("-fx-font-size: 13px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 5px;");
        backButton.setPadding(new Insets(10));
        // Posicionar el boton en la esquina superior izquierda
        BorderPane.setAlignment(backButton, Pos.TOP_LEFT);
        BorderPane.setMargin(backButton, new Insets(10));

        HBox buttonBox = new HBox(10, enviarButton);
        buttonBox.setAlignment(Pos.CENTER);
        buttonBox.setPadding(new Insets(20));

        // Layout principal

        HBox layout = new HBox(100, opcionesBox, botonesPenalidad);
        layout.setPadding(new Insets(20));
        layout.setAlignment(Pos.CENTER);

        BorderPane root = new BorderPane();
        root.setTop(preguntaLabelVBox);
        root.setCenter(layout);
        root.setBottom(buttonBox);
        root.setLeft(backButton);

        Scene scene = new Scene(root, 1000, 650);
        scene.getStylesheets().add(getClass().getResource("/styles.css").toExternalForm());

        return scene;
    }
}
