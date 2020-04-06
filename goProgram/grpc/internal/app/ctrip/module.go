package ctrip

import "strconv"

type CtripBaseRequest struct {
	Cid          string `json:"cid" validate:"gt=4"`                                        // 接口身份标识用户名
	TripType     string `json:"tripType" validate:"oneof=1 2 3"` // 行程类型，1：单程；2：往返；3：多程（暂不支持）；
	AdultNumber  int32  `json:"adultNumber"`                                                // 成人人数1-9
	ChildNumber  int32  `json:"childNumber"`                                                // 儿童人数0-9
	InfantNumber int32  `json:"infantNumber"`                                               // 婴儿人数0-9
}

/**
 * 搜索 0，接口响应成功； 3，其他失败原因； 4，ctrip请求参数错误； 5，程序异常； -1，网络异常(ctrip使用) -2，response数据异常（ctrip使用）
 * 验价  0，成功； 1，舱位失败（请求舱位数大于查询的最大舱位数）；2，满舱（舱位已售完）；3，其他失败原因；4，ctrip 请求参数错误（即将废弃）；5，程序异常；6，币种不一致；7，航司或 GDS 超时； 111，无效的日期范围（第二段的航段时间早于第一段或航段重复）；112，参数验证不通过（如：没有传乘客或航班、乘客信息错误等）；200，航司或 GDS 结果异常；201，航司或 GDS 无可用的运价；202，航司或 GDS 指定的票价不可用；203，航司或 GDS QTE 出错（航班限制：比如不同时间价格不一样的航班未被过滤）；204，航司或 GDS 无联运协议
 */
type CtripBaseResponse struct {
	Status int32  `json:"status"` //
	Msg    string `json:"msg"`    // 提示信息
}

type CtripSearchRequest struct {
	CtripBaseRequest
	FromCity          string `json:"fromCity"`          // 出发地城市三字码
	ToCity            string `json:"toCity"`            // 目的地城市三字码
	FromDate          string `json:"fromDate"`          // 出发日期，格式为YYYYMMDD
	RetDate           string `json:"retDate"`           // 回程日期，格式为 YYYYMMDD
	Channel           string `json:"channel"`           // 搜索请求渠道来源； F：FlightIntlOnline； M：Mobile ; E：EnglishSite； K: 积分票 ( 对于积分票的查询请求 ，查询返回报文的productType务必赋值为JFP ,否则过滤 )
	MainChannel       string `json:"mainChannel"`       //主渠道，如 FlightIntlOnline/ Mobile/ EnglishSite 等
	SubChannelID      string `json:"subChannelId"`      //子渠道号 ,英网子渠道列表参考FAQ文件
	IsCompressEncode  string `json:"isCompressEncode"`  // 标识供应商查询返回报文是否需要压缩 1) 默认不压缩；如果为T，压缩编码；2）若需要压缩，请联系我们处理。
	IsPriceComparison bool   `json:"isPriceComparison"` //是否比价  ,默认false
}

type CtripSearchResponse struct {
	CtripBaseResponse
	Routings []*CtripRouting `json:"routings"` // 报价信息
}

func (rs *CtripSearchResponse) ResponseStatus() string {
	return strconv.Itoa(int(rs.Status))
}

func (rs *CtripSearchResponse) ErrorMsg() string {
	return rs.Msg
}

type CtripVerifyRequest struct {
	CtripBaseRequest
	ReferenceId string        `json:"referenceId"` // 携程关联 ID；携程用来查问题用的。
	Requesttype string        `json:"requesttype"`                  // 请求类型，全部为 1：验价；
	Routing     *CtripRouting `json:"routing" validate:"required"`  // 报价信息，参考搜索返回结果中的 Routing Elements。 1）只含航班信息，航班信息不含经停城市/机场，机型； 2）不含价格信息、退改签信息、行李额信息等
}

type CtripVerifyResponse struct {
	CtripBaseResponse
	CerrorCode string        `json:"cerrorCode"` // 供应商实际外部错误码
	CerrorMsg  string        `json:"cerrorMsg"`  // 供应商实际外部错误信息描述，长度小于 300
	SessionId  string        `json:"sessionId"`  // 会话标识：标记服务接口调用的唯一标识； 数字或字母，长度小于 20 个字符
	MaxSeats   int32         `json:"maxSeats"`   // 可预订的座位数，最大为 9； 供应商需要确保下 maxSeats 不小于验价请求人数
	Routing    *CtripRouting `json:"routing"`    // 报价信息 1）参考搜索返回结果中的 Routing Elements2）不含：airlineAncillaries 及 Rule3）posCode 必须返回
	Rule       *CtripRule    `json:"rule"`       // 退改签信息及行李额信息 1）参考搜索返回结果中的 Rule Element；2）Verify 接口，如果产品是公布运价，退改签获取异常和未获取到结果，需要将价格校验接口返回参数的 status 标注为6。
}

