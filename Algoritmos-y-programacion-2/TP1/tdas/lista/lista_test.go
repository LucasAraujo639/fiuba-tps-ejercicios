package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_VOLUMEN = 1000
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())

	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
}

func TestUnicoValorInsercionPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Prueba insertando un elemento al inicio.
	lista.InsertarPrimero(20)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 20, lista.VerUltimo())

	doble := 0
	lista.Iterar(func(x int) bool {
		doble = 2 * x
		return true
	})

	require.Equal(t, 40, doble)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 20, lista.VerUltimo())
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 20, iter.VerActual())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())

	require.Equal(t, 20, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())

	iter = lista.Iterador()

	require.False(t, iter.HaySiguiente())
}

func TestUnicoValorInsercionUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Prueba insertando un elemento al final.
	lista.InsertarUltimo(33)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 33, lista.VerPrimero())
	require.Equal(t, 33, lista.VerUltimo())

	doble := 0
	lista.Iterar(func(x int) bool {
		doble = 2 * x
		return true
	})

	require.Equal(t, 66, doble)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 33, lista.VerPrimero())
	require.Equal(t, 33, lista.VerUltimo())
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 33, iter.VerActual())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.Equal(t, 33, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())

	iter = lista.Iterador()

	require.False(t, iter.HaySiguiente())
}

