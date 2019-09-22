package matematica

// arquivo de teste: tem sufixo "[...]_test.go"
// método de teste: começa com "Test[...]"
import "testing"

const erroPadrao = "Valor espera %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
