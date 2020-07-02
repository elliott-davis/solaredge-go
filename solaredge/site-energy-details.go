package solaredge

import (
	"errors"
	"fmt"
)

type SiteEnergyDetailsRequest struct {
	StartTime DateTime `url:"startTime"`
	EndTime DateTime `url:"endTime"`
	Meters []Meter `url:"meters"`
	TimeUnit *TimeUnit `url:"timeUnit"` // Should probably be an enum
}

type SiteEnergyDetails struct {
	TimeUnit string `json:"timeUnit"`
	Unit string `json:"unit"`
	Meters []Meters `json:"meters"`
}

type SiteEnergyDetailsResponse struct {
	EnergyDetails SiteEnergyDetails `json:"energyDetails"`
}

func (s *SiteService) EnergyDetails(siteId int64, request SiteEnergyDetailsRequest) (SiteEnergyDetails,  error) {
	// Ensure start and end are defined
	if request.EndTime.IsZero() || request.StartTime.IsZero() {
		return SiteEnergyDetails{}, errors.New("start and End times are required")
	}

	if request.TimeUnit == nil {
		timeUnit := Day
		request.TimeUnit = &timeUnit
	}

	u, err := addOptions(fmt.Sprintf("/site/%d/energyDetails/", siteId), request)
	if err != nil {
		return SiteEnergyDetails{}, err
	}
	// Ensure delta between start and end is one month or less
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return SiteEnergyDetails{}, err
	}
	var siteEnergyDetailsResponse SiteEnergyDetailsResponse
	_, err = s.client.do(req, &siteEnergyDetailsResponse)
	return siteEnergyDetailsResponse.EnergyDetails, err
}
