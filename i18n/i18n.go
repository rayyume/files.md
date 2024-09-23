package i18n

// lang            *i18n.Bundle
var emojisByKeyword map[string]string

// LoadLangFile only supports single language for now
func LoadLangFile(path string) error {
	//lang = i18n.NewBundle(language.English)
	//_, err := lang.LoadMessageFile(path)
	//if err != nil {
	//	return fmt.Errorf("i18n.Load: %w", err)
	//}

	return nil
}

func Tr(str string) string {
	return str
}
