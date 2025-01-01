package diccionario

import (
	TDAPila "tdas/pila"
)

const (
	_ZEROCMP          = 0
	_CANTIDAD_INICIAL = 0
)

type arbol[K comparable, V any] struct {
	cant    int
	funcCmp func(K, K) int
	nodos   *arbolNodo[K, V]
}

type arbolNodo[K comparable, V any] struct {
	izq, der *arbolNodo[K, V]
	clave    K
	dato     V
}

// CrearABB devuelve un ABB funcional.
func CrearABB[K comparable, V any](funcCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	nuevoArbol := new(arbol[K, V])
	nuevoArbol.funcCmp = funcCmp

	return nuevoArbol
}

// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
func (abb *arbol[K, V]) Guardar(clave K, dato V) {

	/*if abb.cant == _CANTIDAD_INICIAL {
		abb.cant++
		abb.nodos = abb.crearNodo(clave, dato)
		return
	}*/

	unionPadreHijo := abb.nodos.buscarNodo(clave, &abb.nodos, abb.funcCmp)

	if *unionPadreHijo != nil {
		(*unionPadreHijo).dato = dato
	} else {
		nuevoNodo := abb.crearNodo(clave, dato)
		(*unionPadreHijo) = nuevoNodo
		abb.cant++
	}
}

// Pertenece determina si una clave ya se encuentra en el diccionario, o no
func (abb *arbol[K, V]) Pertenece(clave K) bool {
	return *abb.nodos.buscarNodo(clave, &abb.nodos, abb.funcCmp) != nil
}

// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje 'La clave no pertenece al diccionario'
func (abb *arbol[K, V]) Obtener(clave K) V {
	unionPadreHijo := abb.nodos.buscarNodo(clave, &abb.nodos, abb.funcCmp)

	if *unionPadreHijo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*unionPadreHijo).dato
}

// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'
func (abb *arbol[K, V]) Borrar(clave K) V {

	unionPadreHijo := abb.nodos.buscarNodo(clave, &abb.nodos, abb.funcCmp)

	if *unionPadreHijo == nil {
		panic("La clave no pertenece al diccionario")
	}

	if abb.Cantidad() == 1 {
		datoBorrado := abb.nodos.dato
		abb.nodos = nil
		abb.cant--
		return datoBorrado
	}

	abb.cant--
	if *unionPadreHijo == abb.nodos {
		return abb.nodos.borrarRaiz()
	}

	return abb.nodos.borrarNodo(unionPadreHijo)
}

// Cantidad devuelve la cantidad de elementos dentro del diccionario
func (abb *arbol[K, V]) Cantidad() int {
	return abb.cant
}

// Iterar itera internamente el diccionario, aplicando la función pasada por parámetro a todos los elementos del mismo
func (abb *arbol[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	if abb.Cantidad() == _CANTIDAD_INICIAL {
		return
	}
	abb.nodos.iterarRangoRec(nil, nil, visitar, abb.funcCmp)

}

// IterarRango itera sólo incluyendo a los elementos que se encuentren comprendidos en el rango indicado, incluyéndolos en caso de encontrarse
func (abb *arbol[K, V]) IterarRango(desde, hasta *K, visitar func(clave K, dato V) bool) {
	if abb.Cantidad() == _CANTIDAD_INICIAL {
		return
	}
	abb.nodos.iterarRangoRec(desde, hasta, visitar, abb.funcCmp)
}

func (nodo *arbolNodo[K, V]) iterarRangoRec(desde, hasta *K, visitar func(clave K, dato V) bool, funcCmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}

	if desde != nil && funcCmp(nodo.clave, *desde) < _ZEROCMP {
		return nodo.der.iterarRangoRec(desde, hasta, visitar, funcCmp)
	}

	if hasta != nil && funcCmp(nodo.clave, *hasta) > _ZEROCMP {
		return nodo.izq.iterarRangoRec(desde, hasta, visitar, funcCmp)
	}

	if !nodo.izq.iterarRangoRec(desde, hasta, visitar, funcCmp) || !visitar(nodo.clave, nodo.dato) {
		return false
	}
	return nodo.der.iterarRangoRec(desde, hasta, visitar, funcCmp)
}

