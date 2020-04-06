package utils

import (
	"fmt"
	"strings"
)

type routingDataPos int

const (
	TraceIdPos routingDataPos = iota
	UidPos
	TripInfoPos
	HitCachePos
)

type RoutingData struct {
	TraceId  string
	Uid      string
	TripInfo string
	CacheHit bool
}

func FetchTraceIdFromData(data string) string {
	traceId, _, err := ResolveKeysFromData(data)
	if nil != err {
		return ""
	}
	return traceId
}

func FetchRoutingKeyFromData(data string) string {
	_, routingKey, err := ResolveKeysFromData(data)
	if nil != err {
		return ""
	}
	return routingKey
}

func ResolveKeysFromData(data string) (traceId, routingKey string, err error) {
	ds := strings.Split(data, "#")
	if len(ds) < 2 {
		err = fmt.Errorf("invalid routing data: %s", data)
		return "", "", err
	}
	traceId = ds[TraceIdPos]
	routingKey = ds[UidPos]
	return
}

func (rd *RoutingData) BuildData() string {
	ds := make([]string, 4)
	ds[TraceIdPos] = rd.TraceId
	ds[UidPos] = rd.Uid
	ds[TripInfoPos] = rd.TripInfo
	if rd.CacheHit {
		ds[HitCachePos] = "1"
	} else {
		ds[HitCachePos] = "0"
	}
	return strings.Join(ds, "#")
}
