package edu.fiuba.algo3.modelo;

import edu.fiuba.algo3.modelo.Opcion.OpcionGrupo;
import edu.fiuba.algo3.modelo.Opcion.OpcionSimple;
import edu.fiuba.algo3.modelo.Penalidad.ConPenalidad;
import edu.fiuba.algo3.modelo.Penalidad.Penalidad;
import edu.fiuba.algo3.modelo.Penalidad.PenalidadParcial;
import edu.fiuba.algo3.modelo.Penalidad.SinPenalidad;
import edu.fiuba.algo3.modelo.TipoDePregunta.*;
import org.json.simple.JSONArray;
import org.json.simple.JSONObject;
import org.json.simple.parser.JSONParser;
import org.json.simple.parser.ParseException;

import java.io.FileReader;
import java.io.IOException;
import java.util.*;

public class CreadorDePreguntas {
    HashMap<String, TipoDePregunta> tipoDePreguntas;

    public CreadorDePreguntas(){
    }

    private static OpcionSimple obtenerOpcionPorId(ArrayList<OpcionSimple> opciones, int id){
        for(OpcionSimple opcion : opciones){
            if(opcion.getId() == id){
                return opcion;
            }
        }
        return null;
    }

    private static ArrayList<OpcionSimple> obtenerOpcionesCorrectas(ArrayList<OpcionSimple> opciones, String correctas){
        ArrayList<OpcionSimple> opcionesCorrectas = new ArrayList<>();
        String[] correctasArray = correctas.split(",");
        for(String correcta : correctasArray) {
            opcionesCorrectas.add(obtenerOpcionPorId(opciones,Integer.parseInt(correcta)));
        }
        return opcionesCorrectas;
    }

    private static HashSet<OpcionSimple> obtenerOpcionesCorrectasGrupo(ArrayList<OpcionSimple> opciones,String correctas){
        HashSet<OpcionSimple> opcionesCorrectas = new HashSet<>();
        String[] correctasArray = correctas.split(":")[1].trim().split(",");
        for(String correcta : correctasArray){
            opcionesCorrectas.add(obtenerOpcionPorId(opciones,Integer.parseInt(correcta)));
        }
        return opcionesCorrectas;
    }

    private static ArrayList<OpcionSimple> obtenerOpciones(JSONObject preguntaJSON){
        ArrayList<OpcionSimple> opciones = new ArrayList<>();
        OpcionSimple opcion1 = new OpcionSimple(String.valueOf(preguntaJSON.get("Opcion 1")), 1);
        opciones.add(opcion1);
        OpcionSimple opcion2 = new OpcionSimple(String.valueOf(preguntaJSON.get("Opcion 2")), 2);
        opciones.add(opcion2);
        if(preguntaJSON.containsKey("Opcion 3")){
            OpcionSimple opcion3 = new OpcionSimple(String.valueOf(preguntaJSON.get("Opcion 3")), 3);
            opciones.add(opcion3);
        }
        if(preguntaJSON.containsKey("Opcion 4")){
            OpcionSimple opcion4 = new OpcionSimple(String.valueOf(preguntaJSON.get("Opcion 4")), 4);
            opciones.add(opcion4);
        }
        if(preguntaJSON.containsKey("Opcion 5")){
            OpcionSimple opcion5 = new OpcionSimple(String.valueOf(preguntaJSON.get("Opcion 5")), 5);
            opciones.add(opcion5);
        }
        if(preguntaJSON.containsKey("Opcion 6")){
            OpcionSimple opcion6 = new OpcionSimple(String.valueOf(preguntaJSON.get("Opcion 6")), 6);
            opciones.add(opcion6);
        }
        return opciones;
    }

    private static Pregunta crearPreguntaGroupChoice(JSONObject preguntaJSON){
        String enunciado = (String) preguntaJSON.get("Pregunta");
        String nombreGrupoA = (String) preguntaJSON.get("Grupo A");
        String nombreGrupoB = (String) preguntaJSON.get("Grupo B");
        String respuestaCorrecta = (String) preguntaJSON.get("Respuesta");
        String tema = (String) preguntaJSON.get("Tema");
        ArrayList<OpcionSimple> opciones = obtenerOpciones(preguntaJSON);
        String respuestaCorrectaGrupoA = respuestaCorrecta.split(";")[0];
        String respuestaCorrectaGrupoB = respuestaCorrecta.split(";")[1];
        HashSet<OpcionSimple> opcionesCorrectasGrupoA = obtenerOpcionesCorrectasGrupo(opciones, respuestaCorrectaGrupoA);
        HashSet<OpcionSimple> opcionesCorrectasGrupoB = obtenerOpcionesCorrectasGrupo(opciones, respuestaCorrectaGrupoB);
        OpcionGrupo opcionGrupoA = new OpcionGrupo(nombreGrupoA, opcionesCorrectasGrupoA);
        OpcionGrupo opcionGrupoB = new OpcionGrupo(nombreGrupoB, opcionesCorrectasGrupoB);
        ArrayList<OpcionGrupo> gruposCorrectos = new ArrayList<>();
        gruposCorrectos.add(opcionGrupoA);
        gruposCorrectos.add(opcionGrupoB);
        GroupChoice groupChoice = new GroupChoice(opciones,gruposCorrectos);
        Penalidad simple = new SinPenalidad();
        return new Pregunta(groupChoice,simple,enunciado,tema);
    }

