package ftdClient

import (
	"fmt"
	"net/url"
	"strings"
)

func (c *Client) GetApplication(name string) (*[]Application, error) {
	applications := Items[Application]{}
	var apps []Application
	var err error
	if name == "*" {
		err = doFTDRequest(&applications, "object/applications?limit=500", "GET", c)
		apps = append(apps, applications.Items...)
		for len(applications.Paging.Next) > 0 {
			link := strings.Split(applications.Paging.Next[0], "?")[1]
			err = doFTDRequest(&applications, fmt.Sprintf("object/applications?%s", link), "GET", c)
			apps = append(apps, applications.Items...)
		}
	} else {
		err = doFTDRequest(&applications, fmt.Sprintf("object/applications?filter=name:%s", url.QueryEscape(name)), "GET", c)
		apps = applications.Items
	}

	return &apps, err
}

func (c *Client) GetApplicationCategory(name string) (*ApplicationCategory, error) {
	applicationcategories := Items[ApplicationCategory]{}
	err := doFTDRequest(&applicationcategories, fmt.Sprintf("object/applicationcategories?filter=name:%s", url.QueryEscape(name)), "GET", c)
	if len(applicationcategories.Items) == 1 {
		appCategory := applicationcategories.Items[0]
		return &appCategory, err
	}
	return nil, err
}
