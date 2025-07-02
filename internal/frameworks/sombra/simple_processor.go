package sombra

import (
	"bytes"
	"github.com/sombrahq/sombra-cli/internal/core/entities"
	"github.com/sombrahq/sombra-cli/internal/core/usecases"
	"regexp"
	"strings"
)

type Processor struct {
}

func (l *Processor) ProcessString(target string, mapping entities.MapList) string {
	var old, replacement string
	for _, item := range mapping {
		old = item.Key
		replacement = item.Value
		target = l.stringReplace(target, old, replacement)
	}
	return target
}

func (l *Processor) ProcessContent(content []byte, mapping entities.MapList) []byte {
	var old, replacement string
	for _, item := range mapping {
		old = item.Key
		replacement = item.Value
		content = l.bytesReplace(content, old, replacement)
	}
	return content
}

func (l *Processor) bytesReplace(content []byte, old string, replacement string) []byte {
	parts := strings.SplitN(old, ":", 2)
	if len(parts) > 0 && parts[0] == "re" {
		prev := regexp.MustCompile(parts[1])
		return prev.ReplaceAll(content, ([]byte)(replacement))
	}
	// looks like `old` is not a regex, let's do a string replace then
	return bytes.ReplaceAll(content, ([]byte)(old), ([]byte)(replacement))
}

func (l *Processor) stringReplace(content string, old string, replacement string) string {
	parts := strings.SplitN(old, ":", 2)
	if len(parts) > 0 && parts[0] == "re" {
		prev := regexp.MustCompile(parts[1])
		return prev.ReplaceAllString(content, replacement)
	}
	// looks like `old` is not a regex, let's do a string replace then
	return strings.ReplaceAll(content, old, replacement)
}

func NewProcessor() *Processor {
	return &Processor{}
}

var _ usecases.SombraStringsPort = (*Processor)(nil)
