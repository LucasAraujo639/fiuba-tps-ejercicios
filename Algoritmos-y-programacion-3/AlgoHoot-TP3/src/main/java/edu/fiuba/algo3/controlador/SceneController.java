package edu.fiuba.algo3.controlador;

import edu.fiuba.algo3.vista.*;
import edu.fiuba.algo3.vista.PreguntasScene.*;
import edu.fiuba.algo3.vista.componentes.BonificadorBoton;
import edu.fiuba.algo3.modelo.Bonificador.BonificadorDecorador;
import edu.fiuba.algo3.modelo.Juego;
import edu.fiuba.algo3.modelo.Jugador;
import edu.fiuba.algo3.modelo.Pregunta;
import edu.fiuba.algo3.modelo.TipoDePregunta.GroupChoice;
import edu.fiuba.algo3.modelo.TipoDePregunta.MultipleChoice;
import edu.fiuba.algo3.modelo.TipoDePregunta.OrderedChoice;
import edu.fiuba.algo3.modelo.TipoDePregunta.VerdaderoFalso;
import javafx.geometry.Insets;
import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Alert;
import javafx.scene.control.Button;
import javafx.scene.layout.VBox;
import javafx.stage.Stage;

import java.util.ArrayList;
import java.util.List;

public class SceneController {

    private Stage primaryStage;
    private Juego juego;

    public SceneController(Stage primaryStage) {
        this.primaryStage = primaryStage;
    }

    //Escenas de preguntas
    public void switchToJugarVerdaderoFalsoScene() {
        EscenaDePregunta jugarScene = new JugarVerdaderoFalsoScene(this, juego);
        Scene jScene = jugarScene.getScene();
        primaryStage.setScene(jScene);
    }

    public void switchToJugarGroupChociceScene() {
        JugarGroupChoiceScene jugarScene = new JugarGroupChoiceScene(this, juego);
        Scene jScene = jugarScene.getScene();
        primaryStage.setScene(jScene);
    }

    public void switchToJugarMultipleChoiceScene() {
        JugarMultipleChoiceScene jugarScene = new JugarMultipleChoiceScene(this, juego);
        Scene jScene = jugarScene.getScene();
        primaryStage.setScene(jScene);
    }

    public void switchToJugarOrderChoiceScene() {
        JugarOrderChoiceScene jugarScene = new JugarOrderChoiceScene(this, juego);
        Scene jScene = jugarScene.getScene();
        primaryStage.setScene(jScene);
    }

    public void switchToJugarPregunta(){
        if(juego.getPreguntaActual().getTipo() instanceof VerdaderoFalso){
            switchToJugarVerdaderoFalsoScene();
        }
        if(juego.getPreguntaActual().getTipo() instanceof GroupChoice){
            switchToJugarGroupChociceScene();
        }
        if(juego.getPreguntaActual().getTipo() instanceof MultipleChoice){
            switchToJugarMultipleChoiceScene();
        }
        if(juego.getPreguntaActual().getTipo() instanceof OrderedChoice){
            switchToJugarOrderChoiceScene();
        }
    }

    // Menu + Reglas + Añadir Jugadores
    public void switchToMenuScene() {
        MenuScene menuScene = new MenuScene(this, juego);
        Scene mScene = menuScene.getScene();
        primaryStage.setScene(mScene);
    }

    public void switchToReglasScene(Juego juego) {
        this.juego = juego;
        ReglasScene reglasScene = new ReglasScene(this);
        Scene rScene = reglasScene.getScene();
        primaryStage.setScene(rScene);
    }

    public void switchToAñadirJugadorScene(Juego juego) {
        this.juego = juego;
        AñadirJugadorScene añadirJugadorScene = new AñadirJugadorScene(this);
        Scene ajScene = añadirJugadorScene.getScene(juego);
        primaryStage.setScene(ajScene);
    }

    public void iniciarJuego(int cantidadRondas, int puntajeGanador){
        juego.iniciarJuego(cantidadRondas, puntajeGanador);
        switchToTurnosScene();
    }


    // la idea seria, mandar una pregunta y que respondan los cuatro
    public void siguienteTurno(){
        Pregunta preguntaActual = juego.getPreguntaActual();
        if(preguntaActual.respondieronTodos(juego.devolverJugadores().size())){
            preguntaActual.validarRespuestas();
            switchToMostrarPuntajesScene();
            juego.setearPreguntaActual();
        } else {
            switchToTurnosScene();
        }
    }

    // Turnos
    public void switchToTurnosScene() {
        List<Jugador> jugadores = juego.devolverJugadores();
        for(Jugador jugador: jugadores){
            // Verificar si se cumple condicion de corte
            if (jugador.getPuntajeTotal() >= 3) {
                this.switchToGanadorScene();
                return;
            }
        }
        TurnosScene turnosScene = new TurnosScene(this, juego);
        Scene scene = turnosScene.getScene();
        System.out.println("EVENTO PUNTAJES ACITVADO");
        primaryStage.setScene(scene);
    }

    public void switchToMostrarPuntajesScene(){
        MostrarPuntajesScene mostrarPuntajesScene = new MostrarPuntajesScene(this, juego);
        Scene scene = mostrarPuntajesScene.getScene();
        primaryStage.setScene(scene);
    }

    public void switchToGanadorScene(){
        MostrarGanadorScene mostrarGanadorScene = new MostrarGanadorScene(this, juego);
        Scene scene = mostrarGanadorScene.getScene();
        primaryStage.setScene(scene);
    }

    public void addPlayer(String playerName) {
        Jugador jugador = new Jugador(playerName);
        juego.agregarJugador(jugador);
    }

    public void switchToTheStartScene() {
        this.juego = new Juego();
        MenuScene menuScene = new MenuScene(this, juego);
        Scene mScene = menuScene.getScene();
        primaryStage.setScene(mScene);
    }

    public VBox MostrarBonificadores(Pregunta pregunta, Jugador jugador){
        ArrayList<BonificadorDecorador> bonificadores = pregunta.obtenerBonificadoresDisponibles(jugador);
        VBox botonesPenalidad = new VBox(10);
        for(BonificadorDecorador bonificador : bonificadores){
            BonificadorBoton boton = new BonificadorBoton(bonificador, "-fx-font-size: 14px; -fx-background-color: #000000; -fx-text-fill: White;");
            boton.setOnAction(e -> {
                // Usar bonificador
                jugador.usarBonificador(bonificador, pregunta);

                // Mostrar ventana emergente
                Alert alert = new Alert(Alert.AlertType.INFORMATION);
                alert.setTitle("Bonificador Usado");
                alert.setHeaderText(null);
                alert.setContentText("Has usado el bonificador: " + bonificador.getNombreBonificador());
                alert.showAndWait();

                // Deshabilitar todos los botones
                for (javafx.scene.Node node : botonesPenalidad.getChildren()) {
                    if (node instanceof Button) {
                        node.setDisable(true);
                    }
                }
            });
            botonesPenalidad.getChildren().add(boton);
        }
        botonesPenalidad.setAlignment(Pos.CENTER_LEFT);
        botonesPenalidad.setPadding(new Insets(10));
        return botonesPenalidad;
    }

}
