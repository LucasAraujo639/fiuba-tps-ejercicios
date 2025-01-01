package diccionario_test

import (
	"math/rand"
	"strings"
	"tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func funcionComparacionEnteros(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
func funcionComparacionStrings(a, b string) int {
	return strings.Compare(a, b)
}
func TestAbbVacio(t *testing.T) {
	t.Log("Comprueba que diccionario vacio no tiene claves")
	abb := diccionario.CrearABB[string, int](funcionComparacionStrings)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece("pepe"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("pepe") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("pepe") })
}
func TestInsercionAbb(t *testing.T) {
	t.Log("Inserta varios elementos en el ABB y verifica que los elementos se inserten correctamente ")
	abb := diccionario.CrearABB[string, int](funcionComparacionStrings)
	abb.Guardar("F", 1)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("F"))
	require.False(t, abb.Pertenece("A"))
	require.EqualValues(t, 1, abb.Obtener("F"))

	abb.Guardar("B", 2)
	require.EqualValues(t, 2, abb.Cantidad())
	require.True(t, abb.Pertenece("B"))
	require.False(t, abb.Pertenece("A"))
	require.EqualValues(t, 2, abb.Obtener("B"))

	abb.Guardar("A", 3)
	require.EqualValues(t, 3, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	require.False(t, abb.Pertenece("Z"))
	require.EqualValues(t, 3, abb.Obtener("A"))

	abb.Guardar("Z", 4)
	require.EqualValues(t, 4, abb.Cantidad())
	require.True(t, abb.Pertenece("Z"))
	require.False(t, abb.Pertenece("J"))
	require.EqualValues(t, 4, abb.Obtener("Z"))

	abb.Guardar("K", 5)
	require.EqualValues(t, 5, abb.Cantidad())
	require.True(t, abb.Pertenece("Z"))
	require.EqualValues(t, 5, abb.Obtener("K"))

	abb.Guardar("N", 6)
	require.EqualValues(t, 6, abb.Cantidad())
	require.True(t, abb.Pertenece("Z"))
	require.EqualValues(t, 6, abb.Obtener("N"))

	abb.Guardar("M", 7)
	require.EqualValues(t, 7, abb.Cantidad())
	require.True(t, abb.Pertenece("Z"))
	require.EqualValues(t, 7, abb.Obtener("M"))

	abb.Guardar("H", 8)
	require.EqualValues(t, 8, abb.Cantidad())
	require.True(t, abb.Pertenece("Z"))
	require.EqualValues(t, 8, abb.Obtener("H"))

	require.EqualValues(t, 1, abb.Obtener("F"))
	require.EqualValues(t, 2, abb.Obtener("B"))
	require.EqualValues(t, 3, abb.Obtener("A"))
	require.EqualValues(t, 4, abb.Obtener("Z"))
	require.True(t, abb.Pertenece("F"))
	require.True(t, abb.Pertenece("B"))
	require.True(t, abb.Pertenece("A"))
	require.True(t, abb.Pertenece("Z"))
	require.False(t, abb.Pertenece("L"))
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que funcione correctamente con un solo nodo")
	abb := diccionario.CrearABB[int, string](funcionComparacionEnteros)
	abb.Guardar(100, "pepe")
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(200))
	require.True(t, abb.Pertenece(100))
	require.EqualValues(t, "pepe", abb.Obtener(100))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(30) })
	abb.Borrar(100)
	require.EqualValues(t, 0, abb.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(100) })
}

