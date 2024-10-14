# WebhookNacEventsEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** | random mac | [optional] [default to null]
**AuthType** | **string** | authentication type, e.g. \&quot;eap-tls\&quot;, \&quot;peap-tls\&quot;, \&quot;eap-ttls\&quot;, \&quot;eap-teap\&quot;, \&quot;mab\&quot;, \&quot;psk\&quot;, \&quot;device-auth\&quot; | [optional] [default to null]
**Bssid** | **string** | BSSID | [optional] [default to null]
**DryrunNacruleId** | **string** | NAC Policy Dry Run Rule ID, if present and matched | [optional] [default to null]
**DryrunNacruleMatched** | **bool** | True - if dryrun rule present and matched with priority, False - if not matched or not present | [optional] [default to null]
**IdpId** | **string** | SSO ID, if present and used | [optional] [default to null]
**IdpRole** | **[]string** | IDP returned roles/groups for the user | [optional] [default to null]
**Mac** | **string** | MAC address | [optional] [default to null]
**NacruleId** | **string** | NAC Policy Rule ID, if matched | [optional] [default to null]
**NacruleMatched** | **bool** | NAC Policy Rule Matched | [optional] [default to null]
**NasVendor** | **string** | vendor of NAS device | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**RandomMac** | **bool** | AP MAC | [optional] [default to null]
**RespAttrs** | **[]string** | Radius attributes returned by NAC to NAS Devive | [optional] [default to null]
**SiteId** | **string** | site id if assigned, null if not assigned | [optional] [default to null]
**Ssid** | **string** | SSID | [optional] [default to null]
**Timestamp** | **float64** | start time, in epoch | [optional] [default to null]
**Type_** | **string** | event type, e.g. NAC_CLIENT_PERMIT | [optional] [default to null]
**Username** | **string** | Username presented by the client | [optional] [default to null]
**Vlan** | **string** | Vlan ID | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

