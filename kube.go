package main

import (
	"encoding/base64"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// CreateKubeClient -- Creates KubeClient
func (p Plugin) CreateKubeClient() (*kubernetes.Clientset, error) {
	ca, err := base64.StdEncoding.DecodeString(p.KubeConfig.Ca)
	config := clientcmdapi.NewConfig()
	config.Clusters["drone"] = &clientcmdapi.Cluster{
		Server:                   p.Config.Server,
		CertificateAuthorityData: ca,
	}
	config.AuthInfos["drone"] = &clientcmdapi.AuthInfo{
		Token: p.Config.Token,
	}
	config.Contexts["drone"] = &clientcmdapi.Context{
		Cluster:  "drone",
		AuthInfo: "drone",
	}
	//config.Clusters["drone"].CertificateAuthorityData = ca
	config.CurrentContext = "drone"
	clientBuilder := clientcmd.NewNonInteractiveClientConfig(*config, "drone", &clientcmd.ConfigOverrides{}, nil)
	actualCfg, err := clientBuilder.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}
	return kubernetes.NewForConfig(actualCfg)
}