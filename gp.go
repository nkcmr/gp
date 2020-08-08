package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

type app struct {
}

func (a *app) oauthFlowBegin() error {
	fmt.Println("signing in with google...")
	return nil
}

func gpComplete(a *app) prompt.Completer {
	return func(d prompt.Document) []prompt.Suggest {
		s := []prompt.Suggest{
			{Text: "help", Description: "get some help using this tool"},
			{Text: "signin", Description: "authorize this tool to access your google photos library"},
			{Text: "whoami", Description: "get details of if you are logged in or not"}
			{Text: "quit", Description: "exit gp"},
		}
		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}
}

func gpExec(a *app) prompt.Executor {
	return func(input string) {
		switch ip := strings.ToLower(strings.TrimSpace(input)); ip {
		case "signin":
			if err := a.oauthFlowBegin(); err != nil {
				fmt.Printf("error: failed to sign in with google: %s", err.Error())
			}
		case "quit", "exit":
			os.Exit(0)
		default:
			fmt.Printf("error: unknown command: %s\n", ip)
		}
	}
}

func main() {
	a := &app{}
	p := prompt.New(
		gpExec(a),
		gpComplete(a),
		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(_ *prompt.Buffer) {
				fmt.Println("bye!")
				os.Exit(0)
			},
		}),
	)
	p.Run()
}
