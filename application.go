package ftdClient

import "fmt"

func (c *Client) GetApplication(name string) (*[]Application, error) {
	applications := Items[Application]{}
	var err error
	if name == "*" {
		err = doFTDRequest(&applications, "object/applications/", "GET", c)
	} else {
		err = doFTDRequest(&applications, fmt.Sprintf("object/applications?filter=name:%s", name), "GET", c)
	}

	return &applications.Items, err
}
