# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgUserMacs**](OrgsUserMACsApi.md#CreateOrgUserMacs) | **Post** /api/v1/orgs/{org_id}/usermacs | createOrgUserMacs
[**DeleteOrgUserMac**](OrgsUserMACsApi.md#DeleteOrgUserMac) | **Delete** /api/v1/orgs/{org_id}/usermacs/{usermac_id} | deleteOrgUserMac
[**GetOrgUserMac**](OrgsUserMACsApi.md#GetOrgUserMac) | **Get** /api/v1/orgs/{org_id}/usermacs/{usermac_id} | getOrgUserMac
[**ImportOrgUserMacs**](OrgsUserMACsApi.md#ImportOrgUserMacs) | **Post** /api/v1/orgs/{org_id}/usermacs/import | importOrgUserMacs
[**SearchOrgUserMacs**](OrgsUserMACsApi.md#SearchOrgUserMacs) | **Get** /api/v1/orgs/{org_id}/usermacs/search | searchOrgUserMacs
[**UpdateOrgUserMac**](OrgsUserMACsApi.md#UpdateOrgUserMac) | **Put** /api/v1/orgs/{org_id}/usermacs/{usermac_id} | updateOrgUserMac

# **CreateOrgUserMacs**
> InlineResponse200106 CreateOrgUserMacs(ctx, orgId, optional)
createOrgUserMacs

Create Org User MACs  ### Usermacs import CSV file format mac,labels,vlan,notes  921b638445cd,”bldg1,flor1”,vlan-100  721b638445ef,”bldg2,flor2”,vlan-101,Canon Printers  721b638445ee,”bldg3,flor3”,vlan-102  921b638445ce,”bldg4,flor4”,vlan-103  921b638445cf,”bldg5,flor5”,vlan-104

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsUserMACsApiCreateOrgUserMacsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsUserMACsApiCreateOrgUserMacsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserMac**](UserMac.md)|  | 

### Return type

[**InlineResponse200106**](inline_response_200_106.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgUserMac**
> DeleteOrgUserMac(ctx, orgId, usermacId)
deleteOrgUserMac

Delete Org User MAC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **usermacId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgUserMac**
> InlineResponse200108 GetOrgUserMac(ctx, orgId, usermacId)
getOrgUserMac

Get Org User MAC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **usermacId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200108**](inline_response_200_108.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportOrgUserMacs**
> ImportOrgUserMacs(ctx, orgId, optional)
importOrgUserMacs

Import Org User MACs  ### CSV Import example ```csv  mac,labels,vlan,notes,name,radius_group 921b638445cd,\"bldg1,flor1\",vlan-100 721b638445ef,\"bldg2,flor2\",vlan-101,Canon Printers 721b638445ee,\"bldg3,flor3\",vlan-102,Printer2,VIP 921b638445ce,\"bldg4,flor4\",vlan-103 921b638445cf,\"bldg5,flor5\",vlan-104 ````

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsUserMACsApiImportOrgUserMacsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsUserMACsApiImportOrgUserMacsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgUserMacs**
> InlineResponse200107 SearchOrgUserMacs(ctx, orgId, optional)
searchOrgUserMacs

Search Org User MACs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsUserMACsApiSearchOrgUserMacsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsUserMACsApiSearchOrgUserMacsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| partial/full MAC addres | 
 **labels** | [**optional.Interface of []string**](string.md)| optional, array of strings of labels | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200107**](inline_response_200_107.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgUserMac**
> InlineResponse200108 UpdateOrgUserMac(ctx, orgId, usermacId, optional)
updateOrgUserMac

Update Org User MAC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **usermacId** | [**string**](.md)|  | 
 **optional** | ***OrgsUserMACsApiUpdateOrgUserMacOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsUserMACsApiUpdateOrgUserMacOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UserMac**](UserMac.md)|  | 

### Return type

[**InlineResponse200108**](inline_response_200_108.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

