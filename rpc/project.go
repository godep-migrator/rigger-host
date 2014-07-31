package rpc

import (
	"net/http"

	"github.com/rigger-dot-io/rigger-host/models"
)

type ProjectCreateArgs struct {
	Name  string
	Owner string
	Host  string
	SCM   string
	URL   string
}

type ProjectReply struct {
	Id    string
	Error string
}

type ProjectService struct {
}

func (p *ProjectService) Create(r *http.Request, args *ProjectCreateArgs, reply *ProjectReply) error {
	project, err := models.NewProject(args.Name, args.Owner, args.Host, args.SCM, args.URL)
	if err != nil {
		return err
	}
	reply.Id = project.Id
	return nil
}
