package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ToUpper(s string) string {
	a := []rune(s)
	for i, v := range a {
		if v >= 'a' && v <= 'z' {
			a[i] = a[i] - 32
		}
	}
	return string(a)
}

func Ispecial(c rune) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	x := []rune(s)
	vrai := true
	for i := range x {
		if Ispecial(x[i]) && vrai {
			if x[i] >= 'a' && x[i] <= 'z' {
				x[i] = x[i] - 32
			}
			vrai = false
		} else if x[i] >= 'A' && x[i] <= 'Z' {
			x[i] = x[i] + 32
		} else if !(Ispecial(x[i])) {
			vrai = true
		}
	}
	return string(x)
}

func ToLower(s string) string {
	a := []rune(s)
	for i, v := range a {
		if v >= 'A' && v <= 'Z' {
			a[i] = a[i] + 32
		}
	}
	return string(a)
}

func convertbin(b string) int {
	nbr, _ := strconv.Atoi(b)
	c := 1
	var n int
	var res int
	for nbr != 0 {
		n = nbr % 10
		res = res + (n * c)
		c = c * 2
		nbr = nbr / 10
	}
	return res
}

func converthex(b string) int {
	tab := []rune(b)
	c := 1
	n0 := ""
	var n int
	var res int
	for i := len(tab) - 1; i >= 0; i-- {
		if tab[i] == 'A' || tab[i] == 'a' {
			n = 10
		}
		if tab[i] == 'B' || tab[i] == 'b' {
			n = 11
		}
		if tab[i] == 'C' || tab[i] == 'c' {
			n = 12
		}
		if tab[i] == 'D' || tab[i] == 'd' {
			n = 13
		}
		if tab[i] == 'E' || tab[i] == 'e' {
			n = 14
		}
		if tab[i] == 'F' || tab[i] == 'f' {
			n = 15
		}
		if tab[i] >= '0' && tab[i] <= '9' {
			n0 = string(tab[i])
			n1, _ := strconv.Atoi(n0)
			n = n1
		}
		res = res + (n * c)
		c = c * 16
	}

	return res
}

func Indice(S string) int {
	value := []rune{}
	value2 := []rune(S)
	for i := 0; i < len(value2)-2; i++ {
		value = append(value, value2[i])
	}
	val := string(value)
	valint, _ := strconv.Atoi(val)
	return valint
}

func Spaces(S string) string {
	get := []rune(S)
	put := []rune{}
	for i := 0; i < len(get)-1; i++ {
		put = append(put, get[i])
	}
	return string(put)

}

func OnlyPonctuation(S string) bool {
	tab := []rune(S)
	if tab[len(tab)-1]== ' '{
		for i := 0; i < len(tab)-1; i++ {
			if tab[i] != '.' && tab[i] != ',' && tab[i] != ';' && tab[i] != ':' && tab[i] != '!' && tab[i] != '?' {
				return false
				break
			}
		}
	}else{
		for i := 0; i < len(tab); i++ {
			if tab[i] != '.' && tab[i] != ',' && tab[i] != ';' && tab[i] != ':' && tab[i] != '!' && tab[i] != '?' {
				return false
				break
			}
		}
	}
	return true
}

func FirstPonctuation(S string) bool {
	tab := []rune(S)
	if tab[0] == '.' || tab[0] == ',' || tab[0] == ';' || tab[0] == ':' || tab[0] == '!' || tab[0] == '?' {
		return true
	}
	return false
}

func MiddlePonctuation(S string) bool {
	tab := []rune(S)
	for i := 1; i < len(tab)-2; i++ {
		if tab[i] == '.' || tab[i] == ',' || tab[i] == ';' || tab[i] == ':' || tab[i] == '!' || tab[i] == '?' {
			return true
			break
		}
	}
	return false
}

