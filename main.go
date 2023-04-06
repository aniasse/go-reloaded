package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)
func spaceaposs(s string) string{
	tab:=[]rune(s)
	tab1:=[]rune{}

	for i:=0; i<len(tab); i++{
		if i==0 && tab[i]==39{
			tab1=append(tab1, tab[i])
			tab1=append(tab1, ' ')
		}else{
		if i< len(tab)-2 && i>0 && tab[i+1]==39 && tab[i]!=' ' && tab[i+2]==' '{
			tab1=append(tab1, tab[i])
			tab1=append(tab1, ' ')
		}else if i< len(tab)-1 && i>0 && tab[i]==39 && tab[i+1]!=' ' && tab[i-1]==' '{
			tab1=append(tab1, tab[i])
			tab1=append(tab1, ' ')
		}else{
			tab1=append(tab1, tab[i])
		}
		}
	}
	return string(tab1)
}
func double(s string) string{
	tab:=[]rune(s)
	tab1:=[]rune{}

	for i:=0; i<len(tab); i++{
		if i>0 && tab[i]==' ' && tab[i-1]==' '{
			continue
		}else{
			tab1=append(tab1, tab[i])
		}
	}
	return string(tab1)
}

func spaceapoint(s string) string{
	tab:=[]rune(s)
	tab1:=[]rune{}

	for i:=0; i<len(tab); i++{
		if i<len(tab)-1 && (tab[i] == '.' || tab[i] == ',' || tab[i] == ';' || tab[i] == ':' || tab[i] == '!' || tab[i] == '?' ) && tab[i+1]!=' '{
			tab1=append(tab1, tab[i])
			tab1=append(tab1, ' ')
		}else{
			tab1=append(tab1, tab[i])
		}
	}
	return string(tab1)
}
func supspace(s string)string{
	t:=[]rune(s)
	t1:=[]rune{}
	for i:=0; i<len(t); i++{
		if i==1 && t[0]==' ' && t[i] == 39{
			continue
		}
		if i==1 && t[0]==39 && t[i] ==' ' {
			continue
		}
		if i>0 && i<len(t)-1 && t[i]==' ' && (t[i+1] == '.' || t[i+1] == ',' || t[i+1] == ';' || t[i+1] == ':' || t[i+1] == '!' || t[i+1] == '?' ){
			continue
		}else{
			t1=append(t1, t[i])
		}
	}
	return string(t1)
}
func SeparePonctuation(s string, index int , t []string){
	tab:=[]rune(s)
	tabponctuation:=[]rune{}
	reste:=[]rune{}
	ind:=0
	for i:=0; i<len(tab); i++{
		if tab[i]==',' || tab[i]==';' || tab[i]=='!' || tab[i]=='?' || tab[i]=='.'|| tab[i]==':'{
			tabponctuation=append(tabponctuation, tab[i])
			ind++
		}else{
			break
		}
	}
	for j:=ind; j<len(tab); j++{
		reste=append(reste, tab[j])
	}
	
	t[index-1]+=string(tabponctuation)
	t[index]=string(reste)
}

func Ponctuat3(S string, indexe int, result []string) {
	get := []rune(S)
	put1 := []rune{}
	put2 := []rune{}
	c := 0
	for i := 0; i < len(get)-1; i++ {
		put1 = append(put1, get[i])
		c++
		if get[i] == '.' || get[i] == ',' || get[i] == ';' || get[i] == ':' || get[i] == '!' || get[i] == '?' || get[i]==':'{
			if get[i+1] == '.' || get[i+1] == ',' || get[i+1] == ';' || get[i+1] == ':' || get[i+1] == '!' || get[i+1] == '?' {
				continue
			} else {
				break
			}
		}
	}
	for j := c; j < len(get); j++ {
		put2 = append(put2, get[j])
	}
	result[indexe] = string(put1) + string(put2)
}

