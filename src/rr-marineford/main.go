package main

import (
	corectx "rr-factory.gloryholiday.com/yuetu/golang-core-context"
	grpc_service "rr-factory.gloryholiday.com/yuetu/marineford/grpc-service"
)


func main() {
	cts,_ := corectx.WithCancelAndSignalHandler()
	grpc_service.GrpcStartup()
	<-ctx.Done()
}