func (rs *CtripVerifyResponse) ResponseStatus() string {
	return strconv.Itoa(int(rs.Status))
}

func (rs *CtripVerifyResponse) ErrorMsg() string {
	return rs.Msg
}

type CtripRouting struct {
	Data                 string                  `json:"data" validate:"gt=10"` // 可保存必要信息，验价时会放在请求报文中传给供应商；
	PublishPrice         int32                   `json:"publishPrice"`          // 成人公布价（以CNY为单位），不含税；不校验请赋值0
	AdultPrice           int32                   `json:"adultPrice"`            // 成人单价 不含税
	AdultTax             int32                   `json:"adultTax"`              // 成人税费 1）不能是0，若存在为0的情况，请与我们联系
	ChildPublishPrice    int32                   `json:"childPublishPrice"`     // 儿童公布价，不含税； 不校验请赋值0
	ChildPrice           int32                   `json:"childPrice"`            // 儿童单价，不含税 1）对于含儿童的查询，必须返回 2）不得大于成人不含税单价，否则过滤
	ChildTax             int32                   `json:"childTax"`              // 儿童税费 1）对于含儿童的查询，必须返回；2）不能是0，若存在为0的情况，请与我们联系
	InfantPublishPrice   int32                   `json:"infantPublishPrice"`    // 婴儿公布价；不校验请赋值0
	InfantPrice          int32                   `json:"infantPrice"`           // 婴儿单价，不含税 1）对于含婴儿的查询，必须返回
	InfantTax            int32                   `json:"infantTax"`             // 婴儿税费 1）对于含婴儿的查询，必须返回； 2）可以为0
	AdultTaxType         int32                   `json:"adultTaxType"`          // 成人税费类型，0：未含税 / 1：已含税 正常赋0，如赋1请提前告知
	ChildTaxType         int32                   `json:"childTaxType"`          // 儿童税费类型，0：未含税 / 1：已含税 正常赋0，如赋1请提前告知
	PriceType            int32                   `json:"priceType"`             // 报价类型，0：普通价 / 1：留学生价 请全部赋0
	ApplyType            int32                   `json:"applyType"`             // 报价类型，0：预定价 / 1：申请价 请全部赋0
	Exchange             float64                 `json:"exchange"`              // 汇率；不校验请赋值0
	AdultAgeRestriction  string                  `json:"adultAgeRestriction"`   // 适用年龄区间 1）使用“-”表示“至”，例如*-12，表示12 岁及以下；2）置空表示无限制
	Eligibility          string                  `json:"eligibility"`           // 1）旅客资质，标准三字码： NOR：普通成人 (携程退改签强校验) LAB：劳务人员 SEA：海员 SNR：老年人 STU：学生 YOU：青年 2）如果投放非NOR的政策，请提前告知我们
	Nationality          string                  `json:"nationality"`           // 允许国籍，使用标准国家二字码 1）置空表示无限制（一般都是置空的）； 2）若多个使用“/”分割； 3）与forbiddenNationality 只能2 选1，若同时出现，为非 法数据，将被过滤
	ForbiddenNationality string                  `json:"forbiddenNationality"`  // 禁用国籍，使用标准国家二字码 1）置空表示无限制（一般都是置空的）； 2）若多个使用“/”分割； 3）与nationality 只能2选1，若同时出现，为非法数据，将被过滤
	PlanCategory         int32                   `json:"planCategory"`          // 产品类型，0：旅行套餐 / 1：商务优选 / 2：特惠推荐 新上线供应商请赋值为0
	InvoiceType          string                  `json:"invoiceType"`           // 报销凭证，T：行程单 / F：发票 / E：电子发票 1）默认F发票； 2）廉航票台可赋值为E电子发票； 3）非Eterm出票则无法打印行程单，不可赋值T行程单；否则出票后系统无法打印，将自动转为F发票开票； 4）赋值T行程单或者E电子发票，请提前通知我们调整配置
	MinStay              string                  `json:"minStay"`               // 最短停留时间，单位：天
	MaxStay              string                  `json:"maxStay"`               // 最长停留时间，单位：天
	MinPassengerCount    int32                   `json:"minPassengerCount"`     // 最小出行人数 如果没有返回，默认为1
	MaxPassengerCount    int32                   `json:"maxPassengerCount"`     // 最大出行人数 如果没有返回，默认为9
	BookingOfficeNo      string                  `json:"bookingOfficeNo"`       // 订位Office号，可为空
	TicketingOfficeNo    string                  `json:"ticketingOfficeNo"`     // 出票Office号，可为空
	ValidatingCarrier    string                  `json:"validatingCarrier"`     // 出票航司 1）整个行程只能赋一个出票航司； 2）如不赋值会取行程第一航段的carrier作为出票航司； 3）此字段非常重要，请务必准确赋值
	PosCode              string                  `json:"posCode"`               // 出票地国家2字码，多个出票地用‘|’隔开
	ComplexTerm          int32                   `json:"complexTerm"`           // 特殊产品复合类型
	ReservationType      string                  `json:"reservationType"`       // 政策来源 1）非公布运价此字段可不赋值； 2）公布运价此字段必须按要求返回，否则产品将按照未知订座系统，输出到旅行套餐； 3）使用IATA标准2字代码 1E：TravelSky 1A：Amadeus 1B：Abacus 1S：Sabre 1P：WorldSpan 1G：Galileo OT：未知订座系统来源
	ProductType          string                  `json:"productType"`           // 产品类型， 公布运价产品：PUB 控位产品：KWP 私有运价产品：PRV 积分票产品：JFP(只有在积分票查询请求时，返回积分票产品才有意义)
	FareBasis            string                  `json:"fareBasis"`             // 1）fareBasis数量必须和航段数一致 2）同一旅行方向的farebasis可以不一致（多段） 3）不同旅行方向farebasis 可以不一致 4）每航段1 个，使用“ ; ”分割
	AirlineAncillaries   *CtripAirlineAncillarie `json:"airlineAncillaries"`    // 增值服务信息
	FromSegments         []*CtripSegment         `json:"fromSegments"`          // 去程航段按顺序
	RetSegments          []*CtripSegment         `json:"retSegments"`           // 回程航段按顺序 单程 OW搜索时，此节点返回空
	Rule                 *CtripRule              `json:"rule"`                  // 免费行李信息，退改签信息
}