func Ponctuat1(S string, indexe int, result []string) {
	get := []rune(S)
	put := []rune{}
	if get[len(get)-1]== ' '{
		for i := 0; i < len(get)-1; i++ {
			put = append(put, get[i])
		}
	}else{
		for i := 0; i < len(get); i++ {
			put = append(put, get[i])
		}
	}
	
	result[indexe] = " "
	temp := []rune(result[indexe-1])
	put2 := []rune{}
	for i := 0; i < len(temp)-1; i++ {
		put2 = append(put2, temp[i])
	}
	result[indexe-1] = string(put2) + string(put)
}

func Ponctuat2(S string, indexe int, result []string) {
	get := []rune(S)
	put := []rune{}
	put3 := []rune{}
	c := 0
	for i := 0; i < len(get); i++ {
		if get[i] == '.' || get[i] == ',' || get[i] == ';' || get[i] == ':' || get[i] == '!' || get[i] == '?' {
			put = append(put, get[i])
			c++
		} else {
			break
		}
	}
	for j := c; j < len(get); j++ {
		put3 = append(put3, get[j])
	}
	result[indexe] = string(put3)
	temp := []rune(result[indexe-1])
	put2 := []rune{}
	for i := 0; i < len(temp)-1; i++ {
		put2 = append(put2, temp[i])
	}
	result[indexe-1] = string(put2) + string(put) + " "
}

func Ponctuat3(S string, indexe int, result []string) {
	get := []rune(S)
	put1 := []rune{}
	put2 := []rune{}
	c := 0
	for i := 0; i < len(get)-1; i++ {
		put1 = append(put1, get[i])
		c++
		if get[i] == '.' || get[i] == ',' || get[i] == ';' || get[i] == ':' || get[i] == '!' || get[i] == '?' {

			if get[i+1] == '.' || get[i+1] == ',' || get[i+1] == ';' || get[i+1] == ':' || get[i+1] == '!' || get[i+1] == '?' {
				continue
			} else {
				break
			}
		}
	}
	for j := c; j < len(get)-1; j++ {
		put2 = append(put2, get[j])
	}
	result[indexe] = string(put1) + " " + string(put2) + " "
}

func Apposs(s string) string {

	s2 := ""

	cp := false
	for i := 0; i < len(s); i++ {

		if s[i] == 39 && i < len(s)-1 && !cp {
			s2 += string(s[i])
			if s[i+1] == ' ' {
				i += 2
			}

			cp = true
		}
		if s[i] == ' ' && i < len(s)-1 && cp && s[i+1] == 39 {
			i++
			cp = false
		} else if s[i] == 39 && i < len(s)-1 && cp {
			cp = false
		}
		s2 += string(s[i])

	}

	return s2
}

