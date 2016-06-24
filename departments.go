package liveagent

import "fmt"

type Department struct {
	Departmentid string `json:"departmentid"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Onlinestatus string `json:"onlinestatus"`
	Presetstatus string `json:"presetstatus"`
	Deleted      string `json:"deleted"`
}

type DepartmentOnlineState struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	OnlineStatus       string `json:"onlineStatus"`
	PresetStatus       string `json:"presetStatus"`
	ChatCount          string `json:"chat_count"`
	NewCount           string `json:"new_count"`
	CustomerReplyCount string `json:"customer_reply_count"`
	TotalCount         string `json:"total_count"`
	MaxCount           string `json:"max_count"`
}

type DepartmentsOnlineStatusResponse struct {
	Response struct {
		DepartmentsOnlineStates []DepartmentOnlineState `json:"departmentsOnlineStates"`
	} `json:"response"`
}

type DepartmentsResponse struct {
	Response struct {
		Departments []Department `json:"departments"`
	} `json:"response"`
}

type DepartmentResponse struct {
	Response Department `json:"response"`
}

func (c *Client) DepartmentsOnlineStatus() ([]DepartmentOnlineState, error) {
	var r DepartmentsOnlineStatusResponse
	err := c.get("onlinestatus/departments", nil, &r)
	if err != nil {
		return nil, err
	}
	return r.Response.DepartmentsOnlineStates, nil
}

func (c *Client) Departments() ([]Department, error) {
	var r DepartmentsResponse
	err := c.get("departments", nil, &r)
	if err != nil {
		return nil, err
	}
	return r.Response.Departments, nil
}

func (c *Client) Department(id string) (*Department, error) {
	var r DepartmentResponse
	err := c.get(fmt.Sprintf("departments/%s", id), nil, &r)
	if err != nil {
		return nil, err
	}
	return &(r.Response), nil
}
