package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rotors struct {
	Table  []int //alfabet
	Flaga  int   //kiedy on robi obrót
	CurPos int   //pozycja początkowa
	Turn   bool  //czy robimy obrót
}

func InputRotors() (rtr, pos string) {
	fmt.Println("Masz 5 rotorów jakie możesz wykorzystać do szyfrowania")
	fmt.Println("1 rotor: ZNCWMSDIYOGAQBKFJEXLRUHTVP")
	fmt.Println("2 rotor: BKMWAEPTCVFOHNXDSYRUQGZILJ")
	fmt.Println("3 rotor: HPIVJESRFAYOZGKXMNLTUDCWQB")
	fmt.Println("4 rotor: GOZFWSUMCDLEBKAITPNXJYHQRV")
	fmt.Println("5 rotor: UIBJRDLCZPMKHTAYNSVFEQXOGW")
	fmt.Println("Napisz 3 rotory, które chcesz wykorzystać (Naprzykład: 123): ")
	rtr = strings.ReplaceAll(InputText(), " ", "")
	fmt.Println("Teraż napisz pozycje początkowe (Naprzykład: 2 11 23): ")
	pos = InputText() //strings.ReplaceAll(InputText(), " ", "")
	return
}

func MakeRotors(rtr, pos string, rOne, rTwo, rThree *Rotors) {
	var choice []int

	posValues := strings.Fields(pos)

	for _, char := range posValues {
		a, _ := strconv.Atoi(char)
		choice = append(choice, a)
	}

	numbers := []byte(rtr)
	r1r2r3 := []*Rotors{rOne, rTwo, rThree}

	for i := 0; i < 3; i++ {
		r := r1r2r3[i]
		switch numbers[i] - 48 {
		case 1:
			r.Table = []int{25, 13, 2, 22, 12, 18, 3, 8, 24, 14, 6, 0, 16, 1, 10, 5, 9, 4, 23, 11, 17, 20, 7, 19, 21, 15}
			r.Flaga = 23
			r.CurPos = choice[i]
			r.Turn = false
		case 2:
			r.Table = []int{1, 10, 12, 22, 0, 4, 15, 19, 2, 21, 5, 14, 7, 13, 23, 3, 18, 24, 17, 20, 16, 6, 25, 8, 11, 9}
			r.Flaga = 11
			r.CurPos = choice[i]
			r.Turn = false
		case 3:
			r.Table = []int{7, 15, 8, 21, 9, 4, 18, 17, 5, 0, 24, 14, 25, 6, 10, 23, 12, 13, 11, 19, 20, 3, 2, 22, 16, 1}
			r.Flaga = 7
			r.CurPos = choice[i]
			r.Turn = false
		case 4:
			r.Table = []int{6, 14, 25, 5, 22, 18, 20, 12, 2, 3, 11, 4, 1, 10, 0, 8, 19, 15, 13, 23, 9, 24, 7, 16, 17, 21}
			r.Flaga = 3
			r.CurPos = choice[i]
			r.Turn = false
		case 5:
			r.Table = []int{20, 8, 1, 9, 17, 3, 11, 2, 25, 15, 12, 10, 7, 19, 0, 24, 13, 18, 21, 5, 4, 16, 23, 14, 6, 22}
			r.Flaga = 17
			r.CurPos = choice[i]
			r.Turn = false
		}

		table := r.Table
		r.Table = append(table[r.CurPos:], table[:r.CurPos]...)
	}
	rOne.Turn = true
}

func Rotation(rOne, rTwo, rThree *Rotors) {
	r1r2r3 := []*Rotors{rOne, rTwo, rThree}

	for i := 0; i < 3; i++ {
		r := r1r2r3[i]
		if i != 2 {
			r1 := r1r2r3[i+1]
			if r.CurPos == r1.Flaga {
				r1.CurPos++
				if r1.CurPos > 25 {
					r1.CurPos -= 25
				}
				r1.Turn = true
			}
		}

		if r.Turn == true {
			firstEle := r.Table[0]
			r.Table = r.Table[1:]
			r.Table = append(r.Table, firstEle)
			rOne.CurPos++
			r.Turn = false
		}
	}
	rOne.Turn = true
}