func TestReemplazoDato(t *testing.T) {
	t.Log("Se inserta un par de nodos con la misma clave y se comprueba que se remplaze correctamente")
	abb := diccionario.CrearABB[int, string](funcionComparacionEnteros)
	abb.Guardar(100, "A")
	abb.Guardar(80, "B")
	abb.Guardar(90, "C")
	abb.Guardar(60, "D")
	abb.Guardar(120, "E")
	require.EqualValues(t, 5, abb.Cantidad())
	require.EqualValues(t, "A", abb.Obtener(100))
	require.EqualValues(t, "B", abb.Obtener(80))
	require.EqualValues(t, "C", abb.Obtener(90))
	require.EqualValues(t, "D", abb.Obtener(60))
	require.EqualValues(t, "E", abb.Obtener(120))
	abb.Guardar(100, "AA")
	abb.Guardar(80, "BB")
	abb.Guardar(90, "CC")
	abb.Guardar(60, "DD")
	abb.Guardar(120, "EE")
	require.EqualValues(t, 5, abb.Cantidad())
	require.True(t, abb.Pertenece(100))
	require.True(t, abb.Pertenece(80))
	require.True(t, abb.Pertenece(90))
	require.True(t, abb.Pertenece(60))
	require.True(t, abb.Pertenece(120))
	require.EqualValues(t, "AA", abb.Obtener(100))
	require.EqualValues(t, "BB", abb.Obtener(80))
	require.EqualValues(t, "CC", abb.Obtener(90))
	require.EqualValues(t, "DD", abb.Obtener(60))
	require.EqualValues(t, "EE", abb.Obtener(120))

}

func TestBorrarRaizSinElementos(t *testing.T) {
	dic := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	dic.Guardar(100, 100)

	require.Equal(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(100))
	require.Equal(t, 100, dic.Obtener(100))

	require.Equal(t, 100, dic.Borrar(100))

	require.Equal(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(100))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(100) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(100) })
}

func TestBorrarHojasAgregarElemento(t *testing.T) {
	t.Log("Se crea un nodo con dos hijos, se borran los hijos y se vuelve a agregar un elemento.")
	dic := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	dic.Guardar(100, 100)
	dic.Guardar(400, 400)
	dic.Guardar(50, 50)
	require.Equal(t, 3, dic.Cantidad())
	require.True(t, dic.Pertenece(100))
	require.True(t, dic.Pertenece(50))
	require.True(t, dic.Pertenece(400))

	require.Equal(t, 400, dic.Borrar(400))

	require.Equal(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(100))
	require.True(t, dic.Pertenece(50))
	require.False(t, dic.Pertenece(400))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(400) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(400) })

	require.Equal(t, 50, dic.Borrar(50))

	require.Equal(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(100))
	require.False(t, dic.Pertenece(50))
	require.False(t, dic.Pertenece(400))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(50) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(50) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(400) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(400) })

	dic.Guardar(70, 70)

	require.Equal(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(100))
	require.True(t, dic.Pertenece(70))
}

func TestBorrarRaizConUnHijoYVolverAgregar(t *testing.T) {
	t.Log("Se borra la raíz de un árbol con un hijo y luego se agrega ese mismo par clave-dato para luego borrar nuevamente la raíz." +
		"Debería quedar como al inicio.")
	dic := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	dic.Guardar(200, 20)
	dic.Guardar(100, 10)

	require.Equal(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(200))
	require.True(t, dic.Pertenece(100))
	require.Equal(t, 20, dic.Obtener(200))
	require.Equal(t, 10, dic.Obtener(100))

	require.Equal(t, 20, dic.Borrar(200))

	require.Equal(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(200))
	require.True(t, dic.Pertenece(100))
	require.Equal(t, 10, dic.Obtener(100))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(200) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(200) })

	dic.Guardar(200, 20)

	require.Equal(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(200))
	require.True(t, dic.Pertenece(100))
	require.Equal(t, 20, dic.Obtener(200))
	require.Equal(t, 10, dic.Obtener(100))

	require.Equal(t, 10, dic.Borrar(100))

	require.Equal(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(100))
	require.True(t, dic.Pertenece(200))
	require.Equal(t, 20, dic.Obtener(200))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(100) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(100) })

	dic.Guardar(100, 10)

	require.Equal(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(200))
	require.True(t, dic.Pertenece(100))
	require.Equal(t, 20, dic.Obtener(200))
	require.Equal(t, 10, dic.Obtener(100))
}

