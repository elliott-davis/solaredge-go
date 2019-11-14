package solaredge

import "fmt"

type SiteDetailsResponse struct {
	Details Site `json:"details"`
}

func (s *SiteService) Details(siteId int64) (Site, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("/site/%d/details/",siteId), nil)
	if err != nil {
		return Site{}, err
	}
	var siteDetails SiteDetailsResponse
	_, err = s.client.do(req, &siteDetails)
	return siteDetails.Details, err
}
