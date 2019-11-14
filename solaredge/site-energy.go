package solaredge

import (
	"fmt"
)

type SiteEnergyRequest struct {
	TimePeriodRequest
	TimeUnit string `json:"timeUnit"`
}

type SiteEnergyValue struct {
	Date DateTime `json:"date"`
	Value *float64 `json:"value"`
}

type SiteEnergyResponse struct {
	Energy struct {
		TimeUnit string `json:"timeUnit"`
		Unit string `json:"unit"`
		Values []SiteEnergyValue
	} `json:"energy"`
}

type TimeFrameEnergy struct {
	Energy float64 `json:"energy"`
	Unit string `json:"unit"`
}

type TimeFrameEnergyResponse struct {
	TimeFrameEnergy TimeFrameEnergy `json:"timeFrameEnergy"`
}

func (s *SiteService) Energy(siteId int64, energyOptions TimePeriodRequest) ([]SiteEnergyValue, error) {
	u, err := addOptions(fmt.Sprintf("/site/%d/energy/", siteId), energyOptions)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return []SiteEnergyValue{}, err
	}
	var siteEnergyResponse SiteEnergyResponse
	_, err = s.client.do(req, &siteEnergyResponse)
	return siteEnergyResponse.Energy.Values, err
}

func (s *SiteService) TimeFrameEnergy(siteId int64, energyOptions *SiteEnergyRequest) (TimeFrameEnergy, error) {
	u, err := addOptions(fmt.Sprintf("/site/%d/timeFrameEnergy/", siteId), energyOptions)
	if err != nil {
		return TimeFrameEnergy{}, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return TimeFrameEnergy{}, err
	}
	var timeFrameEnergyResponse TimeFrameEnergyResponse
	_, err = s.client.do(req, &timeFrameEnergyResponse)
	return timeFrameEnergyResponse.TimeFrameEnergy, err
}