func main() {
	if len(os.Args)==3{
	file := os.Args[1]
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	output := string(data)
	out := []rune(output)
	tabstring := []string{}
	word := []rune{}

	for i := 0; i < len(out); i++ {
		word = append(word, out[i])
		if out[i] == ' ' || i == len(out)-1 {
			word1 := string(word)
			tabstring = append(tabstring, word1)
			word1 = ""
			word = nil
		}
	}
	result := tabstring
	for i := 0; i < len(result); i++ {
		if result[i] == "(cap) " {
			result[i-1] = Capitalize(result[i-1])
		}
		if result[i] == "(up) " {
			result[i-1] = ToUpper(result[i-1])
		}
		if result[i] == "(low) " {
			result[i-1] = ToLower(result[i-1])
		}
		if result[i] == "(hex) " {
			result[i-1] = strconv.Itoa(converthex(Spaces(result[i-1]))) + " "
		}
		if result[i] == "(bin) " {
			result[i-1] = strconv.Itoa(convertbin(Spaces(result[i-1]))) + " "
		}
		if result[i] == "(cap, " {
			temp := Indice(result[i+1])
			for j := i - 1; j >= i-temp; j-- {
				result[j] = Capitalize(result[j])
			}
		}
		if result[i] == "(up, " {
			temp := Indice(result[i+1])
			for j := i - 1; j >= i-temp; j-- {
				result[j] = ToUpper(result[j])
			}
		}
		if result[i] == "(low, " {
			temp := Indice(result[i+1])
			for j := i - 1; j >= i-temp; j-- {
				result[j] = ToLower(result[j])
			}
		}
		if result[i] == "(hex, " {
			temp := Indice(result[i+1])
			for j := i - 1; j >= i-temp; j-- {
				result[j] = strconv.Itoa(converthex(Spaces(result[j])))
			}
		}
		if result[i] == "(bin, " {
			temp := Indice(result[i+1])
			for j := i - 1; j >= i-temp; j-- {
				result[j] = strconv.Itoa(convertbin(Spaces(result[j])))
			}
		}
		if OnlyPonctuation(result[i]) {
			Ponctuat1(result[i], i, result)
		}
		if FirstPonctuation(result[i]) {
			Ponctuat2(result[i], i, result)
		}
		if MiddlePonctuation(result[i]) {
			Ponctuat3(result[i], i, result)
		}
	}

	Semifinal := ""
	for i := 0; i < len(result); i++ {
		Semifinal += result[i]
	}
	RuneFinal := []rune(Semifinal)
	PutFinal := []rune{}
	for i := 0; i < len(RuneFinal); i++ {
		if RuneFinal[i] == ' ' && (RuneFinal[i-1] == 'a' || RuneFinal[i-1] == 'A') && (RuneFinal[i+1] == 'a' || RuneFinal[i+1] == 'e' || RuneFinal[i+1] == 'i' || RuneFinal[i+1] == 'o' || RuneFinal[i+1] == 'u' || RuneFinal[i+1] == 'h' || RuneFinal[i+1] == 'A' || RuneFinal[i+1] == 'E' || RuneFinal[i+1] == 'I' || RuneFinal[i+1] == 'O' || RuneFinal[i+1] == 'U' || RuneFinal[i+1] == 'H') && (RuneFinal[i-2] == ' ') {
			PutFinal = append(PutFinal, 'n')
		}
		PutFinal = append(PutFinal, RuneFinal[i])
	}

	final := ""
	for i := 0; i < len(PutFinal); i++ {
		for PutFinal[i] == '(' {
			for PutFinal[i] != ')' {
				i += 1
			}
			if PutFinal[i] == ')' && PutFinal[i+1] != ' '{
				i++
			}
			if PutFinal[i] == ')' && PutFinal[i+1] == ' '{
				i+=2
			}

		}
		final += string(PutFinal[i])
	}
	out1 := []rune(final)
	tabstring1 := []string{}
	worl2:= []rune{}
	for i := 0; i < len(out1); i++ {
		worl2= append(worl2, out1[i])
		if out1[i] == ' ' || i == len(out1)-1 {
			worl3 := string(worl2)
			tabstring1 = append(tabstring1, worl3)
			worl3 = ""
			worl2= nil
		}
	}
	for i := 0; i < len(tabstring1); i++ {
		if OnlyPonctuation(tabstring1[i]) {
			Ponctuat1(tabstring1[i], i, tabstring1)
		}
		if FirstPonctuation(tabstring1[i]) {
			Ponctuat2(tabstring1[i], i, tabstring1)
		}
		if MiddlePonctuation(tabstring1[i]) {
			Ponctuat3(tabstring1[i], i, tabstring1)
		}

	}
	finalement := ""
	for i := 0; i < len(tabstring1); i++ {
		finalement += string(tabstring1[i])
	}
    // Ouverture du fichier en mode écriture
    correct, err := os.Create(os.Args[2])
    if err != nil {
        fmt.Println(err)
        return
    }

    // Écriture de texte dans le fichier
    _, err = correct.WriteString(Apposs(finalement))
    if err != nil {
        fmt.Println(err)
    	correct.Close()
        return
    }

    // Fermeture du fichier
    err = correct.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

}
