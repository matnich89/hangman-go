package util

import "math/rand"

/*
 The last one is the fear of long words
 How cruel...
*/
var words = []string{"keyboard", "cat", "rainbow", "computer", "hippopotomonstrosesquippedaliophobia"}

func GenerateWord() string {
	randomIndex := rand.Intn(len(words))
	word := words[randomIndex]
	return word
}
