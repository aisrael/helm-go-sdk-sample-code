package main

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

func main() {
	settings := cli.New()

	actionConfig := new(action.Configuration)
	// You can pass an empty string instead of settings.Namespace() to list
	// all namespaces
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}

	chart, err := loader.LoadDir("helm/charts/nginx")
	if err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}

	install := action.NewInstall(actionConfig)
	install.ReleaseName = "nginx"
	vals := make(map[string]interface{})
	release, err := install.Run(chart, vals)
	if err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}

	log.Printf("release.Name: %s\n", release.Name)
	log.Printf("release.Version: %d\n", release.Version)

	list := action.NewList(actionConfig)
	// Only list deployed
	list.Deployed = true
	results, err := list.Run()
	if err != nil {
		log.Fatalf("%+v", err)
		os.Exit(1)
	}

	for _, rel := range results {
		log.Printf("%+v", rel)
	}
}
