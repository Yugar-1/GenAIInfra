// filepath: /home/ubuntu/GenAIInfra/kubernetes-addons/services-operator/helmtest.go
package main

import (
	"fmt"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/registry"
)

func main() {
	settings := cli.New()
	ociURL := "oci://ghcr.io/opea-project/charts/vllm"
	version := "1.1.0"
	releaseName := "myvllm"
	namespace := "default"

	registryClient, err := registry.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	chartDownloader := downloader.ChartDownloader{
		Out:              os.Stdout,
		RegistryClient:   registryClient,
		Getters:          getter.All(settings),
		RepositoryConfig: settings.RepositoryConfig,
		RepositoryCache:  settings.RepositoryCache,
	}
	destination := "/home/ubuntu/GenAIInfra/kubernetes-addons/services-operator/"
	chartPath, _, err := chartDownloader.DownloadTo(ociURL, version, destination)
	if err != nil {
		log.Fatalf("failed to download chart from OCI: %v\n", err)
	}
	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Fatalf("load failed: %v\n", err)
	}
	vals := map[string]interface{}{
		"image": map[string]interface{}{
			"repository": "opea/vllm",
		},
		"nodeSelector": map[string]interface{}{
			"kubernetes.io/hostname": "gnr-server05",
		},
		"global": map[string]interface{}{
			"HUGGINGFACEHUB_API_TOKEN": "", // need to config
			"HF_ENDPOINT":              "https://hf-mirror.com",
			"http_proxy":               "http://proxy-dmz.intel.com:912",
			"https_proxy":              "http://proxy-dmz.intel.com:912",
			"no_proxy":                 ".zxyqna.svc.cluster.local,.intel.com,10.0.0.0/8",
		},
	}
	fmt.Println("values: ", vals)
	fmt.Println("chartpath: ", chartPath)
	fmt.Printf("charts: %v/n", chart)

	actionConfig := new(action.Configuration)
	err = actionConfig.Init(settings.RESTClientGetter(), "default", "secret", log.Printf)
	if err != nil {
		log.Fatal("xxxxx", err)
	}
	install := action.NewInstall(actionConfig)
	install.ReleaseName = releaseName
	install.Namespace = namespace
	install.RepoURL = ociURL
	install.Version = version
	install.Timeout = 30e9
	install.CreateNamespace = true
	// install.Wait = true
	result, err := install.Run(chart, vals)
	if err != nil {
		log.Fatal("yyyyyyyyyyy", err)
	}
	log.Printf("Installed release: %s", result.Name)
}