func ForwardPath(letter *int, rOne, rTwo, rThree *Rotors) {
	r1r2r3 := []*Rotors{rOne, rTwo, rThree}

	for i := 0; i < 3; i++ {
		r := r1r2r3[i]
		for j := range r.Table {
			if *letter == j {
				*letter = r.Table[j]
				break
			}
		}
	}
}

func Miror(letter *int) {
	MirrorR := map[int]int{
		0: 8, 1: 9, 2: 4, 3: 6, 4: 2, 5: 13, 6: 3, 7: 15,
		8: 0, 9: 1, 10: 20, 11: 14, 12: 24, 13: 5, 14: 11,
		15: 7, 16: 25, 17: 22, 18: 23, 19: 21, 20: 10,
		21: 19, 22: 17, 23: 18, 24: 12, 25: 16}

	for key, value := range MirrorR {
		if *letter == key {
			*letter = value
			break
		}
	}
}

func ReturnPath(letter *int, rOne, rTwo, rThree *Rotors) {
	r1r2r3 := []*Rotors{rThree, rTwo, rOne}

	for i := 0; i < 3; i++ {
		r := r1r2r3[i]
		for j := range r.Table {
			if *letter == r.Table[j] {
				*letter = j
				break
			}
		}
	}
}

func FindByte(alfabet map[byte]int, j byte) int {

	for key, value := range alfabet {
		if j == key {
			return value
		}
	}
	return 0
}

func FindLetter(alfabet map[byte]int, j int) byte {
	for key, value := range alfabet {
		if j == value {
			return key
		}
	}
	return 0
}

func Enigma(txt string) {
	letters := []byte(txt)
	ciphertext := make([]byte, len(letters))
	rotor1 := &Rotors{}
	rotor2 := &Rotors{}
	rotor3 := &Rotors{}

	trans := map[byte]int{
		'A': 14, 'B': 7, 'C': 13, 'D': 4, 'E': 10, 'F': 22,
		'G': 0, 'H': 6, 'I': 17, 'J': 25, 'K': 18, 'L': 9,
		'M': 20, 'N': 23, 'O': 15, 'P': 1, 'Q': 24, 'R': 19,
		'S': 2, 'T': 8, 'U': 11, 'V': 21, 'W': 5, 'X': 12, 'Y': 3, 'Z': 16}

	rtrs, poss := InputRotors()
	MakeRotors(rtrs, poss, rotor1, rotor2, rotor3)

	for i := range letters {
		leter := FindByte(trans, letters[i])
		ForwardPath(&leter, rotor1, rotor2, rotor3)
		Miror(&leter)
		ReturnPath(&leter, rotor1, rotor2, rotor3)
		Rotation(rotor1, rotor2, rotor3)

		ciphertext[i] = FindLetter(trans, leter)
	}

	result := string(ciphertext)
	fmt.Println("Wynik: ", result)
}

func InputText() (text string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	return
}

func ReadText(text *string) bool {
	fmt.Println("Podaj nazwę pliku: ")
	plikName := InputText()            // nazwa
	data, err := os.ReadFile(plikName) // dane byte
	if err != nil {
		fmt.Println("Błąd odczytu pliku: ", err)
		return false
	}
	*text = string(data) // byte -> string
	return true
}

func Menu() (sentence string) {
	fmt.Println("___Enigma___")
	fmt.Println("Jakie dane chcesz podać: ")
	fmt.Println(" > 1: tekst ")
	fmt.Println(" > 2: plik ")
	fmt.Println("Twój wybór: ")

	mode := InputText()
	switch mode {
	case "1":
		fmt.Println("Podaj tekst:")
		sentence = InputText()
	case "2":
		if ReadText(&sentence) != true {
			sentence = ""
		}
	}
	return
}

func main() {
	text := Menu()

	if text != "" {
		fmt.Println("Podany text: ", text)
	} else {
		return
	}

	Enigma(text)
}
