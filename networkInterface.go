package ftdClient

import (
	"fmt"
)

func (c *Client) GetNetworkInterface(ID string) (*NetworkInterface, error) {
	i := NetworkInterface{}
	err := doFTDRequest(&i, fmt.Sprintf("/devices/default/interfaces/%s", ID), "GET", c)
	return &i, err
}

func (c *Client) GetNetworkInterfaceSecurityZone(interfaceID string) (*ReferenceModel, error) {
	secz := ReferenceModel{}
	securityZones := Items[SecurityZone]{}
	doFTDRequest(&securityZones, "/object/securityzones", "GET", c)
	for _, sz := range securityZones.Items {
		for _, iface := range sz.Interfaces {
			if iface.ID == interfaceID {
				secz.ID = sz.ID
				secz.Name = sz.Name
				secz.Type = sz.Type
				// found security zone for this interface
				//fmt.Println(secz)
			}
		}
	}
	return &secz, nil
}

func (c *Client) CreateNetworkInterface(name string, securityZone *ReferenceModel) (*Items[NetworkInterface], *ReferenceModel, error) {
	i := Items[NetworkInterface]{}
	sz := &ReferenceModel{}
	err := doFTDRequest(&i, fmt.Sprintf("/devices/default/interfaces?filter=name:%s", name), "GET", c)
	if err != nil {
		return nil, nil, err
	}
	//fmt.Println(i.Items)

	if securityZone.ID == "" {
		sz, err = c.GetNetworkInterfaceSecurityZone(i.Items[0].ID)
		if err != nil {
			return nil, nil, err
		}
		fmt.Println(securityZone)
	}
	return &i, sz, err
}

func (c *Client) UpdateNetworkInterface(hardwareName string, securityZone ReferenceModel) (*NetworkInterface, error) {
	i := NetworkInterface{}

	err := doFTDRequest(&i, fmt.Sprintf("/devices/default/interfaces?filter=name:%s", hardwareName), "PUT", c)

	return &i, err
}
