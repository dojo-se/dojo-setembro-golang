package roleta_romana

func RoletaRomana(n int, k int, i int) (indice_sobrevivente int, err string) {
	if n < 1 {
		err = "n menor que 1"
		return
	}
	if k < 1 {
		err = "k menor que 1"
		return
	}
	if i < 1 {
		err = "i menor que 1"
		return
	}

	pessoas := make([]int, n)
	for index := 0; index < n; index++ {
		pessoas[index] = (index + 1)
	}

	indice_sobrevivente = roletaRomana(pessoas, k, i-1)

	return
}

func roletaRomana(sobreviventes []int, k int, i int) int {
	if len(sobreviventes) == 1 {
		return sobreviventes[0]
	} else {
		quem_morre := (i + k) % len(sobreviventes)
		sobreviventes = append(sobreviventes[0:quem_morre], sobreviventes[(quem_morre+1):]...)
		return roletaRomana(sobreviventes, k, quem_morre)
	}
}
