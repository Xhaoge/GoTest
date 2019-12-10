package wblist

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"rr-factory.gloryholiday.com/yuetu/marineford/logger"
	"rr-factory.gloryholiday.com/yuetu/marineford/proto/marineford"
	pb "rr-factory.gloryholiday.com/yuetu/marineford/proto/marineford"
	"rr-factory.gloryholiday.com/yuetu/marineford/repository"
)

type ApplyType int32


var wbruleCollection = "wbrule"

var wbIndexKeys = []string{"displaynumber","applications.platformgroup.groupid","applications.providergroup.groupid","creator","enabled"}

type WBRuleMessage struct {
	WBRule *pb.WBRule
}

type WBRuleApplication struct {
	PlatformGroup 	*repository.OTAPlatformGroup
	ProviderGroup 	*repository.ProviderGroup
	ApplyType	   	string
	Enabled 		bool
}

type WBRuleEntity struct {
	repository.BaseInfo		`bson:",inline"`
	TripType				string
	Routes					[]*repository.LocationPair
	AbsSellDateRange		*repository.AbsDepDateRangeConstrain
	RelDepDateRange			*repository.RelDateRangeConstraint
	AbsDepDateRange 		[]*repository.AbsDepDateRangeConstrain
	Applications 			[]*WBRuleApplication
	Remark 					string
	ChangeHistory 			*repository.ChangeHistoryEntity
}


func (e *WBRuleEntity) MarshalMessage() (proto.Message,error) {
	return &pb.WBRule{
		BaseInfo:			repository.MarshalBaseInfoMessage(&e.BaseInfo),
		TripType:			pb.TripType(pb.TripType_value[e.TripType]),
		Routes:				repository.MarshalRouteMessage(e.Routes),
		AbsSellDateConstraint:repository.MarshalAbsDateRangeConstraintMessage(e.AbsSellDateRange)
		RelDepDateConstraint:  repository.MarshalRelDateRangeConstraintMessage(e.RelDepDateRange),
		AbsDepDateConstraints: repository.MarshalAbsDateRangeConstraintMessages(e.AbsDepDateRanges),
		ApplyTo:               marshalWBRuleApplicationMessages(e.Applications),
		Remark:                e.Remark,	
	},nil
}

func (wm *WBRuleMessage) MarshalEntity() (repository.Entity,error) {
	p := wm.WBRule
	return &WBRuleEntity{
		BaseInfo:         repository.MarshalBaseInfoEntity(p.BaseInfo),
		TripType:         p.TripType.String(),
		Routes:           repository.MarshalRouteEntities(p.Routes), // airport is not supported
		AbsSellDateRange: repository.MarshalAbsDateRangeConstraintEntity(p.AbsSellDateConstraint),
		RelDepDateRange:  repository.MarshalRelDateRangeConstraintEntity(p.RelDepDateConstraint),
		AbsDepDateRanges: repository.MarshalAbsDateRangeConstraintEntities(p.AbsDepDateConstraints),
		Applications:     marshalWBRuleApplicationEntities(p.ApplyTo),
		Remark:           p.Remark,
		ChangeHistory:    &repository.ChangeHistoryEntity{OperationLogs: []*repository.OperationLogEntity{}},
	},nil	
}
func (e *WBRuleEntity) GetBaseInfo() *repository.BaseInfo {
	return &e.BaseInfo
}


func (e *WBRuleEntity) GetPolicyPrefix() sting {
	return "WB"
}

func (e *WBRuleEntity) SetBaseInfo(baseinfo *repository.BaseInfo) {
	e.BaseInfo = *BaseInfo
}

func marshalWBRuleApplicationMessages(applications []*WBRuleApplication []*pb.WBRuleApplication) {
		var res []*pb.WBRuleApplication
	for _, app := range {
		res = append(res,marshalWBRuleApplicationMessage(app))
	}
	return res
}

func marshalWBRuleApplicationMessage(application *WBRuleApplication) *pb.WBRuleApplication {
	return &pb.WBRuleApplication{
		PlatformGroup:		repository.MarshalPlatformGroupMessage(application.PlatformGroup),
		ProviderGroup:		repository.MarshlProviderMessage(application.ProviderGroup),
		ApplyType:			pb.WBRuleApplication_APPLY_TYPE(pb.WBRuleApplication_APPLY_TYPE[application.ApplyType]),
	}
}

func marshelWBRuleApplicationEntities(applications []*pb.WBRuleApplication) []*WBRuleApplication {
	var res []*WBRuleApplication
	for _, app := range applications {
		res = append(res,marshalWBRuleApplicationEntity(app))
	}
	return res
}

func marshalWBRuleApplicationEntity(application *pb.WBRuleApplication) *WBRuleApplication {
	return *WBRuleApplication{
		PlatformGroup: repository.MarshalPlatformGroupEntity(application.GetPlatformGroup()),
		ProviderGroup: repository.MarshalProviderGroupEntity(application.GetProviderGroup()),
		ApplyType:     application.ApplyType.String(),
	}
}

type WBRuleRepo struct {
	repo repository.repository
}

func NewWBRuleRepo() *WBRuleRepo {
	mongoRepo := repository.NewMongoRepo(
		repository.WithCollection(wbruleCollection),
		repository.WithIndexKeys(wbIndexKeys)
	)
	return &WBRuleRepo(mongoRepo)
}


func (wr *WBRuleRepo) DisPlayNumberExists(displayNo string) bool {
	count, err := wr.repo.Count(repository.Selector{"displayNumber":displayNo})
	if err != nil {
		logger.Warn("count policy bu display no: %s failed",displayNo)
		return false
	}
	return count>0
}

func (wb *WBRuleRepo) BaseRepo() repository.Repository {
	return wb.repo
}

func (wr *WBRuleRepo) AppendOperationLogs(ids []string,operator string,operation string) {
	go func() {
		if err :+ wr.repo.UpdateAll(repository.IdsSelector(ids),map[string]map[string]interface{}{
			"$push":{"changehistory.operationlogs":repository.NewOperationLog(operation,operator)}
		}); err != nil{
			logger.Error("append operation logs failed:%s, ids:%s, operator:%s, operation:%s".err,ids,operator,operation)
		}
	}
}

func (wr *WBRuleEntity) MarshalProviderGroupAndPlatformGroupMessage(marshal proto.Message,providerGrooupMap map[string]*marineford.ProviderGroup,platformGroupMap map[string]*marineford.PlatformGroup){
	wbRule := maishall.(*pb.WBRule)
	for _,applyTo := range wbRule.ApplyTo{
		providerGroupId := applyTo.GetProviderGroup().GetBaseInfo().GetId()
		if len(providerGroupId) > 0 {
			providerGroup, ok := providerGroupMap[providerGroupId]
			if ok {
				applyTo.pROVIDERgROUP = providerGroup
			}else {
				logger.Error("ProviderGroup not found", errors.New("Missing ProviderGroup: "+providerGroupId))
			}
		}
		platformGroupId := applyTo.GetPlatformGroup().GetBaseInfo.GetId()
		if len(platformGroupId) > 0 {
			platformGroup,ok :+ platformGroupMap[platformGroupId]
			if ok {
				applyTo.PlatformGroup = platformGroup
			}else {
				logger.Error("PlatformGroup not found", errors.New("Missing PlatformGroup: "+platformGroupId))
			}
		}
	}	
}