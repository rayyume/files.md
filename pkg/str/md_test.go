package str

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/require"
)

func TestMultilineTextWithMarkdown(t *testing.T) {
	r := require.New(t)
	text := "header\nSome text with two italic paragraphs\n\nAlso italic\n\nheader2\nitalic\ncode\n\nheader3\njust text"

	var messageEntities = []tgbotapi.MessageEntity{
		{Type: "italic", Offset: 7, Length: 51},
		{Type: "bold", Offset: 58, Length: 8},
		{Type: "italic", Offset: 66, Length: 7},
		{Type: "code", Offset: 73, Length: 6},
	}

	markdown := EntitiesToMarkdown(text, messageEntities)
	expectedMarkdown := "header\n_Some text with two italic paragraphs\n\nAlso italic\n\n_*header2\n*_italic\n_`code\n\n`header3\njust text"
	r.Equal(expectedMarkdown, markdown)
}

func TestSpacedItalic(t *testing.T) {
	r := require.New(t)
	text := "Header\nLeverage one Minute Praising instead"

	var messageEntities = []tgbotapi.MessageEntity{
		{Type: "italic", Offset: 16, Length: 20},
	}

	markdown := EntitiesToMarkdown(text, messageEntities)
	expectedMarkdown := "Header\nLeverage _one Minute Praising _instead"
	r.Equal(expectedMarkdown, markdown)
}
