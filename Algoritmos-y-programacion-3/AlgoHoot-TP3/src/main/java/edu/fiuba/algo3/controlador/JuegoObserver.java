package edu.fiuba.algo3.controlador;
import edu.fiuba.algo3.modelo.Observable;
import edu.fiuba.algo3.modelo.Observer;

public class JuegoObserver implements Observer {
    private SceneController sceneController;
    public JuegoObserver(SceneController sceneController){
        this.sceneController = sceneController;
    }
    @Override
    public void update(Observable o) {
          sceneController.switchToGanadorScene();
    }
}