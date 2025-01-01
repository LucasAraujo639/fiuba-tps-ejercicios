package cola

// Implementación del nodo.
type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

// Implementación de la cola enlazada.
type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

// crearNodo crea un nodo que almacena datos de tipo T y retorna un puntero al mismo.
func crearNodo[T any](dato T) *nodo[T] {
	nuevoNodo := new(nodo[T])
	nuevoNodo.dato = dato
	return nuevoNodo
}

// CrearColaEnlazada crea e inicializa una cola enlazada.
func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaEnlazada[T])
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil && c.ultimo == nil
}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	return c.primero.dato
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (c *colaEnlazada[T]) Encolar(elem T) {
	nuevoNodo := crearNodo[T](elem)

	if c.EstaVacia() {
		c.primero = nuevoNodo
	} else {
		c.ultimo.siguiente = nuevoNodo
	}
	c.ultimo = nuevoNodo

}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	aux := c.primero

	if c.primero == c.ultimo {
		c.ultimo = nil
	}
	c.primero = c.primero.siguiente

	return aux.dato
}
