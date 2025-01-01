package edu.fiuba.algo3.vista;

import edu.fiuba.algo3.controlador.SceneController;
import javafx.geometry.Insets;
import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.layout.BorderPane;
import javafx.scene.layout.VBox;
import javafx.scene.control.ScrollPane;

public class ReglasScene {

    private SceneController sceneController;

    public ReglasScene(SceneController sceneController) {
        this.sceneController = sceneController;
    }

    public Scene getScene() {
        // Título
        Label titleLabel = new Label("Reglas");
        titleLabel.setStyle("-fx-font-size: 36px; -fx-font-weight: bold; -fx-text-fill: #010101;");

        // Texto de las reglas
        String reglasTexto =
                "1. Objetivo del Juego:\n" +
                        " Lograr más puntos que el otro jugador respondiendo correctamente las preguntas.\n" +

                        " \n2. Que preguntas me pueden tocar?:\n" +
                        "  - Verdadero/Falso clásico: Asigna un punto por respuesta correcta.\n" +
                        "  - Verdadero/Falso con penalidad: Asigna un punto por respuesta correcta y resta un punto por respuesta incorrecta.\n" +
                        "  - Multiple choice clásico: Asigna un punto por todas las opciones correctas.\n" +
                        "  - Multiple choice con puntaje parcial: Asigna un punto por cada opción correcta siempre que no se seleccione ninguna opción incorrecta.\n" +
                        "  - Multiple choice con penalidad: Asigna un punto por cada opción correcta y resta un punto por cada opción incorrecta.\n" +
                        "  - Ordered choice: Asigna un punto por ordenar las opciones correctamente.\n" +
                        "  - Group choice: Asigna un punto por colocar las opciones en el grupo correcto.\n" +

                        " \n3. Que es un multiplicadores y como usarlo:\n" +
                        "  Cada jugador tiene dos multiplicadores (\"x2\" y \"x3\") que pueden usar en preguntas con penalidad.\n" +
                        "  A la derecha de las preguntas con penalidad te va a aparecer la posibilidad de usarlos.\n" +

                        " \n4. Que es exclusividad de puntaje:\n" +
                        " Cada jugador tiene dos opciones de exclusividad de puntaje para usar en preguntas sin penalidad.\n" +
                        " El jugador podrá asignar la exclusividad de puntaje a la pregunta cuando le llegue.\n" +
                        " Al calcular los puntos, se asignará el doble del puntaje, pero solo en caso de que solo uno de los jugadores haya realizado la opción correcta. Es decir, si los N jugadores eligen la opción correcta, no se asignará puntaje a ninguno. Si uno de los jugadores eligió la opción correcta, ese jugador conseguirá el doble del puntaje.\n" +
                        " Alcanza con que uno de los jugadores asigne la exclusividad de puntaje para que la regla afecte a todos los jugadores en esa pregunta.\n" +
                        " Si por lo menos 2 jugadores asignan exclusividad de puntaje, entonces el efecto se multiplica por la cantidad de jugadores que la asignaron y se asignará el puntaje al jugador que elija la opción correcta (sólo si ningún otro jugador no eligió la opción correcta a su vez).\n" +

                        " \n5. Anulador de puntaje:\n" +
                        "  Cada jugador puede usar una vez un anulador de puntaje para que los otros jugadores no reciban puntos si responden correctamente.";

        Label reglasLabel = new Label(reglasTexto);
        reglasLabel.setStyle("-fx-font-size: 18px; -fx-text-fill: #010101;");
        reglasLabel.setPadding(new Insets(20));
        reglasLabel.setWrapText(true);

        // Botón para volver al menú
        Button backButton = new Button("Volver al Menú");
        backButton.setOnAction(e -> sceneController.switchToMenuScene());
        backButton.setStyle("-fx-font-size: 18px; -fx-background-color: #010101; -fx-text-fill: White; -fx-border-color: #010101; -fx-border-width: 10px;");

        VBox layout = new VBox(20, titleLabel, reglasLabel, backButton);
        layout.setPadding(new Insets(20));
        layout.setAlignment(Pos.CENTER);

        ScrollPane scrollPane = new ScrollPane(layout);
        scrollPane.setFitToWidth(true);

        BorderPane root = new BorderPane();
        root.setCenter(scrollPane);

        Scene scene = new Scene(root, 1000, 650);
        scene.getStylesheets().add(getClass().getResource("/styles.css").toExternalForm());

        return scene;
    }
}
