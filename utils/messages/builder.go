package messages

import (
	"bytes"
	"text/template"

	"github.com/gofiber/fiber/v2"
)

// CreateMsg is a helper function for creating message with context
func CreateMsg(ctx *fiber.Ctx, messageId string, templateData ...map[string]string) string {

	// Lookup template from package-level Templates (defined in messages.go)
	tplText, ok := Templates[messageId]
	if !ok {
		// fallback to returning the messageId directly
		return messageId
	}

	data := map[string]string{}
	if len(templateData) > 0 && templateData[0] != nil {
		data = templateData[0]
	}

	tpl, err := template.New("msg").Parse(tplText)
	if err != nil {
		return tplText
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return tplText
	}

	return buf.String()
}
