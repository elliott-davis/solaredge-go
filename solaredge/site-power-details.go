package solaredge

import (
	"errors"
	"fmt"
)

type Meters struct {
	Type Meter `json:"type"`
	Values []SitePowerDetailValues `json:"values"`
}

type SitePowerDetailsRequest struct {
	StartTime DateTime `url:"startTime"`
	EndTime DateTime `url:"endTime"`
	Meters []Meter `url:"meter"`
}

type SitePowerDetails struct {
	TimeUnit string `json:"timeUnit"`
	Unit string `json:"unit"`
	Meters []Meters `json:"meters"`
}

type SitePowerDetailsResponse struct {
	PowerDetails SitePowerDetails `json:"powerDetails"`
}

func  (s *SiteService) PowerDetails(siteId int64, request SitePowerDetailsRequest) (SitePowerDetails, error) {
	// Ensure start and end are defined
	if request.EndTime.IsZero() || request.StartTime.IsZero() {
		return SitePowerDetails{}, errors.New("start and End times are required")
	}

	u, err := addOptions(fmt.Sprintf("/site/%d/powerDetails/", siteId), request)
	if err != nil {
		return SitePowerDetails{}, err
	}
	// Ensure delta between start and end is one month or less
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return SitePowerDetails{}, err
	}
	var sitePowerDetailsResponse SitePowerDetailsResponse
	_, err = s.client.do(req, &sitePowerDetailsResponse)
	return sitePowerDetailsResponse.PowerDetails, err
}
