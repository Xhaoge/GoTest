package ctrip

import (
	"context"
	"fmt"

	"Pro_golang/Golang/internal/api"
	"Pro_golang/Golang/internal/app/ctrip"
)

func NewCtripApiDefinition() {
	fmt.Println("this is package ctrip")
	apiDef := api.NewApiDefinition(api.WithName("Ctrip"), api.WithPrifix("ctrip"))
	ctripApiServer := ctrip.NewCtripApiServer(context.Background())

	return apiDef.
		WithHandler(api.GET, "ping", api.Ping).
		WithHandler(api.POST, "search", ctripApiServer.Search).
		WithHandler(api.POST, "Verify", ctripApiServer.Verify)

}
