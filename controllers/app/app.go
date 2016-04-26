package appcontroller

import (
	"fmt"
	"net/http"

	"github.com/CiscoCloud/shipped-utils/controllers"
	f "github.com/MustWin/gomarathon"
	"github.com/gorilla/mux"
)

func GetApp(w http.ResponseWriter, r *http.Request) {
	client, err := controllers.NewMarathonClient()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	status := http.StatusOK
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Input Id", id)
	if len(id) < 1 {
		http.Error(w, "Input param id is missing", http.StatusInternalServerError)
		return
	}
	resp, err := client.GetApp(id)
	if err != nil {
		fmt.Println("Invalid Details : " + err.Error())
		status = http.StatusServiceUnavailable
	}

	if response, er := ParseResponse(resp); er != nil {
		http.Error(w, fmt.Sprintf("App Error: %s for id %s", err.Error(), id), http.StatusInternalServerError)
	} else {
		controllers.ServeJsonResponseWithCode(w, response, status)
	}

}

func ParseResponse(resp *f.Response) (controllers.App, error) {
	var out controllers.App
	fmt.Println("Error ", resp)
	if resp == nil {
		return out, fmt.Errorf("Found Empty Response")
	}
	if resp.App == nil {
		return out, fmt.Errorf("No App is found")
	}

	out.Id = resp.App.ID
	if len(resp.App.Labels) > 0 {
		if pn, k := resp.App.Labels["project_name"]; k {
			out.ProjectName = pn
		}
		if pi, k := resp.App.Labels["project_id"]; k {
			out.ProjectId = pi
		}
		if en, k := resp.App.Labels["env_name"]; k {
			out.EnvName = en
		}
		if ei, k := resp.App.Labels["env_id"]; k {
			out.EnvId = ei
		}

		if sn, k := resp.App.Labels["service_name"]; k {
			out.ServiceName = sn
		}
		if si, k := resp.App.Labels["service_id"]; k {
			out.ServiceId = si
		}
	}

	if resp.App.Tasks != nil {
		var tasks []*controllers.Task
		for _, tsks := range resp.App.Tasks {
			task := new(controllers.Task)
			task.Host = tsks.Host
			task.Ports = tsks.Ports
			task.Id = tsks.ID
			task.AppId = tsks.AppID

			tasks = append(tasks, task)
		}
		if len(tasks) > 0 {
			out.Tasks = tasks
		}
	}
	return out, nil
}
