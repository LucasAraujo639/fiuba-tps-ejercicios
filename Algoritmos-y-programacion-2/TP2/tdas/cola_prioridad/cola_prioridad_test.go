package cola_prioridad_test

import (
	"math/rand"
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"

	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_VOLUMEN = 10000
)

func comparacionEnterosMax(a, b int) int    { return a - b }
func comparacionEnterosMin(a, b int) int    { return b - a }
func comparacionStringsMax(a, b string) int { return strings.Compare(a, b) }
func comparacionStringsMin(a, b string) int { return strings.Compare(b, a) }
func TestColaPrioridadVacia(t *testing.T) {
	t.Log("Se crea un heap vacio de maximos de enteros a partir de un arreglo")
	cola := TDAColaPrioridad.CrearHeap[int](comparacionEnterosMax)
	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaPrioridadDeArregloVacia(t *testing.T) {
	t.Log("Se crea un heap vacio de maximos de enteros a partir de un arreglo")
	arreglo := []int{}
	cola := TDAColaPrioridad.CrearHeapArr[int](arreglo, comparacionEnterosMax)

	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestHeapMaximosEnteros(t *testing.T) {
	t.Log("Se crea un heap de maximos de enteros a partir de un arreglo y se verifica que estee correcto")
	arr := []int{50, 10, 60, 100, 20, 90, 30, 80, 40, 70}
	cola := TDAColaPrioridad.CrearHeapArr[int](arr, comparacionEnterosMax)
	require.Equal(t, 10, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())
	require.EqualValues(t, 100, cola.Desencolar())
	require.Equal(t, 9, cola.Cantidad())
	require.EqualValues(t, 90, cola.VerMax())
	require.EqualValues(t, 90, cola.Desencolar())
	require.Equal(t, 8, cola.Cantidad())
	require.EqualValues(t, 80, cola.VerMax())
	require.EqualValues(t, 80, cola.Desencolar())
	require.Equal(t, 7, cola.Cantidad())
	require.EqualValues(t, 70, cola.VerMax())
	require.EqualValues(t, 70, cola.Desencolar())
	require.Equal(t, 6, cola.Cantidad())
	require.EqualValues(t, 60, cola.VerMax())
	require.EqualValues(t, 60, cola.Desencolar())
	require.Equal(t, 5, cola.Cantidad())
	require.EqualValues(t, 50, cola.VerMax())
	require.EqualValues(t, 50, cola.Desencolar())
	require.Equal(t, 4, cola.Cantidad())
	require.EqualValues(t, 40, cola.VerMax())
	require.EqualValues(t, 40, cola.Desencolar())
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 30, cola.VerMax())
	require.EqualValues(t, 30, cola.Desencolar())
	require.Equal(t, 2, cola.Cantidad())
	require.EqualValues(t, 20, cola.VerMax())
	require.EqualValues(t, 20, cola.Desencolar())
	require.Equal(t, 1, cola.Cantidad())
	require.EqualValues(t, 10, cola.VerMax())
	require.EqualValues(t, 10, cola.Desencolar())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}
func TestHeapMinimosEnteros(t *testing.T) {
	t.Log("Se crea un heap de minimos de enteros a partir de un arreglo y se verifica que se desencolen los elementos correctamente")
	arr := []int{50, 10, 60, 100, 20, 90, 30, 80, 40, 70}
	cola := TDAColaPrioridad.CrearHeapArr[int](arr, comparacionEnterosMin)
	require.Equal(t, 10, cola.Cantidad())
	require.EqualValues(t, 10, cola.VerMax())
	require.EqualValues(t, 10, cola.Desencolar())
	require.Equal(t, 9, cola.Cantidad())
	require.EqualValues(t, 20, cola.VerMax())
	require.EqualValues(t, 20, cola.Desencolar())
	require.Equal(t, 8, cola.Cantidad())
	require.EqualValues(t, 30, cola.VerMax())
	require.EqualValues(t, 30, cola.Desencolar())
	require.Equal(t, 7, cola.Cantidad())
	require.EqualValues(t, 40, cola.VerMax())
	require.EqualValues(t, 40, cola.Desencolar())
	require.Equal(t, 6, cola.Cantidad())
	require.EqualValues(t, 50, cola.VerMax())
	require.EqualValues(t, 50, cola.Desencolar())
	require.Equal(t, 5, cola.Cantidad())
	require.EqualValues(t, 60, cola.VerMax())
	require.EqualValues(t, 60, cola.Desencolar())
	require.Equal(t, 4, cola.Cantidad())
	require.EqualValues(t, 70, cola.VerMax())
	require.EqualValues(t, 70, cola.Desencolar())
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 80, cola.VerMax())
	require.EqualValues(t, 80, cola.Desencolar())
	require.Equal(t, 2, cola.Cantidad())
	require.EqualValues(t, 90, cola.VerMax())
	require.EqualValues(t, 90, cola.Desencolar())
	require.Equal(t, 1, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())
	require.EqualValues(t, 100, cola.Desencolar())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestInserccionHeapMaximosEnteros(t *testing.T) {
	t.Log("Se encolan elementos a un heap de maximos vacio y se verifica que desencolen los elementos correctamente")
	cola := TDAColaPrioridad.CrearHeap[int](comparacionEnterosMax)
	require.Equal(t, 0, cola.Cantidad())
	cola.Encolar(50)
	require.Equal(t, 1, cola.Cantidad())
	require.EqualValues(t, 50, cola.VerMax())
	cola.Encolar(60)
	require.Equal(t, 2, cola.Cantidad())
	require.EqualValues(t, 60, cola.VerMax())
	cola.Encolar(100)
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())
	cola.Encolar(20)
	require.Equal(t, 4, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())
	cola.Encolar(90)
	require.Equal(t, 5, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())
	cola.Encolar(30)
	require.Equal(t, 6, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())
	cola.Encolar(80)
	require.Equal(t, 7, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())
	cola.Encolar(40)
	require.Equal(t, 8, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())
	cola.Encolar(70)
	require.Equal(t, 9, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())
	cola.Encolar(10)
	require.Equal(t, 10, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())

	require.EqualValues(t, 100, cola.Desencolar())
	require.EqualValues(t, 90, cola.Desencolar())
	require.EqualValues(t, 80, cola.Desencolar())
	require.EqualValues(t, 70, cola.Desencolar())
	require.EqualValues(t, 60, cola.Desencolar())
	require.EqualValues(t, 50, cola.Desencolar())
	require.EqualValues(t, 40, cola.Desencolar())
	require.EqualValues(t, 30, cola.Desencolar())
	require.EqualValues(t, 20, cola.Desencolar())
	require.EqualValues(t, 10, cola.Desencolar())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}
