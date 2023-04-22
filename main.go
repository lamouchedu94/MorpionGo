package main

import (
	"fmt"
)

func main() {
	jeu()
}

type coup struct {
	x int
	y int
}

func jeu() {
	//var tab = []int
	grille := construction_grille()
	run := true
	joueur := 1
	for run {
		affichage(grille)
		var input string
		fmt.Scanln(&input)
		if input == "q" {
			run = false
			fmt.Println("arrêt")
		}
		c := coup_joueur(input)

		if case_libre(grille, c) {
			if joueur == 1 {
				grille[c.x][c.y] = 1
				joueur = 2
			} else {
				grille[c.x][c.y] = 2
				joueur = 1
			}

		}

		gagnant, j := verif(grille)
		if gagnant {
			mess := ""
			if j != 3 {
				mess = fmt.Sprintf("Le joueur %d à gagné", j)
			} else {
				mess = "partie nulle"
			}
			fmt.Println(mess)
			run = false
		}

	}

}

func verif(grille [][]int) (bool, int) {
	for i := 0; i < 3; i++ {
		if grille[i][0] == grille[i][1] && grille[i][1] == grille[i][2] && grille[i][2] != 0 {
			return true, grille[i][0]
		}
		if grille[0][i] == grille[1][i] && grille[1][i] == grille[2][i] && grille[2][i] != 0 {
			return true, grille[0][i]
		}
	}
	//diag 1
	if grille[0][0] == grille[1][1] && grille[1][1] == grille[2][2] && grille[2][2] != 0 {
		return true, grille[0][0]
	}
	//diag 2
	if grille[0][2] == grille[1][1] && grille[1][1] == grille[2][0] && grille[2][0] != 0 {
		return true, grille[0][2]
	}
	//Si tout plein
	rempli := true
out:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grille[i][j] == 0 {
				rempli = false
				break out
			}
		}
	}
	if rempli {
		return true, 3
	}

	return false, 0
}

func coup_joueur(input string) coup {
	c := coup{}
	fmt.Sscanf(input, "%d,%d", &c.x, &c.y)
	return c
}

func case_libre(grille [][]int, c coup) bool {
	return grille[c.x][c.y] == 0
}

func affichage(grille [][]int) {
	for _, tab := range grille {
		fmt.Println(tab)
	}
}

func construction_grille() [][]int {
	var grille [][]int

	for i := 0; i < 3; i++ {
		tab := make([]int, 3)
		grille = append(grille, tab)
	}
	return grille
}
