# ApitokensApitokenIdBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | **string** | email of the token creator / null if creator is deleted | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**Key** | **string** |  | [optional] [default to null]
**LastUsed** | **float64** |  | [optional] [default to null]
**Name** | **string** | name of the token | [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**Privileges** | [**[]PrivilegeOrg**](privilege_org.md) | list of privileges the token has on the orgs/sites | [optional] [default to null]
**SrcIps** | **[]string** | list of allowed IP addresses from where the token can be used from. At most 10 IP addresses can be specified, cannot be changed once the API Token is created. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

