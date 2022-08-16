package solaredge

type SiteListResponse struct {
	Sites struct {
		Count int64  `json:"count"`
		List  []Site `json:"site"`
	}
}

func (s *SiteService) List(listOptions *ListOptions) ([]Site, error) {
	u, err := addOptions("/sites/list/", listOptions)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	var siteList SiteListResponse
	_, err = s.client.do(req, &siteList)
	return siteList.Sites.List, err
}
