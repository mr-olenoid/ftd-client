package ftdClient

import "fmt"

func (c *Client) GetAccessRule(apId string, ID string) (*AccessRule, error) {
	ar := AccessRule{}
	ar.Type = "accessrule"
	err := doFTDRequest(&ar, fmt.Sprintf("policy/accesspolicies/%s/accessrules/%s", apId, ID), "GET", c)
	return &ar, err
}

func (c *Client) CreateAccessRule(apId string, ar AccessRule) (*AccessRule, error) {
	ar.Type = "accessrule"
	err := doFTDRequest(&ar, fmt.Sprintf("policy/accesspolicies/%s/accessrules", apId), "POST", c)
	return &ar, err
}

func (c *Client) UpdateAccessRule(apId string, ar AccessRule) (*AccessRule, error) {
	ar.Type = "accessrule"
	err := doFTDRequest(&ar, fmt.Sprintf("policy/accesspolicies/%s/accessrules/%s", apId, ar.ID), "PUT", c)
	return &ar, err
}

func (c *Client) DeleteAccessRule(apId string, ar AccessRule) error {
	ar.Type = "accessrule"
	err := doFTDRequest(&ar, fmt.Sprintf("policy/accesspolicies/%s/accessrules/%s", apId, ar.ID), "DELETE", c)
	return err
}
