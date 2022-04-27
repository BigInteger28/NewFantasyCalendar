package main

var maandenNormaal = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func isSchrikkeljaar(jaar int) bool {
	if jaar < 0 {
		jaar *= -1
	}
	return jaar%4 == 0
}

func aantalJaarDagen(jaar int) int {
	var aantalDagen int = 0
	if jaar < start.jaar {
		for ; jaar < start.jaar; jaar++ {
			if !isSchrikkeljaar(jaar) {
				aantalDagen -= 365
			} else {
				aantalDagen -= 366
			}
		}
	} else {
		for ; jaar > start.jaar; jaar-- {
			if !isSchrikkeljaar(jaar) {
				aantalDagen += 365
			} else {
				aantalDagen += 366
			}
		}
	}
	return aantalDagen
}

func berekenVerschilDagen(input Datum) int {
	var dagnr int = 0
	//zet om in dagen in het jaar
	if input.maand > 1 {
		for i := 0; i < input.maand-1; i++ {
			dagnr += maandenNormaal[i]
		}
		if isSchrikkeljaar(input.jaar) && input.maand > 2 {
			dagnr += 1
		}
	}
	dagnr += input.dag
	var dagenVerschil int = dagnr - startDagnr
	dagenVerschil += aantalJaarDagen(input.jaar)
	return dagenVerschil
}

func naarNieuweDatum(input int) Datum {
	var aantalJaar int
	var nieuweDatum Datum
	nieuweDatum.dag = ((input%maand + maand) % maand) + 1
	if input < 0 {
		aantalJaar = (input / jaar) - 1
		input = input - ((aantalJaar + 1) * jaar)
		input = jaar + input
	} else {
		aantalJaar = (input / jaar) + 1
		input = input - ((aantalJaar - 1) * jaar)
	}
	nieuweDatum.jaar = aantalJaar
	nieuweDatum.maand = (input / maand) + 1
	return nieuweDatum
}