func TestBorrarRaizConDosHijosYVolverAAgregar(t *testing.T) {
	t.Log("Se borra la raíz de un árbol de 3 nodos donde la raíz tiene dos hijos. Luego se vuelven a agregar elementos.")

	dic := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	dic.Guardar(200, 20)
	dic.Guardar(300, 30)
	dic.Guardar(100, 10)

	require.Equal(t, 3, dic.Cantidad())
	require.True(t, dic.Pertenece(300))
	require.True(t, dic.Pertenece(200))
	require.True(t, dic.Pertenece(100))
	require.Equal(t, 30, dic.Obtener(300))
	require.Equal(t, 20, dic.Obtener(200))
	require.Equal(t, 10, dic.Obtener(100))

	require.Equal(t, 20, dic.Borrar(200))

	require.Equal(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(300))
	require.False(t, dic.Pertenece(200))
	require.True(t, dic.Pertenece(100))
	require.Equal(t, 30, dic.Obtener(300))
	require.Equal(t, 10, dic.Obtener(100))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(200) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(200) })

	dic.Guardar(400, 40)

	require.Equal(t, 3, dic.Cantidad())
	require.True(t, dic.Pertenece(300))
	require.True(t, dic.Pertenece(400))
	require.True(t, dic.Pertenece(100))
	require.Equal(t, 30, dic.Obtener(300))
	require.Equal(t, 40, dic.Obtener(400))
	require.Equal(t, 10, dic.Obtener(100))

	dic.Guardar(200, 20)

	require.Equal(t, 4, dic.Cantidad())
	require.True(t, dic.Pertenece(300))
	require.True(t, dic.Pertenece(400))
	require.True(t, dic.Pertenece(100))
	require.True(t, dic.Pertenece(200))
	require.Equal(t, 30, dic.Obtener(300))
	require.Equal(t, 40, dic.Obtener(400))
	require.Equal(t, 10, dic.Obtener(100))
	require.Equal(t, 20, dic.Obtener(200))
}

func TestBorrarNodo2Hijos(t *testing.T) {
	t.Log("Se borra un nodo con dos hijos.")
	dic := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	dic.Guardar(100, 10)
	dic.Guardar(200, 20)
	dic.Guardar(150, 15)
	dic.Guardar(250, 25)

	require.Equal(t, 4, dic.Cantidad())
	require.True(t, dic.Pertenece(100))
	require.True(t, dic.Pertenece(150))
	require.True(t, dic.Pertenece(200))
	require.True(t, dic.Pertenece(250))
	require.Equal(t, 10, dic.Obtener(100))
	require.Equal(t, 15, dic.Obtener(150))
	require.Equal(t, 20, dic.Obtener(200))
	require.Equal(t, 25, dic.Obtener(250))

	require.Equal(t, 20, dic.Borrar(200))

	require.Equal(t, 3, dic.Cantidad())
	require.True(t, dic.Pertenece(100))
	require.True(t, dic.Pertenece(150))
	require.False(t, dic.Pertenece(200))
	require.True(t, dic.Pertenece(250))
	require.Equal(t, 10, dic.Obtener(100))
	require.Equal(t, 15, dic.Obtener(150))
	require.Equal(t, 25, dic.Obtener(250))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(200) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(200) })
}

