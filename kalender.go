package main

var maandenNormaal = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func isSchrikkeljaar(jaar int) bool {
	if jaar < 0 {
		jaar *= -1
	}
	if jaar%100 == 0 && jaar%400 != 0 {
		return false
	} else {
		return jaar%4 == 0
	}
}

func aantalJaarDagen(jaar int, maand int) int {
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
		if isSchrikkeljaar(input.jaar) && input.maand < 2 {
			dagnr += 1
		}
	}
	dagnr += input.dag
	var dagenVerschil int = dagnr - startDagnr
	dagenVerschil += aantalJaarDagen(input.jaar, input.maand)
	return dagenVerschil
}

func naarNieuweDatumCHI(input int) Datum {
	var aantalJaar int
	var nieuweDatum Datum
	nieuweDatum.dag = ((input%maand + maand) % maand) + 1
	if input < 0 {
		aantalJaar = (input / jaar) - 1
		input = input - ((aantalJaar + 1) * jaar)
		input = jaar + input
	} else {
		aantalJaar = input / jaar
		input = input - (aantalJaar * jaar)
	}
	nieuweDatum.jaar = aantalJaar
	nieuweDatum.maand = (input / maand) + 1
	return nieuweDatum
}

func naarNieuweDatumGRE(input Datum) Datum {
	aantalDagen := (input.dag - 1) + ((input.maand - 1) * maand) + (input.jaar * jaar)
	var nieuweDatum Datum = start
	var ditjaarschrikkel bool = false
	if aantalDagen > -1 {
		for i := 0; i < aantalDagen; i++ {
			if !ditjaarschrikkel || nieuweDatum.maand != 2 {
				nieuweDatum.dag = (nieuweDatum.dag % (maandenNormaal[nieuweDatum.maand-1])) + 1
			} else {
				nieuweDatum.dag = (nieuweDatum.dag % 29) + 1
			}
			if nieuweDatum.dag == 1 {
				nieuweDatum.maand = (nieuweDatum.maand % 12) + 1
			}
			if nieuweDatum.dag == 1 && nieuweDatum.maand == 1 {
				nieuweDatum.jaar += 1
				if isSchrikkeljaar(nieuweDatum.jaar) {
					ditjaarschrikkel = true
				} else {
					ditjaarschrikkel = false
				}
			}
		}
	} else {
		//TODO
		aantalDagen *= -1
		for i := 0; i < aantalDagen; i++ {
			if nieuweDatum.dag == 1 {
				if nieuweDatum.maand != 1 {
					nieuweDatum.maand -= 1
				} else {
					nieuweDatum.maand = 12
				}
			}
			//TOT HIER bekeken
			if !ditjaarschrikkel || nieuweDatum.maand != 2 {
				nieuweDatum.dag = (nieuweDatum.dag % (maandenNormaal[nieuweDatum.maand-1])) - 1
			} else {
				nieuweDatum.dag = (nieuweDatum.dag % 29) - 1
			}

			if nieuweDatum.dag == 1 && nieuweDatum.maand == 1 {
				nieuweDatum.jaar += 1
				if isSchrikkeljaar(nieuweDatum.jaar) {
					ditjaarschrikkel = true
				} else {
					ditjaarschrikkel = false
				}
			}
		}
	}
	return nieuweDatum
}
