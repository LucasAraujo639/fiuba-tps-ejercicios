package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

// Funciones auxiliares.

func EncolarEnterosConsec(t *testing.T, c TDACola.Cola[int], limite int) TDACola.Cola[int] {
	for i := 1; i <= limite; i++ {
		c.Encolar(i)
		require.False(t, c.EstaVacia())
		require.Equal(t, 1, c.VerPrimero())
	}

	return c
}

func DesencolarEnterosConsec(t *testing.T, c TDACola.Cola[int], limite int) ([]int, []int, []int) {
	resultadoEsperado := make([]int, limite)
	resultadoVerPrimero := make([]int, limite)
	resultadoDesencolar := make([]int, limite)

	for i := 0; i < limite; i++ {
		require.False(t, c.EstaVacia())
		resultadoEsperado[i] = i + 1
		resultadoVerPrimero[i] = c.VerPrimero()
		resultadoDesencolar[i] = c.Desencolar()
	}

	return resultadoEsperado, resultadoVerPrimero, resultadoDesencolar
}

// Pruebas.
func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestUnicoValor(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(50)

	require.Equal(t, 50, cola.VerPrimero())
	require.False(t, cola.EstaVacia())
	require.Equal(t, 50, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolado(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	cola.Encolar(1)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.VerPrimero())
	require.Equal(t, 1, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	cola.Encolar(2)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 2, cola.VerPrimero())

	cola.Encolar(3)
	cola.Encolar(4)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 2, cola.VerPrimero())
	require.Equal(t, 2, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, 3, cola.VerPrimero())

	cola.Encolar(5)
	cola.Encolar(6)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 3, cola.VerPrimero())
	require.Equal(t, 3, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, 4, cola.VerPrimero())
	require.Equal(t, 4, cola.Desencolar())
	require.Equal(t, 5, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, 6, cola.VerPrimero())
	require.Equal(t, 6, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncoladoVolumen(t *testing.T) {
	const LIMITE = 10000
	cola := TDACola.CrearColaEnlazada[int]()

	cola = EncolarEnterosConsec(t, cola, LIMITE)

	require.Equal(t, 1, cola.VerPrimero())

	resultadoEsperado, resultadoVerPrimero, resultadoDesencolar := DesencolarEnterosConsec(t, cola, LIMITE)

	require.Equal(t, resultadoEsperado, resultadoVerPrimero)
	require.Equal(t, resultadoEsperado, resultadoDesencolar)

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestMultiplesDesencolados(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	// Primera prueba:

	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.VerPrimero())

	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 2, cola.Desencolar())
	require.Equal(t, 3, cola.Desencolar())
	require.Equal(t, 4, cola.Desencolar())

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	// Segunda prueba:

	cola.Encolar(5)
	cola.Encolar(6)
	cola.Encolar(7)
	cola.Encolar(8)
	cola.Encolar(9)
	cola.Encolar(10)
	cola.Encolar(11)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 5, cola.VerPrimero())

	require.Equal(t, 5, cola.Desencolar())
	require.Equal(t, 6, cola.Desencolar())
	require.Equal(t, 7, cola.Desencolar())
	require.Equal(t, 8, cola.Desencolar())
	require.Equal(t, 9, cola.Desencolar())
	require.Equal(t, 10, cola.Desencolar())
	require.Equal(t, 11, cola.Desencolar())

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	// Tercera prueba:

	cola.Encolar(12)
	cola.Encolar(13)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 12, cola.VerPrimero())

	require.Equal(t, 12, cola.Desencolar())
	require.Equal(t, 13, cola.Desencolar())

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestMultiplesDesencoladosVolumen(t *testing.T) {
	const LIMITE1 int = 5000
	const LIMITE2 int = 10000
	const LIMITE3 int = 20000
	resultadoEsperado, resultadoVerPrimero, resultadoDesencolar := []int{}, []int{}, []int{}
	cola := TDACola.CrearColaEnlazada[int]()

	// Primera prueba:

	cola = EncolarEnterosConsec(t, cola, LIMITE1)

	require.Equal(t, 1, cola.VerPrimero())

	resultadoEsperado, resultadoVerPrimero, resultadoDesencolar = DesencolarEnterosConsec(t, cola, LIMITE1)

	require.Equal(t, resultadoEsperado, resultadoVerPrimero)
	require.Equal(t, resultadoEsperado, resultadoDesencolar)

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	// Segunda prueba:

	cola = EncolarEnterosConsec(t, cola, LIMITE2)

	require.Equal(t, 1, cola.VerPrimero())

	resultadoEsperado, resultadoVerPrimero, resultadoDesencolar = DesencolarEnterosConsec(t, cola, LIMITE2)

	require.Equal(t, resultadoEsperado, resultadoVerPrimero)
	require.Equal(t, resultadoEsperado, resultadoDesencolar)

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	// Tercera prueba:

	cola = EncolarEnterosConsec(t, cola, LIMITE3)

	require.Equal(t, 1, cola.VerPrimero())

	resultadoEsperado, resultadoVerPrimero, resultadoDesencolar = DesencolarEnterosConsec(t, cola, LIMITE3)

	require.Equal(t, resultadoEsperado, resultadoVerPrimero)
	require.Equal(t, resultadoEsperado, resultadoDesencolar)

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolarStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	elementos := []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "sed", "do",
		"eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua"}
	resultadoDesencolar, resultadoVerPrimero := make([]string, len(elementos)), make([]string, len(elementos))

	// Encolar elementos en la cola.
	for i := 0; i < len(elementos); i++ {
		cola.Encolar(elementos[i])
	}

	// Verificar que el primer elemento sea el esperado.
	require.Equal(t, "Lorem", cola.VerPrimero())

	// Desencolar elementos.
	for i := 0; i < len(elementos); i++ {
		resultadoVerPrimero[i] = cola.VerPrimero()
		resultadoDesencolar[i] = cola.Desencolar()
	}

	// Comparar resultados obtenidos.
	require.Equal(t, elementos, resultadoVerPrimero)
	require.Equal(t, elementos, resultadoDesencolar)

	// Verificar que la cola está vacía.
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}
