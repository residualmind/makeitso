package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	flavFlag := flag.String("f", "senior", "flavor: To make the generated code match the existing code-base, pick a flavor from [ junior | senior | bigbrain ]")
	langFlag := flag.String("l", "", "language: The (programming) language the source is in. If no language is provided, a guess is made.")

	flag.CommandLine.Usage = func() {
		stderr("Usage:")
		stderr(fmt.Sprintf("\t%s [flags] FILE", os.Args[0]))
		stderr("\nAvailable flags:")
	}
	flag.Parse()

	token, ok := os.LookupEnv("OPENAI_KEY")
	if !ok || token == "" {
		fatal("No OpenAI API key provided. Please set it via the environment variable OPENAI_KEY")
	}

	if len(flag.Args()) != 1 {
		picard()
		flag.CommandLine.Usage()
		flag.PrintDefaults()
		os.Exit(1)
	}

	fileName := flag.Arg(0)
	buf, err := os.ReadFile(fileName)
	if err != nil {
		fatal(err)
	}

	file := string(buf)
	lang := *langFlag
	if lang == "" {
		lang = "a programming language you need to determine from the given file before following other instructions."
	}

	outFile := os.Stdout

	theMegaBrain := &megaBrain{
		client: openai.NewClient(token),
		flavor: flavorType(*flavFlag),
	}

	var output string
	output, err = theMegaBrain.thinkRealHard(theMegaBrain.engineerPrompt(basePromptPass1, lang, file))
	if err != nil {
		fatal(err)
	}

	output, err = theMegaBrain.thinkRealHard(theMegaBrain.engineerPrompt(basePromptPass2, lang, output))
	if err != nil {
		fatal(err)
	}

	// i know how to use regex. go away
	output = strings.ReplaceAll(output, "/* "+begin+" */", "")
	output = strings.ReplaceAll(output, "/* "+end+" */", "")
	output = strings.ReplaceAll(output, "// "+begin, "")
	output = strings.ReplaceAll(output, "// "+end, "")
	output = strings.ReplaceAll(output, "/*"+begin+"*/", "")
	output = strings.ReplaceAll(output, "/*"+end+"*/", "")
	output = strings.ReplaceAll(output, "//"+begin, "")
	output = strings.ReplaceAll(output, "//"+end, "")

	_, _ = outFile.WriteString(output + "\n")
}

func picard() {
	stderr(strings.Join([]string{
		"",
		"  .-'---`-.",
		",'          `.",
		"|             \\",
		"|              \\",
		"\\           _  \\",
		",\\  _    ,'-,/-)\\",
		"( * \\ \\,' ,' ,'-)",
		" `._,)     -',-')",
		"   \\/         ''/",
		"    )        / /",
		"   /       ,'-'",
		"", "The computer's attempt to execute the designated task has regrettably failed.", "",
	}, "\n"))
}

func fatal(err any) {
	picard()
	stderr(fmt.Sprintf("%v", err))
	os.Exit(1)
}
func stderr(s string) {
	_, _ = os.Stderr.WriteString(s + "\n")
}
