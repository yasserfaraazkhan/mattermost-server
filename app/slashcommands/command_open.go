// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package slashcommands

import (
	"github.com/mattermost/mattermost-server/v5/app"
	"github.com/mattermost/mattermost-server/v5/corelibs/i18n"
	"github.com/mattermost/mattermost-server/v5/model"
)

type OpenProvider struct {
	JoinProvider
}

const (
	CmdOpen = "open"
)

func init() {
	app.RegisterCommandProvider(&OpenProvider{})
}

func (open *OpenProvider) GetTrigger() string {
	return CmdOpen
}

func (open *OpenProvider) GetCommand(a *app.App, T i18n.TranslateFunc) *model.Command {
	cmd := open.JoinProvider.GetCommand(a, T)
	cmd.Trigger = CmdOpen
	cmd.DisplayName = T("api.command_open.name")
	return cmd
}
