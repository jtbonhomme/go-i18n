package main

import (
    "embed"
    "fmt"
    "github.com/BurntSushi/toml"
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "golang.org/x/text/language"

)

//go:embed active.*.toml
var LocaleFS embed.FS

func main() {
    fmt.Printf("test:\n")

    bundle := i18n.NewBundle(language.English)

    bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
    bundle.LoadMessageFileFS(LocaleFS, "active.es.toml")
    bundle.LoadMessageFileFS(LocaleFS, "active.en.toml")
    bundle.LoadMessageFileFS(LocaleFS, "active.fr.toml")

    name:="Bob"
    unreadEmailCount:=2

    lang := "fr"
    accept := ""
    localizer := i18n.NewLocalizer(bundle, lang, accept)
    msg, err := localizer.Localize(&i18n.LocalizeConfig{
        DefaultMessage: &i18n.Message{
                ID:          "PersonUnreadEmails",
                Description: "The number of unread emails a person has",
                One:         "{{.Name}} has {{.UnreadEmailCount}} unread email.",
                Other:       "{{.Name}} has {{.UnreadEmailCount}} unread emails.",
            },
            PluralCount: unreadEmailCount,
            TemplateData: map[string]interface{}{
                "Name":             name,
                "UnreadEmailCount": unreadEmailCount,
            },
    })

    if err != nil {
        fmt.Printf("error: %w\n", err)
    }
    fmt.Printf("%s\n", msg)
}