func FirstPonctuation(s string) bool{
	tab:=[]rune(s)
	//count:=0
	if (tab[0]== ',' || tab[0]==';' || tab[0]=='!' || tab[0]=='?' || tab[0]=='.'|| tab[0]==':') && len(tab)>1{
		for i:=1; i<len(tab); i++{
			if tab[i]!= ',' && tab[i]!=';' && tab[i]!='.' && tab[i]!='!' && tab[i]!='?' && tab[i]!= ':'{
				return true
				break
			}
		}
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

//Retourne true si le string est constitiuer que de ponctuations
func UnikPontuation(s string) bool{
	tab:=[]rune(s)
	for _, v:= range tab{
		if v!='.' && v!=',' && v!=';' && v!='!' && v!='?' && v!=':' {
			return false
			break
		}
	}
	return true
}

//Retourne le int que se trouve a l'interieur d'un string comme "3)"
func trier(s string) int{
	tab:=[]rune(s)
	tab2:=[]rune{}

	for i:=0; i<len(tab)-1;i++{
		tab2=append(tab2, tab[i])
	}
	str:=string(tab2)
	entier, _:=strconv.Atoi(str)

	return entier
}
func parenthese(s string) bool {
	t:=[]rune(s)
	if t[len(t)-1]==')'{
		return true
	}else{
		return false
	}
}

func Stringtab(t []string) string{
	world:=""
	for _, v := range t {
		if world==""{
			world+=v
		}else{
			world=world+" "+v
		}
	}
	return world
}
func supprime2(t []string, ind int) []string{

	tab:=[]string{}
	for i:=0; i<len(t); i++ {
		if i!=ind {
			tab=append(tab, t[i])
		}
	}
	return tab
}
func supprime3(t []string, ind int) []string{

	tab:=[]string{}
	for i:=0; i<len(t); i++ {
		if i!=ind && i!=ind+1{
			tab=append(tab, t[i])
		}
	}
	return tab
}

func sup(t2 []string) []string{
	t:=[]string{}
	for i:=0; i<len(t2); i++{
		if t2[i] == "(cap)" || t2[i] == "(up)" || t2[i] == "(low)" || t2[i] == "(hex)" || t2[i] == "(bin)" {
			continue
		}else if t2[i] == "(cap," && len(t2)>1 && isNumb(t2[i+1]) && parenthese(t2[i+1]) || t2[i] == "(up," && len(t2)>1 && isNumb(t2[i+1]) && parenthese(t2[i+1])  || t2[i] == "(low," && len(t2)>1 && isNumb(t2[i+1]) && parenthese(t2[i+1]) || t2[i] == "(hex," && len(t2)>1 && isNumb(t2[i+1]) && parenthese(t2[i+1]) || t2[i] == "(bin," && len(t2)>1 && isNumb(t2[i+1]) && parenthese(t2[i+1])  {
			i+=1
			continue
		} else{
			t=append(t, t2[i])
		}

	}
	return t
}
func correct(t []string) []string{
	str:=Stringtab(t)
	Sortie:=[]rune(str)
	PutFinal := []rune{}
	for i := 0; i < len(Sortie); i++ {
		if i-1==0 && Sortie[i] == ' ' && (Sortie[i-1] == 'a' || Sortie[i-1] == 'A') && (Sortie[i+1] == 'a' || Sortie[i+1] == 'e' || Sortie[i+1] == 'i' || Sortie[i+1] == 'o' || Sortie[i+1] == 'u' || Sortie[i+1] == 'h' || Sortie[i+1] == 'A' || Sortie[i+1] == 'E' || Sortie[i+1] == 'I' || Sortie[i+1] == 'O' || Sortie[i+1] == 'U' || Sortie[i+1] == 'H') {
			PutFinal = append(PutFinal, 'n')
		}

		if i >2  && Sortie[i] == ' ' && (Sortie[i-1] == 'a' || Sortie[i-1] == 'A') && (Sortie[i+1] == 'a' || Sortie[i+1] == 'e' || Sortie[i+1] == 'i' || Sortie[i+1] == 'o' || Sortie[i+1] == 'u' || Sortie[i+1] == 'h' || Sortie[i+1] == 'A' || Sortie[i+1] == 'E' || Sortie[i+1] == 'I' || Sortie[i+1] == 'O' || Sortie[i+1] == 'U' || Sortie[i+1] == 'H') && (Sortie[i-2] == ' ') {
			PutFinal = append(PutFinal, 'n')
		}
		PutFinal = append(PutFinal, Sortie[i])
	}

	return strings.Split(string(PutFinal), " ")
}

func Apposs(s []string) string {
	s2 := ""
	count := false
	for i := 0; i < len(s); i++ {
		if s[i] == "'" && !(count) {
			s2 += s[i]
			count = true
		} else if s[i] == "'" && (count) {
			s2 += s[i] + " "
			count = false
		} else {
			if i < len(s)-1 && s[i+1] != "'"  {
				s2 += s[i] + " "
			} else if i < len(s)-1 && s[i+1] == "'" && !count {
				s2 += s[i] + " "
			} else {
				s2 += s[i]
			}
		}
	}
	return s2
}

func instruction1(s string) bool{
	if s=="(cap)" || s=="(up)" || s=="(low)" || s=="(hex)" || s=="(bin)" {
		return true
	}
	return false
}

func instruction2(s string) bool{
	if s=="(cap," || s=="(up," || s=="(low," || s=="(hex," || s=="(bin," {
		return true
	}
	return false
}
func isNumb(s string) bool {
	tab:=[]rune(s)
	tab1:=tab[:len(tab)-1]
	for _, v:=range tab1{
		if v<'0' || v > '9'{
			return false
			break
		}
	}
	return true
}
func uniksapce(s string) bool {
	tab:=[]rune(s)
	for _, v:= range tab{
		if v!=' ' && v!='\n' && v!='\t'{
			return false
			break
		}
	}
	return true
}

func main() {
	//Recuperer le contenue du fichier en byte dans un variable
	txt ,er:=ioutil.ReadFile(os.Args[1])
	if er != nil{
		fmt.Println("le fichier rencontre des erreurs")
	}
	conv:=string(txt)
	if !(uniksapce(conv)){
//	conv:=double(conv1)

	tabString := strings.Split(conv, " ")
	tabString=correct(tabString)
	if (len(tabString)==1 && instruction1(tabString[0])){
		tabString=nil
	}
	if (len(tabString)==2 && instruction2(tabString[0]) && isNumb(tabString[1]) && parenthese(tabString[1])){
		tabString=nil
	}
	// tabString=firsrinstance(tabString)
	for i:=0; i<len(tabString); i++{
		if i==0 && instruction1(tabString[i]){
			tabString=supprime2(tabString, i)
		}
		if i > 0 && tabString[i]=="(bin)"{
			entier, _:=strconv.ParseInt(tabString[i-1], 2, 64)
			entier2:=strconv.FormatInt(entier, 10)
			if entier2=="0"{
				tabString[i-1]=tabString[i-1]
			}else{
			tabString[i-1]=entier2
			}
			tabString=supprime2(tabString, i)
			i-=1
		}
		if i > 0 && tabString[i]=="(hex)"{
			entier, _:=strconv.ParseInt(tabString[i-1], 16, 64)
			entier2:=strconv.FormatInt(entier, 10)
			if entier2=="0"{
				tabString[i-1]=tabString[i-1]
			}else{
			tabString[i-1]=entier2
			}
			tabString=supprime2(tabString, i)
			i-=1
		}
		if i > 0 && tabString[i]=="(up)"{
			tabString[i-1]=strings.ToUpper(tabString[i-1])
			tabString=supprime2(tabString, i)
			i-=1
		}
		if i > 0 && tabString[i]=="(cap)"{
			tabString[i-1]=Capitalize(tabString[i-1])
			tabString=supprime2(tabString, i)
			i-=1
		}
		if i > 0 && tabString[i]=="(low)"{
			tabString[i-1]=strings.ToLower(tabString[i-1])
			tabString=supprime2(tabString, i)
			i-=1
		}
		if i>0 && tabString[i]=="(cap," && isNumb(tabString[i+1]) && parenthese(tabString[i+1]){
			limit:=trier(tabString[i+1])
			for j:=i-1; j>i-limit-1; j--{
				if j<0{
					break
				}
				tabString[j]=Capitalize(tabString[j])
			}
			tabString=supprime3(tabString, i)
			i-=2
		}
		if i > 0 && tabString[i]=="(up," && isNumb(tabString[i+1]) && parenthese(tabString[i+1]){
			limit:=trier(tabString[i+1])
			for j:=i-1; j>i-limit-1; j--{
				if j<0{
					break
				}
				tabString[j]=strings.ToUpper(tabString[j])
			}
			tabString=supprime3(tabString, i)
			i-=2
		}
		if i > 0 && tabString[i]=="(low," && isNumb(tabString[i+1]) && parenthese(tabString[i+1]){
			limit:=trier(tabString[i+1])
			for j:=i-1; j>i-limit-1; j--{
				if j<0{
					break
				}
				tabString[j]=strings.ToLower(tabString[j])
			}
			tabString=supprime3(tabString, i)
			i-=2
		}
		if i > 0 && tabString[i]=="(hex," && isNumb(tabString[i+1]) && parenthese(tabString[i+1]){
			limit:=trier(tabString[i+1])
			for j:=i-1; j>i-limit-1; j--{
				if j<0{
					break
				}
				entier, _:=strconv.ParseInt(tabString[j], 16, 64)
				entier2:=strconv.FormatInt(entier, 10)
				tabString[j]=entier2
			}
			tabString=supprime3(tabString, i)
			i-=2
		}
		if i > 0 && tabString[i]=="(bin," && isNumb(tabString[i+1]) && parenthese(tabString[i+1]){
			limit:=trier(tabString[i+1])
			for j:=i-1; j>i-limit-1; j--{
				if j<0{
					break
				}
				entier, _:=strconv.ParseInt(tabString[j], 2, 64)
				entier2:=strconv.FormatInt(entier, 10)
				tabString[j]=entier2
			}
			tabString=supprime3(tabString, i)
			i-=2
		}
		if i > 0 && UnikPontuation(tabString[i]){
			tabString[i-1]=tabString[i-1]+tabString[i]
			tabString=supprime2(tabString, i)
			i-=1
		}else{
			if i > 0 && MiddlePonctuation(tabString[i]) {
				Ponctuat3(tabString[i], i, tabString)
			}
			if i > 0 && FirstPonctuation(tabString[i]){
				SeparePonctuation(tabString[i], i, tabString)
			}
		}
		
	 }
	
	tabString=correct(tabString)
	tabString=strings.Split(spaceapoint(Stringtab(tabString)), " ")
  	tabString=strings.Split(spaceaposs(Stringtab(tabString)), " ")
	tabString=strings.Split(Apposs(tabString), " ")
	tabString=sup(tabString)
	phrase:=Stringtab(tabString)
	phrase1:=supspace(phrase)

	// fmt.Println(string(corrected))
	correct, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = correct.WriteString(string(phrase1))
	if err != nil {
		fmt.Println(err)
		correct.Close()
		return
	}
	err = correct.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}else{
	correct, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = correct.WriteString(conv)
	if err != nil {
		fmt.Println(err)
		correct.Close()
		return
	}
	err = correct.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

}