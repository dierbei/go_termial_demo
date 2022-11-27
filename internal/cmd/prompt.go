package cmd

import (
	"fmt"
	prompt "github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

func init() {
	RootCmd.AddCommand(promptCmd)
}

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "进入交互终端",
	RunE: func(cmd *cobra.Command, args []string) error {
		p := prompt.New(
			executorCmd(cmd),
			completer,
			prompt.OptionPrefix(">>> "),
		)
		p.Run()
		return nil
	},
}

func executorCmd(cmd *cobra.Command) func(in string) {
	return func(in string) {
		in = strings.TrimSpace(in)
		blocks := strings.Split(in, " ")
		//log.Println(blocks)
		//if err := cmd.ParseFlags(blocks); err != nil {
		//	log.Fatalln(err)
		//}
		//cmd.ParseFlags(blocks[1:])
		switch blocks[0] {
		case "exit":
			fmt.Println("Bye!")
			os.Exit(0)
		case "list":
			if err := podCmd.RunE(cmd, blocks); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

var suggestions = []prompt.Suggest{
	// Command
	{"test", "this is test"},
	{"exit", "Exit prompt"},
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}
