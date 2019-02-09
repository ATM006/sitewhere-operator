# SiteWhere Operator

## Install the operator

### Create Service Account

```console
kubectl create -f deploy/service_account.yaml
```

### Setup RBAC

```console
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml
```

### Setup the CRD

```console
kubectl create -f deploy/crds/sitewhere_v1alpha1_sitewhere_crd.yaml
```

### Deploy SiteWhere Operator

```console
kubectl create -f deploy/operator.yaml
```

## Create SiteWhere Custom Resource

### Create a SiteWhere CR

The default controller will watch for SiteWhere objects and create a pod for each CR

```console
kubectl create -f deploy/crds/sitewhere_v1alpha1_sitewhere_cr.yaml
```

## Cleanup

```console
kubectl delete -f deploy/crds/sitewhere_v1alpha1_sitewhere_cr.yaml
kubectl delete -f deploy/operator.yaml
kubectl delete -f deploy/role.yaml
kubectl delete -f deploy/role_binding.yaml
kubectl delete -f deploy/service_account.yaml
kubectl delete -f deploy/crds/sitewhere_v1alpha1_sitewhere_crd.yaml
```
