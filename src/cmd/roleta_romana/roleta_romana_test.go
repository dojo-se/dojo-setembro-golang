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

func TestMataOPrimeiro(t *testing.T) {
	indice_sobrevivente, _ := RoletaRomana(2, 1, 2)

	if indice_sobrevivente != 2 {
		t.Errorf("indice_sobrevivente deveria ser 2, mas foi %d", indice_sobrevivente)
	}
}

func TestComCincoPessoas(t *testing.T) {
	indice_sobrevivente, _ := RoletaRomana(5, 2, 1)

	if indice_sobrevivente != 4 {
		t.Errorf("indice_sobrevivente deveria ser 4, mas foi %d", indice_sobrevivente)
	}
}

func TestDerradeiro(t *testing.T) {
	indice_sobrevivente, _ := RoletaRomana(10, 1, 1)

	if indice_sobrevivente != 5 {
		t.Errorf("indice_sobrevivente deveria ser 5, mas foi %d", indice_sobrevivente)
	}
}

func TestGenocidio(t *testing.T) {
	indice_sobrevivente, _ := RoletaRomana(10000, 2, 1)

	if indice_sobrevivente != 2692 {
		t.Errorf("indice_sobrevivente deveria ser 2692, mas foi %d", indice_sobrevivente)
	}
}
