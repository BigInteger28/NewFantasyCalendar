package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Datum struct {
	dag   int
	maand int
	jaar  int
}

var start = Datum{21, 03, 2021}
var startDagnr int = 80

var week = 5
var maand = 20
var maanden = 5
var jaar = 100

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func inputNaarDatum(input string) Datum {
	split := strings.Split(input, "/")
	dag, err := strconv.Atoi(split[0])
	check(err)
	maand, err := strconv.Atoi(split[1])
	check(err)
	jaar, err := strconv.Atoi(split[2])
	check(err)
	var datum = Datum{dag, maand, jaar}
	return datum
}

func main() {
	var input string
	var menu int
	for {
		fmt.Println("\n1. Naar CHI")
		fmt.Println("2. Naar GREGORIAANS")
		fmt.Println("3. CHI Maand kalender")
		fmt.Print("Uw keuze: ")
		fmt.Scanln(&menu)
		if menu == 1 {
			fmt.Print("Gregoriaanse datum: ")
			fmt.Scanln(&input)
			datum := inputNaarDatum(input)
			verschil := berekenVerschilDagen(datum)
			var uitkomst Datum = naarNieuweDatumCHI(verschil)
			fmt.Println(uitkomst.dag, "/", uitkomst.maand, "/", uitkomst.jaar)
		} else if menu == 2 {
			fmt.Print("Chi datum: ")
			fmt.Scanln(&input)
			datum := inputNaarDatum(input)
			var uitkomst Datum = naarNieuweDatumGRE(datum)
			fmt.Println(uitkomst.dag, "/", uitkomst.maand, "/", uitkomst.jaar)
		} else if menu == 3 {
			var startchi string
			var chijaar int
			fmt.Print("Start CHI-datum: ")
			fmt.Scanln(&startchi)
			fmt.Print("Aantal CHI-jaar: ")
			fmt.Scanln(&chijaar)
			datum := inputNaarDatum(startchi)
			gregDat := naarNieuweDatumGRE(datum)
			for m := 0; m < chijaar; m++ {
				for i := 0; i < maanden; i++ {
					fmt.Println(datum.dag, "/", datum.maand, "/", datum.jaar, "\t-\t", gregDat.dag, "/", gregDat.maand, "/", gregDat.jaar)
					datum.maand = ((datum.maand%maanden + maanden) % maanden) + 1
					if datum.maand == 1 {
						datum.jaar++
					}
					gregDat = naarNieuweDatumGRE(datum)
				}
				fmt.Println("")
			}
		}
	}
}
