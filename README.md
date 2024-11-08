# Analyze Korifi Resources

This tool expects a valid kubeconfig in $HOME/.kube/config pointing to a Korifi cluster and a prepared `cf` CLI. After each step, all resources within namespaces prefixed with `cf` will be printed out. These steps are performed:

1. Empty cluster
1. Create org via CLI
1. Create space via CLI
1. Push [sample](./sample/) app 1
1. Push [sample](./sample/) app 2
1. Delete org

## Results

### Single Cluster

```
Empty cluster ...
Namespace: cf
core/v1/configmaps: 1
core/v1/secrets: 1
core/v1/serviceaccounts: 3
rbac.authorization.k8s.io/v1/rolebindings: 3
rbac.authorization.k8s.io/v1/roles: 1
korifi.cloudfoundry.org/v1alpha1/cfdomains: 1

Create organization ...
Namespace: cf
core/v1/configmaps: 1
core/v1/secrets: 1
core/v1/serviceaccounts: 3
rbac.authorization.k8s.io/v1/rolebindings: 3
rbac.authorization.k8s.io/v1/roles: 1
coordination.k8s.io/v1/leases: 1
korifi.cloudfoundry.org/v1alpha1/cforgs: 1
korifi.cloudfoundry.org/v1alpha1/cfdomains: 1
Namespace: cf-org-5ac2c439-f468-4e80-ac52-9a59394fcada
core/v1/configmaps: 1
core/v1/secrets: 1
core/v1/serviceaccounts: 1
rbac.authorization.k8s.io/v1/rolebindings: 1

Create space ...
Namespace: cf
core/v1/configmaps: 1
core/v1/secrets: 1
core/v1/serviceaccounts: 3
rbac.authorization.k8s.io/v1/rolebindings: 3
rbac.authorization.k8s.io/v1/roles: 1
coordination.k8s.io/v1/leases: 1
korifi.cloudfoundry.org/v1alpha1/cforgs: 1
korifi.cloudfoundry.org/v1alpha1/cfdomains: 1
Namespace: cf-org-5ac2c439-f468-4e80-ac52-9a59394fcada
core/v1/configmaps: 1
core/v1/secrets: 1
core/v1/serviceaccounts: 1
rbac.authorization.k8s.io/v1/rolebindings: 2
coordination.k8s.io/v1/leases: 1
korifi.cloudfoundry.org/v1alpha1/cfspaces: 1
Namespace: cf-space-3e5366d2-423f-44ab-8fa9-566225a59c2c
core/v1/configmaps: 1
core/v1/secrets: 1
core/v1/serviceaccounts: 3
rbac.authorization.k8s.io/v1/rolebindings: 3

Push app 1 ...
Namespace: cf
core/v1/configmaps: 1
core/v1/secrets: 1
core/v1/serviceaccounts: 3
rbac.authorization.k8s.io/v1/rolebindings: 3
rbac.authorization.k8s.io/v1/roles: 1
coordination.k8s.io/v1/leases: 2
korifi.cloudfoundry.org/v1alpha1/cforgs: 1
korifi.cloudfoundry.org/v1alpha1/cfdomains: 1
Namespace: cf-org-5ac2c439-f468-4e80-ac52-9a59394fcada
core/v1/configmaps: 1
core/v1/secrets: 1
core/v1/serviceaccounts: 1
rbac.authorization.k8s.io/v1/rolebindings: 2
coordination.k8s.io/v1/leases: 1
korifi.cloudfoundry.org/v1alpha1/cfspaces: 1
Namespace: cf-space-3e5366d2-423f-44ab-8fa9-566225a59c2c
core/v1/configmaps: 1
core/v1/endpoints: 1
core/v1/events: 29
core/v1/persistentvolumeclaims: 1
core/v1/pods: 2
core/v1/secrets: 4
core/v1/serviceaccounts: 4
core/v1/services: 1
apps/v1/controllerrevisions: 1
apps/v1/statefulsets: 1
events.k8s.io/v1/events: 29
rbac.authorization.k8s.io/v1/rolebindings: 3
coordination.k8s.io/v1/leases: 1
discovery.k8s.io/v1/endpointslices: 1
gateway.networking.k8s.io/v1/httproutes: 1
gateway.networking.k8s.io/v1beta1/httproutes: 1
korifi.cloudfoundry.org/v1alpha1/cfapps: 1
korifi.cloudfoundry.org/v1alpha1/cfroutes: 1
korifi.cloudfoundry.org/v1alpha1/cfprocesses: 1
korifi.cloudfoundry.org/v1alpha1/appworkloads: 1
korifi.cloudfoundry.org/v1alpha1/buildworkloads: 1
korifi.cloudfoundry.org/v1alpha1/cfbuilds: 1
korifi.cloudfoundry.org/v1alpha1/cfpackages: 1
kpack.io/v1alpha2/images: 1
kpack.io/v1alpha2/builds: 1
kpack.io/v1alpha2/sourceresolvers: 1
kpack.io/v1alpha1/images: 1
kpack.io/v1alpha1/builds: 1
kpack.io/v1alpha1/sourceresolvers: 1
metrics.k8s.io/v1beta1/pods: 1

Push app 2 ...
Namespace: cf
core/v1/configmaps: 1
core/v1/secrets: 1
core/v1/serviceaccounts: 3
rbac.authorization.k8s.io/v1/rolebindings: 3
rbac.authorization.k8s.io/v1/roles: 1
coordination.k8s.io/v1/leases: 3
korifi.cloudfoundry.org/v1alpha1/cforgs: 1
korifi.cloudfoundry.org/v1alpha1/cfdomains: 1
Namespace: cf-org-5ac2c439-f468-4e80-ac52-9a59394fcada
core/v1/configmaps: 1
core/v1/secrets: 1
core/v1/serviceaccounts: 1
rbac.authorization.k8s.io/v1/rolebindings: 2
coordination.k8s.io/v1/leases: 1
korifi.cloudfoundry.org/v1alpha1/cfspaces: 1
Namespace: cf-space-3e5366d2-423f-44ab-8fa9-566225a59c2c
core/v1/configmaps: 1
core/v1/endpoints: 2
core/v1/events: 61
core/v1/persistentvolumeclaims: 2
core/v1/pods: 4
core/v1/secrets: 7
core/v1/serviceaccounts: 4
core/v1/services: 2
apps/v1/controllerrevisions: 2
apps/v1/statefulsets: 2
events.k8s.io/v1/events: 61
rbac.authorization.k8s.io/v1/rolebindings: 3
coordination.k8s.io/v1/leases: 2
discovery.k8s.io/v1/endpointslices: 2
gateway.networking.k8s.io/v1/httproutes: 2
gateway.networking.k8s.io/v1beta1/httproutes: 2
korifi.cloudfoundry.org/v1alpha1/cfapps: 2
korifi.cloudfoundry.org/v1alpha1/cfroutes: 2
korifi.cloudfoundry.org/v1alpha1/cfprocesses: 2
korifi.cloudfoundry.org/v1alpha1/appworkloads: 2
korifi.cloudfoundry.org/v1alpha1/buildworkloads: 2
korifi.cloudfoundry.org/v1alpha1/cfbuilds: 2
korifi.cloudfoundry.org/v1alpha1/cfpackages: 2
kpack.io/v1alpha2/images: 2
kpack.io/v1alpha2/builds: 2
kpack.io/v1alpha2/sourceresolvers: 2
kpack.io/v1alpha1/images: 2
kpack.io/v1alpha1/builds: 2
kpack.io/v1alpha1/sourceresolvers: 2
metrics.k8s.io/v1beta1/pods: 2
```
