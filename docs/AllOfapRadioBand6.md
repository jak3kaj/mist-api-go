# AllOfapRadioBand6

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRrmDisable** | **bool** |  | [optional] [default to false]
**AntGain** | **int32** |  | [optional] [default to 0]
**AntennaMode** | [***Object**](.md) |  | [optional] [default to null]
**Bandwidth** | [***Object**](.md) |  | [optional] [default to null]
**Channel** | **int32** | For Device. (primary) channel for the band, 0 means using the Site Setting | [optional] [default to null]
**Channels** | **[]int32** | For RFTemplates. List of channels, null or empty array means auto | [optional] [default to []]
**Disabled** | **bool** | whether to disable the radio | [optional] [default to false]
**Power** | **int32** | TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / … | [optional] [default to null]
**PowerMax** | **int32** | when power&#x3D;0, max tx power to use, HW-specific values will be used if not set | [optional] [default to 18]
**PowerMin** | **int32** | when power&#x3D;0, min tx power to use, HW-specific values will be used if not set | [optional] [default to 8]
**Preamble** | [***Object**](.md) |  | [optional] [default to null]
**StandardPower** | **bool** | for 6GHz Only, standard-power operation, AFC (Automatic Frequency Coordination) will be performed and we&#x27;ll fallback to Low Power Indoor if AFC failed | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