type CtripSegment struct {
	Carrier           string `json:"carrier"`           // 销售航司 IATA 二字码，必须与 flightNumber 航司相同
	FlightNumber      string `json:"flightNumber"`      // 航班号，如：CA123。 航班号数字前若有多余的数字 0，必须去掉；如 CZ006 需返回 CZ6
	DepAirport        string `json:"depAirport"`        // 出发机场；IATA 三字码
	DepTerminal       string `json:"depTerminal"`       // 出发航站楼；使用简写，例如T1
	DepTime           string `json:"depTime"`           // 起飞日期时间（当地时间） ，格式：YYYYMMDDHHMM  例：201203100300 表示 2012 年 3 月 10 日 3 时 0 分
	ArrAirport        string `json:"arrAirport"`        // 到达机场；IATA 三字码
	ArrTerminal       string `json:"arrTerminal"`       // 到达航站楼；使用简写，例如T1
	ArrTime           string `json:"arrTime"`           // 到达日期时间（当地时间） ，格式：YYYYMMDDHHMM  例：201203101305 表示 2012 年 3 月 10 日 13 时 5 分
	StopCities        string `json:"stopCities"`        // 经停城市 IATA三字码
	StopAirports      string `json:"stopAirports"`      // 经停机场 IATA三字码
	CodeShare         bool   `json:"codeShare"`         // 代码共享标识（true 代码共享/false 非代码共享）
	OperatingCarrier  string `json:"operatingCarrier"`  // 实际承运航司 若codeShare=true， operatingCarrier 不能为空。
	OperatingFlightNo string `json:"operatingFlightNo"` // 实际承运航班号
	Cabin             string `json:"cabin"`             // 子舱位，1-2 个字符
	CabinGrade        string `json:"cabinGrade"`        // 舱等； 头等：F；商务：C；超经：S；经济：Y 目前仅支持返回Y、S
	CabinCount        int32  `json:"cabinCount"`        // 舱位剩余数量，如果舱位数量超过9，请返回9 如果没有返回，默认为9
	AircraftCode      string `json:"aircraftCode"`      // 机型 ，IATA标准3 字码，并过滤下列机型运价信息 BUS|ICE|LCH|LMO|MTL|RFS|TGV|THS|THT|TRN|TSL|
	Duration          int32  `json:"duration"`          // 飞行时长；单位为分钟，通过时差转换后的结果
}