func TestBorrarNodo2HijosConReemplazanteConHijo(t *testing.T) {
	t.Log("Se borra un nodo con dos hijos con un reemplazante con un hijo y se verifica que todavía exista el hijo.")
	dic := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	dic.Guardar(100, 10)
	dic.Guardar(300, 30)
	dic.Guardar(200, 20)
	dic.Guardar(400, 40)
	dic.Guardar(150, 15)
	dic.Guardar(180, 18)

	require.Equal(t, 6, dic.Cantidad())
	require.True(t, dic.Pertenece(100))
	require.True(t, dic.Pertenece(300))
	require.True(t, dic.Pertenece(200))
	require.True(t, dic.Pertenece(400))
	require.True(t, dic.Pertenece(150))
	require.True(t, dic.Pertenece(180))
	require.Equal(t, 10, dic.Obtener(100))
	require.Equal(t, 30, dic.Obtener(300))
	require.Equal(t, 20, dic.Obtener(200))
	require.Equal(t, 40, dic.Obtener(400))
	require.Equal(t, 15, dic.Obtener(150))
	require.Equal(t, 18, dic.Obtener(180))

	require.Equal(t, 30, dic.Borrar(300))

	require.Equal(t, 5, dic.Cantidad())
	require.True(t, dic.Pertenece(100))
	require.False(t, dic.Pertenece(300))
	require.True(t, dic.Pertenece(200))
	require.True(t, dic.Pertenece(400))
	require.True(t, dic.Pertenece(150))
	require.True(t, dic.Pertenece(180))
	require.Equal(t, 10, dic.Obtener(100))
	require.Equal(t, 20, dic.Obtener(200))
	require.Equal(t, 40, dic.Obtener(400))
	require.Equal(t, 15, dic.Obtener(150))
	require.Equal(t, 18, dic.Obtener(180))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(300) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(300) })
}

func TestAbbBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el abb, y se los borra, revisando que en todo momento " +
		"el abb se comporte de manera adecuada")
	abb := diccionario.CrearABB[int, string](funcionComparacionEnteros)
	abb.Guardar(100, "A")
	abb.Guardar(80, "B")
	abb.Guardar(90, "C")
	abb.Guardar(60, "D")
	abb.Guardar(120, "E")
	require.EqualValues(t, 5, abb.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(101) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(119) })

	abb.Borrar(60)
	require.EqualValues(t, 4, abb.Cantidad())

	abb.Borrar(120)
	require.EqualValues(t, 3, abb.Cantidad())

	abb.Borrar(100)
	require.EqualValues(t, 2, abb.Cantidad())

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(60) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(120) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(100) })
	require.False(t, abb.Pertenece(60))
	require.False(t, abb.Pertenece(120))
	require.False(t, abb.Pertenece(100))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(60) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(120) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(100) })
	abb.Guardar(120, "C")
	abb.Guardar(60, "D")
	abb.Guardar(100, "A")
	require.True(t, abb.Pertenece(60))
	require.True(t, abb.Pertenece(120))
	require.True(t, abb.Pertenece(100))
	require.EqualValues(t, 5, abb.Cantidad())
}
func TestClaveVacia(t *testing.T) {
	t.Log("Probamos guardar una clave vacia (deberia funcionar)")
	abb := diccionario.CrearABB[string, string](funcionComparacionStrings)
	clave := ""
	abb.Guardar(clave, clave)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, clave, abb.Obtener(clave))
}
func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := diccionario.CrearABB[string, *int](funcionComparacionStrings)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

// -------------------------------------------------------------------------------------------
// ----------------------------Iterador interno-----------------------------------------------
// -------------------------------------------------------------------------------------------
func TestIteradorInterno(t *testing.T) {
	t.Log("Probamos con una suma que el iterador interno funcione")
	abb := diccionario.CrearABB[string, int](funcionComparacionStrings)
	abb.Guardar("F", 100)
	abb.Guardar("A", 80)
	abb.Guardar("B", 90)
	abb.Guardar("L", 60)
	abb.Guardar("E", 120)
	require.EqualValues(t, 5, abb.Cantidad())

	suma := 0
	ptrSuma := &suma
	abb.Iterar(func(clave string, dato int) bool {
		*ptrSuma += dato
		return true
	})
	require.EqualValues(t, 450, suma)

}

func TestRecorrerTodoIteradorInterno(t *testing.T) {
	t.Log("Verificamos que podamos recorrer el arbol de principio a fin")
	abb := diccionario.CrearABB[string, int](funcionComparacionStrings)
	abb.Guardar("A", 100)
	abb.Guardar("B", 80)
	abb.Guardar("C", 90)
	abb.Guardar("D", 60)
	abb.Guardar("H", 150)
	require.EqualValues(t, 5, abb.Cantidad())

	suma := 0
	ptrSuma := &suma
	abb.Iterar(func(_ string, _ int) bool {
		*ptrSuma++
		return true
	})
	require.EqualValues(t, abb.Cantidad(), suma)

}

