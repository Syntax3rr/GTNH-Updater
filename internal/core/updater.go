package core

import (
	"fmt"
)

type UpdateArgs struct {
	Force bool
	Channel string
	Path string
	Instance string
}

func (args *UpdateArgs) Update() error {
	fmt.Println("Updating", args.Instance, "to", args.Channel, "on", args.Path)
	return nil
}

type InstallArgs struct {
	Path string
	Target string
}	

func (args *InstallArgs) Install() error {
	fmt.Println("Installing", args.Target, "to", args.Path)
	return nil
}