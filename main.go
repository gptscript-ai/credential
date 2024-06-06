package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strconv"

	"github.com/gptscript-ai/go-gptscript"
	"github.com/tidwall/gjson"
)

type input struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
	Env     string `json:"env"`
}

type sysPromptInput struct {
	Message   string `json:"message"`
	Fields    string `json:"fields"`
	Sensitive string `json:"sensitive"`
}

func main() {
	// Set up signal handler
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if len(os.Args) != 2 {
		fmt.Println("Missing input string")
		os.Exit(1)
	}

	inputStr := os.Args[1]

	var in input
	if err := json.Unmarshal([]byte(inputStr), &in); err != nil {
		fmt.Println("Error parsing input JSON:", err)
		os.Exit(1)
	}

	// Check for the "sensitive" value in the JSON input.
	// We don't unmarshal this value because we want the default to be true instead of false.
	// So we just check for it here manually instead, and only flip it to false if the user provided
	// exactly the string "false" or the boolean false.
	sensitive := true
	sensitiveVal := gjson.Get(inputStr, "sensitive")
	if sensitiveVal.Type == gjson.False || (sensitiveVal.Type == gjson.String && sensitiveVal.String() == "false") {
		sensitive = false
	}

	client, err := gptscript.NewGPTScript()
	if err != nil {
		fmt.Println("Error creating GPTScript client:", err)
		os.Exit(1)
	}
	defer client.Close()

	sysPromptIn, err := json.Marshal(sysPromptInput{
		Message:   in.Message,
		Fields:    in.Field,
		Sensitive: strconv.FormatBool(sensitive),
	})
	if err != nil {
		fmt.Println("Error marshalling sys prompt input:", err)
		os.Exit(1)
	}

	run, err := client.Run(ctx, "sys.prompt", gptscript.Options{
		Input: string(sysPromptIn),
	})
	if err != nil {
		fmt.Println("Error running GPTScript:", err)
		os.Exit(1)
	}

	res, err := run.Text()
	if err != nil {
		fmt.Println("Error getting GPTScript response:", err)
		os.Exit(1)
	}

	k := gjson.Get(res, in.Field).String()
	fmt.Printf(`{"env": {"%s": "%s"}}`, in.Env, k)
}
