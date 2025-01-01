package edu.fiuba.algo3.vista.PreguntasScene;

import edu.fiuba.algo3.controlador.SceneController;
import edu.fiuba.algo3.modelo.Juego;
import edu.fiuba.algo3.modelo.Jugador;
import edu.fiuba.algo3.modelo.Opcion.Opcion;
import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Penalidad.Penalidad;
import edu.fiuba.algo3.modelo.Pregunta;
import edu.fiuba.algo3.modelo.Respuesta;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.geometry.Insets;
import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.ListCell;
import javafx.scene.control.ListView;
import javafx.scene.input.MouseEvent;
import javafx.scene.layout.BorderPane;
import javafx.scene.layout.HBox;
import javafx.scene.layout.VBox;

import java.util.ArrayList;

public class JugarOrderChoiceScene {

    private SceneController sceneController;
    private Juego juego;

    public JugarOrderChoiceScene(SceneController sceneController, Juego juego) {
        this.sceneController = sceneController;
        this.juego = juego;
    }

    public Scene getScene() {
        Jugador jugador = juego.getJugadorActual();
        Pregunta pregunta = juego.getPreguntaActual();
        ArrayList<OpcionSimple> opciones = pregunta.obtenerOpciones();
        Respuesta respuesta = new Respuesta(jugador);
        Penalidad penalidadDeLaPregunta = pregunta.getPenalidad();

        // Crear la etiqueta para el enunciado de la pregunta
        Label preguntaLabel = new Label(pregunta.getEnunciado());
        preguntaLabel.setStyle("-fx-font-size: 24px; -fx-font-weight: bold; -fx-text-fill: #333;");
        preguntaLabel.setPadding(new Insets(20));
        preguntaLabel.setWrapText(true);
        VBox preguntaLabelVBox = new VBox(preguntaLabel);
        preguntaLabelVBox.setAlignment(Pos.CENTER);

        // Crear la lista de opciones con el ListView
        ObservableList<OpcionSimple> opcionesObservableList = FXCollections.observableArrayList(opciones);
        ListView<OpcionSimple> opcionesListView = new ListView<>(opcionesObservableList);
        opcionesListView.setCellFactory(param -> new ListCell<OpcionSimple>() {
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

        opcionesListView.setStyle("-fx-font-size: 18px; -fx-padding: 10px;");
        opcionesListView.setPrefWidth(400);

        // Botón para mover la opción seleccionada hacia arriba
        Button moveUpButton = new Button("⬆");
        moveUpButton.setOnAction(e -> {
            int selectedIndex = opcionesListView.getSelectionModel().getSelectedIndex();
            if (selectedIndex > 0) {
                OpcionSimple selectedItem = opcionesObservableList.remove(selectedIndex);
                opcionesObservableList.add(selectedIndex - 1, selectedItem);
                opcionesListView.getSelectionModel().select(selectedIndex - 1);
            }
        });
        moveUpButton.setStyle("-fx-font-size: 24px; -fx-background-color: #010101; -fx-font-weight: bold; -fx-text-fill: white;");
        moveUpButton.setPadding(new Insets(40, 40, 40, 40));
        moveUpButton.setWrapText(true);

        // Botón para mover la opción seleccionada hacia abajo
        Button moveDownButton = new Button("⬇");
        moveDownButton.setOnAction(e -> {
            int selectedIndex = opcionesListView.getSelectionModel().getSelectedIndex();
            if (selectedIndex < opcionesObservableList.size() - 1) {
                OpcionSimple selectedItem = opcionesObservableList.remove(selectedIndex);
                opcionesObservableList.add(selectedIndex + 1, selectedItem);
                opcionesListView.getSelectionModel().select(selectedIndex + 1);
            }
        });
        moveDownButton.setStyle("-fx-font-size: 24px; -fx-background-color: #010101; -fx-font-weight: bold; -fx-text-fill: white;");
        moveDownButton.setPadding(new Insets(40, 40, 40, 40));
        moveDownButton.setWrapText(true);

        VBox opcionesBox = new VBox(10, opcionesListView, moveUpButton, moveDownButton);
        opcionesBox.setPadding(new Insets(20));
        opcionesBox.setAlignment(Pos.CENTER);

        // Botón para enviar la respuesta
        Button enviarButton = new Button("Enviar");
        enviarButton.setOnAction(e -> {
            ArrayList<OpcionSimple> opcionesSeleccionadas = new ArrayList<>(opcionesListView.getItems());

            for (Opcion opcion : opcionesSeleccionadas) {
                respuesta.agregarOpcion(opcion);
            }

            jugador.responder(pregunta, respuesta);

            juego.siguienteTurno();
            sceneController.siguienteTurno();
        });

        enviarButton.setStyle("-fx-font-size: 18px; -fx-background-color: #010101; -fx-text-fill: white; -fx-border-color: #010101; -fx-border-width: 10px;");
        // Boton para volver al menú
        Button backButton = new Button("Volver al Menú");
        backButton.setOnAction(e -> sceneController.switchToTheStartScene());
        backButton.setStyle("-fx-font-size: 13px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 5px;");
        backButton.setPadding(new Insets(10));
        // Posicionar el boton en la esquina superior izquierda
        BorderPane.setAlignment(backButton, Pos.TOP_LEFT);
        BorderPane.setMargin(backButton, new Insets(10));

        // Mostrar las penalidades si las hay
        VBox botonesPenalidad = sceneController.MostrarBonificadores(pregunta, jugador);
        botonesPenalidad.setAlignment(Pos.CENTER_RIGHT);
        botonesPenalidad.setPadding(new Insets(20));

        HBox buttonBox = new HBox(10, enviarButton);
        buttonBox.setAlignment(Pos.CENTER);
        buttonBox.setPadding(new Insets(20));

        // Organizar la disposición de la interfaz
        HBox mainLayout = new HBox(20, opcionesBox, botonesPenalidad);
        mainLayout.setPadding(new Insets(20));
        mainLayout.setAlignment(Pos.CENTER);

        VBox layout = new VBox(20, mainLayout);
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
