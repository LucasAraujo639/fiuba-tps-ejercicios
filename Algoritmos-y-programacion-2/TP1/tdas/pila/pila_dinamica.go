package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

// CONSTANTES:

// Tamaño inicial de la pila.
const TAMANO_INICIAL int = 10

// Factor de agrandamiento del tamaño del arreglo de la pila.
const FACTOR_AGRANDAMIENTO_ARREGLO int = 2

// Factor de reducción del tamaño del arreglo de la pila.
const FACTOR_REDUCCION_ARREGLO int = 2

// Factor que relaciona la cantidad de datos que debe tener la pila con el tamaño del arreglo de la pila
// para determinar si hay que reducir el arreglo.
const RELACION_ELEM_LONG_ARREGLO int = 4

// CÓDIGO:

// CrearPilaDinamica crea un nuevo elemento con interfaz Pila.
func CrearPilaDinamica[T any]() Pila[T] {
	nuevaPila := new(pilaDinamica[T])
	nuevaPila.datos = make([]T, TAMANO_INICIAL)

	return nuevaPila
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	return p.datos[p.cantidad-1]
}

// Redimensiona el arreglo que contiene los datos.
func (p *pilaDinamica[T]) redimensionarArreglo(longitud int) {
	nuevo := make([]T, longitud)
	copy(nuevo, p.datos)
	p.datos = nuevo
}

// Apilar agrega un nuevo elemento a la pila.
func (p *pilaDinamica[T]) Apilar(elem T) {
	if p.cantidad == len(p.datos) {
		p.redimensionarArreglo(FACTOR_AGRANDAMIENTO_ARREGLO * len(p.datos))
	}
	p.cantidad++

	p.datos[p.cantidad-1] = elem
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	p.cantidad--

	if RELACION_ELEM_LONG_ARREGLO*p.cantidad == len(p.datos) && len(p.datos) > TAMANO_INICIAL {
		p.redimensionarArreglo(len(p.datos) / FACTOR_REDUCCION_ARREGLO)
	}

	return p.datos[p.cantidad]
}
