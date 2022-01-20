package mocks

import "github.com/Shivansh-yadva13/todo_react_go/api/models"

var TodoList = []models.Todos{
	{
		Id:     1,
		Title:  "Complete Networking DevOps",
		Desc:   "CCNA Part 1 Netwroking, router, NiC, switches, LAN, WAN.",
		Status: true,
	},
	{
		Id:     2,
		Title:  "React + Go Application",
		Desc:   "Build Go API and React App and fetch from go api.",
		Status: false,
	},
}
