package solaredge

import (
	"fmt"
)

type SiteDataResponse struct {
	DataPeriod TimePeriodResponse `json:"dataPeriod"`
}

func (s *SiteService) Data(siteId int64) (TimePeriodResponse, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("/site/%d/dataPeriod/",siteId), nil)
	if err != nil {
		return TimePeriodResponse{}, err
	}
	var siteDataResponse SiteDataResponse
	_, err = s.client.do(req, &siteDataResponse)
	return siteDataResponse.DataPeriod, err
}
