package main

//go:generate easyi18n extract . ./locales/en.json
//go:generate easyi18n update ./locales/en.json ./locales/zh-Hans.json
//go:generate easyi18n generate --pkg=catalog ./locales ./catalog/catalog.go
//go:generate go build -o example

import (
	"fmt"
	"os"

	_ "github.com/rain-31/easy-i18n/example/catalog"
	"github.com/rain-31/easy-i18n/i18n"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/text/language"
)

func main() {
	sessionId := uuid.NewV4().String()
	i18n.RegistPrinter(sessionId, language.SimplifiedChinese)

	i18n.Printf(sessionId, `hello world!`)
	fmt.Println()

	name := `Lukin`

	i18n.Printf(sessionId, `hello %s!`, name)
	fmt.Println()

	i18n.Printf(sessionId, `%s has %d cat.`, name, 1)
	fmt.Println()

	i18n.Printf(sessionId, `%s has %d cat.`, name, 2, i18n.Plural(
		`%[2]d=1`, `%s has %d cat.`,
		`%[2]d>1`, `%s has %d cats.`,
	))
	fmt.Println()

	i18n.Fprintf(sessionId, os.Stderr, `%s have %d apple.`, name, 2, i18n.Plural(
		`%[2]d=1`, `%s have an apple.`,
		`%[2]d=2`, `%s have two apples.`,
		`%[2]d>2`, `%s have %d apples.`,
	))
	fmt.Println()
}
