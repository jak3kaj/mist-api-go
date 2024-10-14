# AllOfnetworkTemplateRadiusConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctInterimInterval** | **int32** | how frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled | [optional] [default to 0]
**AcctServers** | [**[]RadiusAcctServer**](radius_acct_server.md) |  | [optional] [default to null]
**AuthServers** | [**[]RadiusAuthServer**](radius_auth_server.md) |  | [optional] [default to null]
**AuthServersRetries** | **int32** | radius auth session retries | [optional] [default to 3]
**AuthServersTimeout** | **int32** | radius auth session timeout | [optional] [default to 5]
**CoaEnabled** | **bool** |  | [optional] [default to false]
**CoaPort** | **int32** |  | [optional] [default to 3799]
**Network** | **string** | use &#x60;network&#x60;or &#x60;source_ip&#x60; which network the RADIUS server resides, if there&#x27;s static IP for this network, we&#x27;d use it as source-ip | [optional] [default to null]
**SourceIp** | **string** | use &#x60;network&#x60;or &#x60;source_ip&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

