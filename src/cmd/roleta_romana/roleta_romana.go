package roleta_romana

func RoletaRomana(n int, k int, i int) (indice_sobrevivente int, err string) {
	if n < 1 {
		err = "n menor que 1"
	}
	if k < 1 {
		err = "k menor que 1"
	}
	if i < 1 {
		err = "i menor que 1"
	}
	indice_sobrevivente = 1
	return
}