// --------------------------------------------
// ----- Funciones de manejo de los nodos -----
// --------------------------------------------

// crearNodo crea e inicializa un nodo.
func (abb *arbol[K, V]) crearNodo(clave K, dato V) *arbolNodo[K, V] {
	nuevoNodo := new(arbolNodo[K, V])
	nuevoNodo.clave = clave
	nuevoNodo.dato = dato

	return nuevoNodo
}

// buscarNodo busca la unión entre el nodo padre y el hijo con la clave pasada por argumento.
func (nodo *arbolNodo[K, V]) buscarNodo(clave K, mejorCandidato **arbolNodo[K, V], funcCmp func(K, K) int) **arbolNodo[K, V] {
	if nodo == nil {
		return mejorCandidato
	}

	if funcCmp(nodo.clave, clave) > 0 {
		return nodo.izq.buscarNodo(clave, &nodo.izq, funcCmp)
	}
	if funcCmp(nodo.clave, clave) < 0 {
		return nodo.der.buscarNodo(clave, &nodo.der, funcCmp)
	}

	return mejorCandidato
}

// reemplazarDatos reemplaza la clave y el dato del nodo que invoca la primitiva por los almacenados del nodo pasado por argumento.
func (nodo *arbolNodo[K, V]) reemplazarDatos(nodoReemplazante *arbolNodo[K, V]) {
	nodo.clave = nodoReemplazante.clave
	nodo.dato = nodoReemplazante.dato
}

// reemplazarHijos reemplaza los hijos del nodo que invoca la función por los del nodo pasado por argumento.
func (nodo *arbolNodo[K, V]) reemplazarHijos(nodoReemplazante *arbolNodo[K, V]) {
	nodo.izq = nodoReemplazante.izq
	nodo.der = nodoReemplazante.der
}

// Funciones de borrado de la raiz.
func (nodo *arbolNodo[K, V]) borrarRaiz() V {

	datoBorrado := nodo.dato

	if nodo.izq != nil && nodo.der == nil {
		nodo.reemplazarDatos(nodo.izq)
		nodo.reemplazarHijos(nodo.izq)
	} else if nodo.izq == nil && nodo.der != nil {
		nodo.reemplazarDatos(nodo.der)
		nodo.reemplazarHijos(nodo.der)
	} else {
		reemplazante := nodo.der.buscarReemplazante(&nodo.der)

		//nodo.modificarAlturas((*reemplazante).clave, -1)

		nodo.reemplazarDatos((*reemplazante))

		nodo.borrarNodo(reemplazante)
	}
	return datoBorrado
}

// Funciones de borrado de nodos distintos a la raiz.
func (nodo *arbolNodo[K, V]) borrarNodo(unionPadreHijo **arbolNodo[K, V]) V {
	if (*unionPadreHijo).izq == nil && (*unionPadreHijo).der == nil {
		return nodo.borrarHoja(unionPadreHijo)
	} else if (*unionPadreHijo).izq != nil && (*unionPadreHijo).der != nil {
		return nodo.borrarNodoDosHijos(unionPadreHijo)
	}

	return nodo.borrarNodoUnHijo(unionPadreHijo)
}

// borrarHoja borra la hoja almacenada en la posición de memoria pasada por parámetro y devuelve el dato almacenado en ella.
func (nodo *arbolNodo[K, V]) borrarHoja(unionPadreHijo **arbolNodo[K, V]) V {
	datoBorrado := (*unionPadreHijo).dato

	(*unionPadreHijo) = nil

	return datoBorrado
}

// borrarNodoUnHijoDer borra un nodo con un único hijo, que es el nodo almacenado en la posición de memoria pasada por parámetro,
// y devuelve el dato almacenado en él.
func (nodo *arbolNodo[K, V]) borrarNodoUnHijo(unionPadreHijo **arbolNodo[K, V]) V {
	datoBorrado := (*unionPadreHijo).dato

	if (*unionPadreHijo).izq != nil && (*unionPadreHijo).der == nil {
		(*unionPadreHijo) = (*unionPadreHijo).izq
	} else {
		(*unionPadreHijo) = (*unionPadreHijo).der
	}

	return datoBorrado
}

