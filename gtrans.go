package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

var apiKey string

func main() {
	loop := true
	apiKey = getInput("Enter API Server Key:")
	tgtLang := getInput("Enter Language code to translate to (eg: nb for Norwegian Bokmål, en for english):")

	for loop {
		text := getInput("Text to translate (type quit to exit):")
		if text == "quit" {
			loop = false
		} else {
			translated, err := translateText(tgtLang, text)
			if err != nil {
				log.Fatalf("Wrong language code, exiting\nError: %v\n", err)
			}
			fmt.Println(translated)
		}
	}
}

func getInput(in string) string {
	fmt.Println(in)
	reader := bufio.NewReader(os.Stdin)
	out, _ := reader.ReadString('\n')
	out = strings.Replace(out, "\n", "", -1)
	out = strings.Replace(out, "\r", "", -1)
	return out
}

func translateText(tgtLang, text string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(tgtLang)
	if err != nil {
		return "", err
	}

	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", err
	}
	return resp[0].Text, nil
}
