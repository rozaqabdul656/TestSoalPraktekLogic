package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var Dadu int
	var BanyakPemain int
	Score := []int{}
	KetersediaanDadu := []int{}
	fmt.Println("Masukan Banyaknya  Dadu: ")

	fmt.Scanln(&Dadu)

	fmt.Println("Masukan Banyaknya Pemain: ")

	fmt.Scanln(&BanyakPemain)

	for i := 0; i < BanyakPemain; i++ {
		KetersediaanDadu = append(KetersediaanDadu, Dadu) // asign ketersediaan dadu setiap pemain
		Score = append(Score, 0)
	}
	giliran := 0
	for {
		giliran++
		hasilRandomTotalDadu := [][]int{}

		for i := 0; i < BanyakPemain; i++ {
			RandomDadu := []int{}

			for d := 0; d < KetersediaanDadu[i]; d++ {
				RandomDadu = append(RandomDadu, rand.Intn(6)+1)

			}
			hasilRandomTotalDadu = append(hasilRandomTotalDadu, RandomDadu)

		}
		fmt.Printf(" Giliran %d lempar Dadu  : \n", giliran)

		for i := 0; i < BanyakPemain; i++ {
			fmt.Printf("Pemain #%d  (%d): ", i+1, Score[i])
			for x := 0; x < KetersediaanDadu[i]; x++ {
				fmt.Print(hasilRandomTotalDadu[i][x])
				if x != KetersediaanDadu[i]-1 {
					fmt.Print(",")
				}

			}
			if KetersediaanDadu[i] == 0 {
				fmt.Print("_ (Berhenti bermain karena tidak memiliki dadu)")
			}
			fmt.Println("")

		}
		fmt.Println("Setelah Evaluasi ")
		exceptIndex := [][]int{}
		inisialiSasi := []int{}
		for i := 0; i < BanyakPemain; i++ {
			exceptIndex = append(exceptIndex, inisialiSasi) // asign ketersediaan dadu setiap pemain

		}
		for i := 0; i < BanyakPemain; i++ {
			for x := 0; x < KetersediaanDadu[i]; x++ {
				if hasilRandomTotalDadu[i][x] == 6 {
					KetersediaanDadu[i] -= 1
					hasilRandomTotalDadu[i] = RemoveIndex(hasilRandomTotalDadu[i], x)
					Score[i] += 1
				} else if hasilRandomTotalDadu[i][x] == 1 {
					isexcept := false
					for y := 0; y < len(exceptIndex[i]); y++ {
						if exceptIndex[i][y] == x {
							isexcept = true
						}
					}
					if isexcept == false {

						hasilRandomTotalDadu[i] = RemoveIndex(hasilRandomTotalDadu[i], x)
						KetersediaanDadu[i] -= 1
						// fmt.Println(exceptIndex)
						if i < BanyakPemain-1 {
							KetersediaanDadu[i+1] += 1
							exceptIndex[i+1] = append(exceptIndex[i+1], len(hasilRandomTotalDadu[i+1])) // assign except checking

							hasilRandomTotalDadu[i+1] = append(hasilRandomTotalDadu[i+1], 1)

						}
						if i == BanyakPemain-1 {
							KetersediaanDadu[0] += 1
							exceptIndex[0] = append(exceptIndex[0], len(hasilRandomTotalDadu[0])) // assign except checking

							hasilRandomTotalDadu[0] = append(hasilRandomTotalDadu[0], 1)

						}
					}

				}
			}

		}
		checkingisOver := 0
		indexCheckisOver := 0
		for i := 0; i < BanyakPemain; i++ {

			fmt.Printf("Pemain #%d  (%d): ", i+1, Score[i])
			for x := 0; x < KetersediaanDadu[i]; x++ {
				fmt.Print(hasilRandomTotalDadu[i][x])
				if x != KetersediaanDadu[i]-1 {
					fmt.Print(",")
				}
			}
			if KetersediaanDadu[i] == 0 {
				fmt.Print("_ (Berhenti bermain karena tidak memiliki dadu)")
			} else {
				checkingisOver++
				indexCheckisOver = i
			}
			fmt.Println("")

		}
		if checkingisOver == 1 {
			fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu. \n", indexCheckisOver+1)
			var max int
			max = Score[0]
			for _, value := range Score {

				if value > max {
					max = value
				}
			}

			for c := 0; c < len(Score); c++ { // checking bila pemeneang lebih dari satu
				if max == Score[c] {
					fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya. \n", c+1)
				}
			}

			break
		}

	}

}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// func randomNumber(rand int) int {
// 	return min + rand.Intn(rand)
// }
