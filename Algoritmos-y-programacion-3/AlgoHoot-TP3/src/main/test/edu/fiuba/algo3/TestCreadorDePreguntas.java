package edu.fiuba.algo3;

import org.json.simple.parser.ParseException;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import edu.fiuba.algo3.modelo.CreadorDePreguntas;
import edu.fiuba.algo3.modelo.Pregunta;

import java.io.IOException;
import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.*;

public class TestCreadorDePreguntas {
    @BeforeEach
    public void setUp() throws IOException, ParseException {
        CreadorDePreguntas.leerArchivo();
    }


    @Test
    public void testLeerArchivo() throws IOException, ParseException {
        ArrayList<Pregunta> preguntas = CreadorDePreguntas.leerArchivo();
        for (Pregunta pregunta : preguntas) {
            String enunciado = pregunta.getEnunciado();
            System.out.println(enunciado);
        }
        assertFalse(preguntas.isEmpty(), "La lista de preguntas no debería estar vacía");
    }

}

