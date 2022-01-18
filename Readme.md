## Code Generation
```
execDir=~/go/pkg/mod/k8s.io/code-generator@v0.23.1

$execDir/generate-groups.sh all github.com/apaarshrm39/Kluster/pkg/generated github.com/apaarshrm39/Kluster/pkg/apis "apaarshrm.dev:v1alpha1" --output-base "$(dirname "${BASH_SOURCE[0]}")/../../.."  --go-header-file ./hack/boilerplate.go.txt

```

## Controller-gen to create CRDs 
```
binary=~/go/bin

binary/controller-gen paths=github.com/apaarshrm39/Kluster/pkg/apis/apaarshrm.dev/v1alpha1 crd:trivialVersions=true output:crd:artifacts:config=manifests
```
