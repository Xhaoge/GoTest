package location

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
	"rr-factory.gloryholiday.com/yuetu/golang-core/utils"
)

type Locales struct {
	En   string `json:"en"`
	ZhTW string `json:"zh_TW"`
	ZhCN string `json:"zh_CN"`
	Ko   string `json:"ko"`
}

type City struct {
	Code     string         `json:"code"`
	Country  *Country       `json:"-"`
	TimeZone string         `json:"-"`
	Location *time.Location `json:"-"`
}

func (city *City) WithDefault(def *City) *City {
	if city == nil {
		return def
	}
	return city
}

type Country struct {
	Code        string `json:"code"`
	Area        *Area  `json:"area"`
	RoutingCode string `json:"routingCode"`
}

type Area struct {
	Code string `json:"code"`
}

type Airport struct {
	City     *City          `json:"city"`
	Code     string         `json:"code"`
	Country  *Country       `json:"country"`
	TimeZone string         `json:"timeZone"`
	Location *time.Location `json:"-"`
}

type LocationTime struct {
	LocationCode string
	Time         string
	Layout       string
}

type Location interface {
	Duration(LocationTime, LocationTime) (int32, error)
	GetAirport(string) *Airport
	GetCity(string) *City
	GetCountry(string) *Country
	GetArea(string) *Area
	BelongsTo(string, string) bool
}

const (
	DefaultLayout        string = "2006-01-02T15:04:05"
	DepArrDateTimeLayout string = "200601021504"
	areaEncodingMask     int64  = 0x7ffff00000000000 // The top rest bits to represent areas
	countryEncodingMask  int64  = 0x7fffffff00000000 // Use 12 bits to represent countries, 4096 is enough for all countries on earth
	cityEncodingMask     int64  = 0x7fffffffffff0000 // Use 16 bits to represent cities, and the lowest 16 bits to represent airports

	AirportWeight = 8
	CityWeight    = 4
	CountryWeight = 2
	AreaWeight    = 1
)

type loc struct {
	airportsEncoding  map[string]int64
	airportsMap       map[string]*Airport
	citiesEncoding    map[string]int64
	citiesMap         map[string]*City
	countriesEncoding map[string]int64
	countriesMap      map[string]*Country
	areasEncoding     map[string]int64
	areasMap          map[string]*Area
}

var location Location
var initializer sync.Once = sync.Once{}

func init() {
	location = NewLocation()
}

func NewLocation() Location {
	initializer.Do(func() {
		var airports []Airport
		err := json.Unmarshal([]byte(location_data), &airports)
		if err != nil {
			logger.Fatal(logger.Message("Failed to decode your airports data: %s", err.Error()))
		}

		ape := map[string]int64{}
		ap := map[string]*Airport{}
		cpe := map[string]int64{}
		cp := map[string]*City{}
		cme := map[string]int64{}
		cm := map[string]*Country{}
		ame := map[string]int64{}
		am := map[string]*Area{}

		for i, _ := range airports {
			airport := &airports[i]

			airportLocation, err := time.LoadLocation(airport.TimeZone)
			if err != nil {
				logger.ErrorNt("Failed to load location for airport: "+airport.Code, err)
				continue
			}

			airport.Location = airportLocation
			city := airport.City
			city.TimeZone = airport.TimeZone
			city.Location = airportLocation
			country := airport.Country
			area := country.Area
			city.Country = country

			airportCodeL := strings.ToLower(airport.Code)
			airportCodeU := strings.ToUpper(airport.Code)
			cityCodeL := strings.ToLower(city.Code)
			cityCodeU := strings.ToUpper(city.Code)
			countryCodeL := strings.ToLower(country.Code)
			countryCodeU := strings.ToUpper(country.Code)
			areaCodeL := strings.ToLower(area.Code)
			areaCodeU := strings.ToUpper(area.Code)

			ap[airportCodeL] = airport
			ap[airportCodeU] = airport
			cp[cityCodeL] = city
			cp[cityCodeU] = city
			cm[countryCodeL] = country
			cm[countryCodeU] = country
			am[areaCodeL] = area
			am[areaCodeU] = area

			areaEncoding, ok := ame[areaCodeL]

			if !ok {
				areaEncoding = int64(1+len(ame)) << 55
				ame[areaCodeL] = areaEncoding
				ame[areaCodeU] = areaEncoding
			}

			ctyEncoding, ok := cme[countryCodeL]
			if !ok {
				ctyEncoding = int64(1+len(cme))<<40 | areaEncoding
				cme[countryCodeL] = ctyEncoding
				cme[countryCodeU] = ctyEncoding
			}

			cityEncoding, ok := cpe[cityCodeL]
			if !ok {
				cityEncoding = int64(1+len(cpe))<<20 | ctyEncoding
				cpe[cityCodeL] = cityEncoding
				cpe[cityCodeU] = cityEncoding
			}

			_, ok = ape[airportCodeL]
			if !ok {
				airportEncoding := int64(1+len(ape)) | cityEncoding
				ape[airportCodeL] = airportEncoding
				ape[airportCodeU] = airportEncoding
			}
		}

		location = &loc{
			airportsEncoding:  ape,
			airportsMap:       ap,
			citiesEncoding:    cpe,
			citiesMap:         cp,
			countriesEncoding: cme,
			countriesMap:      cm,
			areasEncoding:     ame,
			areasMap:          am,
		}
	})

	return location
}

