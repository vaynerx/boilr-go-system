// Package {{ Project }}
// Copyright (c) {{ Year }} {{ if Owner }}{{ Owner }}{{ end }}
// {{ Title }}

package main

import (
	context "context"

	app "{{ if Owner }}{{ Owner }}{{ end }}.{{ Project }}/internal/app"

	logx "github.com/ory/x/logrusx"
	log "github.com/sirupsen/logrus"
)

func init() {
	logx.New()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	var err error

	app := app.NewApp(ctx)
	if err = app.InitApp(ctx); err != nil {
		return err
	}

	return nil
}
