package main

import (
	"fmt"
)

func main() {
	grille := [][]int{[]int{1, 0, 1}, []int{0, 2, 0}, []int{0, 0, 0}}
	fmt.Println(mini_max(grille, 2, 8, true))
	//jeu()
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
	//1 pour joueur 2 pour IA
	type_j1 := 1
	type_j2 := 2
	for run {
		affichage(grille)
		var input string
		if type_j1 == 1 && joueur == 1 || type_j2 == 1 && joueur == 2 {
			fmt.Scanln(&input)
			//input = "1,1"
			if input == "q" {
				run = false
				fmt.Println("arrêt")
			}
		}
		c := coup{}
		if joueur == 2 && type_j2 == 2 {
			_, c = mini_max(grille, 2, 8, true)
			fmt.Println(c)
		} else {
			c = coup_joueur(input)
		}

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

type maximum struct {
	val_coup int
	c        coup
}

func mini_max(grille [][]int, joueur int, profondeur int, maxi bool) (int, coup) {
	gagnant, j := verif(grille)
	if gagnant && j == joueur {
		return 100 + profondeur, coup{}
	}
	if gagnant && j == 3 || profondeur == 0 {
		return 0, coup{}
	}
	if gagnant && j != joueur {
		return 100 - profondeur, coup{}
	}
	if maxi {
		max := maximum{}
		max.val_coup = -200
		for _, coup := range coup_possible(grille) {
			grille_c := copy_moi(grille)
			grille_c[coup.x][coup.y] = joueur
			score, c := mini_max(grille_c, joueur, profondeur-1, !maxi)

			if score > max.val_coup {
				max.val_coup = score
				max.c = c
			}
		}
		return max.val_coup, max.c
	} else {
		min := maximum{}
		min.val_coup = 200
		for _, coup := range coup_possible(grille) {
			grille_c := copy_moi(grille)
			grille_c[coup.x][coup.y] = joueur*-1 + 3
			score, c := mini_max(grille_c, joueur, profondeur-1, !maxi)

			if score < min.val_coup {
				min.val_coup = score
				min.c = c
			}
		}
		return min.val_coup, min.c
	}
}

func copy_moi(grille [][]int) [][]int {
	grille_c := construction_grille()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grille_c[i][j] = grille[i][j]
		}
	}
	return grille_c
}

func coup_possible(grille [][]int) []coup {
	res := []coup{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grille[i][j] == 0 {
				res = append(res, coup{i, j})
			}
		}
	}
	return res
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
