package solaredge

import "fmt"

type OverviewData struct {
	Energy float64 `json:"energy"`
	Revenue  float32  `json:"revenue"`
}

type  SiteOverview struct {
	LifetimeData OverviewData `json:"lifetimeData"`
	LastYearData OverviewData `json:"lastYearData"`
	LastMonthData OverviewData `json:"lastMonthData"`
	LastDayData OverviewData `json:"lastYearData"`
	CurrentPower struct {
		Power float64 `json:"power"`
	} `json:"currentPower"`
}

type SiteOvewrviewResponse struct {
	Overview SiteOverview `json:"overview"`
}

func (s *SiteService) Overview(siteId int64) (SiteOverview, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("/site/%d/overview/",siteId), nil)
	if err != nil {
		return SiteOverview{}, err
	}
	var siteOverview SiteOvewrviewResponse
	_, err = s.client.do(req, &siteOverview)
	return siteOverview.Overview, err
}
