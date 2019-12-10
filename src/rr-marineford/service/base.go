package service

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
	"rr-factory.gloryholiday.com/yuetu/marineford/repository/ppconfig"
	repo "rr-factory.gloryholiday.com/yuetu/marineford/repository"
	"rr-factory.gloryholiday.com/yuetu/marineford"
	"rr-factory.gloryholiday.com/yuetu/marineford/utils"
)


type displayNoGenerator struct {
	prefix 			string
	checkIfExists	func(string)bool
}

func (d *displayNoGenerator) generateDisplayNo() string {
	generated := generateDisplayNoWithoutCheck(d.prefix)
	if d.checkIfExists(generated) {
		logger.WarnNt(logger.Message("display no %s already exists try to generate another oone",generated))
		return d.generateDisplayNo)()
	}
	return generated
}


func generateDisplayNoWithoutCheck(prefix string) string {
	return fmt.Sprintf("%s%s%s",prefix,time.Now().Format("060102"),utils.RandomAlphaNumber(4))
}

func revisionEntity(repository repo.Repository,traceId string,newVerdion repo.IBaseInfo,dbEntity repo.IBaseInfo) error {
	copied,err := ebEntity.GetBaseInfo().NewRevisionCopy(newVerdion.GetBaseInfo().Creator)
	newVersion.SetBaseInfo(copied)

	if err != nil {
		return err
	}

	updateId := ebEntity.GetBaseInfo().ID()
	dbEntity.GetBaseInfo().Id = bson.NewObjectId()
	dbEntity.GetBaseInfo().Archive()

	err = repository.Insert(dbEntity)
	displayNumber := dbEntity.GetBaseInfo().displayNumber
	if err != nil {
		logger.Error(traceId,logger.Message("insert back up entity %s %s failed",displayNumber,ebEntity.GetBaseInfo().Id.Hex(),err))
		return err
	}
	err = repository.UpdateById(updateId,newVerdion)
	if err != nil {
		logger.Error(traceId,logger.Message("update old version entity %s %s failed",displayNumber,updateId),err)
		if re := repository.RemoveId(dbEntity.GetBaseInfo().ID()); re != nil {
			logger.Error(traceId,logger.Message("rollback new backup entity when revision %s failed", dbEntity.GetBaseInfo().ID()),err)
		}
		return err 
	}

	return nil
}

func initBeforeCreate(displayNoGenerator *displayNoGenerator,ibe repo.IBaseInfo) {
	ibe.GetBaseInfo().Id = bson.NewObjectId()

	displayNumber := displayNoGenerator.generateDisplayNo()
	ibe.GetBaseInfo().DisplayNumber = displayNumber
	ibe.GetBaseInfo().CreatorAt = time.Now()
}


func initBeforeCreateWithoutDisplayNumber(ibe repo.IBaseInfo){
	ibe.GetBaseInfo().Id =  bson.NewObjectId()
	ibe.GetBaseInfo().CreatorAt = time.Now()
}

func tracesformProviderGroupMap(providerGroups []*ppconfig.ProviderGroupEntity) map[string]*marineford.ProviderGroup {
	var providerGroupMap = make(map[string]*marineford.ProviderGroup)
	for _, providerGroup := range providerGroups {
		id := providerGroup.GetBaseInfo().Id.Hex()
		providerGroupMap[id] = ppconfig.MarshalProviderGroupMessage(providerGroup)
	}
	return providerGroupMap
}

func tracesformPlatformGroupMap(platformGroups []*ppconfig.PlatformGroupEntity) map[string]*marineford.PlatformGroup {
	var platformGroupMap = make(map[string]*marineford.PlatformGroup)
	for _, platformGroup := range platformGroups {
		id := platformGroup.GetBaseInfo().Id.Hex()
		platformGroupMap[id] = ppconfig.MarshalPlatformGroupMessage(platformGroup)
	}
	return platformGroupMap
	
}