func TestMultiplesInsercionesUnicoValor(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarPrimero(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarUltimo(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 2, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarUltimo(3)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarPrimero(4)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
	require.Equal(t, 4, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarUltimo(5)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 5, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

}

func TestInsercionPrimeroPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarPrimero(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	lista.InsertarPrimero(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsercionPrimeroUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarPrimero(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	lista.InsertarUltimo(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsercionUltimoPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarUltimo(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	lista.InsertarPrimero(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsercionUltimoUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarUltimo(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	lista.InsertarUltimo(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func VerPrimeroYUltimo(t *testing.T) {
	t.Log("Insertamos Primero y Despues a lo ultimo y vemos si cumple con el orden deseado")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())

	lista.InsertarPrimero(2)
	lista.InsertarUltimo(3)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
}

func TestInsertarPocosValores(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarPrimero(1)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	require.Equal(t, 1, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	lista.InsertarPrimero(2)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	lista.InsertarUltimo(3)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	lista.InsertarPrimero(4)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	require.Equal(t, 4, lista.BorrarPrimero())

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 4, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 6, lista.VerUltimo())

	lista.InsertarPrimero(7)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 5, lista.Largo())
	require.Equal(t, 7, lista.VerPrimero())
	require.Equal(t, 6, lista.VerUltimo())

	require.Equal(t, 7, lista.BorrarPrimero())
	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 6, lista.VerUltimo())

	require.Equal(t, 5, lista.BorrarPrimero())
	require.Equal(t, 6, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestVolumen(t *testing.T) {
	t.Log("Hacemos pruebas de volumen ")
	lista := TDALista.CrearListaEnlazada[int]()

	//Insertar primero
	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, i, lista.Largo())
		lista.InsertarPrimero(i)
		require.False(t, lista.EstaVacia())

		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, 0, lista.VerUltimo())

	}
	for i := 999; i >= 0; i-- {
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, i, lista.BorrarPrimero())
		require.Equal(t, i, lista.Largo())
	}
	// Verificar si se vacio correctamente
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	//Insertar ultimo
	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, i, lista.Largo())
		lista.InsertarUltimo(i)
		require.EqualValues(t, 0, lista.VerPrimero())
		require.EqualValues(t, i, lista.VerUltimo())

	}
	for i := 0; i < _VOLUMEN; i++ {
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
	// Verificar si se vacio correctamente
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}
func TestListaFloats(t *testing.T) {
	t.Log("Insertamos a la lista floats ")
	lista := TDALista.CrearListaEnlazada[float32]()
	lista.InsertarPrimero(1.13)
	require.EqualValues(t, 1.13, lista.VerPrimero())
	lista.InsertarPrimero(1.15)
	require.EqualValues(t, 1.15, lista.VerPrimero())
	require.EqualValues(t, 1.15, lista.BorrarPrimero())
	require.EqualValues(t, 1.13, lista.VerPrimero())
	require.EqualValues(t, 1.13, lista.BorrarPrimero())

	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
}
func TestListaStrings(t *testing.T) {
	t.Log("Insertamos a la lista strings")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("asd")
	require.EqualValues(t, "asd", lista.VerPrimero())
	lista.InsertarPrimero("asd1")
	require.EqualValues(t, "asd1", lista.VerPrimero())

	require.EqualValues(t, "asd1", lista.BorrarPrimero())
	require.EqualValues(t, "asd", lista.VerPrimero())

	require.EqualValues(t, "asd", lista.BorrarPrimero())

	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
}

// ----Tests iterador externo----
func TestIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIterInsertarYBorrarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	// Se inserta un elemento con la lista vacía.
	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	iter.Insertar(77)

	// Se verifica que el iterador se comporte como se espera
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 77, iter.VerActual())

	// Se itera hasta el final.
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Se verifica el estado de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 77, lista.VerPrimero())
	require.Equal(t, 77, lista.VerUltimo())

	// Se crea un iterador para borrar el elemento guardado.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 77, iter.VerActual())

	// Se borra el elemento guardado.
	require.Equal(t, 77, iter.Borrar())

	// Se verifica que el iterador se comporte como en una lista vacía.
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Se verifica que la lista se comporte como una lista vacía.
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestIterInsertarInicioLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	// Se inserta un elemento con la lista vacía.
	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	iter.Insertar(1)

	// Se verifica que el iterador se comporte como se espera
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	// Se itera hasta el final.
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Se verifica el estado de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	// Se genera otro iterador.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	// Se inserta otro elemento y se verifica que se comporta como se debería.
	iter.Insertar(2)

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	// Se itera hasta el final.
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Se verifica el estado de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	// Se genera otro iterador.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	// Se inserta otro elemento y se verifica que se comporta como se debería.
	iter.Insertar(3)

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())

	// Se itera hasta el final.
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Se verifica el estado de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	// Se recorre una última vez la lista.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIterInsertarFinalLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	// Se inserta un elemento con la lista vacía.
	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	iter.Insertar(1)

	// Se verifica que el iterador se comporte como se espera
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	// Se itera hasta el final.
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Se verifica el estado de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	// Se genera otro iterador.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	// Se itera hasta el final.

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Se inserta otro elemento y se verifica que se comporta como se debería.
	iter.Insertar(2)

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	// Se itera hasta el final nuevamente.

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Se verifica el estado de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	// Se genera otro iterador.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	// Se itera hasta el final.
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Se inserta otro elemento y se verifica que se comporta como se debería.
	iter.Insertar(3)

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())

	// Se itera hasta el final nuevamente.

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Se verifica el estado de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Se recorre una última vez la lista.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestInsercionesYBorrados(t *testing.T) {
	const LIMITE int = 10
	lista := TDALista.CrearListaEnlazada[int]()

	iter := lista.Iterador()

	// Insertar elementos con el iterador
	for i := 1; i <= LIMITE; i++ {
		iter.Insertar(i)
		require.Equal(t, i, iter.VerActual())
		iter.Borrar()
		require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
		require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
		require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	}
}

func TestIterInsercionesPrincipioVolumen(t *testing.T) {
	const LIMITE int = 10000
	lista := TDALista.CrearListaEnlazada[int]()
	elementosEsperados := make([]int, LIMITE)
	elementosVistos := make([]int, LIMITE)
	elementosBorrados := make([]int, LIMITE)

	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Insertar elementos.
	for i := 1; i <= LIMITE; i++ {
		iter.Insertar(i)
		require.True(t, iter.HaySiguiente())
		require.Equal(t, i, iter.VerActual())
		elementosEsperados[i-1] = i
	}

	// Iterar hasta el final.
	for i := LIMITE; i > 0; i-- {
		require.True(t, iter.HaySiguiente())
		require.Equal(t, i, iter.VerActual())
		iter.Siguiente()
	}

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Verifico lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, LIMITE, lista.Largo())
	require.Equal(t, LIMITE, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	// Borrar elementos.
	iter = lista.Iterador()

	for i := LIMITE; i > 0; i-- {
		require.True(t, iter.HaySiguiente())
		require.Equal(t, i, iter.VerActual())
		elementosVistos[i-1] = iter.VerActual()
		elementosBorrados[i-1] = iter.Borrar()
	}

	// Comparar resultados.
	require.Equal(t, elementosEsperados, elementosBorrados)
	require.Equal(t, elementosEsperados, elementosVistos)

	// Verifico que el iterador haya llegado al final.
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Verifico estado de la lista.
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestIterInsercionesFinalVolumen(t *testing.T) {
	const LIMITE int = 10000
	lista := TDALista.CrearListaEnlazada[int]()

	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Insertar elementos.
	for i := 1; i <= LIMITE; i++ {
		iter.Insertar(i)
		require.True(t, iter.HaySiguiente())
		require.Equal(t, i, iter.VerActual())
	}

	//Itero hasta el final
	iter = lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	// Inserto elementos al final del iterador
	for i := LIMITE + 1; i <= LIMITE; i++ {
		iter.Insertar(i)
		iter.Siguiente()
		require.Equal(t, i, iter.VerActual())
	}

	// Verifico que el iterador haya llegado al final.
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

}

func TestIteradorInsertarAlInicio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verifico condiciones iniciales de lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Verifico estar al inicio de la lista:
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	// Inserto elemento al inicio de la lista:
	iter.Insertar(55)
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 55, iter.VerActual())

	// Itero el resto de la lista:
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()

	// Verifico cambios en lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 4, lista.Largo())
	require.Equal(t, 55, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 55, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestIteradorInsertarAlFinal(t *testing.T) {
	t.Log("Inserto al final un elemento cuando el iterador esta al final")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verifico condiciones iniciales de lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Itero:
	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())

	//Inserto al final del iterador
	iter.Insertar(7)
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 7, iter.VerActual())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())

	// Verifico cambios en lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 4, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 7, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 7, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestIteradorInsertarEnElMedio(t *testing.T) {
	t.Log("Inserto en el medio y comprobamos que haya sido entre el anterior y el actual")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verifico condiciones iniciales de lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Itero hasta la mitad de la lista:
	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	//Inserto en el medio
	iter.Insertar(100)
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 100, iter.VerActual())

	// Itero hasta el final:
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())

	// Verifico cambios en lista:
	require.False(t, lista.EstaVacia())
	require.Equal(t, 4, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 100, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}
