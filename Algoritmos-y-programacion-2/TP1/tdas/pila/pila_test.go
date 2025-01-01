package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

// Funciones auxiliares:

// Apila los números naturales del 1 a "cantidad". Devuelve la pila con elementos apilados.
func ApilarEnterosConsec(pila TDAPila.Pila[int], cantidad int) TDAPila.Pila[int] {
	for i := 1; i <= cantidad; i++ {
		pila.Apilar(i)
	}

	return pila
}

func DesapilarEnterosConsec(pila TDAPila.Pila[int], cantidad int) ([]int, []int, []int) {
	resultadoEsperado := make([]int, cantidad)
	resultadoVerTope := make([]int, cantidad)
	resultadoDesapilar := make([]int, cantidad)

	for i := cantidad; i > 0; i-- {
		resultadoEsperado[i-1] = i
		resultadoVerTope[i-1] = pila.VerTope()
		resultadoDesapilar[i-1] = pila.Desapilar()
	}

	return resultadoEsperado, resultadoVerTope, resultadoDesapilar
}

// Pruebas:

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilarElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)

	require.False(t, pila.EstaVacia())
	require.Equal(t, 3, pila.VerTope())

	pila.Apilar(4)
	pila.Apilar(5)

	require.False(t, pila.EstaVacia())
	require.Equal(t, 5, pila.VerTope())
	require.Equal(t, 5, pila.Desapilar())
	require.Equal(t, 4, pila.Desapilar())
	require.Equal(t, 3, pila.Desapilar())
	require.False(t, pila.EstaVacia())
	require.Equal(t, 2, pila.VerTope())

	pila.Apilar(6)
	pila.Apilar(7)
	pila.Apilar(8)
	pila.Apilar(9)
	pila.Apilar(10)
	pila.Apilar(11)

	require.False(t, pila.EstaVacia())
	require.Equal(t, 11, pila.VerTope())
	require.Equal(t, 11, pila.Desapilar())
	require.Equal(t, 10, pila.Desapilar())
	require.Equal(t, 9, pila.Desapilar())

	pila.Apilar(12)
	pila.Apilar(13)
	pila.Apilar(14)

	require.False(t, pila.EstaVacia())
	require.Equal(t, 14, pila.VerTope())
	require.Equal(t, 14, pila.Desapilar())
	require.Equal(t, 13, pila.Desapilar())
	require.Equal(t, 12, pila.Desapilar())
	require.Equal(t, 8, pila.Desapilar())
	require.Equal(t, 7, pila.Desapilar())
	require.Equal(t, 6, pila.Desapilar())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilarElementosVolumen(t *testing.T) {
	const LIMITE int = 5000
	resultadoEsperado, resultadoVerTope, resultadoDesapilar := []int{}, []int{}, []int{}
	pila := TDAPila.CrearPilaDinamica[int]()

	pila = ApilarEnterosConsec(pila, LIMITE)

	// Revisar el tope de la pila.
	require.Equal(t, LIMITE, pila.VerTope())

	// Desapilar pila y verificar tope:
	resultadoEsperado, resultadoVerTope, resultadoDesapilar = DesapilarEnterosConsec(pila, LIMITE)

	// Ver si los resultados concuerdan con lo esperado:
	require.Equal(t, resultadoEsperado, resultadoDesapilar)
	require.Equal(t, resultadoEsperado, resultadoVerTope)

	// Ver si se comporta como pila vacía.
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestMultiplesVaciados(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	// Primer apilado y desapilado:

	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)

	require.Equal(t, 4, pila.VerTope())

	require.Equal(t, 4, pila.Desapilar())
	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.Desapilar())

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	// Segundo apilado y desapilado:

	pila.Apilar(5)
	pila.Apilar(6)
	pila.Apilar(7)
	pila.Apilar(8)
	pila.Apilar(9)
	pila.Apilar(10)
	pila.Apilar(11)

	require.Equal(t, 11, pila.VerTope())

	require.Equal(t, 11, pila.Desapilar())
	require.Equal(t, 10, pila.Desapilar())
	require.Equal(t, 9, pila.Desapilar())
	require.Equal(t, 8, pila.Desapilar())
	require.Equal(t, 7, pila.Desapilar())
	require.Equal(t, 6, pila.Desapilar())
	require.Equal(t, 5, pila.Desapilar())

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	// Tercer apilado y desapilado:

	pila.Apilar(12)
	pila.Apilar(13)

	require.Equal(t, 13, pila.VerTope())

	require.Equal(t, 13, pila.Desapilar())
	require.Equal(t, 12, pila.Desapilar())

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestMultiplesVaciadosVolumen(t *testing.T) {
	const LIMITE1 int = 5000
	const LIMITE2 int = 10000
	const LIMITE3 int = 20000
	resultadoEsperado, resultadoVerTope, resultadoDesapilar := []int{}, []int{}, []int{}
	pila := TDAPila.CrearPilaDinamica[int]()

	// Primera prueba:

	pila = ApilarEnterosConsec(pila, LIMITE1)

	require.Equal(t, LIMITE1, pila.VerTope())

	resultadoEsperado, resultadoVerTope, resultadoDesapilar = DesapilarEnterosConsec(pila, LIMITE1)

	require.Equal(t, resultadoEsperado, resultadoVerTope)
	require.Equal(t, resultadoEsperado, resultadoDesapilar)

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	// Segunda prueba:

	pila = ApilarEnterosConsec(pila, LIMITE2)

	require.Equal(t, LIMITE2, pila.VerTope())

	resultadoEsperado, resultadoVerTope, resultadoDesapilar = DesapilarEnterosConsec(pila, LIMITE2)

	require.Equal(t, resultadoEsperado, resultadoVerTope)
	require.Equal(t, resultadoEsperado, resultadoDesapilar)

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	// Tercera prueba:

	pila = ApilarEnterosConsec(pila, LIMITE3)

	require.Equal(t, LIMITE3, pila.VerTope())

	resultadoEsperado, resultadoVerTope, resultadoDesapilar = DesapilarEnterosConsec(pila, LIMITE3)

	require.Equal(t, resultadoEsperado, resultadoVerTope)
	require.Equal(t, resultadoEsperado, resultadoDesapilar)

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilarStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	elementos := []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "sed", "do",
		"eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua"}
	resultadoDesapilar, resultadoVerTope := make([]string, len(elementos)), make([]string, len(elementos))

	// Apilar elementos en la pila.
	for i := 0; i < len(elementos); i++ {
		pila.Apilar(elementos[i])
	}

	// Verificar que el último elemento sea el esperado.
	require.Equal(t, "aliqua", pila.VerTope())

	// Desapilar pila.
	for i := len(elementos); i > 0; i-- {
		resultadoVerTope[i-1] = pila.VerTope()
		resultadoDesapilar[i-1] = pila.Desapilar()
	}

	// Comparar resultados obtenidos.
	require.Equal(t, elementos, resultadoVerTope)
	require.Equal(t, elementos, resultadoDesapilar)

	// Verificar que la pila está vacía.
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}
