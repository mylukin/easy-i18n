package i18n

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

// Generate catalog.go
func Generate(id, pkgName string, paths []string, outFile string) error {
	if len(paths) == 0 {
		paths = []string{"."}
	}

	goFile, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]*Message{}
	for _, path := range paths {
		if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			messages, err := unmarshal(id, path)
			if err != nil {
				return err
			}

			lang := info.Name()[0 : len(info.Name())-5]
			data[lang] = messages
			fmt.Printf("Generate %+v ...\n", path)

			return nil
		}); err != nil {
			return err
		}
	}
	err = os.MkdirAll(path.Dir(outFile), os.ModePerm)
	if err != nil {
		return err
	}
	err = i18nTmpl.Execute(goFile, struct {
		Data      map[string]*Message
		BackQuote string
		Package   string
	}{
		data,
		"`",
		pkgName,
	})

	return err
}

var funcs = template.FuncMap{
	"funcName": func(lang string) string {
		lang = strings.ReplaceAll(lang, "_", "")
		lang = strings.ReplaceAll(lang, "-", "")
		lang = strings.ToUpper(lang[:1]) + lang[1:]
		return lang
	},
	"quote": func(text string) string {
		return strconv.Quote(text)
	},
}

var i18nTmpl = template.Must(template.New("i18n").Funcs(funcs).Parse(`package {{.Package}}

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// init
func init() {
	{{- range $k, $v := .Data }}
	init{{ funcName $k }}(language.Make("{{ $k }}"))
	{{- end }}
}

{{- range $k, $v := .Data }}
// init{{ funcName $k }} will init {{ $k }} support.
func init{{ funcName $k }}(tag language.Tag) {
	{{- range $k, $v := $v }}
	message.SetString(tag, {{quote $k}}, {{quote $v}})
	{{- end }}
}
{{- end }}
`))
