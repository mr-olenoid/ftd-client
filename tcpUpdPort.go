package ftdClient

import "fmt"

func (c *Client) GetTcpUdpPort(ID string, portType string) (*TcpUdpPort, error) {
	var tcpUdpPort TcpUdpPort
	var err error
	switch portType {
	case "TCP":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/tcpports/%s", ID), "GET", c)
	case "UDP":
		err = doFTDRequest(&tcpUdpPort, fmt.Sprintf("object/udpports/%s", ID), "GET", c)
	default:
		err = fmt.Errorf("expect TCP or UDP, got: %s", portType)
	}

	return &tcpUdpPort, err
}
