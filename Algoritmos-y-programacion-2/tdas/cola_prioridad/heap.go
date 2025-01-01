package cola_prioridad

const (
	TAMANO_INICIAL               = 10
	FACTOR_AGRANDAMIENTO_ARREGLO = 2
	FACTOR_REDUCCION_ARREGLO     = 2
	RELACION_ELEM_LONG_ARREGLO   = 4
	ZEROCMP                      = 0
)

// ESTRUCTURAS:

type heapImplementacion[T any] struct {
	nodos    []T
	cantidad int
	funcCmp  func(T, T) int
}

// CÓDIGO:

// CrearHeap crea un heap vacío.
func CrearHeap[T any](funcCmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heapImplementacion[T])
	heap.nodos = make([]T, TAMANO_INICIAL)
	heap.funcCmp = funcCmp

	return heap
}

// CrearHeapArr crea un heap a partir de un slice.
func CrearHeapArr[T any](arr []T, funcCmp func(T, T) int) ColaPrioridad[T] {
	if len(arr) == 0 {
		heap := CrearHeap(funcCmp)
		return heap
	}
	//Para que el arreglo original no sea modificado
	nuevoArr := make([]T, len(arr))
	copy(nuevoArr, arr)

	nuevoHeap := new(heapImplementacion[T])
	heapify(nuevoArr, len(arr), funcCmp)
	nuevoHeap.nodos = nuevoArr
	nuevoHeap.cantidad = len(arr)
	nuevoHeap.funcCmp = funcCmp

	return nuevoHeap
}

// HeapSort implementa el algoritmo de ordenamiento por heap.
func HeapSort[T any](elementos []T, funcCmp func(T, T) int) {

	heapify(elementos, len(elementos), funcCmp)

	for i := len(elementos) - 1; i >= 0; i-- {
		swap(elementos, i, 0)
		downheap(elementos, 0, i, funcCmp)
	}

}

// PRIMITIVAS:

// EstaVacia devuelve true si la la cola se encuentra vacía, false en caso contrario.
func (heap *heapImplementacion[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

// Encolar Agrega un elemento al heap.
func (heap *heapImplementacion[T]) Encolar(nuevoDato T) {
	if heap.cantidad == cap(heap.nodos) {
		heap.redimensionarArreglo(FACTOR_AGRANDAMIENTO_ARREGLO * cap(heap.nodos))
	}

	heap.cantidad++
	heap.nodos[heap.cantidad-1] = nuevoDato

	heap.upheap(heap.cantidad - 1)
}

// VerMax devuelve el elemento con máxima prioridad. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (heap *heapImplementacion[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}

	return heap.nodos[0]
}

// Desencolar elimina el elemento con máxima prioridad, y lo devuelve. Si está vacía, entra en pánico con un
// mensaje "La cola esta vacia"
func (heap *heapImplementacion[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}

	swap(heap.nodos, 0, heap.cantidad-1)
	heap.cantidad--
	downheap(heap.nodos, 0, heap.cantidad, heap.funcCmp)

	if RELACION_ELEM_LONG_ARREGLO*heap.cantidad == cap(heap.nodos) && cap(heap.nodos) > TAMANO_INICIAL {
		heap.redimensionarArreglo(len(heap.nodos) / FACTOR_REDUCCION_ARREGLO)
	}

	return heap.nodos[heap.cantidad]
}

// Cantidad devuelve la cantidad de elementos que hay en la cola de prioridad.
func (heap *heapImplementacion[T]) Cantidad() int {
	return heap.cantidad
}

// FUNCIONES AUXILIARES:

// Redimensiona el arreglo que contiene los datos.
func (heap *heapImplementacion[T]) redimensionarArreglo(longitud int) {
	nuevo := make([]T, longitud)
	copy(nuevo, heap.nodos)
	heap.nodos = nuevo
}

func swap[T any](nodos []T, pos1, pos2 int) {
	nodos[pos1], nodos[pos2] = nodos[pos2], nodos[pos1]
}

// upheap realiza dicha acción para el elemento de una determinada posición en el arreglo de nodos.
func (heap *heapImplementacion[T]) upheap(posActual int) {
	if posActual == 0 {
		return
	}

	posPadre := (posActual - 1) / 2

	if heap.funcCmp(heap.nodos[posPadre], heap.nodos[posActual]) < ZEROCMP {
		swap(heap.nodos, posPadre, posActual)
		heap.upheap(posPadre)
	}
}

// downheap realiza dicha acción para el elemento de una determinada posición en el arreglo de nodos.
func downheap[T any](nodos []T, posActual int, cantidad int, funcCmp func(T, T) int) {
	posHijoIzq := 2*posActual + 1
	posHijoDer := 2*posActual + 2

	if posHijoIzq >= cantidad {
		return
	}
	datoAct := nodos[posActual]
	datoIzq := nodos[posHijoIzq]

	// Si hay dos hijos.
	if posHijoDer < cantidad {
		datoDer := nodos[posHijoDer]

		if funcCmp(datoIzq, datoDer) > ZEROCMP || funcCmp(datoIzq, datoDer) == ZEROCMP {
			if funcCmp(datoIzq, datoAct) > ZEROCMP {
				swap(nodos, posActual, posHijoIzq)
				downheap(nodos, posHijoIzq, cantidad, funcCmp)
			}

		} else {
			if funcCmp(datoDer, datoAct) > ZEROCMP {
				swap(nodos, posActual, posHijoDer)
				downheap(nodos, posHijoDer, cantidad, funcCmp)
			}
		}
		// Si hay un solo hijo (que por ser árbol izquierdista es el izquierdo).
	} else if funcCmp(datoAct, datoIzq) < ZEROCMP {
		swap(nodos, posActual, posHijoIzq)
	}

}

// heapify se fija en todas las posiciones que el elemento cumpla la propiedad de heap, si no lo cumple hace downheap en dichas posiciones
func heapify[T any](arr []T, cantidad int, funcCmp func(T, T) int) {
	for i := len(arr) - 1; i >= 0; i-- {
		downheap(arr, i, cantidad, funcCmp)
	}
}
