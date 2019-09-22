package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto         string
		parte         string
		indexEsperado int
	}{
		{"Udemy é show", "Udemy", 1},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Leonardo", "o", 2},
	}

	for _, teste := range testes {
		//t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.indexEsperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.indexEsperado, atual)
		}
	}
}
