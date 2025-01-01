package votos

import (
	"bufio"
	"os"
	"rerepolez/errores"
	"strconv"
	TDACola "tdas/cola"
)

type campoPadron struct {
	dni    int
	yaVoto bool
}

type padronElectoralImplementacion struct {
	padron *[]campoPadron
}

func CrearPadronElectoral() PadronElectoral {
	return new(padronElectoralImplementacion)
}

func (padronElectoral *padronElectoralImplementacion) CargarPadron(archivo *os.File) error {
	if archivo == nil {
		return errores.ErrorLeerArchivo{}
	}

	padronElectoral.padron = obtenerPadronElectoral(archivo)

	return nil
}

func (padronElectoral *padronElectoralImplementacion) PerteneceAlPadron(dni int) bool {
	res, _ := buscarDNI(padronElectoral.padron, dni)
	return res
}

func (padronElectoral *padronElectoralImplementacion) YaVoto(dni int) bool {
	res, campo := buscarDNI(padronElectoral.padron, dni)
	if !res {
		return false
	}

	return campo.yaVoto
}

func (padronElectoral *padronElectoralImplementacion) FinVoto(dni int) error {
	res, campo := buscarDNI(padronElectoral.padron, dni)
	if !res {
		return errores.DNIFueraPadron{}
	}
	campo.yaVoto = true
	return nil
}

// obtenerPadronElectoral devuelve un puntero a un slice con los votantes. Se asume que los elementos del archivo pasado son correctos y no se repiten.
func obtenerPadronElectoral(archivoPadron *os.File) *[]campoPadron {
	// Cerrar archivo luego de terminar de usar la función.
	defer archivoPadron.Close()

	// Leer archivo.
	escaner := bufio.NewScanner(archivoPadron)

	padronElectoral := make([]campoPadron, 0)

	for escaner.Scan() {
		dni, _ := strconv.Atoi(escaner.Text())
		campoNuevo := new(campoPadron)
		campoNuevo.dni = dni
		padronElectoral = append(padronElectoral, *campoNuevo)
	}

	// Ordenar padrón electoral.
	ordenarPadronElectoral(&padronElectoral)

	return &padronElectoral
}

// -------------------------------------------------------------------------------------------------------------
//                               Función de búsqueda DNI en padrón electoral
// -------------------------------------------------------------------------------------------------------------

// buscarDNI busca el DNI en el padrón utilizando una búsqueda binaria.
func buscarDNI(padronElectoral *[]campoPadron, dniBuscado int) (bool, *campoPadron) {
	inicio := 0
	fin := len(*padronElectoral) - 1
	for inicio <= fin {
		medio := (inicio + fin) / 2
		campoPadronMedio := &(*padronElectoral)[medio]

		if campoPadronMedio.dni < dniBuscado {
			inicio = medio + 1
		} else if campoPadronMedio.dni > dniBuscado {
			fin = medio - 1
		} else {
			return true, campoPadronMedio
		}
	}

	return false, nil
}

// -------------------------------------------------------------------------------------------------------------
//                          Función de ordenamiento de datos de padrón electoral
// -------------------------------------------------------------------------------------------------------------

const (
	CANT_ELEMENTOS_ORDENABLES_DNI = 10
	CANT_MAX_DIGITOS_DNI          = 8
)

// Función que ordena el padrón electoral utilizando el algoritmo de ordenamiento Radix.
func ordenarPadronElectoral(padronElectoral *[]campoPadron) {
	if len(*padronElectoral) == 1 {
		return
	}

	// Iterar el proceso de ordenamiento por la cantidad de dígitos posibles del DNI.
	for i := 0; i < CANT_MAX_DIGITOS_DNI; i++ {
		countingSortNumeroDNI(padronElectoral, i)
	}
}

// Función de ordenamiento auxiliar.
func countingSortNumeroDNI(padronElectoral *[]campoPadron, posicionDNI int) {
	if len(*padronElectoral) == 1 {
		return
	}

	// Creo variables auxiliares.
	frecuencias := make([]TDACola.Cola[campoPadron], CANT_ELEMENTOS_ORDENABLES_DNI)

	for i := 0; i < CANT_ELEMENTOS_ORDENABLES_DNI; i++ {
		frecuencias[i] = TDACola.CrearColaEnlazada[campoPadron]()
	}

	// Calculo potencia de 10 a partir de la cual se comienza a
	base := 1
	for i := 0; i < posicionDNI; i++ {
		base *= 10
	}

	// Calcular frecuencias:
	for _, campo := range *padronElectoral {
		frecuencias[dniAPosicion(campo.dni, base)].Encolar(campo)
	}

	// Ordenar nuevamente los DNI:
	indiceFrec := 0
	for i := 0; i < len(*padronElectoral); i++ {
		for frecuencias[indiceFrec].EstaVacia() {
			indiceFrec++
		}
		(*padronElectoral)[i] = frecuencias[indiceFrec].Desencolar()
	}
}

// byteAPosicion le asigna a cada caracter numérico del DNI una posición en el arreglo de frecuencias.
func dniAPosicion(dni int, base int) int {
	return (dni / base) % CANT_ELEMENTOS_ORDENABLES_DNI
}
