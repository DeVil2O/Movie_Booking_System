package models

import "time"

type Admin struct {
	AdminId  string    `json:"adminid"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Tickets  []*Ticket `json:"tickets"`
	Token    string    `json:"token"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

type Ticket struct {
	TicketId    uint64    `json:"ticketid"`
	Username    string    `json:"username"`
	Phonenumber string    `json:"phonenumber"`
	StartTime   time.Time `json:"starttime"`
	EndTime     time.Time `json:endtime`
	CreatedAt   time.Time `json:createdat`
	Expired     bool      `json:expired`
}
