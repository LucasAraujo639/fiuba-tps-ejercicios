package lista

// Nodo de la lista:

// Implementación del nodo.
type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

// crearNodo crea un nodo que almacena datos de tipo T y retorna un puntero al mismo.
func crearNodo[T any](dato T) *nodo[T] {
	nuevoNodo := new(nodo[T])
	nuevoNodo.dato = dato
	return nuevoNodo
}

// Lista:

// Implementación de la lista.
type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

// CrearListaEnlazada crea e inicializa una lista enlazada
func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

// EstaVacia informa si la lista está vacía (true) o no (false).
func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil && lista.largo == 0
}

func (lista *listaEnlazada[T]) validarVacio() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

// InsertarPrimero inserta al inicio de la lista el elemento pasado por argumento.
func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevoNodo := crearNodo[T](elem)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
	}
	lista.primero = nuevoNodo
	lista.largo++
}

// InsertarUltimo inserta al final de la lista el elemento pasado por argumento.
func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	nuevoNodo := crearNodo[T](elem)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++

}

// BorrarPrimero elimina el primer elemento de la lista y retorna su valor.
func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.validarVacio()
	auxPrimero := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--

	if lista.largo == 0 {
		lista.ultimo = nil
	}
	return auxPrimero
}

// VerPrimero muestra el valor del primer elemento de la lista.
func (lista *listaEnlazada[T]) VerPrimero() T {
	lista.validarVacio()
	return lista.primero.dato
}

// VerUltimo muestra el valor del último elemento de la lista.
func (lista *listaEnlazada[T]) VerUltimo() T {
	lista.validarVacio()
	return lista.ultimo.dato
}

// Largo muestra el largo de la lista.
func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

// Iterar recorre internamente la lista y aplica una función 'visitar' la cual
// debe cumplir que: devuelve 'true' si desea seguir iterando y 'false' en caso contrario.
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	punteroIter := lista.primero
	for punteroIter != nil {
		if !visitar(punteroIter.dato) {
			break
		}
		punteroIter = punteroIter.siguiente
	}
}

// Iterador devuelve un elemento IteradorLista el cual es un iterador externo de la lista.
// Las primitivas asociadas al mismo se encuentran más adelante.
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	nuevoIter := new(iteradorLista[T])
	nuevoIter.actual = lista.primero
	nuevoIter.listaAsociada = lista
	return nuevoIter
}

// Iterador externo de la lista:

// Implementación del iterador.
type iteradorLista[T any] struct {
	anterior      *nodo[T]
	actual        *nodo[T]
	listaAsociada *listaEnlazada[T]
}

// VerActual devuelve el dato contenido en el elemento la lista en el que se encuentra el iterador.
func (iter *iteradorLista[T]) VerActual() T {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

// HaySiguiente indica si hay algún elemento para ver.
func (iter *iteradorLista[T]) HaySiguiente() bool {
	return iter.actual != nil
}

// Siguiente hace que el iterador avance al siguiente elemento de la lista.
func (iter *iteradorLista[T]) Siguiente() {

	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente

}

// Insertar inserta un elemento nuevo en la lista entre el elemento en el que se encuentra el iterador y el anterior a ese.
// Luego de insertarlo el iterador se posiciona sobre el elemento insertado.
func (iter *iteradorLista[T]) Insertar(elem T) {
	nuevoNodo := crearNodo[T](elem)

	if iter.listaAsociada.EstaVacia() { // Si la lista está vacía.
		iter.listaAsociada.primero = nuevoNodo
		iter.listaAsociada.ultimo = nuevoNodo
	} else if iter.anterior == nil && iter.actual != nil { // Si se inserta al comienzo de la lista.
		nuevoNodo.siguiente = iter.listaAsociada.primero
		iter.listaAsociada.primero = nuevoNodo
	} else { // Si se inserta en otro lado.
		iter.anterior.siguiente = nuevoNodo

		if iter.actual == nil { // Si se inserta en el final.
			iter.listaAsociada.ultimo = nuevoNodo
		} else { // Si no se inserta en los extremos.
			nuevoNodo.siguiente = iter.actual
		}
	}
	iter.actual = nuevoNodo

	// Se aumenta la cantidad de elementos de la lista.
	iter.listaAsociada.largo++
}

// Borrar elimina el elemento de la lista sobre el que se encuentra el iterador.
// Luego de borrarlo el iterador se posiciona sobre el elemento siguiente al elemento que se borra de la lista.
func (iter *iteradorLista[T]) Borrar() T {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}

	auxNodo := iter.actual

	if iter.anterior == nil { // Si se encuentra al inicio de la lista.
		iter.listaAsociada.primero = iter.actual.siguiente

		if iter.actual == iter.listaAsociada.ultimo { // Si la lista tiene un único elemento
			iter.listaAsociada.ultimo = iter.actual.siguiente
		}
	} else { // Si no se encuentra al inicio de la lista.
		iter.anterior.siguiente = iter.actual.siguiente

		if iter.actual == iter.listaAsociada.ultimo { // Si se encuentra al final de la lista.
			iter.listaAsociada.ultimo = iter.anterior
		}
	}

	iter.actual = auxNodo.siguiente

	// Disminuye la cantidad de elementos de la lista.
	iter.listaAsociada.largo--

	return auxNodo.dato
}
