package roleta_romana

import "testing"

func TestPrimeiroTeste(t *testing.T) {
	indice_sobrevivente, _ := RoletaRomana(1, 1, 1)

	if indice_sobrevivente != 1 {
		t.Errorf("indice_sobrevivente deveria ser 1, mas foi %d", indice_sobrevivente)
	}
}

func TestValidacaoDados(t *testing.T) {
	_, err := RoletaRomana(0, 1, 1)
	if err != "n menor que 1" {
		t.Errorf("err deve exibir messagem de erro: %s", err)
	}

	_, err = RoletaRomana(1, 0, 1)
	if err != "k menor que 1" {
		t.Errorf("err deve exibir messagem de erro: %s", err)
	}
	_, err = RoletaRomana(1, 1, 0)
	if err != "i menor que 1" {
		t.Errorf("err deve exibir messagem de erro: %s", err)
	}
}