func TestInserccionHeapMinimosEnteros(t *testing.T) {
	t.Log("Se encolan elementos a un heap de minimos vacio y se verifica que desencolen los elementos correctamente")
	cola := TDAColaPrioridad.CrearHeap[int](comparacionEnterosMin)
	require.Equal(t, 0, cola.Cantidad())
	cola.Encolar(50)
	require.Equal(t, 1, cola.Cantidad())
	require.EqualValues(t, 50, cola.VerMax())
	cola.Encolar(60)
	require.Equal(t, 2, cola.Cantidad())
	require.EqualValues(t, 50, cola.VerMax())
	cola.Encolar(100)
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 50, cola.VerMax())
	cola.Encolar(20)
	require.Equal(t, 4, cola.Cantidad())
	require.EqualValues(t, 20, cola.VerMax())
	cola.Encolar(90)
	require.Equal(t, 5, cola.Cantidad())
	require.EqualValues(t, 20, cola.VerMax())
	cola.Encolar(30)
	require.Equal(t, 6, cola.Cantidad())
	require.EqualValues(t, 20, cola.VerMax())
	cola.Encolar(80)
	require.Equal(t, 7, cola.Cantidad())
	require.EqualValues(t, 20, cola.VerMax())
	cola.Encolar(40)
	require.Equal(t, 8, cola.Cantidad())
	require.EqualValues(t, 20, cola.VerMax())
	cola.Encolar(70)
	require.Equal(t, 9, cola.Cantidad())
	require.EqualValues(t, 20, cola.VerMax())
	cola.Encolar(10)
	require.Equal(t, 10, cola.Cantidad())
	require.EqualValues(t, 10, cola.VerMax())

}
func TestHeapMaximosStrings(t *testing.T) {
	t.Log("Se crea un heap de maximos de strings a partir de un arreglo y se verifica que estee correcto")
	arr := []string{"F", "K", "M", "A", "B", "Z"}
	cola := TDAColaPrioridad.CrearHeapArr[string](arr, comparacionStringsMax)
	require.Equal(t, 6, cola.Cantidad())
	require.EqualValues(t, "Z", cola.VerMax())
	require.EqualValues(t, "Z", cola.Desencolar())
	require.Equal(t, 5, cola.Cantidad())
	require.EqualValues(t, "M", cola.VerMax())
	require.EqualValues(t, "M", cola.Desencolar())
	require.Equal(t, 4, cola.Cantidad())
	require.EqualValues(t, "K", cola.VerMax())
	require.EqualValues(t, "K", cola.Desencolar())
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, "F", cola.VerMax())
	require.EqualValues(t, "F", cola.Desencolar())
	require.Equal(t, 2, cola.Cantidad())
	require.EqualValues(t, "B", cola.VerMax())
	require.EqualValues(t, "B", cola.Desencolar())
	require.Equal(t, 1, cola.Cantidad())
	require.EqualValues(t, "A", cola.VerMax())
	require.EqualValues(t, "A", cola.Desencolar())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}