func newError(message string) error {
	return errors.New(message)
}

func notFound(code string) error {
	return newError(fmt.Sprintf("%s not found", code))
}

func (l *loc) getLocation(code string) *time.Location {
	airport := l.GetAirport(code)
	if nil != airport {
		return airport.Location
	}

	city, ok := l.citiesMap[code]
	if ok {
		return city.Location
	}

	return nil
}

var unknownCodeErr = errors.New("ignore err")

func (l *loc) Duration(depAt LocationTime, arrAt LocationTime) (int32, error) {
	depLoc := l.getLocation(depAt.LocationCode)
	if nil == depLoc {
		if unknownCodes[depAt.LocationCode] {
			return 0, unknownCodeErr
		}
		return 0, notFound(depAt.LocationCode)
	}

	arrLoc := l.getLocation(arrAt.LocationCode)
	if nil == arrLoc {
		if unknownCodes[arrAt.LocationCode] {
			return 0, unknownCodeErr
		}
		return 0, notFound(arrAt.LocationCode)
	}

	depTime, err := time.ParseInLocation(depAt.Layout, depAt.Time, depLoc)
	if err != nil {
		return 0, err
	}

	arrTime, err := time.ParseInLocation(arrAt.Layout, arrAt.Time, arrLoc)
	if err != nil {
		return 0, err
	}

	return int32(arrTime.Sub(depTime).Minutes()), nil
}

func (l *loc) GetAirport(code string) *Airport {
	return l.airportsMap[code]
}

/**
* Get city by airport code or city code
 */
func (l *loc) GetCity(code string) *City {
	city, ok := l.citiesMap[code]
	if ok {
		return city
	}

	airport := l.GetAirport(code)
	if airport != nil {
		return airport.City
	}

	return nil
}

func (l *loc) GetCountry(code string) *Country {
	country, ok := l.countriesMap[code]
	if ok {
		return country
	}

	city := l.GetCity(code)
	if city != nil {
		return city.Country
	}
	return nil
}

func (l *loc) GetArea(code string) *Area {
	area, ok := l.areasMap[code]
	if ok {
		return area
	}

	country := l.GetCountry(code)
	if country != nil {
		return country.Area
	}

	return nil
}

func (l *loc) getEncoding(code string) int64 {
	if e, ok := l.airportsEncoding[code]; ok {
		return e
	}
	if e, ok := l.citiesEncoding[code]; ok {
		return e
	}
	if e, ok := l.countriesEncoding[code]; ok {
		return e
	}
	if e, ok := l.areasEncoding[code]; ok {
		return e
	}

	return 0
}

func (l *loc) BelongsTo(code string, pattern string) bool {
	codeEncoding := l.getEncoding(code)
	patternEncoding := l.getEncoding(pattern)

	if codeEncoding == 0 || patternEncoding == 0 {
		return false
	}

	logger.Debug(logger.Message("Code[%s:%s], Pattern[%s:%s]", code, strconv.FormatInt(codeEncoding, 2), pattern, strconv.FormatInt(patternEncoding, 2)))

	return codeEncoding == patternEncoding ||
		(cityEncodingMask&codeEncoding) == patternEncoding ||
		(countryEncodingMask&codeEncoding) == patternEncoding ||
		(areaEncodingMask&codeEncoding) == patternEncoding
}

func GetAirport(code string) *Airport {
	return location.GetAirport(code)
}

func GetCity(code string) *City {
	return location.GetCity(code)
}

func GetNotNilCity(code string) *City {
	city := GetCity(code)
	if city == nil {
		return &City{
			Code: code,
			Country: &Country{
				Area: &Area{},
			},
		}
	}
	return city
}

func GetCountry(code string) *Country {
	return location.GetCountry(code)
}

func GetArea(code string) *Area {
	return location.GetArea(code)
}

func Duration(depCode, depTime, arrCode, arrTime, format string) int32 {
	var layout string
	if len(format) == 0 {
		layout = DepArrDateTimeLayout
	} else {
		layout = format
	}
	duration, err := location.Duration(LocationTime{
		LocationCode: depCode,
		Time:         depTime,
		Layout:       layout,
	}, LocationTime{
		LocationCode: arrCode,
		Time:         arrTime,
		Layout:       layout,
	})

	if err != nil && err != unknownCodeErr {
		logger.Warn(logger.NT, logger.Message("Failed to parse duration. Error: %v", err))
	}

	return duration
}

func DurationInSameTz(depTime, arrTime string) int32 {
	// Any valid airport code is ok to make this function call return the right duration
	dep := utils.ParseToYYYYMMDDHHMM(depTime)
	arr := utils.ParseToYYYYMMDDHHMM(arrTime)
	return int32(arr.Sub(dep).Minutes())
}

func BelongsTo(code string, pattern string) bool {
	if pattern == "*" {
		return true
	}
	return location.BelongsTo(code, pattern)
}

func GetWeight(code string) int32 {
	loc, _ := location.(*loc)
	if _, ok := loc.airportsEncoding[code]; ok {
		return AirportWeight
	}
	if _, ok := loc.citiesEncoding[code]; ok {
		return CityWeight
	}
	if _, ok := loc.countriesEncoding[code]; ok {
		return CountryWeight
	}
	if _, ok := loc.areasEncoding[code]; ok {
		return AreaWeight
	}

	return 0
}
