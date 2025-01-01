package edu.fiuba.algo3.vista;

import edu.fiuba.algo3.controlador.SceneController;
import edu.fiuba.algo3.modelo.Bonificador.BonificadorDecorador;
import edu.fiuba.algo3.modelo.Juego;
import edu.fiuba.algo3.modelo.Jugador;
import javafx.geometry.Insets;
import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.layout.BorderPane;
import javafx.scene.layout.VBox;

import java.util.List;

public class MostrarPuntajesScene {

    private SceneController sceneController;
    private Juego juego;

    public MostrarPuntajesScene(SceneController sceneController, Juego juego) {
        this.sceneController = sceneController;
        this.juego = juego;
    }

    public Scene getScene() {
        // Crear el layout principal
        VBox layout = new VBox(20);
        layout.setPadding(new Insets(20));
        layout.setAlignment(Pos.CENTER);

        // Crear la etiqueta del título
        Label tituloLabel = new Label("Puntajes de Jugadores");
        tituloLabel.setStyle("-fx-font-size: 24px; -fx-font-weight: bold; -fx-text-fill: #010101;");
        layout.getChildren().add(tituloLabel);

        // Obtener la lista de jugadores y mostrar sus puntajes
        List<Jugador> jugadores = juego.devolverJugadores();
        for (Jugador jugador : jugadores) {
            Label puntajeLabel = new Label(jugador.getNombre() + ": " + jugador.getPuntajeTotal() + " puntos");
            puntajeLabel.setStyle("-fx-font-size: 18px; -fx-padding: 10px;");
            layout.getChildren().add(puntajeLabel);
            BonificadorDecorador ultimoBonificador = jugador.getUltimoBonificador();
            if (ultimoBonificador != null) {
                Label bonificadorLabel = new Label("Bonificador aplicado: " + ultimoBonificador.getNombreBonificador());
                bonificadorLabel.setStyle("-fx-font-size: 18px; -fx-padding: 10px;");
                layout.getChildren().add(bonificadorLabel);
                jugador.setUltimoBonificador(null);
            }
        }

        // Botón para volver al menú
        Button continuarButton = new Button("Continuar");
        continuarButton.setOnAction(e -> sceneController.switchToTurnosScene());
        continuarButton.setStyle("-fx-font-size: 18px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 10px;");

        layout.getChildren().add(continuarButton);

        // Crear la escena y agregar el layout principal
        BorderPane root = new BorderPane();
        root.setCenter(layout);

        Scene scene = new Scene(root, 1000, 650);
        scene.getStylesheets().add(getClass().getResource("/styles.css").toExternalForm());

        return scene;
    }
}
