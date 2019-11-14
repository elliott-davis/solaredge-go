package solaredge

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
	"time"
)

// SiteService handles communication with the site related
// methods of the SolarEdge API.
type SiteService service

type Location struct {
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
	Address string `json:"address"`
	Address2 string `json:"address2"`
	Zip string `json:"zip"`
	TimeZone string `json:"timeZone"`
}

type PublicSettings struct {
	Name string `json:"name"`
	IsPublic bool `json:"isPublic"`
}

type PrimaryModule struct {
	ManufacturerName string `json:"ManufacturerName"`
	ModelName string `json:"modelName"`
	MaximumPower float32 `json:"maximumPower"`
	TemperatureCoef float32 `json:"temperatureCoef"`
}

type Site struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	AccountID int64 `json:"accountId"`
	Status string `json:"status"` //might be an enum
	PeakPower float32 `json:"peakPower"`
	Currency string `json:"currency"`
	LastUpdateTime string `json:"lastUpdateTime"` // should be a time.Time probably - format: 2019-11-14
	InstallationDate string `json:"installationDate"` //should be a time.Time probably
	PTODate string `json:"ptoDate"`
	Notes string `json:"notes"`
	Type string `json:"type"`
	Location Location `json:"location"`
	AlertQuantity int64 `json:"alertQuantity"`
	AlertSeverity string `json:"alertSeverity"`
	PublicSettings PublicSettings `json:"publicSettings"`
	PrimaryModule PrimaryModule `json:"primaryModule"`
}

type TimePeriodRequest struct {
	StartDate YMDTime `json:"startDate" url:"startDate"`
	EndDate YMDTime `json:"endDate" url:"endDate"`
}

type TimePeriodResponse struct {
	StartDate time.Time `json:"startDate"`
	EndDate time.Time `json:"endDate"`
}

type DateTime struct {
	time.Time
}

type Meter int

const (
	Production Meter = iota
	Consumption
	SelfConsumption
	FeedIn
	Purchased
)

type TimeUnit int

const (
	QuarterOfAnHour TimeUnit = iota
	Hour
	Day
	Week
	Month
	Year
)

func (t *DateTime) UnmarshalJSON(b []byte) error {
	strInput := string(b)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006-01-02 15:04:05", strInput)
	if err != nil {
		return err
	}

	t.Time = newTime
	return nil
}

func (t DateTime) EncodeValues(key string, v *url.Values) error {
	v.Set(key, t.Format("2006-01-02 15:04:05"))
	return nil
}

type YMDTime time.Time

func (t YMDTime) EncodeValues(key string, v *url.Values) error {
	v.Set(key, (time.Time(t)).Format("2006-01-02"))
	return nil
}

func (m Meter) String() string {
	return [...]string{"Production", "Consumption", "SelfConsumption", "FeedIn", "Purchased"}[m]
}

var toMeterID = map[string]Meter {
	"Production": Production,
	"Consumption": Consumption,
	"SelfConsumption": SelfConsumption,
	"FeedIn": FeedIn,
	"Purchased": Purchased,
}

// MarshalJSON marshals the enum as a quoted json string
func (m Meter) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(m.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (m *Meter) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*m = toMeterID[j]
	return nil
}

func (t TimeUnit) String() string {
	return [...]string{"QUARTER_OF_AN_HOUR", "HOUR", "DAY", "WEEK", "MONTH", "YEAR"}[t]
}

var toTimeUnitID =  map[string]TimeUnit {
	"QUARTER_OF_AN_HOUR": QuarterOfAnHour,
	"HOUR": Hour,
	"DAY": Day,
	"WEEK": Week,
	"MONTH": Month,
	"YEAR": Year,
}

// MarshalJSON marshals the enum as a quoted json string
func (m TimeUnit) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(m.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (m *TimeUnit) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*m = toTimeUnitID[j]
	return nil
}