func TestInserccionHeapMinimosStrings(t *testing.T) {
	t.Log("Se encolan strings a un heap de minimos vacio y se verifica que desencolen los elementos correctamente")
	cola := TDAColaPrioridad.CrearHeap[string](comparacionStringsMin)
	require.Equal(t, 0, cola.Cantidad())
	cola.Encolar("F")
	require.Equal(t, 1, cola.Cantidad())
	require.EqualValues(t, "F", cola.VerMax())
	cola.Encolar("M")
	require.Equal(t, 2, cola.Cantidad())
	require.EqualValues(t, "F", cola.VerMax())
	cola.Encolar("B")
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, "B", cola.VerMax())
	cola.Encolar("Z")
	require.Equal(t, 4, cola.Cantidad())
	require.EqualValues(t, "B", cola.VerMax())
	cola.Encolar("K")
	require.Equal(t, 5, cola.Cantidad())
	require.EqualValues(t, "B", cola.VerMax())
	cola.Encolar("A")
	require.Equal(t, 6, cola.Cantidad())
	require.EqualValues(t, "A", cola.VerMax())

}
func TestInserccionYExtraccionMultiple(t *testing.T) {
	t.Log("Se realiza una secuencia de insercciones y extracciones, comprobando que funcione correctamente en todo momento")
	cola := TDAColaPrioridad.CrearHeap[int](comparacionEnterosMax)
	require.Equal(t, 0, cola.Cantidad())

	//Encolo tres elementos
	cola.Encolar(50)
	require.Equal(t, 1, cola.Cantidad())
	require.EqualValues(t, 50, cola.VerMax())
	cola.Encolar(60)
	require.Equal(t, 2, cola.Cantidad())
	require.EqualValues(t, 60, cola.VerMax())
	cola.Encolar(100)
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 100, cola.VerMax())

	//Empiezo con la inserccion y extraccion multiple
	require.EqualValues(t, 100, cola.Desencolar())
	require.Equal(t, 2, cola.Cantidad())

	cola.Encolar(20)
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 60, cola.VerMax())
	require.EqualValues(t, 60, cola.Desencolar())
	require.Equal(t, 2, cola.Cantidad())

	cola.Encolar(90)
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 90, cola.VerMax())
	require.EqualValues(t, 90, cola.Desencolar())
	require.Equal(t, 2, cola.Cantidad())
	cola.Encolar(30)
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 50, cola.VerMax())
	require.EqualValues(t, 50, cola.Desencolar())
	require.Equal(t, 2, cola.Cantidad())
	cola.Encolar(80)
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 80, cola.VerMax())
	require.EqualValues(t, 80, cola.Desencolar())
	require.Equal(t, 2, cola.Cantidad())
	cola.Encolar(10)
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 30, cola.VerMax())
	require.EqualValues(t, 30, cola.Desencolar())
	require.Equal(t, 2, cola.Cantidad())
	cola.Encolar(5)
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 20, cola.VerMax())
	require.EqualValues(t, 20, cola.Desencolar())
	require.Equal(t, 2, cola.Cantidad())

	require.EqualValues(t, 10, cola.Desencolar())
	require.EqualValues(t, 5, cola.Desencolar())

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}
func TestElementosRepetidos(t *testing.T) {
	t.Log("Se inserta elementos repetidos en el heap y verifica que se mantengan en el orden correcto despues de las extracciones")
	cola := TDAColaPrioridad.CrearHeap[int](comparacionEnterosMax)
	require.Equal(t, 0, cola.Cantidad())
	cola.Encolar(50)
	cola.Encolar(50)

	require.EqualValues(t, 50, cola.VerMax())
	require.EqualValues(t, 50, cola.Desencolar())
	cola.Encolar(60)
	cola.Encolar(60)

	require.EqualValues(t, 60, cola.VerMax())
	require.EqualValues(t, 60, cola.Desencolar())
	cola.Encolar(100)
	cola.Encolar(100)

	require.EqualValues(t, 100, cola.VerMax())
	require.EqualValues(t, 100, cola.Desencolar())
	cola.Encolar(20)
	cola.Encolar(20)

	require.EqualValues(t, 100, cola.VerMax())
	require.EqualValues(t, 100, cola.Desencolar())
	cola.Encolar(90)
	cola.Encolar(90)

	require.EqualValues(t, 90, cola.VerMax())
	require.EqualValues(t, 90, cola.Desencolar())
	cola.Encolar(30)
	cola.Encolar(30)

	require.EqualValues(t, 90, cola.VerMax())
	require.EqualValues(t, 90, cola.Desencolar())

	require.EqualValues(t, 60, cola.Desencolar())
	require.EqualValues(t, 50, cola.Desencolar())
	require.EqualValues(t, 30, cola.Desencolar())
	require.EqualValues(t, 30, cola.Desencolar())
	require.EqualValues(t, 20, cola.Desencolar())
	require.EqualValues(t, 20, cola.Desencolar())
	require.Equal(t, 0, cola.Cantidad())

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}
func TestHeapMaximosFlotantes(t *testing.T) {
	t.Log("Se crea un heap de maximos flotantes mediante insercciones y se verifica que se desencolen los elementos correctamente")
	cola := TDAColaPrioridad.CrearHeap[float32](func(a, b float32) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	require.Equal(t, 0, cola.Cantidad())
	cola.Encolar(0.5)
	require.Equal(t, 1, cola.Cantidad())
	require.EqualValues(t, 0.5, cola.VerMax())
	cola.Encolar(0.6)
	require.Equal(t, 2, cola.Cantidad())
	require.EqualValues(t, 0.6, cola.VerMax())
	cola.Encolar(1)
	require.Equal(t, 3, cola.Cantidad())
	require.EqualValues(t, 1, cola.VerMax())
	cola.Encolar(0.2)
	require.Equal(t, 4, cola.Cantidad())
	require.EqualValues(t, 1, cola.VerMax())
	cola.Encolar(0.9)
	require.Equal(t, 5, cola.Cantidad())
	require.EqualValues(t, 1, cola.VerMax())
	cola.Encolar(0.3)
	require.Equal(t, 6, cola.Cantidad())
	require.EqualValues(t, 1, cola.VerMax())
	cola.Encolar(0.8)
	require.Equal(t, 7, cola.Cantidad())
	require.EqualValues(t, 1, cola.VerMax())
	cola.Encolar(0.4)
	require.Equal(t, 8, cola.Cantidad())
	require.EqualValues(t, 1, cola.VerMax())
	cola.Encolar(0.7)
	require.Equal(t, 9, cola.Cantidad())
	require.EqualValues(t, 1, cola.VerMax())
	cola.Encolar(0.1)
	require.Equal(t, 10, cola.Cantidad())
	require.EqualValues(t, 1, cola.VerMax())

	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 0.9, cola.Desencolar())
	require.EqualValues(t, 0.8, cola.Desencolar())
	require.EqualValues(t, 0.7, cola.Desencolar())
	require.EqualValues(t, 0.6, cola.Desencolar())
	require.EqualValues(t, 0.5, cola.Desencolar())
	require.EqualValues(t, 0.4, cola.Desencolar())
	require.EqualValues(t, 0.3, cola.Desencolar())
	require.EqualValues(t, 0.2, cola.Desencolar())
	require.EqualValues(t, 0.1, cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestPruebasVolumenHeap(t *testing.T) {
	t.Log("Se crea un heap y se realizan pruebas de volumen insertando muchos elementos y desencolandolos")
	cola := TDAColaPrioridad.CrearHeap[int](comparacionEnterosMax)
	require.Equal(t, 0, cola.Cantidad())
	for i := 0; i <= _VOLUMEN; i++ {
		cola.Encolar(i)
		require.Equal(t, i+1, cola.Cantidad())
		require.EqualValues(t, i, cola.VerMax())
		require.False(t, cola.EstaVacia())
	}
	for i := _VOLUMEN; i > 0; i-- {
		require.EqualValues(t, i, cola.Desencolar())
		require.Equal(t, i, cola.Cantidad())
		require.EqualValues(t, i-1, cola.VerMax())
		require.False(t, cola.EstaVacia())

	}

}
func TestPruebasVolumenHeapArr(t *testing.T) {
	t.Log("Se crea un heap y se realizan pruebas de volumen insertando muchos elementos y desencolandolos")
	arr := []int{}
	cola := TDAColaPrioridad.CrearHeapArr[int](arr, comparacionEnterosMax)
	require.Equal(t, 0, cola.Cantidad())
	for i := 0; i <= _VOLUMEN; i++ {
		cola.Encolar(i)
		require.Equal(t, i+1, cola.Cantidad())
		require.EqualValues(t, i, cola.VerMax())
		require.False(t, cola.EstaVacia())
	}
	for i := _VOLUMEN; i > 0; i-- {
		require.EqualValues(t, i, cola.Desencolar())
		require.Equal(t, i, cola.Cantidad())
		require.EqualValues(t, i-1, cola.VerMax())
		require.False(t, cola.EstaVacia())

	}

}
func TestHeapSort(t *testing.T) {
	t.Log("Se verifica que HeapSort ordene el arreglo de enteros correctamente")
	arr_ordenado := []int{10, 20, 20, 30, 40, 50, 60, 70, 80, 80, 90, 100, 100}
	arr := []int{50, 10, 60, 100, 20, 90, 30, 80, 40, 70, 20, 80, 100}
	TDAColaPrioridad.HeapSort(arr, comparacionEnterosMax)
	require.EqualValues(t, arr, arr_ordenado)
}
func TestHeapSortStrings(t *testing.T) {
	t.Log("Se verifica que HeapSort ordene el arreglo de strings correctamente")
	arr := []string{"F", "K", "A", "C", "B", "E", "M", "Z", "R", "T", "U", "V", "D"}
	TDAColaPrioridad.HeapSort(arr, comparacionStringsMax)
	require.EqualValues(t, 13, len(arr))
	require.EqualValues(t, "A", arr[0])
	require.EqualValues(t, "B", arr[1])
	require.EqualValues(t, "C", arr[2])
	require.EqualValues(t, "D", arr[3])
	require.EqualValues(t, "E", arr[4])
	require.EqualValues(t, "F", arr[5])
	require.EqualValues(t, "K", arr[6])
	require.EqualValues(t, "M", arr[7])
	require.EqualValues(t, "R", arr[8])
	require.EqualValues(t, "T", arr[9])
	require.EqualValues(t, "U", arr[10])
	require.EqualValues(t, "V", arr[11])
	require.EqualValues(t, "Z", arr[12])

}

func TestHeapSortVolumen(t *testing.T) {
	t.Log("Prueba de volumen HEAPSORT")
	arr := []int{}
	for i := 0; i < _VOLUMEN; i++ {
		arr = append(arr, rand.Intn(_VOLUMEN))
	}
	TDAColaPrioridad.HeapSort(arr, comparacionEnterosMax)
	for i := 0; i < len(arr)-1; i++ {
		require.True(t, arr[i] <= arr[i+1])
	}
}