func TestRemoverElementoInicioIterador(t *testing.T) {
	t.Log("Remover elemento cuando el iterador es creado")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verificar condiciones iniciales de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Crear iterador y verificar estar en el inicio de la pila.
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())

	// Borrar elemento.
	iter.Borrar()

	// Iterar hasta el final de la lista.
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()

	// Verificar condiciones finales.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}
func TestRemoverElementoFinalIterador(t *testing.T) {
	t.Log("Remover elemento cuando el iterador es llega a su fin")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verificar condiciones iniciales de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Crear iterador e iterar hasta estar en el final de la lista (sobre el nodo que contiene el número 3).
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())

	// Borrar elemento.
	iter.Borrar()
	// Verificar que se haya borrado.
	require.False(t, iter.HaySiguiente())

	// Verificar condiciones finales de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}
func TestRemoverElementoDelMedio(t *testing.T) {
	t.Log("Remover elemento cuando esta en el medio y verificar que no esta")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	// Verificar condiciones iniciales de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Crear iterador e iterar hasta estar en el medio de la lista (en el nodo que contiene el número 2).
	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())

	// Borrar elemento.
	iter.Borrar()

	// Recorrer el resto de la lista.
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())

	// Verificar condiciones finales de la lista.
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	// Iterar lista para verificar que se haya borrado el elemento.
	iter = lista.Iterador()

	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestIteradorInternoListaCompleta(t *testing.T) {
	const LIMITE int = 10
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 1; i <= LIMITE; i++ {
		lista.InsertarUltimo(i)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
		require.Equal(t, 1, lista.VerPrimero())
		require.Equal(t, i, lista.VerUltimo())
	}

	// Itero lista con una función que suma el dato almacenado en el elemento de la lista.
	suma := 0
	lista.Iterar(func(dato int) bool {
		suma += dato
		return true
	})

	// Comparo la suma con el resultado esperado.
	sumaEsperada := LIMITE * (LIMITE + 1) / 2

	require.Equal(t, sumaEsperada, suma)
}

func TestIteradorInternoBuscarElementos(t *testing.T) {
	const LIMITE int = 10
	lista := TDALista.CrearListaEnlazada[int]()
	elementosBuscados := []int{7, 5, 18, 0, 1, -3, 5, 6}
	resultadosBusquedaEsperados := []bool{true, true, false, false, true, false, true, true}
	resultadosBusqueda := []bool{}

	for i := 1; i <= LIMITE; i++ {
		lista.InsertarUltimo(i)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
		require.Equal(t, 1, lista.VerPrimero())
		require.Equal(t, i, lista.VerUltimo())
	}

	// Se le pide buscar ciertos elementos específicos almacenados en un slice.
	for _, elem := range elementosBuscados {
		encontrado := false

		lista.Iterar(func(datoAlmacenado int) bool {
			if datoAlmacenado == elem {
				encontrado = true
				return false
			}
			return true
		})

		resultadosBusqueda = append(resultadosBusqueda, encontrado)
	}

	// Se comparan los resultados de las búsquedas con los esperados.
	require.Equal(t, resultadosBusquedaEsperados, resultadosBusqueda)
}
