package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type megaBrain struct {
	client *openai.Client
	flavor flavorType
}

type flavorType string

const (
	junior   flavorType = "junior"
	senior   flavorType = "senior"
	bigbrain flavorType = "bigbrain"
)

func (mb *megaBrain) engineerPrompt(base []string, lang string, file string) string {
	flavor, ok := flavors[mb.flavor]
	if !ok {
		fatal("don't like this flavor")
	}

	p := strings.Join(base, " ")
	p = strings.ReplaceAll(p, "{{LANG}}", lang)
	p = strings.ReplaceAll(p, "{{FLAVOR}}", flavor)
	p = strings.ReplaceAll(p, "{{FILE}}", file)

	return p
}

func (mb *megaBrain) thinkRealHard(prompt string) (string, error) {
	for tries := 10; tries > 0; tries-- {
		res, err := mb.think(prompt)
		if err != nil {
			continue
		}

		return res, nil
	}

	return "", errors.New("max gpt tries reached, giving up")
}
func (mb *megaBrain) think(prompt string) (string, error) {
	resp, err := mb.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: prompt,
				},
				// {
				// 	Role:    openai.ChatMessageRoleAssistant,
				// 	Content: `{"output": "`,
				// },
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("GTP error: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("no messages returned from GPT")
	}

	var out struct {
		Output string `json:"output"`
	}

	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &out); err != nil {
		return "", err
	}

	return out.Output, nil
}