type CtripAirlineAncillarie struct {
	BaggageService bool `json:"baggageService"` // 行李增值服务，true：包含 / false：不包含
	UnFreeBaggage  bool `json:"unFreeBaggage"`  // 有无免费行李额，true：无免费行李额 / false：全部有免费行李额或部分有免费行李额 1）默认false
}

type CtripRule struct {
	FormatBaggageInfoList []*CtripFreeBaggage `json:"formatBaggageInfoList"` // 免费行李信息
	RefundInfoList        []*CtripRefund      `json:"refundInfoList"`        // 退票信息
	ChangesInfoList       []*CtripChange      `json:"changesInfoList"`       // 改期信息
	Note                  string              `json:"note"`                  // 备注信息，最大300个字符
	IsUseCtripRule        bool                `json:"isUseCtripRule"`        // 是否要使用携程退改签，true ：使用 / false：不使用 1）默认false
	TariffNo              string              `json:"tariffNo"`              // 公布运价相关参数，地理区间见运价集群编码 因供应商无法确定准确的tariffNo，此节点传运价类型， 公布运价PUB，私有运价PRI，强校验情况下，不传会过滤运价
	RuleNo                string              `json:"ruleNo"`                // 公布运价相关参数
	FareRuleMatchMode     int                 `json:"fareRuleMatchMode"`     // 退改签匹配模式，0：准确匹配 / 1：模糊匹配，1）默认1
}

type CtripFreeBaggage struct {
	SegmentNo     int32  `json:"segmentNo"`     // 航段序号，从1开始 按航段赋值
	PassengerType int32  `json:"passengerType"` // 乘客类型，0：成人 / 1：儿童 / 2：婴儿 1）对于多乘客类型的查询、验价，必须按乘客类型返回；如成人+儿童的查询，成人和儿童的行李额都要有
	BaggagePiece  int32  `json:"baggagePiece"`  // 托运行李额件数，单位PC，枚举值如下： 0：表示无免费托运行李，baggageWeight 需赋值-1； -1：表示计重制，baggageWeight表示每人可携带的总重量，baggageWeight必须赋正值 n （ n> 0 ） ：表示计件制，每人可携带n件行李，baggageWeight表示每件行李重量，baggageWeight必须赋正值
	BaggageWeight int32  `json:"baggageWeight"` // 托运行李额重量，单位KG，跟BaggagePiece配合使用
	CnBaggage     string `json:"cnBaggage"`     // 中文行李备注
	EnBaggage     string `json:"enBaggage"`     // 英文行李备注
}

type CtripRefund struct {
	RefundType         int32             `json:"refundType"`         // 退票类型 0：客票全部未使用； 1：客票部分使用（即去程已使用，在往返行程中使用，代表回程的退票信息） 单程时0；往返时0、1都要有
	ConditionList      []*CtripCondition `json:"conditionList"`      // 具体规定；参考下面的Condition Element
	RefundStatus       string            `json:"refundStatus"`       // 退票标识 T：不可退 H：有条件退 F：免费退 E：按航司客规商务优选产品类型专用，其他产品类型传E统一过滤
	RefundFee          float64           `json:"refundFee"`          // 退票费 1）若refundStatus =H，必须赋值； 2）若refundStatus =T/F，此字段可不赋值
	Currency           string            `json:"currency"`           // 退票费币种 1）若refundStatus =H，必须赋值 2）目前仅支持人民币：CNY 3）默认CNY
	PassengerType      int32             `json:"passengerType"`      // 乘客类型，0：成人 / 1：儿童 / 2：婴儿 对于多乘客类型的查询、验价，必须按乘客类型返回；如成人+儿童的查询，成人和儿童的退改签都要有
	RefNoshow          string            `json:"refNoshow"`          // 是否允许NoShow退票 T：不可退； H：有条件退； F：免费退； E：按航司客规为准，商务优选产品类型专用，其他产品类型传E统一过滤
	RefNoShowCondition int32             `json:"refNoShowCondition"` // 退票时航班起飞前多久算NoShow，单位：小时 1）若无法确认此时间，请默认赋0。
	RefNoshowFee       float64           `json:"refNoshowFee"`       // NoShow退票费用 1）当refNoshow =H，必须赋值； 2）当 refundStatus = H 且 refNoshow = H 时，展示给客人的noshow退票费= refundFee+ refNoshowFee 3）当 refundStatus = T 且 refNoshow = H时，展示给客人的noshow退票费 = 0+refNoshowFee
	CnRefRemark        string            `json:"cnRefRemark"`        // 中文退票备注
	EnRefRemark        string            `json:"enRefRemark"`        // 英文退票备注
}