// borrarNodoDosHijos borra un nodo con dos hijos, que es el nodo almacenado en la posición de memoria pasada por parámetro,
// y devuelve el dato almacenado en él.
func (nodo *arbolNodo[K, V]) borrarNodoDosHijos(unionPadreHijo **arbolNodo[K, V]) V {

	// Guarda el dato del nodo a borrar.
	datoBorrado := (*unionPadreHijo).dato

	// Se busca un nodo reemplazante al nodo a borrar.
	unionReemplazante := (*unionPadreHijo).der.buscarReemplazante(&(*unionPadreHijo).der)

	// Copia el dato y la clave del mejor reemplazante en el nodo a sustituir.
	(*unionPadreHijo).clave = (*unionReemplazante).clave
	(*unionPadreHijo).dato = (*unionReemplazante).dato

	// Borra el nodo mejor reemplazante.
	if (*unionReemplazante).izq == nil && (*unionReemplazante).der == nil {
		nodo.borrarHoja(unionReemplazante)
	} else {
		nodo.borrarNodoUnHijo(unionReemplazante)
	}

	return datoBorrado
}

// buscarReemplazante busca el nodo que reempĺazará al
func (nodo *arbolNodo[K, V]) buscarReemplazante(candidato **arbolNodo[K, V]) **arbolNodo[K, V] {
	if nodo.izq == nil {
		return candidato
	}

	return nodo.izq.buscarReemplazante(&nodo.izq)
}

// ----------------------------------------
// ------------- iter externo-------------
// ----------------------------------------

type iterAbb[K comparable, V any] struct {
	pila    TDAPila.Pila[*arbolNodo[K, V]]
	desde   *K
	hasta   *K
	funcCmp func(K, K) int
}

// Iterador devuelve un IterDiccionario para este Diccionario
func (abb *arbol[K, V]) Iterador() IterDiccionario[K, V] {
	iter := abb.IteradorRango(nil, nil)
	return iter
}

// Iterador Rango crea un IterDiccionario que sólo itere por las claves que se encuentren en el rango indicado
func (abb *arbol[K, V]) IteradorRango(desde, hasta *K) IterDiccionario[K, V] {
	iteradorRango := new(iterAbb[K, V])

	iteradorRango.pila = TDAPila.CrearPilaDinamica[*arbolNodo[K, V]]()
	iteradorRango.desde = desde
	iteradorRango.hasta = hasta
	iteradorRango.funcCmp = abb.funcCmp

	if abb.cant == _CANTIDAD_INICIAL {
		return iteradorRango
	}
	iteradorRango.apilarTodoIzqRango(abb.nodos)
	return iteradorRango
}

// Apilar todos los nodos hacia la izquierda si esta en el rango
func (iter *iterAbb[K, V]) apilarTodoIzqRango(nodo *arbolNodo[K, V]) {
	if nodo == nil {
		return
	}

	if iter.desde != nil && iter.funcCmp(nodo.clave, *iter.desde) < _ZEROCMP {
		iter.apilarTodoIzqRango(nodo.der)
		return
	}
	if iter.hasta != nil && iter.funcCmp(nodo.clave, *iter.hasta) > _ZEROCMP {
		iter.apilarTodoIzqRango(nodo.izq)
		return
	}
	iter.pila.Apilar(nodo)
	iter.apilarTodoIzqRango(nodo.izq)

}

// HaySiguiente devuelve si hay más datos para ver. Esto es, si en el lugar donde se encuentra parado el iterador hay un elemento.
func (iter iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()

}

// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador. Si no HaySiguiente, debe entrar en pánico con el mensaje 'El iterador termino de iterar'
func (iter iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

// Siguiente si HaySiguiente, devuelve la clave actual (equivalente a VerActual, pero únicamente la clave), y además avanza al siguiente elemento en el diccionario. Si no HaySiguiente, entonces debe entrar en pánico con mensaje 'El iterador termino de iterar'
func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodo := iter.pila.Desapilar()
	if nodo.der != nil {
		iter.apilarTodoIzqRango(nodo.der)
	}

}
