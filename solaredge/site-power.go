package solaredge

import (
	"errors"
	"fmt"
	"time"
)

type SitePowerDetailValues struct {
	Date DateTime `json:"date"`
	Value *float32 `json:"value"`
}

type SitePowerResponse struct {
	Power SitePower `json:"power"`
}

type SitePower struct {
	TimeUnit string `json:"timeUnit"`
	Unit string `json:"unit"`
	Values []SitePowerDetailValues
}

type SitePowerRequest struct {
	StartTime DateTime `url:"startTime"`
	EndTime DateTime `url:"endTime"`
}

func (s *SiteService) Power(siteId int64, request SitePowerRequest) (SitePower, error) {
	// Ensure start and end are defined
	if request.EndTime.IsZero() || request.StartTime.IsZero() {
		return SitePower{}, errors.New("start and End times are required")
	}
	if request.EndTime.Sub(request.StartTime.Time) > time.Hour * 24 *  31 {
		return SitePower{}, errors.New("duration between StartDate and EndDate should be less than one month")
	}

	u, err := addOptions(fmt.Sprintf("/site/%d/power/", siteId), request)
	if err != nil {
		return SitePower{}, err
	}
	// Ensure delta between start and end is one month or less
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return SitePower{}, err
	}
	var sitePowerResponse SitePowerResponse
	_, err = s.client.do(req, &sitePowerResponse)
	return sitePowerResponse.Power, err
}
