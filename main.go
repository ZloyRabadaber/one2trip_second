package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type pos struct {
	x int
	y int
}

type matrix struct {
	m        [][]int
	N        int
	k        pos
	location pos
}

func printM(m [][]int) {
	for i := range m {
		for j := range m[i] {
			if j == 0 {
				fmt.Println()
			}
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%-3X", m[i][j])
		}
	}
	fmt.Println()
}

func main() {
	var err error

	if len(os.Args) < 2 {
		fmt.Println("Отсутствует обязательный аргумент.")
		fmt.Println("Используйте: ", os.Args[0], " N")
		fmt.Println("\t, где N - размер матрицы.")
		return
	}
	rand.Seed(time.Now().UnixNano())

	var M matrix

	M.N, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Аргумент должен быть числом.")
		return
	}
	M.m = make([][]int, M.N)
	M.k = pos{-1, 1}
	M.location = pos{M.N / 2, M.N / 2}

	//коррекция начальной позиции для чётной размерности
	if M.N%2 == 0 {
		M.location.y--
	}

	//заполнение исходной матрицы случайными числами
	for i := range M.m {
		M.m[i] = make([]int, M.N)
		for j := range M.m[i] {
			M.m[i][j] = rand.Intn(255)
		}
	}
	printM(M.m)

	//отступ для читабельности
	fmt.Println()

	for step := 1; step <= M.N; step++ {
		M.moveX(step)
		if !M.moveY(step) {
			break
		}
	}
	fmt.Println()
}

func (M *matrix) isOutRange() bool {
	return M.location.x < 0 || M.location.x >= M.N || M.location.y < 0 || M.location.y >= M.N

}

func (M *matrix) moveX(step int) bool {
	for j := 0; j < step; j++ {
		if M.isOutRange() {
			return false
		}
		fmt.Printf("%-3X ", M.m[M.location.y][M.location.x])
		M.location.x += M.k.x
	}
	M.k.x *= -1
	return true
}

func (M *matrix) moveY(step int) bool {
	for j := 0; j < step; j++ {
		if M.isOutRange() {
			return false
		}
		fmt.Printf("%-3X ", M.m[M.location.y][M.location.x])
		M.location.y += M.k.y
	}
	M.k.y *= -1
	return true
}
