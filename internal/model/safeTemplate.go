package model

import (
	"bytes"
	"html/template"
	"strings"
	"sync"

	"golang.org/x/text/cases"
)

type SafeTemplate struct {
	*template.Template
	current string
	mu      sync.RWMutex
}

var (
	templateName = "prom"
	defaultFuncs = map[string]interface{}{
		"toUpper": strings.ToUpper,
		"toLower": strings.ToLower,
		"title":   cases.Title,
		//"markdown": markdownEscapeString,
	}
	isMarkdownSpecial [128]bool
)

func init() {
	for _, c := range "_*`" {
		isMarkdownSpecial[c] = true
	}
}

func (t *SafeTemplate) UpdateTemplate(newTpl string) (err error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	tpl, err := template.New(templateName).
		Funcs(defaultFuncs).
		Option("missingkey=zero").
		Parse(newTpl)
	if err != nil {
		return
	}

	_ = t.current // old template
	t.Template = tpl
	t.current = newTpl
	return
}

func (t *SafeTemplate) Clone() (*template.Template, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.Template.Clone()
}

func markdownEscapeString(s string) string {
	b := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(b)

	for _, c := range s {
		if c < 128 && isMarkdownSpecial[c] {
			buf.WriteByte('\\')
		}
		buf.WriteRune(c)
	}
	return buf.String()
}
