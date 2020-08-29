package models

type UsersTicket struct {
	TicketId    uint64 `json:"ticketid"`
	Username    string `json:"username"`
	Phonenumber string `json:"phonenumber"`
	Timings     string `json:"timings"`
}
