package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

var morsemap = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
	".": ".-.-.-",
	",": "--..--",
	"?": "..--..",
	"!": "-.-.--",
	"-": "-....-",
	"/": "-..-.",
	"@": ".--.-.",
	"(": "-.--.",
	")": "-.--.-",
	" ": "/",
}

func encodeToMorse(message, letterSplitter string) string {
	var output string
	message = strings.ToUpper(message)
	words := strings.Split(message, " ")
	for i, word := range words {
		word := encodeWord(word, letterSplitter)
		if word != "" {
			if i > 0 {
				output += letterSplitter + "/" + letterSplitter
			}
			output += word
		}
	}
	return output
}

func encodeWord(word, letterSplitter string) string {
	var morse string
	for i := 0; i < len(word); i++ {
		code := morsemap[word[i:i+1]]
		if code != "" {
			morse += code + letterSplitter
		}
	}

	if morse != "" {
		morse = morse[:len(morse)-1]
	}
	return morse
}

func encode(ctx *fiber.Ctx) error {
	str := ctx.Query("text")
	if str == "" {
		return ctx.SendStatus(fiber.StatusOK)
	}
	return ctx.SendString(encodeToMorse(str, " "))
}

func setHandlers(app *fiber.App) {
	app.Get("/encode", encode)
}

func main() {
	app := fiber.New()
	setHandlers(app)
	app.Listen(":8080")
}
