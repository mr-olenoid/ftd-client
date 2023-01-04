package ftdClient

import "fmt"

func (c *Client) GetApplicationFilter(ID string) (*ApplicationFilter, error) {
	var af ApplicationFilter
	err := doFTDRequest(&af, fmt.Sprintf("object/applicationfilters/%s", ID), "GET", c)
	return &af, err
}

func (c *Client) CreateApplicationFilter(af ApplicationFilter) (*ApplicationFilter, error) {
	err := doFTDRequest(&af, "object/applicationfilters", "POST", c)
	return &af, err
}

func (c *Client) UpdateApplicationFilter(af ApplicationFilter) (*ApplicationFilter, error) {
	err := doFTDRequest(&af, fmt.Sprintf("object/applicationfilters/%s", af.ID), "PUT", c)
	return &af, err
}

func (c *Client) DeleteApplicationFilter(af ApplicationFilter) (*ApplicationFilter, error) {
	err := doFTDRequest(&af, fmt.Sprintf("object/applicationfilters/%s", af.ID), "DELETE", c)
	return &af, err
}
