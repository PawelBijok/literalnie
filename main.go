package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func containsWord(words *[]string, word string) bool {
	for _, w := range *words {
		if w == word {
			return true
		}
	}
	return false
}

func findLetterIndexes(word string, letter string) []int {
	var positions []int
	for i, l := range word {
		if string(l) == letter {
			positions = append(positions, i)
		}
	}
	return positions
}

func containsLetterAtAnyPosition(word string, letters []string) bool {
	for _, letter := range letters {
		if strings.Contains(word, letter) {
			return true
		}
	}
	return false
}

func containsLetterAtSpecyficPosition(word string, letter string, position int) bool {
	if string(word[position]) == letter {
		return true
	}
	return false

}

func containsLetterAnywhereExceptSpecyficPosition(word string, letter string, position int) bool {
	foundPositions := findLetterIndexes(word, letter);
	if(len(foundPositions) == 0) {
		return false;
	}
	for _, foundPosition := range foundPositions {
		if(foundPosition == position) {
			return false
		}
	}
	
	return true
}


func main() {
	fmt.Println("Welcome to nut cracker!")
	file, err := os.Open("/Users/pawelbijok/go/bin/tmp/slowa.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var words []string
	// wordsPossible := make([]string, 0,7000)


	//read line from file
	scanner := bufio.NewScanner(file)

	fmt.Println("Wczytuje słówka.... doczkej...")
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	fmt.Printf("Wczytołech! \n\n")

	
	
	for  {

	inputError := true

	//entering word
	var enteredWorld string

	for inputError == true {
	fmt.Print("Wpisz pięcioliterowe słowo: ")
	fmt.Scanln(&enteredWorld) 
	enteredWorld = strings.ToUpper(enteredWorld)
	if len(enteredWorld) != 5 {
		fmt.Println("To nie jest pięcioliterowe słowo debilu!")
		inputError = true
		continue
	}
	if containsWord(&words, enteredWorld) == false {
		fmt.Println("Nie ma takiego słowa w słowniku!")
		inputError = true
		continue
	}
	inputError = false
	}

	inputError = true
	//entering weight
	var enteredWeights string


	for inputError == true {
		fmt.Print("Wpisz kolejne wagi liter (0: nie ma, 1: jest ale nie w tym miejscu, 2: jest na tym miejscu), przykładowo (01020): ")
		fmt.Scanln(&enteredWeights) 
		if len(enteredWeights) != 5 {
			fmt.Println("Waga musi mieć 5 znaków!")
			inputError = true
			continue
		}
		inputError = false
	}
	
	//get timestamp
	start := time.Now();

	//maping letters that are not allowed
	var forbiddenLetters []string

	for i, letter := range enteredWorld {
		if(enteredWeights[i] == '0') {
			//check if this letter is known at other position
			indexes := findLetterIndexes(enteredWorld, string(letter))
			delete := true
			for _, index := range indexes {
				if(index != i && enteredWeights[index] != '0') {
					delete = false
					break
				}
			}
			if(delete) {
				forbiddenLetters = append(forbiddenLetters, string(letter))
			}
		
		}
	}

	wordsLenght := len(words)
	//removing  words that contains forbidden letters
	for i:=0; i<wordsLenght; i++ {
		if(containsLetterAtAnyPosition(words[i], forbiddenLetters)) {
			words = append(words[:i], words[i+1:]...)
			i--;
			wordsLenght--
		}
	}
	//removind words that have forbidden letters at specyfic position
	for i:=0; i<wordsLenght; i++ {
		for j, letter := range enteredWorld {
			if enteredWeights[j]== '0' {
				if containsLetterAtSpecyficPosition(words[i], string(letter), j) == true {
				words = append(words[:i], words[i+1:]...)
				i--;
				wordsLenght--
				break
				}
		}
		}
	}

	//removing words that dont have specyfic letter at specific position
	for i:=0; i<wordsLenght; i++ {
		for j, letter := range enteredWorld {
			if enteredWeights[j]== '2' {
				if containsLetterAtSpecyficPosition(words[i], string(letter), j) == false {
				words = append(words[:i], words[i+1:]...)
				i--;
				wordsLenght--
				break
				}
			}
		}
	}

	//removing words that have letter in a word but not on that position
	for i:=0; i<wordsLenght; i++ {
		for j, letter := range enteredWorld {
			if enteredWeights[j]== '1' {
				if containsLetterAnywhereExceptSpecyficPosition(words[i], string(letter), j) == false {
				words = append(words[:i], words[i+1:]...)
				i--;
				wordsLenght--
				break
				}
			}
		}
	}
	end := time.Now();
	difference := end.Sub(start)
	
	if(wordsLenght == 0) {
		fmt.Printf("\n Nie mamy nic do zaproponowania!\n")
		return 
	}
	if wordsLenght == 1 {
		fmt.Printf("\n***********************************************\n")
		fmt.Printf("\nSłowo dnia to: %s\n", words[0])
		fmt.Printf("\n***********************************************\n")
		return 
	}
	fmt.Println("\nMożliwe słowa:")
	for i, word := range words {
		
		fmt.Println(i+1, word)
	}
	fmt.Printf("\nWyniki znaleziono w: %s\n", difference)
	fmt.Printf("\n\n Chcesz kontynuować (T/N): ")
	exit := "T"
	fmt.Scanln(&exit);
	if exit == string('n') || exit == string('N') {
		return;
	}
	}
	
}
