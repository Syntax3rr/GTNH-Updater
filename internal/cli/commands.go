package cli

import (
	"github.com/Syntax3rr/GTNH-Updater/internal/core"

	"github.com/alecthomas/kong"
)

type CLI struct {
	Version kong.VersionFlag `short:"v" help:"Print version information and quit"`

	Install InstallCmd `cmd:"" help:"Install a specific version of the modpack."`
	Update UpdateCmd `cmd:"" help:"Update the game to the latest version."`
}

type InstallCmd struct {
	Path string `short:"p" help:"The path to the modloader directory."`
	Target string `short:"t" help:"The target version to install."`
}

type UpdateCmd struct {
	Force bool `short:"f" help:"Force update even if the version is up to date."`
	Channel string `short:"c" help:"The channel to update to (e.g. stable, beta, RC). Default is stable."`
	Path string `short:"p" help:"The path to the modloader directory."`
	Instance string `short:"i" help:"The instance to update (e.g. GTNH, GTNHD, GTNH-Dev)."`
}


func (i *InstallCmd) Run() error {
	args := core.InstallArgs{
		Path: i.Path,
		Target: i.Target,
	}
	return args.Install()
	return nil
}

func (u *UpdateCmd) Run() error {
	args := core.UpdateArgs{
		Force: u.Force,
		Channel: u.Channel,
		Path: u.Path,
		Instance: u.Instance,
	}
	return args.Update()
	return nil
}


func Run(args []string) int {
	var cli CLI

	ctx := kong.Parse(&cli,
		kong.Name("gtnh-updater"),
		kong.Description("An unofficial tool to update the GTNH modpack."),
		kong.UsageOnError(),
		kong.Vars{
			"version": "0.0.1",
		},
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}))

	err := ctx.Run(&cli)
	ctx.FatalIfErrorf(err)

	return 0
}