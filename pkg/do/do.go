package do

import (
	"context"
	"fmt"
	"os"

	kluster "github.com/apaarshrm39/Kluster/pkg/apis/apaarshrm.dev/v1alpha1"
	"github.com/digitalocean/godo"
)

func auth(token string) *godo.Client {
	fmt.Println("Auth DO is called")
	client := godo.NewFromToken(token)
	fmt.Println("DO cclient", client)

	return client
}

func Create(token string, spec kluster.KlusterSpec) (string, error) {
	client := auth(token)
	ctx := context.TODO()

	createRequest := &godo.KubernetesClusterCreateRequest{
		Name:        spec.Name,
		RegionSlug:  spec.Region,
		VersionSlug: spec.Version,
		NodePools: []*godo.KubernetesNodePoolCreateRequest{
			&godo.KubernetesNodePoolCreateRequest{
				Name:  spec.Nodepool[0].Name,
				Size:  spec.Nodepool[0].Size,
				Count: spec.Nodepool[0].Count,
			},
		},
	}

	fmt.Fprintln(os.Stdout, createRequest)
	cluster, _, err := client.Kubernetes.Create(ctx, createRequest)
	if err != nil {
		return "", err
	}

	return cluster.ID, nil

}

func GetStatus(id string, token string) string {
	client := auth(token)
	ctx := context.TODO()

	cluster, _, err := client.Kubernetes.Get(ctx, id)
	if err != nil {
		return "could not get the status from DO"
	}

	return string(cluster.Status.State)

}
