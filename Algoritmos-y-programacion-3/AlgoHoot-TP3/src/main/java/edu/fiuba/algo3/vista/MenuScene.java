package edu.fiuba.algo3.vista;

import edu.fiuba.algo3.controlador.SceneController;
import edu.fiuba.algo3.modelo.Juego;
import edu.fiuba.algo3.modelo.Jugador;
import javafx.application.Platform;
import javafx.beans.value.ChangeListener;
import javafx.beans.value.ObservableValue;
import javafx.scene.Scene;
import javafx.scene.control.*;
import javafx.scene.layout.BorderPane;
import javafx.scene.layout.HBox;
import javafx.scene.layout.VBox;
import javafx.geometry.Insets;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;

public class MenuScene {

    private Juego juego;
    private SceneController sceneController;

    public MenuScene(SceneController sceneController, Juego juego) {
        this.sceneController = sceneController;
        this.juego = juego;
    }

    public Scene getScene() {

        Label titleLabel = new Label("AlgoHoot");
        titleLabel.setStyle("-fx-font-size: 36px; -fx-font-weight: bold; -fx-text-fill: #010101;");

        // Sección de participantes
        Label participantsLabel = new Label("Jugadores");
        participantsLabel.setStyle("-fx-font-size: 24px; -fx-font-weight: bold; -fx-text-fill: White;");

        List<Jugador> jugadores = new ArrayList<>();
        jugadores = juego.devolverJugadores();

        List<String> nombres = new ArrayList<>();
        for (Jugador jugador : jugadores) {
            nombres.add(jugador.getNombre());
        }

        ListView<String> participantsList = new ListView<>();
        participantsList.getItems().addAll(nombres);
        participantsList.setStyle("-fx-font-size: 18px; -fx-background-color: White;");


        Label pointsLimitLabel = new Label("Puntaje para ganar");
        pointsLimitLabel.setStyle("-fx-text-fill: White;");
        AtomicInteger puntosBase = new AtomicInteger(10);
        Label pointsLimitInput = new Label(Integer.toString(puntosBase.get()));
        pointsLimitInput.setPadding(new Insets(10, 10, 10, 10));
        pointsLimitInput.setStyle("-fx-text-fill: White;");
        Button addPoint = new Button("+");
        Button substractPoint = new Button("-");
        addPoint.setOnAction(e->{
            pointsLimitInput.setText(Integer.toString(puntosBase.updateAndGet(i -> i >= 25 ? 1 : i+1)));
        });
        substractPoint.setOnAction(e->{
            pointsLimitInput.setText(Integer.toString(puntosBase.updateAndGet(i -> i <= 1 ? 25 : i-1)));
        });

        HBox pointsLimitBox = new HBox(substractPoint, pointsLimitInput, addPoint);
        VBox pointsLimit = new VBox(pointsLimitLabel, pointsLimitBox);

        Label roundsLimitLabel = new Label("Cantidad de Rondas");
        roundsLimitLabel.setStyle("-fx-text-fill: White;");
        AtomicInteger roundsBase = new AtomicInteger(10);
        Label roundsLimitInput = new Label(Integer.toString(puntosBase.get()));
        roundsLimitInput.setPadding(new Insets(10, 10, 10, 10));
        roundsLimitInput.setStyle("-fx-text-fill: White;");
        Button addRound = new Button("+");
        Button substractRound = new Button("-");
        addRound.setOnAction(e->{
            roundsLimitInput.setText(Integer.toString(roundsBase.updateAndGet(i -> i >= 25 ? 1 : i+1)));
        });
        substractRound.setOnAction(e->{
            roundsLimitInput.setText(Integer.toString(roundsBase.updateAndGet(i -> i <= 1 ? 25 : i-1)));
        });

        HBox roundsLimitBox = new HBox(substractRound, roundsLimitInput, addRound);
        VBox roundsLimit = new VBox(roundsLimitLabel, roundsLimitBox);

        VBox participantsBox = new VBox(participantsLabel, participantsList, pointsLimit, roundsLimit);
        participantsBox.setPadding(new Insets(30, 40, 0, 40));
        participantsBox.setSpacing(30);
        participantsBox.setStyle("-fx-background-color: black; -fx-border-color: #010101; -fx-border-width: 3px;");

        // Botones
        Button playButton = new Button("JUGAR");
        playButton.setOnAction(e -> {
            if (juego.devolverJugadores().size() < 2) {
                Alert alert = new Alert(Alert.AlertType.WARNING);
                alert.setTitle("Jugadores insuficientes");
                alert.setHeaderText(null);
                alert.setContentText("Se requieren al menos 2 jugadores para iniciar el juego.");
                alert.showAndWait();
            } else {
                sceneController.iniciarJuego(Integer.parseInt(roundsLimitInput.getText()), Integer.parseInt(pointsLimitInput.getText()));
            }
        });
        playButton.setStyle("-fx-font-size: 18px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 10px;" );

        Button rulesButton = new Button("Reglas");
        rulesButton.setOnAction(e -> sceneController.switchToReglasScene(juego));
        rulesButton.setStyle("-fx-font-size: 18px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 10px;");

        Button addPlayerButton = new Button("Agregar Jugador");
        addPlayerButton.setOnAction(e -> sceneController.switchToAñadirJugadorScene(juego));
        addPlayerButton.setStyle("-fx-font-size: 18px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 10px;");

        Button exitButton = new Button("Salir");
        exitButton.setOnAction(e -> Platform.exit());
        exitButton.setStyle("-fx-font-size: 18px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 10px;");


        playButton.setMaxWidth(Double.MAX_VALUE);
        rulesButton.setMaxWidth(Double.MAX_VALUE);
        addPlayerButton.setMaxWidth(Double.MAX_VALUE);
        exitButton.setMaxWidth(Double.MAX_VALUE);

        VBox buttonsBox = new VBox(playButton, rulesButton, addPlayerButton, exitButton);
        buttonsBox.setPadding(new Insets(50));
        buttonsBox.setSpacing(20);
        BorderPane.setMargin(buttonsBox, new Insets(20, 20, 20, 40));

        // Layout principal
        BorderPane root = new BorderPane();
        root.setTop(titleLabel);
        root.setLeft(participantsBox);
        root.setRight(buttonsBox);
        BorderPane.setMargin(titleLabel, new Insets(20, 20, 0, 400));
        BorderPane.setMargin(participantsBox, new Insets(20, 20, 20, 100));
        BorderPane.setMargin(buttonsBox, new Insets(90));

        Scene scene = new Scene(root, 1000, 650);
        scene.getStylesheets().add(getClass().getResource("/styles.css").toExternalForm());

        return scene;
    }

}
