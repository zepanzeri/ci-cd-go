package soma

import "testing"

func TesteSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado invalido. Resultado: %d. Esperado: %d", total, 30)
	}
}