type CtripChange struct {
	ChangesType        int32             `json:"changesType"`        // 改期类型 0：客票全部未使用； 1：客票部分使用(即去程已使用，在往返行程中使用，代表回程的改期信息) 单程时0；往返时0、1都要有
	ConditionList      []*CtripCondition `json:"conditionList"`      // 具体规定；参考下面的Condition Element
	ChangesStatus      string            `json:"changesStatus"`      // 改期标识 T：不可改期 H：有条件改期 F：免费改期 E：按航司客规，商务优选产品类型专用，其他产品类型传E统一过滤
	ChangesFee         float64           `json:"changesFee"`         // 改期费 1）当 changesStatus =H，必须赋值； 2）若 changesStatus =T/F，此字段可不赋值。
	Currency           string            `json:"currency"`           // 改期费币种 1）当refundStatus =H，必须赋值。 2）目前仅支持人民币：CNY 3）默认CNY
	PassengerType      int32             `json:"passengerType"`      // 乘客类型，0 成人/1 儿童/2 婴儿 1）对于对乘客类型的查询、验价，必须按乘客类型返回；如成人+儿童的查询，成人和儿童的退改签都要有
	RevNoshow          string            `json:"revNoshow"`          // 是否允许NoShow改期 T：不可改期； H：有条件改期； F：免费改期； E：按航司客规为准 商务优选产品类型专用，其他产品类型传E统一过滤
	RevNoShowCondition int32             `json:"revNoShowCondition"` // 改期时航班起飞前多久算NoShow，单位：小时 1）若无法确认此时间，请默认赋0。
	RevNoshowFee       float64           `json:"revNoshowFee"`       // NoShow改期费用 1）当revNoshow = H，必须赋值； 2）当 changesStatus = H 且 revNoshow = H 时，展示给客人的noshow改期费= changesFee+ revNoshowFee 3）当 changesStatus = T 且 revNoshow = H 时，展示给客人的noshow改期费 = 0+ revNoshowFee
	CnRevRemark        string            `json:"cnRevRemark"`        // 中文改期备注
	EnRevRemark        string            `json:"enRevRemark"`        // 英文改期备注
}

type CtripCondition struct {
	Status     string  `json:"status"`     // 改期标识 T：不可改期 H：有条件改期 F：免费改期 E：按航司客规，(商务优选产品类型专用，其他产品类型传E统一过滤)
	Amount     float64 `json:"amount"`     // 费用金额； 1）当 status = H，必须赋值； 2）当 status = T/F，可不赋值
	Percentage int32   `json:"percentage"` // 费用百分比（占票面不含税价）；
	EndMinute  int32   `json:"endMinute"`  // 结束时间； 例如：240 表示起飞前 4 小时结束；0 表示起飞时间结束； 240表示起飞时间后4 小时结束；-1 为保留值，表示不限制
}

type CtripOrderRequest struct {
	Cid                   string                      `json:"cid"`                   //Yes 接口身份标识用户名（渠道唯一标识）
	CtripOrderId          *CtripOrderId               `json:"ctripOrderId"`          //Yes 携程订单号，参考下面的 ctripOrderIdElement
	ReferenceId           string                      `json:"referenceId"`           //Yes 携程关联 ID
	TripType              string                      `json:"tripType"`              //Yes 行程类型，1：单程；2：往返；3：多程。
	SessionId             string                      `json:"sessionId"`             //Yes 会话标识：标记服务接口调用的唯一标识，会将价格校验接口中的 sessionId原值传给供应商
	Routing               string                      `json:"routing"`               //Yes 报价信息，参考搜索返回结果中的 Routing Elements；只含航班信息，不含价格信息、Rule 信息
	Passengers            []*CtripPassenger           `json:"passengers"`            //Yes 乘机人信息，参考下面的 Passenger Element
	Contact               *CtripContact               `json:"contact"`               //Yes 联系人信息，参考下面的 Contact Element
	Passengerbaggages     []*CtripPassengerbaggages   `json:"passengerbaggages"`     //No  乘机人预订行李信息，参考下面的 PassengerBaggage Element
	CtripRefRevServiceFee []*CtripRefRevServiceFee    `json:"ctripRefRevServiceFee"` //No  携程退改费用信息，参考下面的 CtripRefRevServiceFee Element
	ValueAddRequest       []*CtripBookValueAddRequest `json:"valueAddRequest"`       //No  增值服务请求信息，参考下面BookValueAddRequest Element 【JSON 示例未添加此项】
}

