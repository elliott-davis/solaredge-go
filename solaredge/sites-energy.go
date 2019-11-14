package solaredge

import (
	"fmt"
	"strings"
)

type SitesEnergy struct {
	TimeUnit string `json:"timeUnit"`
	Unit string `json:"unit"`
	Count int64 `json:"count"`
	List []struct {
		ID int64 `json:"id"`
		Values []SiteEnergyValue
	}
}

type SitesEnergyResponse struct {
	Energy SitesEnergy `json:"energy"`
}

func (s *SitesService) Energy(siteIDS []int64, energyOptions *SiteEnergyRequest) (SitesEnergy, error) {
	IDS := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(siteIDS)), ","), "[]")
	u, err := addOptions(fmt.Sprintf("/site/%s/energy/", IDS), energyOptions)
	if err != nil {
		return SitesEnergy{}, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return SitesEnergy{}, err
	}
	var siteEnergyResponse SitesEnergyResponse
	_, err = s.client.do(req, &siteEnergyResponse)
	return siteEnergyResponse.Energy, err
}

// This is broken
func (s *SitesService) TimeFrameEnergy(siteIDS []int64, energyOptions *SiteEnergyRequest) (SitesEnergy, error) {
	IDS := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(siteIDS)), ","), "[]")
	u, err := addOptions(fmt.Sprintf("/sites/%s/timeFrameEnergy/", IDS), energyOptions)
	if err != nil {
		return SitesEnergy{}, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return SitesEnergy{}, err
	}
	var siteEnergyResponse SitesEnergyResponse
	_, err = s.client.do(req, &siteEnergyResponse)
	return siteEnergyResponse.Energy, err
}
