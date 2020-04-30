package main

import (
	"flag"
	"fmt"
	"github.com/alexandrslesarenko/sudoku-processor/sudokuboard"
	"os"
	"time"
)

func getBoard(fileName *string, dim *int) (*sudokuboard.Board, error) {
	if *fileName != "" {
		fmt.Println("FileName=" + *fileName)
		board, err := sudokuboard.LoadFromJsonFile(*fileName)
		return board, err
	} else {
		board, err := sudokuboard.New(*dim)
		return &board, err
	}
}

func main() {
	fileName := flag.String("file", "", "Enter JSON filename to load predefined board")
	dim := flag.Int("dim", 3, "Enter dimension of board. Default is 3")
	showProgress := flag.Bool("progress", false, "Show progress")
	flag.Parse()
	fmt.Println()
	fmt.Println(os.Args[0], " Copyright (C) 2020 Alexandr Slessarenko")
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY; for details licence at https://www.gnu.org/licenses/gpl-3.0.html")
	fmt.Println("This is free software, and you are welcome to redistribute it under conditions of licence.")

	board, err := getBoard(fileName, dim)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Board for solution")
	board.Print()
	start := time.Now()
	solvedBoard, solved, err := board.GetRecurseSolution(*showProgress)
	duration1 := time.Since(start)
	start1 := time.Now()
	solvedBoard, solved, err = board.GetParallelSolution(*showProgress)
	duration2 := time.Since(start1)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	} else {
		if solved {
			fmt.Println("Board solution:")
			solvedBoard.Print()
		} else {
			fmt.Println("Board has no solution")
		}
	}
	fmt.Println("Duration of recursive solution: ", duration1)
	fmt.Println("Duration of parallel solution: ", duration2)

}
