package edu.fiuba.algo3.vista;

import edu.fiuba.algo3.controlador.SceneController;
import edu.fiuba.algo3.modelo.Juego;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.TextField;
import javafx.scene.layout.BorderPane;
import javafx.scene.layout.HBox;
import javafx.scene.layout.VBox;
import javafx.geometry.Insets;
import javafx.geometry.Pos;

public class AñadirJugadorScene {

    private SceneController sceneController;

    public AñadirJugadorScene(SceneController sceneController) {
        this.sceneController = sceneController;
    }

    public Scene getScene(Juego juego) {
        // Título
        Label titleLabel = new Label("Añadir Jugador");
        titleLabel.setStyle("-fx-font-size: 36px; -fx-font-weight: bold; -fx-text-fill: #010101;");

        // Etiqueta y campo de texto para el nombre del jugador
        Label instructionLabel = new Label("Ingrese el nombre del nuevo jugador:");
        instructionLabel.setStyle("-fx-font-size: 18px; -fx-text-fill: #010101;");

        TextField nameInput = new TextField();
        nameInput.setPromptText("Nombre del jugador");
        nameInput.setStyle("-fx-font-size: 16px;");

        // Botón para añadir jugador
        Button addButton = new Button("Añadir Jugador");
        addButton.setOnAction(e -> {
            String playerName = nameInput.getText();
            if (!playerName.isEmpty()) {
                sceneController.addPlayer(playerName);
                sceneController.switchToMenuScene();
            }
        });
        addButton.setStyle("-fx-font-size: 18px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 10px;");

        // Botón para volver al menú
        Button backButton = new Button("Volver");
        backButton.setOnAction(e -> sceneController.switchToMenuScene());
        backButton.setStyle("-fx-font-size: 18px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 10px;");

        HBox buttonBox = new HBox(10, addButton, backButton);
        buttonBox.setAlignment(Pos.CENTER);

        // Layout principal
        VBox layout = new VBox(20, titleLabel, instructionLabel, nameInput, buttonBox);
        layout.setPadding(new Insets(20));
        layout.setAlignment(Pos.CENTER);

        BorderPane root = new BorderPane();
        root.setCenter(layout);

        Scene scene = new Scene(root, 1000, 650);
        scene.getStylesheets().add(getClass().getResource("/styles.css").toExternalForm());

        return scene;
    }
}