func TestSumaConCorteIteradorInterno(t *testing.T) {
	t.Log("Verificamos que los nodos se recorran en orden")
	abb := diccionario.CrearABB[string, int](funcionComparacionStrings)
	abb.Guardar("F", 100)
	abb.Guardar("E", 80)
	abb.Guardar("A", 90)
	abb.Guardar("C", 60)
	abb.Guardar("Z", 120)
	abb.Guardar("L", 1230)
	require.EqualValues(t, 6, abb.Cantidad())

	suma := 0
	ptrSuma := &suma
	abb.Iterar(func(_ string, dato int) bool {
		if *ptrSuma >= 0 && *ptrSuma < 200 {
			*ptrSuma += dato
			return true
		}
		return false
	})
	require.EqualValues(t, 230, suma)

}

func TestRecorrerIteradorInternoRango(t *testing.T) {
	t.Log("Verificamos el iterador interno con rango sin condicion de corte")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(100, 100)
	abb.Guardar(50, 50)
	abb.Guardar(25, 25)
	abb.Guardar(10, 10)
	abb.Guardar(150, 150)
	abb.Guardar(200, 200)
	abb.Guardar(75, 75)
	abb.Guardar(60, 60)
	abb.Guardar(80, 80)
	abb.Guardar(30, 30)
	abb.Guardar(110, 110)
	abb.Guardar(130, 130)
	abb.Guardar(190, 190)
	abb.Guardar(300, 300)
	require.EqualValues(t, 14, abb.Cantidad())
	desde := 60
	hasta := 120
	suma := 0
	ptrSuma := &suma
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		*ptrSuma += dato
		return true
	})
	require.EqualValues(t, 425, suma)

}

func TestIteradorConUnNodo(t *testing.T) {
	t.Log("Verificamos funcione el iterador interno con un solo nodo")
	abb := diccionario.CrearABB[string, int](funcionComparacionStrings)
	abb.Guardar("A", 100)
	require.EqualValues(t, 1, abb.Cantidad())

	suma := 0
	ptrSuma := &suma
	abb.Iterar(func(_ string, _ int) bool {
		*ptrSuma++
		return true
	})
	require.EqualValues(t, abb.Cantidad(), suma)

}

func TestIteradorVacio(t *testing.T) {
	t.Log("Verificamos que funcione el iterador interno sin ningun nodo")
	abb := diccionario.CrearABB[string, int](funcionComparacionStrings)
	require.EqualValues(t, 0, abb.Cantidad())

	suma := 0
	ptrSuma := &suma
	abb.Iterar(func(_ string, _ int) bool {
		*ptrSuma++
		return true
	})
	require.EqualValues(t, 0, suma)

}

func TestIteradorRangoConUnNodo(t *testing.T) {
	t.Log("Verificamos si funciona un iterador interno de rango con un solo nodo estando en el rango")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(300, 300)

	require.EqualValues(t, 1, abb.Cantidad())
	desde := 200
	hasta := 400
	suma := 0
	ptrSuma := &suma
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		*ptrSuma += dato
		return true
	})
	require.EqualValues(t, 300, suma)

}

