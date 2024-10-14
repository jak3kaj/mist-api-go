# Alarm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckAdminId** | **string** | UUID of the admin who acked the alarm | [optional] [default to null]
**AckAdminName** | **string** | Name &amp; Email ID of the admin who acked the alarm | [optional] [default to null]
**Acked** | **bool** | Whether the alarm is acked or not | [optional] [default to null]
**AckedTime** | **int32** | Epoch (seconds) when the alarm was acked | [optional] [default to null]
**Aps** | **[]string** | additional information: List of MACs of the APs | [optional] [default to null]
**Bssids** | **[]string** | List of BSSIDs | [optional] [default to null]
**Count** | **int32** | Number of incident within an alarm window | [default to null]
**Gateways** | **[]string** | additional information: List of MACs of the gateways | [optional] [default to null]
**Group** | **string** | Group of the alarm | [default to null]
**Hostnames** | **[]string** | additional information: List of Hostnames of the devices (AP/Switch/Gateway) | [optional] [default to null]
**Id** | **string** | UUID of the alarm | [default to null]
**LastSeen** | **float64** | Epoch (seconds) of the last incident/alarm within an alarm window | [default to null]
**Note** | **string** | Text describing the alarm | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**Severity** | **string** | Severity of the alarm | [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**Ssids** | **[]string** | List of SSIDs | [optional] [default to null]
**Switches** | **[]string** | additional information: List of MACs of the switches | [optional] [default to null]
**Timestamp** | **int32** | Epoch (seconds) of the first incident/alarm | [default to null]
**Type_** | **string** | Key-name of the alarm type | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

