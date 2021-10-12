# \SlurmApi

All URIs are relative to *http://localhost/slurmdb/v0.0.37*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmdbdAddClusters**](SlurmApi.md#SlurmdbdAddClusters) | **Post** /clusters/ | Add clusters
[**SlurmdbdAddWckeys**](SlurmApi.md#SlurmdbdAddWckeys) | **Post** /wckeys/ | Add wckeys
[**SlurmdbdDeleteAccount**](SlurmApi.md#SlurmdbdDeleteAccount) | **Delete** /account/{account_name} | Delete account
[**SlurmdbdDeleteAssociation**](SlurmApi.md#SlurmdbdDeleteAssociation) | **Delete** /association/ | Delete association
[**SlurmdbdDeleteCluster**](SlurmApi.md#SlurmdbdDeleteCluster) | **Delete** /cluster/{cluster_name} | Delete cluster
[**SlurmdbdDeleteQos**](SlurmApi.md#SlurmdbdDeleteQos) | **Delete** /qos/{qos_name} | Delete QOS
[**SlurmdbdDeleteUser**](SlurmApi.md#SlurmdbdDeleteUser) | **Delete** /user/{user_name} | Delete user
[**SlurmdbdDeleteWckey**](SlurmApi.md#SlurmdbdDeleteWckey) | **Delete** /wckey/{wckey} | Delete wckey
[**SlurmdbdDiag**](SlurmApi.md#SlurmdbdDiag) | **Get** /diag/ | Get slurmdb diagnostics
[**SlurmdbdGetAccount**](SlurmApi.md#SlurmdbdGetAccount) | **Get** /account/{account_name} | Get account info
[**SlurmdbdGetAccounts**](SlurmApi.md#SlurmdbdGetAccounts) | **Get** /accounts/ | Get account list
[**SlurmdbdGetAssociation**](SlurmApi.md#SlurmdbdGetAssociation) | **Get** /association/ | Get association info
[**SlurmdbdGetAssociations**](SlurmApi.md#SlurmdbdGetAssociations) | **Get** /associations/ | Get association list
[**SlurmdbdGetCluster**](SlurmApi.md#SlurmdbdGetCluster) | **Get** /cluster/{cluster_name} | Get cluster info
[**SlurmdbdGetClusters**](SlurmApi.md#SlurmdbdGetClusters) | **Get** /clusters/ | Get cluster list
[**SlurmdbdGetDbConfig**](SlurmApi.md#SlurmdbdGetDbConfig) | **Get** /config | Dump all configuration information
[**SlurmdbdGetJob**](SlurmApi.md#SlurmdbdGetJob) | **Get** /job/{job_id} | Get job info
[**SlurmdbdGetJobs**](SlurmApi.md#SlurmdbdGetJobs) | **Get** /jobs/ | Get job list
[**SlurmdbdGetQos**](SlurmApi.md#SlurmdbdGetQos) | **Get** /qos/ | Get QOS list
[**SlurmdbdGetSingleQos**](SlurmApi.md#SlurmdbdGetSingleQos) | **Get** /qos/{qos_name} | Get QOS info
[**SlurmdbdGetTres**](SlurmApi.md#SlurmdbdGetTres) | **Get** /tres/ | Get TRES info
[**SlurmdbdGetUser**](SlurmApi.md#SlurmdbdGetUser) | **Get** /user/{user_name} | Get user info
[**SlurmdbdGetUsers**](SlurmApi.md#SlurmdbdGetUsers) | **Get** /users/ | Get user list
[**SlurmdbdGetWckey**](SlurmApi.md#SlurmdbdGetWckey) | **Get** /wckey/{wckey} | Get wckey info
[**SlurmdbdGetWckeys**](SlurmApi.md#SlurmdbdGetWckeys) | **Get** /wckeys/ | Get wckey list
[**SlurmdbdSetDbConfig**](SlurmApi.md#SlurmdbdSetDbConfig) | **Post** /config | Load all configuration information
[**SlurmdbdUpdateAccount**](SlurmApi.md#SlurmdbdUpdateAccount) | **Post** /accounts/ | Update accounts
[**SlurmdbdUpdateTres**](SlurmApi.md#SlurmdbdUpdateTres) | **Post** /tres/ | Set TRES info
[**SlurmdbdUpdateUsers**](SlurmApi.md#SlurmdbdUpdateUsers) | **Post** /users/ | Update user



## SlurmdbdAddClusters

> Dbv0037ResponseClusterAdd SlurmdbdAddClusters(ctx).Execute()

Add clusters

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdAddClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdAddClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdAddClusters`: Dbv0037ResponseClusterAdd
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdAddClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdAddClustersRequest struct via the builder pattern


### Return type

[**Dbv0037ResponseClusterAdd**](Dbv0037ResponseClusterAdd.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdAddWckeys

> Dbv0037ResponseWckeyAdd SlurmdbdAddWckeys(ctx).Execute()

Add wckeys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdAddWckeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdAddWckeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdAddWckeys`: Dbv0037ResponseWckeyAdd
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdAddWckeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdAddWckeysRequest struct via the builder pattern


### Return type

[**Dbv0037ResponseWckeyAdd**](Dbv0037ResponseWckeyAdd.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteAccount

> Dbv0037ResponseAccountDelete SlurmdbdDeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountName := "accountName_example" // string | Slurm Account Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdDeleteAccount(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteAccount`: Dbv0037ResponseAccountDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Slurm Account Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ResponseAccountDelete**](Dbv0037ResponseAccountDelete.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteAssociation

> Dbv0037ResponseAssociationDelete SlurmdbdDeleteAssociation(ctx).Account(account).User(user).Cluster(cluster).Partition(partition).Execute()

Delete association

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    account := "account_example" // string | Account name
    user := "user_example" // string | User name
    cluster := "cluster_example" // string | Cluster name (optional)
    partition := "partition_example" // string | Partition Name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdDeleteAssociation(context.Background()).Account(account).User(user).Cluster(cluster).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteAssociation`: Dbv0037ResponseAssociationDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **cluster** | **string** | Cluster name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0037ResponseAssociationDelete**](Dbv0037ResponseAssociationDelete.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteCluster

> Dbv0037ResponseClusterDelete SlurmdbdDeleteCluster(ctx, clusterName).Execute()

Delete cluster

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterName := "clusterName_example" // string | Slurm cluster name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdDeleteCluster(context.Background(), clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteCluster`: Dbv0037ResponseClusterDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Slurm cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ResponseClusterDelete**](Dbv0037ResponseClusterDelete.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteQos

> Dbv0037ResponseQosDelete SlurmdbdDeleteQos(ctx, qosName).Execute()

Delete QOS

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    qosName := "qosName_example" // string | Slurm QOS Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdDeleteQos(context.Background(), qosName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteQos`: Dbv0037ResponseQosDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qosName** | **string** | Slurm QOS Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ResponseQosDelete**](Dbv0037ResponseQosDelete.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteUser

> Dbv0037ResponseUserDelete SlurmdbdDeleteUser(ctx, userName).Execute()

Delete user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userName := "userName_example" // string | Slurm User Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdDeleteUser(context.Background(), userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteUser`: Dbv0037ResponseUserDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | Slurm User Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ResponseUserDelete**](Dbv0037ResponseUserDelete.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDeleteWckey

> Dbv0037ResponseWckeyDelete SlurmdbdDeleteWckey(ctx, wckey).Execute()

Delete wckey

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    wckey := "wckey_example" // string | Slurm wckey name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdDeleteWckey(context.Background(), wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDeleteWckey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDeleteWckey`: Dbv0037ResponseWckeyDelete
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wckey** | **string** | Slurm wckey name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ResponseWckeyDelete**](Dbv0037ResponseWckeyDelete.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdDiag

> Dbv0037Diag SlurmdbdDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdDiag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdDiag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdDiag`: Dbv0037Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdDiagRequest struct via the builder pattern


### Return type

[**Dbv0037Diag**](Dbv0037Diag.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetAccount

> Dbv0037AccountInfo SlurmdbdGetAccount(ctx, accountName).Execute()

Get account info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountName := "accountName_example" // string | Slurm Account Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetAccount(context.Background(), accountName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAccount`: Dbv0037AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Slurm Account Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037AccountInfo**](Dbv0037AccountInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetAccounts

> Dbv0037AccountInfo SlurmdbdGetAccounts(ctx).Execute()

Get account list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAccounts`: Dbv0037AccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAccountsRequest struct via the builder pattern


### Return type

[**Dbv0037AccountInfo**](Dbv0037AccountInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetAssociation

> Dbv0037AssociationsInfo SlurmdbdGetAssociation(ctx).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()

Get association info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cluster := "cluster_example" // string | Cluster name (optional)
    account := "account_example" // string | Account name (optional)
    user := "user_example" // string | User name (optional)
    partition := "partition_example" // string | Partition Name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetAssociation(context.Background()).Cluster(cluster).Account(account).User(user).Partition(partition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAssociation`: Dbv0037AssociationsInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **account** | **string** | Account name | 
 **user** | **string** | User name | 
 **partition** | **string** | Partition Name | 

### Return type

[**Dbv0037AssociationsInfo**](Dbv0037AssociationsInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetAssociations

> Dbv0037AssociationsInfo SlurmdbdGetAssociations(ctx).Execute()

Get association list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetAssociations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetAssociations`: Dbv0037AssociationsInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetAssociations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetAssociationsRequest struct via the builder pattern


### Return type

[**Dbv0037AssociationsInfo**](Dbv0037AssociationsInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetCluster

> Dbv0037ClusterInfo SlurmdbdGetCluster(ctx, clusterName).Execute()

Get cluster info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterName := "clusterName_example" // string | Slurm cluster name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetCluster(context.Background(), clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetCluster`: Dbv0037ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Slurm cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037ClusterInfo**](Dbv0037ClusterInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetClusters

> Dbv0037ClusterInfo SlurmdbdGetClusters(ctx).Execute()

Get cluster list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetClusters`: Dbv0037ClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetClustersRequest struct via the builder pattern


### Return type

[**Dbv0037ClusterInfo**](Dbv0037ClusterInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetDbConfig

> Dbv0037ConfigInfo SlurmdbdGetDbConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetDbConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetDbConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetDbConfig`: Dbv0037ConfigInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetDbConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetDbConfigRequest struct via the builder pattern


### Return type

[**Dbv0037ConfigInfo**](Dbv0037ConfigInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetJob

> Dbv0037JobInfo SlurmdbdGetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jobId := int64(789) // int64 | Slurm Job ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetJob`: Dbv0037JobInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037JobInfo**](Dbv0037JobInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetJobs

> Dbv0037JobInfo SlurmdbdGetJobs(ctx).SubmitTime(submitTime).StartTime(startTime).EndTime(endTime).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SkipSteps(skipSteps).DisableWaitForResult(disableWaitForResult).ExitCode(exitCode).Format(format).Group(group).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).State(state).Step(step).Node(node).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    submitTime := "submitTime_example" // string | Filter by submission time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    startTime := "startTime_example" // string | Filter by start time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    endTime := "endTime_example" // string | Filter by end time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] (optional)
    account := "account_example" // string | Comma delimited list of accounts to match (optional)
    association := "association_example" // string | Comma delimited list of associations to match (optional)
    cluster := "cluster_example" // string | Comma delimited list of cluster to match (optional)
    constraints := "constraints_example" // string | Comma delimited list of constraints to match (optional)
    cpusMax := "cpusMax_example" // string | Number of CPUs high range (optional)
    cpusMin := "cpusMin_example" // string | Number of CPUs low range (optional)
    skipSteps := true // bool | Report job step information (optional)
    disableWaitForResult := true // bool | Disable waiting for result from slurmdbd (optional)
    exitCode := "exitCode_example" // string | Exit code of job (optional)
    format := "format_example" // string | Comma delimited list of formats to match (optional)
    group := "group_example" // string | Comma delimited list of groups to match (optional)
    jobName := "jobName_example" // string | Comma delimited list of job names to match (optional)
    nodesMax := "nodesMax_example" // string | Number of nodes high range (optional)
    nodesMin := "nodesMin_example" // string | Number of nodes low range (optional)
    partition := "partition_example" // string | Comma delimited list of partitions to match (optional)
    qos := "qos_example" // string | Comma delimited list of QOS to match (optional)
    reason := "reason_example" // string | Comma delimited list of job reasons to match (optional)
    reservation := "reservation_example" // string | Comma delimited list of reservations to match (optional)
    state := "state_example" // string | Comma delimited list of states to match (optional)
    step := "step_example" // string | Comma delimited list of job steps to match (optional)
    node := "node_example" // string | Comma delimited list of used nodes to match (optional)
    wckey := "wckey_example" // string | Comma delimited list of wckeys to match (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetJobs(context.Background()).SubmitTime(submitTime).StartTime(startTime).EndTime(endTime).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SkipSteps(skipSteps).DisableWaitForResult(disableWaitForResult).ExitCode(exitCode).Format(format).Group(group).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).State(state).Step(step).Node(node).Wckey(wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetJobs`: Dbv0037JobInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitTime** | **string** | Filter by submission time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] | 
 **startTime** | **string** | Filter by start time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] | 
 **endTime** | **string** | Filter by end time  Accepted formats:  HH:MM[:SS] [AM|PM]  MMDD[YY] or MM/DD[/YY] or MM.DD[.YY]  MM/DD[/YY]-HH:MM[:SS]  YYYY-MM-DD[THH:MM[:SS]] | 
 **account** | **string** | Comma delimited list of accounts to match | 
 **association** | **string** | Comma delimited list of associations to match | 
 **cluster** | **string** | Comma delimited list of cluster to match | 
 **constraints** | **string** | Comma delimited list of constraints to match | 
 **cpusMax** | **string** | Number of CPUs high range | 
 **cpusMin** | **string** | Number of CPUs low range | 
 **skipSteps** | **bool** | Report job step information | 
 **disableWaitForResult** | **bool** | Disable waiting for result from slurmdbd | 
 **exitCode** | **string** | Exit code of job | 
 **format** | **string** | Comma delimited list of formats to match | 
 **group** | **string** | Comma delimited list of groups to match | 
 **jobName** | **string** | Comma delimited list of job names to match | 
 **nodesMax** | **string** | Number of nodes high range | 
 **nodesMin** | **string** | Number of nodes low range | 
 **partition** | **string** | Comma delimited list of partitions to match | 
 **qos** | **string** | Comma delimited list of QOS to match | 
 **reason** | **string** | Comma delimited list of job reasons to match | 
 **reservation** | **string** | Comma delimited list of reservations to match | 
 **state** | **string** | Comma delimited list of states to match | 
 **step** | **string** | Comma delimited list of job steps to match | 
 **node** | **string** | Comma delimited list of used nodes to match | 
 **wckey** | **string** | Comma delimited list of wckeys to match | 

### Return type

[**Dbv0037JobInfo**](Dbv0037JobInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetQos

> Dbv0037QosInfo SlurmdbdGetQos(ctx).Execute()

Get QOS list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetQos(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetQos`: Dbv0037QosInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetQos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetQosRequest struct via the builder pattern


### Return type

[**Dbv0037QosInfo**](Dbv0037QosInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetSingleQos

> Dbv0037QosInfo SlurmdbdGetSingleQos(ctx, qosName).Execute()

Get QOS info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    qosName := "qosName_example" // string | Slurm QOS Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetSingleQos(context.Background(), qosName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetSingleQos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetSingleQos`: Dbv0037QosInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qosName** | **string** | Slurm QOS Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037QosInfo**](Dbv0037QosInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetTres

> Dbv0037TresInfo SlurmdbdGetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetTres(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetTres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetTres`: Dbv0037TresInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetTresRequest struct via the builder pattern


### Return type

[**Dbv0037TresInfo**](Dbv0037TresInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetUser

> Dbv0037UserInfo SlurmdbdGetUser(ctx, userName).Execute()

Get user info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userName := "userName_example" // string | Slurm User Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetUser(context.Background(), userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetUser`: Dbv0037UserInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | Slurm User Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037UserInfo**](Dbv0037UserInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetUsers

> Dbv0037UserInfo SlurmdbdGetUsers(ctx).Execute()

Get user list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetUsers`: Dbv0037UserInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetUsersRequest struct via the builder pattern


### Return type

[**Dbv0037UserInfo**](Dbv0037UserInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetWckey

> Dbv0037WckeyInfo SlurmdbdGetWckey(ctx, wckey).Execute()

Get wckey info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    wckey := "wckey_example" // string | Slurm wckey name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetWckey(context.Background(), wckey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetWckey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetWckey`: Dbv0037WckeyInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wckey** | **string** | Slurm wckey name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dbv0037WckeyInfo**](Dbv0037WckeyInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdGetWckeys

> Dbv0037WckeyInfo SlurmdbdGetWckeys(ctx).Execute()

Get wckey list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdGetWckeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdGetWckeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdGetWckeys`: Dbv0037WckeyInfo
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdGetWckeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdGetWckeysRequest struct via the builder pattern


### Return type

[**Dbv0037WckeyInfo**](Dbv0037WckeyInfo.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdSetDbConfig

> Dbv0037ConfigResponse SlurmdbdSetDbConfig(ctx).Execute()

Load all configuration information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdSetDbConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdSetDbConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdSetDbConfig`: Dbv0037ConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdSetDbConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdSetDbConfigRequest struct via the builder pattern


### Return type

[**Dbv0037ConfigResponse**](Dbv0037ConfigResponse.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdUpdateAccount

> Dbv0037AccountResponse SlurmdbdUpdateAccount(ctx).Execute()

Update accounts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdUpdateAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdUpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateAccount`: Dbv0037AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdUpdateAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateAccountRequest struct via the builder pattern


### Return type

[**Dbv0037AccountResponse**](Dbv0037AccountResponse.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdUpdateTres

> Dbv0037ResponseTres SlurmdbdUpdateTres(ctx).Execute()

Set TRES info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdUpdateTres(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdUpdateTres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateTres`: Dbv0037ResponseTres
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdUpdateTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateTresRequest struct via the builder pattern


### Return type

[**Dbv0037ResponseTres**](Dbv0037ResponseTres.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbdUpdateUsers

> Dbv0037ResponseUserUpdate SlurmdbdUpdateUsers(ctx).Execute()

Update user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmdbdUpdateUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmdbdUpdateUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmdbdUpdateUsers`: Dbv0037ResponseUserUpdate
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmdbdUpdateUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbdUpdateUsersRequest struct via the builder pattern


### Return type

[**Dbv0037ResponseUserUpdate**](Dbv0037ResponseUserUpdate.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

