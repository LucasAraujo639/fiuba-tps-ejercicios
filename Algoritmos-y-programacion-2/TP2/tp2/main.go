package main

import (
	TDABD "algogram/bd"
	"algogram/comandos"
	"algogram/datos"
	"fmt"
	"os"
)

func main() {
	// Carga de datos...

	// Obtener datos de invocacion.
	params := datos.ObtenerParametrosEjecucion()
	err := datos.VerificarParametrosEjecucion(params)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}
	//Obtengo el archivo pasandole como parametro el nombre
	archivoUsuarios, err := datos.AbrirArchivo(datos.ObtenerNombreArchivoUsuarios(params))
	if err != nil {
		fmt.Println(err)
		return
	}

	//Cargo todos los usuarios en un hash
	dicUsuarios, err := datos.CargarUsuarios(archivoUsuarios)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Creo el TDA UsuariosBD (Base de datos)
	usuariosBD := TDABD.CrearUsuariosBD(dicUsuarios)

	//Creo el TDA postsBD (Base de datos)
	postsBD := TDABD.CrearPostsBD()

	comandos.LectorComandos(usuariosBD, postsBD)

}
