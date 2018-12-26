package chislopropisiuru

import (
	"errors"
	"strings"
)

const (
	N0 = "ноль"

	N1a = "один"
	N1b = "одна"

	N2a = "два"
	N2b = "две"

	N3 = "три"
	N4 = "четыре"
	N5 = "пять"
	N6 = "шесть"
	N7 = "семь"
	N8 = "восемь"
	N9 = "девять"

	N10 = "десять"
	N11 = "одинадцать"
	N12 = "двенадцать"
	N13 = "тринадцать"
	N14 = "четырнадцать"
	N15 = "пятнадцать"
	N16 = "шестнадцать"
	N17 = "семнадцать"
	N18 = "восемнадцать"
	N19 = "девятнадцать"

	N20 = "двадцать"
	N30 = "тридцать"
	N40 = "сорок"
	N50 = "пятьдесят"
	N60 = "шестьдесят"
	N70 = "семьдесят"
	N80 = "восемьдесят"
	N90 = "девяносто"

	N100 = "сто"
	N200 = "двести"
	N300 = "триста"
	N400 = "четыреста"
	N500 = "пятьсот"
	N600 = "шестьсот"
	N700 = "семьсот"
	N800 = "восемьсот"
	N900 = "девятьсот"

	N1000  = "тысяч"
	N1000a = "тысяча"
	N1000b = "тысячи"

	N1000000       = "миллион" // "а", "ов"
	N1000000000    = "миллиард"
	N1000000000000 = "триллион"
)

func Itoa(number string) (string, error) {

	//	var err error

	if number == "0" {
		return N0, nil
	}

	ret_s := make([]string, 0)

	splitted_number := SplitString(number)

	splitted_number_len := len(splitted_number)
	splitted_number_len_countdown := splitted_number_len

	for index, i := range splitted_number {

		var female bool
		var multiple bool
		var last bool

		if splitted_number_len_countdown == 2 {
			female = true
		}

		if !strings.HasSuffix(i, "1") {
			multiple = true
		}

		if index == len(splitted_number)-1 {
			last = true
		}

		splitted_number_len_countdown = splitted_number_len - index

		t, err := xxx_toa(i, female, multiple)
		if err != nil {
			return "", err
		}

		ret_s = append(ret_s, t)

		switch splitted_number_len_countdown {
		case 5:
			suff := ""
			if multiple {
				suff = "ов"
			}
			ret_s = append(ret_s, N1000000000000+suff)
		case 4:
			suff := ""
			if multiple {
				suff = "ов"
			}
			ret_s = append(ret_s, N1000000000+suff)
		case 3:
			suff := ""
			if multiple {
				suff = "ов"
			}
			ret_s = append(ret_s, N1000000+suff)
		case 2:
			var word string

			if !multiple {
				word = N1000a
			} else {
				if last {
					word = N1000

				} else {
					word = N1000b
				}
			}

			ret_s = append(ret_s, word)
		case 1:
		}
	}

	return strings.Join(ret_s, " "), nil
}

func SplitString(value string) []string {
	ret := make([]string, 0)

	for len(value) != 0 {
		count := 3
		if len(value) < 3 {
			count = len(value)
		}
		t := value[len(value)-count : len(value)]
		value = value[:len(value)-count]
		ret = append([]string{t}, ret...)
	}

	return ret
}

func xxx_toa(xxx string, female bool, multiple bool) (ret string, err error) {
	ret = ""

	if len(xxx) > 3 {
		return "", errors.New("loo long")
	}

	for len(xxx) != 3 {
		xxx = "0" + xxx
	}

	switch xxx[0] {
	case '0':
	case '1':
		ret += N100
	case '2':
		ret += N200
	case '3':
		ret += N300
	case '4':
		ret += N400
	case '5':
		ret += N500
	case '6':
		ret += N600
	case '7':
		ret += N700
	case '8':
		ret += N800
	case '9':
		ret += N900
	}

	switch xxx[1:3] {
	default:

		switch xxx[1] {
		default:
			err = errors.New("error at xxx[1] check: " + string(xxx[1]))
			return
		case '0':
		case '2':
			ret += N20
		case '3':
			ret += N30
		case '4':
			ret += N40
		case '5':
			ret += N50
		case '6':
			ret += N60
		case '7':
			ret += N70
		case '8':
			ret += N80
		case '9':
			ret += N90
		}

		var t string

		if t, err = x_toa(xxx[2], female, multiple); err != nil {
			return
		} else {
			ret += t
		}
	case "00":
	case "10":
		ret += N10
	case "20":
		ret += N20
	case "30":
		ret += N30
	case "40":
		ret += N40
	case "50":
		ret += N50
	case "60":
		ret += N60
	case "70":
		ret += N70
	case "80":
		ret += N80
	case "90":
		ret += N90
	case "11":
		ret += N11
	case "12":
		ret += N12
	case "13":
		ret += N13
	case "14":
		ret += N14
	case "15":
		ret += N15
	case "16":
		ret += N16
	case "17":
		ret += N17
	case "18":
		ret += N18
	case "19":
		ret += N19
	}

	return
}

func x_toa(x byte, female bool, multiple bool) (ret string, err error) {
	err = nil
	switch x {
	default:
		err = errors.New("invalid value")
	case '0':
		ret = N0
	case '1':
		ret = N1a
		if female {
			ret = N1b
		}
	case '2':
		ret = N2a
		if multiple {
			ret = N2b
		}
	case '3':
		ret = N3
	case '4':
		ret = N4
	case '5':
		ret = N5
	case '6':
		ret = N6
	case '7':
		ret = N7
	case '8':
		ret = N8
	case '9':
		ret = N9
	}
	return
}
