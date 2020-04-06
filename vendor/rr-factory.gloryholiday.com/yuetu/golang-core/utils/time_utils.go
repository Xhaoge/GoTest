package utils

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/noaway/dateparse"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
)

const (
	DateTimeFormat          string = "2006-01-02T15:04:05"
	DateTimeFormatSpace_HMS string = "2006-01-02 15:04:05"
	DateTimeFormatSpace     string = "2006-01-02 15:04"
	DateFormat              string = "2006-01-02"
	ApiDepArrDateFormat     string = "20060102"
	ApiDepArrDateTimeFormat string = "200601021504"
)

func ParseFromGoogleTimestamp(gts *timestamp.Timestamp) time.Time {
	return time.Unix(gts.GetSeconds(), int64(gts.GetNanos()))
}
func FormatToDateTimeFormatSpace(time time.Time) string {
	return time.Format(DateTimeFormatSpace)
}

func FormatToDateFormat(time time.Time) string {
	return time.Format(DateFormat)
}

func FormatGoogleTimestamp(format string, gts *timestamp.Timestamp) string {
	return ParseFromGoogleTimestamp(gts).Format(format)
}

func FormatTime(originalTime string, targetFormat string) (string, error) {
	parseR, e := dateparse.ParseAny(originalTime)
	if e != nil {
		return "", e
	}
	return parseR.Format(targetFormat), nil
}

func ParseApiDepArrDate(date string) time.Time {
	return ParseDate(date, ApiDepArrDateFormat)
}

func ParseDate(date, format string) time.Time {
	t, err := time.Parse(format, date)
	if err != nil {
		logger.ErrorNt(logger.Message("invalid api dep/arr date:%s", date), err)
		return time.Time{}
	}
	return t
}

func ParseToYYYYMMDDHHMM(date string) time.Time {
	return ParseDate(date, ApiDepArrDateTimeFormat)
}

func ParseToYYYYMMDD(date string) time.Time {
	return ParseApiDepArrDate(date)
}

func ParseStringToTime(stringTime string) (*time.Time, error) {
	parseR, e := dateparse.ParseAny(stringTime)
	if e != nil {
		return nil, e
	}
	return &parseR, nil
}
func ParseStringToTimeWithoutError(stringTime string) *time.Time {
	toTime, _ := ParseStringToTime(stringTime)
	return toTime
}

func ParseToGoogleTimestamp(t time.Time) *timestamp.Timestamp {
	local, _ := time.LoadLocation("Local")
	t.In(local)
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

// Return the millis duration from some specific time
func DurationFrom(from time.Time) int64 {
	return DurationBetween(from, time.Now())
}

func DurationBetween(from time.Time, to time.Time) int64 {
	return to.Sub(from).Nanoseconds() / int64(time.Millisecond)
}

type timeDuration struct {
	start time.Time
	stop  time.Time
}

func (td *timeDuration) String() string {
	if td.start.IsZero() {
		return ""
	}
	stop := td.stop
	if stop.IsZero() {
		stop = time.Now()
	}
	return fmt.Sprintf("%d ms", stop.Sub(td.start).Nanoseconds()/int64(time.Millisecond))
}

type TimeTracker struct {
	bucket map[string]*timeDuration
}

func NewTimeTracker() *TimeTracker {
	return &TimeTracker{
		bucket: map[string]*timeDuration{},
	}
}

func (tt *TimeTracker) String() string {
	lst := []string{}
	for operation, td := range tt.bucket {
		lst = append(lst, operation+": "+td.String())
	}

	sort.Strings(lst)
	return strings.Join(lst, ",")
}

func (tt *TimeTracker) Start(operation string) {
	_, ok := tt.bucket[operation]
	if !ok {

		tt.bucket[operation] = &timeDuration{
			start: time.Now(),
		}
	} else {
		logger.WarnNt(logger.Message("%s has been started by TimeTracker", operation))
	}
}

func (tt *TimeTracker) Stop(operation string) {
	td, ok := tt.bucket[operation]
	if !ok {
		logger.WarnNt(logger.Message("%s is not started by TimeTracker", operation))
		return
	}
	td.stop = time.Now()
}

func MilliSecondsToTime(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}

func MilliSecondsSince(t time.Time) float64 {
	return float64(time.Now().Sub(t).Nanoseconds() / 1e6)
}

type TraceTimer struct {
	StartAt     time.Time
	StopAt      time.Time
	Application string
	Operation   string
}

func NewTraceTimer(application string, operation string) *TraceTimer {
	return &TraceTimer{
		Application: application,
		Operation:   operation,
	}
}

func (t *TraceTimer) Start() *TraceTimer {
	t.StartAt = time.Now()
	return t
}

func (t *TraceTimer) Stop() *TraceTimer {
	t.StopAt = time.Now()
	return t
}

func (t *TraceTimer) DurationAsMilliSeconds() int64 {
	return int64(t.StopAt.Sub(t.StartAt) / time.Millisecond)
}
