# StatsZone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetsWaits** | [***AllOfstatsZoneAssetsWaits**](AllOfstatsZoneAssetsWaits.md) |  | [optional] [default to null]
**ClientsWaits** | [***AllOfstatsZoneClientsWaits**](AllOfstatsZoneClientsWaits.md) |  | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**Id** | **string** | id of the zone | [default to null]
**MapId** | **string** | map_id of the zone | [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** | name of the zone | [default to null]
**NumAssets** | **int32** | number of assets | [optional] [default to null]
**NumClients** | **int32** | number of wifi clients (unconnected + connected) | [optional] [default to null]
**NumSdkclients** | **int32** | number of sdk clients | [optional] [default to null]
**OccupancyLimit** | **int32** |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**SdkclientsWaits** | [***AllOfstatsZoneSdkclientsWaits**](AllOfstatsZoneSdkclientsWaits.md) |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**Vertices** | [**[]ZoneVertex**](zone_vertex.md) | vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area | [optional] [default to null]
**VerticesM** | [**[]ZoneVertexM**](zone_vertex_m.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

