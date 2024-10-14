# OrgIdServicesBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, ip subnets (e.g. 10.0.0.0/8) | [optional] [default to null]
**AppCategories** | **[]string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;app_categories&#x60;, list of application categories are available through /api/v1/const/app_categories | [optional] [default to null]
**AppSubcategories** | **[]string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;app_categories&#x60;, list of application categories are available through /api/v1/const/app_subcategories | [optional] [default to null]
**Apps** | **[]string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;apps&#x60;, list of applications are available through:   * /api/v1/const/applications   * /api/v1/const/gateway_applications   * /insight/top_app_by-bytes?wired&#x3D;true | [optional] [default to null]
**ClientLimitDown** | **int32** | 0 means unlimited | [optional] [default to 0]
**ClientLimitUp** | **int32** | 0 means unlimited | [optional] [default to 0]
**CreatedTime** | **float64** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Dscp** | [***Object**](.md) |  | [optional] [default to null]
**FailoverPolicy** | [***Object**](.md) |  | [optional] [default to null]
**Hostnames** | **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, web filtering | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**MaxJitter** | [***Object**](.md) |  | [optional] [default to null]
**MaxLatency** | [***Object**](.md) |  | [optional] [default to null]
**MaxLoss** | [***Object**](.md) |  | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**ServiceLimitDown** | **int32** | 0 means unlimited | [optional] [default to 0]
**ServiceLimitUp** | **int32** | 0 means unlimited | [optional] [default to 0]
**SleEnabled** | **bool** | whether to enable measure SLE | [optional] [default to false]
**Specs** | [**[]ServiceSpec**](service_spec.md) | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, optional, if it doesn&#x27;t exist, http and https is assumed | [optional] [default to null]
**SsrRelaxedTcpStateEnforcement** | **bool** |  | [optional] [default to false]
**TrafficClass** | [***Object**](.md) |  | [optional] [default to null]
**TrafficType** | **string** | values from &#x60;/api/v1/consts/traffic_types&#x60; | [optional] [default to data_best_effort]
**Type_** | [***Object**](.md) |  | [optional] [default to null]
**Urls** | **[]string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;urls&#x60;, no need for spec as URL can encode the ports being used | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

