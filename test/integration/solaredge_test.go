package integration

import (
	"github.com/elliott-davis/solaredge-go/solaredge"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

var (
	client *solaredge.Client
	siteID int64
)

func init() {
	token := os.Getenv("SOLAREDGE_AUTH_TOKEN")
	id, err := strconv.ParseInt(os.Getenv("SOLAREDGE_SITE_ID"), 10, 64)
	if err != nil {
		panic("Can't get site ID")
	}
	siteID = id
	client = solaredge.NewClient(nil, token)
}

func TestSitesList(t *testing.T) {
	site, err := client.Site.List(&solaredge.ListOptions{Page: 2, PerPage: 1})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, site, "", "dummy check")
}

func TestSiteDetails(t *testing.T) {
	site, err := client.Site.Details(siteID)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, site, "", "dummy check")
}

