package lista

// Interfaz de la lista.
type Lista[T any] interface {

	// EstaVacia informa si la lista está vacía (true) o no (false).
	EstaVacia() bool

	// InsertarPrimero inserta al inicio de la lista el elemento pasado por argumento.
	InsertarPrimero(T)

	// InsertarUltimo inserta al final de la lista el elemento pasado por argumento.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista y retorna su valor.
	BorrarPrimero() T

	// VerPrimero muestra el valor del primer elemento de la lista.
	VerPrimero() T

	// VerUltimo muestra el valor del último elemento de la lista.
	VerUltimo() T

	// Largo muestra el largo de la lista.
	Largo() int

	// Iterar recorre internamente la lista y aplica una función 'visitar' la cual
	// debe cumplir que: devuelve 'true' si desea seguir iterando y 'false' en caso contrario.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un elemento IteradorLista el cual es un iterador externo de la lista.
	// Las primitivas asociadas al mismo se encuentran más adelante.
	Iterador() IteradorLista[T]
}

// Interfaz del iterador externo de la lista.
type IteradorLista[T any] interface {

	// VerActual devuelve el dato contenido en el elemento la lista en el que se encuentra el iterador.
	// Si se intenta utilizar la función cuando no hay un dato para ver se genera un panic con el mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente indica si hay algún elemento para ver.
	HaySiguiente() bool

	// Siguiente hace que el iterador avance al siguiente elemento de la lista.
	// Si se intenta utilizar la función cuando no hay más elementos para iterar se genera un panic con el mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar inserta un elemento nuevo en la lista entre el elemento en el que se encuentra el iterador y el anterior a ese.
	// Luego de insertarlo el iterador se posiciona sobre el elemento insertado.
	Insertar(T)

	// Borrar elimina el elemento de la lista sobre el que se encuentra el iterador.
	// Luego de borrarlo el iterador se posiciona sobre el elemento siguiente al elemento que se borra de la lista.
	// Si se intenta utilizar la función cuando no hay un dato para borrar se genera un panic con el mensaje "El iterador termino de iterar".
	Borrar() T
}
