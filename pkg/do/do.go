package do

import (
	"context"
	"fmt"
	"os"

	kluster "github.com/apaarshrm39/Kluster/pkg/apis/apaarshrm.dev/v1alpha1"
	"github.com/digitalocean/godo"
)

func Create(token string, spec kluster.KlusterSpec) (string, error) {
	fmt.Println("Create DO cluster is called")
	client := godo.NewFromToken(token)
	ctx := context.TODO()
	fmt.Println("DO cclient", client)

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