type CtripOrderId struct {
	AdultOrderId  float64 `json:"adultOrderId"`  //Yes  成人订单号
	ChildOrderId  float64 `json:"childOrderId"`  //NO   儿童订单号
	InfantOrderId float64 `json:"infantOrderId"` //NO   婴儿订单号
}

type CtripPassenger struct {
	Name           string      `json:"name"`           //Yes  LastName/FirstName MiddleName， 姓/名
	AgeType        int32       `json:"ageType"`        //Yes 乘客类型，0 成人/1 儿童/2 婴儿
	Birthday       string      `json:"birthday"`       //Yes 生日，格式：YYYYMMDD
	Gender         string      `json:"gender"`         //Yes 乘客性别，M 男 / F 女
	CardNum        string      `json:"cardNum"`        //Yes 证件号码，最大 15 个字符
	CardType       string      `json:"cardType"`       //Yes 证件类型： PP - 护 照 GA - 港澳通行证TW - 台湾通行证TB - 台胞证 HX - 回乡证 HY - 国际海员证
	CardIssuePlace string      `json:"cardIssuePlace"` //Yes 证件发行国家，国家二字码
	CardExpired    string      `json:"cardExpired"`    //Yes 证件有效时间，格式：YYYYMMDD
	Nationality    string      `json:"nationality"`    //Yes 乘客国籍，国家二字码
	FfpNo          *CtripFFPNo `json:"ffpNo"`          //No  常旅客 信息，参考下面的 FFP No Element
}

type CtripFFPNo struct {
	CardNo  string `json:"cardNo"`  //No 常旅客卡号
	Carrier string `json:"carrier"` //No 常旅客卡航司
}

type CtripContact struct {
	Name     string `json:"name"`     //Yes 联系人姓名，不单独区分姓和名
	Address  string `json:"address"`  //No  详细地址
	Postcode string `json:"postcode"` //No  邮编
	Email    string `json:"email"`    //No  联系人邮箱
	Mobile   string `json:"mobile"`   //No  联系人手机号
}

type CtripPassengerbaggages struct {
	PassengerName string `json:"passengerName"` //No LastName/FirstName MiddleName， 姓/名
	Baggages      string `json:"baggages"`      //No 行李信息集合参考 Baggage Element （增值服务查询接口中的行李额节点）
}

