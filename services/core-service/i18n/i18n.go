package i18n

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	config "services.core-service/configs"
)

const (
	i18nMessage       = "messages"
	messageFolderName = "messages"
	defaultLanguage   = "en"
)

type I18n struct {
	Bundle       *i18n.Bundle
	MapLocalizer map[string]*i18n.Localizer
}

func NewI18n(c config.I18nConfig) (*I18n, error) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	mapLocalizer := make(map[string]*i18n.Localizer)
	for _, lang := range c.Languages {
		bundle.MustLoadMessageFile(fmt.Sprintf("./%s/%v.%v.json", messageFolderName, i18nMessage, lang))
		mapLocalizer[lang] = i18n.NewLocalizer(bundle, lang)
	}

	return &I18n{
		Bundle:       bundle,
		MapLocalizer: mapLocalizer,
	}, nil
}

func (r *I18n) MustLocalize(lang string, msgId string, templateData map[string]string) string {
	var localizerPtr *i18n.Localizer
	localizerPtr, ok := r.MapLocalizer[lang]
	if !ok {
		localizerPtr = r.MapLocalizer[defaultLanguage]
	}
	return localizerPtr.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    msgId,
		TemplateData: templateData,
	})
}
