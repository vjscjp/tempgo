package controllers

type App struct {
	Id          string  `json:"id"`
	ProjectName string  `json:"projectname"`
	ProjectId   string  `json:"projectid"`
	EnvName     string  `json:"envname"`
	EnvId       string  `json:"envid"`
	ServiceName string  `json:"servicename"`
	ServiceId   string  `json:"serviceid"`
	Tasks       []*Task `json:"tasks,omitempty"`
}

type Task struct {
	Host    string `json:"host"`
	Ports   []int  `json:"ports"`
	Id      string `json:"id"`
	AppId   string `json:"appid"`
	
}