type CtripRefRevServiceFee struct {
	Id                            float64 `json:"id"`                            // No 携程订单号
	Status                        int32   `json:"status"`                        // No 退改签推送状态：0 正常， 1 异常
	OutboundRevalidationFee       float64 `json:"outboundRevalidationFee"`       // No 去程改期费用
	OutNonChg                     string  `json:"outNonChg"`                     // No 去程是否可改
	InboundRevalidationFee        float64 `json:"inboundRevalidationFee"`        // No 回程改期费用
	InNonChg                      string  `json:"inNonChg"`                      // No 回程是否可改
	RefundFeeByConsolidator       float64 `json:"refundFeeByConsolidator"`       // No 合作票台退票服务费
	OutRefundFee                  float64 `json:"outRefundFee"`                  // No 全部未使用退票费用
	OutNonRef                     string  `json:"outNonRef"`                     // No 全部未使用可否退票
	InRefundFee                   float64 `json:"inRefundFee"`                   // No 部分使用退票费用
	InNonRef                      string  `json:"inNonRef"`                      // No 部分使用可否退票
	RevalidationFeeByConsolidator float64 `json:"revalidationFeeByConsolidator"` // No 合作票台改期服务费
	OutRevFeeCurrency             string  `json:"outRevFeeCurrency"`             // No 去程改期费用币种
	InRevFeeCurrency              string  `json:"inRevFeeCurrency"`              // No 回程改期费用币种
	InRefCurrency                 string  `json:"inRefCurrency"`                 // No 部分使用退票费用币种
	OutRefCurrency                string  `json:"outRefCurrency"`                // No 全部未使用退票费用币种
	OutNonRev                     string  `json:"outNonRev"`                     // No 是否允许去程改期
	InNonRev                      string  `json:"inNonRev"`                      // No 是否允许回程改期
	OutRevChdFee                  float64 `json:"outRevChdFee"`                  // No 去程儿童改期费
	InRevChdFee                   float64 `json:"inRevChdFee"`                   // No 回程儿童改期费
	RevHasNoShow                  string  `json:"revHasNoShow"`                  // No 是否允许 No Show 改期
	OutRevNoShowFee               float64 `json:"outRevNoShowFee"`               // No 去程 No Show 改期费用（成人）
	InRevNoShowFee                float64 `json:"inRevNoShowFee"`                // No 回程 No Show 改期费用（成人）
	OutRevNoShowChdFee            float64 `json:"outRevNoShowChdFee"`            // No 去程 No Show 改期儿童费用
	InRevNoShowChdFee             float64 `json:"inRevNoShowChdFee"`             // No 回程 No Show 改期儿童费用
	RevNoShowCondition            float64 `json:"revNoShowCondition"`            // No NoShow 改期条件
	OutRefChdFee                  float64 `json:"outRefChdFee"`                  // No 全部退票儿童费用-去程
	InRefChdFee                   float64 `json:"inRefChdFee"`                   // No 部分退票儿童费用
	RefHasNoShow                  string  `json:"refHasNoShow"`                  // No 是否允许 No Show 退票
	OutRefNoShowFee               float64 `json:"outRefNoShowFee"`               // No 全部 No Show 退票费用（成人）-去程
	InRefNoShowFee                float64 `json:"inRefNoShowFee"`                // No 部分 No Show 退票费用（成人）
	OutRefNoShowChdFee            float64 `json:"outRefNoShowChdFee"`            // No 全部 No Show 退票儿童费用-去程
	InRefNoShowChdFee             float64 `json:"inRefNoShowChdFee"`             // No 部分 No Show 退票儿童费用
	RefNoShowCondition            float64 `json:"refNoShowCondition"`            // No No-Show 退票条件
}

type CtripBookValueAddRequest struct {
	ShoppingCarts  []*CtripValueAddShoppingCart `json:"shoppingCarts"`  //Yes 增值服务购物车参考 ValueAddShoppingCart Element
	OnlineCheckIns []*CtripOnlineCheckInDetail  `json:"onlineCheckIns"` //No  在线值机服务产品列表参考 OnlineCheckInDetail Element
}

type CtripValueAddShoppingCart struct {
	PassengerName string                      `json:"passengerName"` //Yes LastName/FirstName MiddleName 姓/名
	Products      []*CtripBookValueAddProduct `json:"products"`      //Yes 选购的产品列表参考 BookValueAddProduct Element
}

type CtripBookValueAddProduct struct {
	ProductId     string `json:"productId"`     //Yes 产品ID，值机服务映射OnlineCheckInDetail.Id
	Type          int32  `json:"type"`          //Yes 产品类型 1：值机服务
	Count         int32  `json:"count"`         //Yes 购买数量
	Tag           string `json:"tag"`           //No  定制化产品附加信息
	ProcessStatus string `json:"processStatus"` //Yes 处理状态 0： 待处理 1：处理成功 2： 处理失败
	ProcessDetail string `json:"processDetail"` //No  处理结果详细信息
}

type CtripOnlineCheckInDetail struct {
	OpenCheckInTime string `json:"openCheckInTime"` //Yes 开放值机时间格式：yyyyMMddHHmm例	：201203101305 表示2012 年3月10 日 13 时 5 分
	MinAge          int32  `json:"minAge"`          //No  最小值机人年龄
	CostPrice       string `json:"costPrice"`       //Yes 底价
	Id              string `json:"id"`              //Yes 产品 Id
	Flight          string `json:"flight"`          //Yes 航班号
	FromAirport     string `json:"fromAirport"`     //Yes 出发机场
	ToAirport       string `json:"toAirport"`       //Yes 到达机场
	DepTime         string `json:"depTime"`         //Yes 起飞日期时间格式：yyyyMMddHHmm例	：201203101305 表示2012 年3月10 日 13 时 5 分
}
