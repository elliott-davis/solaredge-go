package solaredge

import (
	"fmt"
	"strings"
)


type SitesData struct {
	ID int64 `json:"id"`
	TimePeriodResponse
}

type SitesDataResponse struct {
	DataPeriod struct  {
		Count int64 `json:"count"`
		List []SitesData `json:"list"`
	} `json:"dataPeriod"`
}
func (s *SitesService) Data(siteIDS []int64) ([]SitesData, error) {
	IDS := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(siteIDS)), ","), "[]")
	req, err := s.client.NewRequest("GET", fmt.Sprintf("/site/%s/dataPeriod/", IDS), nil)
	if err != nil {
		return nil, err
	}
	var siteDataBulkResponse SitesDataResponse
	_, err = s.client.do(req, &siteDataBulkResponse)
	return siteDataBulkResponse.DataPeriod.List, err
}