    private static Pregunta crearPreguntaVerdaderoFalso(JSONObject preguntaJSON){
        String enunciado = (String) preguntaJSON.get("Pregunta");
        ArrayList<OpcionSimple> opciones = obtenerOpciones(preguntaJSON);
        String respuesta = (String) preguntaJSON.get("Respuesta");
        String tema = (String) preguntaJSON.get("Tema");
        OpcionSimple opcionCorrecta = obtenerOpcionPorId(opciones, Integer.parseInt(respuesta));
        VerdaderoFalso verdaderoFalso = new VerdaderoFalso(opciones, opcionCorrecta);
        String[] tipo = ((String) preguntaJSON.get("Tipo")).split(" ");
        if(tipo[2].equals("Penalidad")){
            return new Pregunta(verdaderoFalso, new ConPenalidad(), enunciado,tema);
        }
        return new Pregunta(verdaderoFalso, new SinPenalidad(), enunciado,tema);
    }

    private static Pregunta crearPreguntaMultipleChoice(JSONObject preguntaJSON){
        String enunciado = (String) preguntaJSON.get("Pregunta");
        ArrayList<OpcionSimple> opciones = obtenerOpciones(preguntaJSON);
        String respuesta = (String) preguntaJSON.get("Respuesta");
        String tema = (String) preguntaJSON.get("Tema");
        ArrayList<OpcionSimple> opcionesCorrectas = obtenerOpcionesCorrectas(opciones, respuesta);
        TipoDePregunta multipleChoice = new MultipleChoice(opciones, opcionesCorrectas);
        String penalidad = ((String) preguntaJSON.get("Tipo")).split(" ")[2];
        if (penalidad.equals("Simple")){
            return new Pregunta(multipleChoice, new SinPenalidad(), enunciado,tema);
        } else if (penalidad.equals("Penalidad")) {
            return new Pregunta(multipleChoice, new ConPenalidad(), enunciado,tema);
        } else {
            return new Pregunta(multipleChoice, new PenalidadParcial(), enunciado,tema);
        }
    }

    private static Pregunta crearPreguntaOrderedChoice(JSONObject preguntaJSON){
        String enunciado = (String) preguntaJSON.get("Pregunta");
        ArrayList<OpcionSimple> opciones = obtenerOpciones(preguntaJSON);
        String respuesta = (String) preguntaJSON.get("Respuesta");
        String tema = (String) preguntaJSON.get("Tema");
        ArrayList<OpcionSimple> opcionesCorrectas = obtenerOpcionesCorrectas(opciones, respuesta);
        TipoDePregunta orderedChoice = new OrderedChoice(opciones, opcionesCorrectas);
        return new Pregunta(orderedChoice, new SinPenalidad(), enunciado,tema);
    }

    public static ArrayList<Pregunta> leerArchivo() throws IOException, ParseException {
        JSONParser parser = new JSONParser();
        ArrayList<Pregunta> preguntas = new ArrayList<>();
        try {
            JSONArray jsonPreguntas = (JSONArray) parser.parse(new FileReader("src/main/java/edu/fiuba/algo3/assets/preguntas.json"));
            for (Object pregunta : jsonPreguntas) {
                Pregunta nuevaPregunta;
                JSONObject preguntaJSON = (JSONObject) pregunta;
                String[] tipo = ((String) preguntaJSON.get("Tipo")).split(" ");
                String tipoDePregunta = tipo[0]+ " " + tipo[1];
                if(tipoDePregunta.equals("Group Choice")){
                    nuevaPregunta = crearPreguntaGroupChoice(preguntaJSON);
                    preguntas.add(nuevaPregunta);
                    System.out.println(nuevaPregunta.getClass());
                }
                if(tipoDePregunta.equals("Verdadero Falso")){
                    nuevaPregunta = crearPreguntaVerdaderoFalso(preguntaJSON);
                    preguntas.add(nuevaPregunta);
                    System.out.println(nuevaPregunta.getClass());
                }
                if(tipoDePregunta.equals("Ordered Choice") || tipoDePregunta.equals("Ordered choice")){
                    nuevaPregunta = crearPreguntaOrderedChoice(preguntaJSON);
                    preguntas.add(nuevaPregunta);
                    System.out.println(nuevaPregunta.getClass());
                }
                if(tipoDePregunta.equals("Multiple Choice")){
                    nuevaPregunta = crearPreguntaMultipleChoice(preguntaJSON);
                    preguntas.add(nuevaPregunta);
                    System.out.println(nuevaPregunta.getClass());
                }
            }
            return preguntas;

        } catch (Exception e){
            e.printStackTrace();
        }
        return preguntas;
    }
}
