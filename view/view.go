package view

import (
	"net/http"
	"html/template"
	"golang.org/x/net/context"
)

type private struct{}
var reqKey private

func NewContext(ctx context.Context, tmpls *template.Template) context.Context {
	return context.WithValue(ctx, reqKey, tmpls)
}

func FromContext(ctx context.Context) *template.Template {
	return ctx.Value(reqKey).(*template.Template)
}

func Exec(ctx context.Context, w http.ResponseWriter, name string, data interface{}) error {
	return FromContext(ctx).ExecuteTemplate(w, name, data)
}