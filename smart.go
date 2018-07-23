package main

import (
    "bufio"
    "fmt"
    "os"
    "math/rand"
    "io/ioutil"
    "strings"
)
const FOLDER_NAME string = "data/"
const FILE_NAME1 string = FOLDER_NAME+"train.txt"
const FILE_NAME2 string = FOLDER_NAME+"facts.txt"
const FILE_NAME3 string = FOLDER_NAME+"cooking.txt"
const FILE_NAME4 string = FOLDER_NAME+"misc.txt"
const FILE_NAME5 string = FOLDER_NAME+"clown.txt"
const FILE_NAME6 string = FOLDER_NAME+"random.txt"
const FILE_NAME7 string = FOLDER_NAME+"funny.txt"
const RESP_LENGTH int = 10

//Structure to hold the words
type data struct {
	word [2]string
	nextWords []string
}

//Get next word from structure
func (dataS data) next(w []data) string {
	l := len(dataS.nextWords)
	if l == 0 {
		j := 0
		p := len(w)
		var current data
		for j<1 {
			toPick := rand.Intn(p)
			current = w[toPick]
			if len(current.word[1])>0 {
				if (current.word[1])[0]<90 {
					break
				}
			}
		}
		return current.word[1]
	}
	toPick := rand.Intn(l)
	return dataS.nextWords[toPick]
}

//Add word to repository
func (dataS data) addWord(toAdd string) data{
	dataS.nextWords = append(dataS.nextWords, toAdd)
	return dataS
}

//Default error checker
func check(e error) {
    if e != nil {
        panic(e)
    }
}

//Initialize conversational data from file
func initData(w []data, filename string) []data{
	dat, err := ioutil.ReadFile(filename)
    check(err)
    toParse := string(dat)
    toParse = strings.Replace(toParse,"---","",-1)
    ret := learnData(toParse,w)
    return ret
}

//Parse string to add to data
func learnData(toParse string, w []data) []data{
	allWords := strings.Fields(toParse)
	for i, aW := range allWords {
		//Check to ensure that data exists for word
		var t []string
		var a [2]string
		current := data{word: a, nextWords: t}
		bmark := false
		for _, d := range w {
			n := ""
			if i>0 {
				n = allWords[i-1]
			}
			if d.word[1] == aW && d.word[0] == n {
				current = d //Since it takes latest appended entry, stale entries ignored
				bmark = true
			}
		}
		//Add word to wordlist if it doesn't exist
		if !bmark {
			if i>0 {
				var b [2]string
				b[0] = allWords[i-1]
				b[1] = aW
				current = data{word: b, nextWords: t}
			}
			if i==0 {
				a[1] = aW
				current = data{word: a, nextWords: t}
			}
			w = append(w, current)
		}
		//Set the next word for data
		if i<len(allWords)-1 {
			current = current.addWord(allWords[i+1])
			w = append(w, current)
		}
	}
	return w
}

//Structure the response based on data
func respond(w []data) string{
	var t []string
	var e [2]string
	e[1] = "a"
	current := data{word: e, nextWords: t}
	l := len(w)
	if l==0 {
		return ""
	}
	j := 0
	for j<1 {
		toPick := rand.Intn(l)
		current = w[toPick]
		if len(current.word[1])>0 {
			if (current.word[1])[0]<90 {
				break
			}
		}
	}
	response := ""
	for i:=0;i<RESP_LENGTH;i++ {
		response = response + current.word[1]
		if current.word[1] != "" {
			response = response + " "		}
		nextWord := current.next(w)
		rep := current //creates replica
		for _, d := range w {
			if d.word[1] == nextWord {
				current = d
			}
		}
		for _, d := range w {
			if rep.word[1] == d.word[0] && nextWord == d.word[1]{
				current = d
			}
		}
	}
	response = response + current.word[1]
	if current.word[1] != "" {
		response = response + " "
	}
	if !strings.Contains(current.word[1],".") {
		for j<1 {
			nextWord := current.next(w)
			rep := current //creates replica
			for _, d := range w {
				if d.word[1] == nextWord {
					current = d
				}
			}
			for _, d := range w {
				if rep.word[1] == d.word[0] && nextWord == d.word[1]{
					current = d
				}
			}
			response = response + current.word[1]
			if strings.Contains(current.word[1],".") {
				break
			}
			if current.word[1] != "" {
				response = response + " "
			}
		}
	}
	return response
}

//Main function of the program
func main() {
	wordList := make([]data, 1)
	var t []string
	var e [2]string
	temp := data{word: e, nextWords: t}
	wordList[0] = temp
	wordList = initData(wordList,FILE_NAME1)
	//Endlessly loop for user input
	x := 0
	for x<1 {
		reader := bufio.NewReader(os.Stdin)
    	fmt.Print("You: ")
    	text, _ := reader.ReadString('\n')
    	if text == "x"+"\n" {
    		break
    	}
    	if text == "tell me a poem"+"\n" {
    		wordList = initData(wordList,FILE_NAME2)
    	}
    	if text == "teach me to cook"+"\n" {
    		wordList = initData(wordList,FILE_NAME3)
    	}
    	if text == "be more interesting"+"\n" {
    		wordList = initData(wordList,FILE_NAME4)
    	}
    	if text == "you scare me"+"\n" {
    		wordList = initData(wordList,FILE_NAME5)
    	}
    	if text == "teach me to fish"+"\n" {
    		wordList = initData(wordList,FILE_NAME6)
    	}
    	if text == "tell me a joke"+"\n" {
    		wordList = initData(wordList,FILE_NAME7)
    	}
    	if text == "forget everything"+"\n" {
    		wordList = make([]data, 1)
    		var t []string
			var e [2]string
			temp = data{word: e, nextWords: t}
			wordList[0] = temp
			wordList = initData(wordList,FILE_NAME1)
    	}
    	wordList = learnData(text,wordList)
    	fmt.Println("\n"+"Friend: "+respond(wordList)+"\n")
	}
}