func TestIteradorFueraDeRango(t *testing.T) {
	t.Log("Verificamos si funciona un iterador interno de rango en el que todos los nodo esteen fuera de rango")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(100, 100)
	abb.Guardar(80, 80)
	abb.Guardar(90, 90)
	abb.Guardar(60, 60)
	abb.Guardar(150, 150)
	abb.Guardar(20, 20)
	abb.Guardar(120, 120)
	abb.Guardar(115, 115)
	require.EqualValues(t, 8, abb.Cantidad())
	desde := 200
	hasta := 1200
	suma := 0
	ptrSuma := &suma
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		*ptrSuma += dato
		return true
	})
	require.EqualValues(t, 0, suma)

}
func TestIteradorRangoConCorte(t *testing.T) {
	t.Log("Verificamos si funciona correctamente el iterador interno de rango con condicion de corte")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(100, 100)
	abb.Guardar(80, 80)
	abb.Guardar(90, 90)
	abb.Guardar(60, 60)
	abb.Guardar(150, 150)
	abb.Guardar(20, 20)
	abb.Guardar(120, 120)
	abb.Guardar(115, 115)
	require.EqualValues(t, 8, abb.Cantidad())
	desde := 80
	hasta := 120
	contador := 0
	suma := 0
	ptrContador := &contador
	ptrSuma := &suma
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		if *ptrContador < 5 {
			*ptrContador++
			*ptrSuma += dato
			return true
		}
		return false
	})
	require.EqualValues(t, 5, contador)
	require.EqualValues(t, 505, suma)

}

// -------------------------------------------------------------------------------------------
// ----------------------------Iterador Externo-----------------------------------------------
// -------------------------------------------------------------------------------------------

func TestIteradorExterno(t *testing.T) {
	t.Log("Probamos con una suma que el iterador externo funcione")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(100, 100)
	abb.Guardar(80, 80)
	abb.Guardar(90, 90)
	abb.Guardar(60, 60)
	abb.Guardar(120, 120)
	require.EqualValues(t, 5, abb.Cantidad())

	iter := abb.Iterador()
	suma := 0
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		suma += valor
		iter.Siguiente()
	}
	require.EqualValues(t, 450, suma)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestIteradorExternoOrdenAscendente(t *testing.T) {
	t.Log("Probamos con una suma que el iterador externo funcione")
	ordenIteraciones := []int{60, 80, 90, 100, 120}
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(100, 100)
	abb.Guardar(80, 80)
	abb.Guardar(90, 90)
	abb.Guardar(60, 60)
	abb.Guardar(120, 120)
	require.EqualValues(t, 5, abb.Cantidad())

	iter := abb.Iterador()
	i := 0
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		require.EqualValues(t, ordenIteraciones[i], valor)
		i++
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestIteradorExternoUnNodo(t *testing.T) {
	t.Log("Probamos que el iterador externo funcione con un solo nodo")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(100, 100)
	require.EqualValues(t, 1, abb.Cantidad())

	iter := abb.Iterador()
	suma := 0
	i := 0
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		suma += valor
		i++
		iter.Siguiente()
	}
	require.EqualValues(t, abb.Cantidad(), i)
	require.EqualValues(t, 100, suma)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestIteradorExternoVacio(t *testing.T) {
	t.Log("Probamos que el iterador externo funcione con un solo nodo")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	require.EqualValues(t, 0, abb.Cantidad())
	iter := abb.Iterador()
	require.EqualValues(t, false, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestIteradorExternoConCorte(t *testing.T) {
	t.Log("Probamos con una suma que el iterador externo con corte funcione")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(100, 100)
	abb.Guardar(80, 80)
	abb.Guardar(90, 90)
	abb.Guardar(60, 60)
	abb.Guardar(120, 120)
	require.EqualValues(t, 5, abb.Cantidad())

	iter := abb.Iterador()
	suma := 0
	for iter.HaySiguiente() {
		if suma < 130 {
			_, valor := iter.VerActual()
			suma += valor
			iter.Siguiente()
			return
		}
		break

	}
	require.EqualValues(t, 140, suma)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestIteradorExternoRango(t *testing.T) {
	t.Log("Probamos que itere por rangos correctamente")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(100, 100)
	abb.Guardar(80, 80)
	abb.Guardar(90, 90)
	abb.Guardar(60, 60)
	abb.Guardar(120, 120)
	abb.Guardar(150, 150)
	abb.Guardar(70, 70)
	require.EqualValues(t, 7, abb.Cantidad())
	desde := 80
	hasta := 120
	suma := 0
	iter := abb.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		suma += valor
		iter.Siguiente()
	}
	require.EqualValues(t, 390, suma)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExternoRangoVacio(t *testing.T) {
	t.Log("Probamos el iterador externo por rangos vacio")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	require.EqualValues(t, 0, abb.Cantidad())
	desde := 80
	hasta := 120
	iter := abb.IteradorRango(&desde, &hasta)
	require.EqualValues(t, false, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestIteradorExternoRangoUnNodo(t *testing.T) {
	t.Log("Probamos que itere por rangos correctamente")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(100, 100)
	require.EqualValues(t, 1, abb.Cantidad())
	desde := 80
	hasta := 120
	suma := 0
	iter := abb.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		suma += valor
		iter.Siguiente()
	}
	require.EqualValues(t, 100, suma)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestIteradorExternoRangoConCorte(t *testing.T) {
	t.Log("Probamos con una suma que el iterador externo con corte funcione")
	abb := diccionario.CrearABB[int, int](funcionComparacionEnteros)
	abb.Guardar(100, 100)
	abb.Guardar(80, 80)
	abb.Guardar(90, 90)
	abb.Guardar(60, 60)
	abb.Guardar(120, 120)
	require.EqualValues(t, 5, abb.Cantidad())

	desde := 80
	hasta := 120
	suma := 0
	iter := abb.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		if suma < 130 {
			_, valor := iter.VerActual()
			suma += valor
			iter.Siguiente()
			return
		}
		break

	}
	require.EqualValues(t, 170, suma)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

//-----------------------------------------------
//---------------Prueba de volumen---------------
//-----------------------------------------------

func TestPruebaVolumen(t *testing.T) {
	const CANTIDAD = 1000
	dic := diccionario.CrearABB[int, int](funcionComparacionEnteros)

	claves := diccionario.CrearHash[int, int]()
	clavesEnArbol := diccionario.CrearHash[int, int]()
	valores := diccionario.CrearHash[int, int]()

	// agregamos elementos a los arreglos de claves y valores
	for i := 0; i < CANTIDAD; i++ {
		claves.Guardar(i, i)
		valores.Guardar(i, i)
	}

	require.Equal(t, CANTIDAD, claves.Cantidad())

	// desordenamos los arreglos
	for i := 0; i < CANTIDAD; i++ {
		j := rand.Intn(CANTIDAD)

		dato1, dato2 := claves.Obtener(i), claves.Obtener(j)

		claves.Guardar(j, dato1)
		claves.Guardar(i, dato2)

		dato1, dato2 = valores.Obtener(i), valores.Obtener(j)

		valores.Guardar(j, dato1)
		valores.Guardar(i, dato2)
	}

	require.Equal(t, CANTIDAD, claves.Cantidad())
	require.Equal(t, CANTIDAD, valores.Cantidad())

	for i := 0; i < CANTIDAD; i++ {
		dic.Guardar(claves.Obtener(i), valores.Obtener(i))
		clavesEnArbol.Guardar(i, claves.Obtener(i))
	}

	require.Equal(t, CANTIDAD, claves.Cantidad())
	require.Equal(t, CANTIDAD, valores.Cantidad())

	require.EqualValues(t, CANTIDAD, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	ok := true
	for i := 0; i < CANTIDAD; i++ {
		ok = dic.Pertenece(claves.Obtener(i))
		if !ok {
			break
		}
		ok = dic.Obtener(claves.Obtener(i)) == valores.Obtener(i)
		if !ok {
			break
		}
	}

	require.True(t, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(t, CANTIDAD, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	// Borra cada uno de los elementos del árbol.
	for i := 0; i < CANTIDAD; i++ {
		claveDic := claves.Obtener(i)
		require.True(t, dic.Pertenece(claveDic))
		require.Equal(t, valores.Obtener(i), dic.Obtener(claveDic))
		require.Equal(t, valores.Obtener(i), dic.Borrar(claveDic))
		require.Equal(t, CANTIDAD-i-1, dic.Cantidad())
		require.False(t, dic.Pertenece(claveDic))
	}

	// Se verifica estado final.
	require.Equal(t, 0, dic.Cantidad())
}
