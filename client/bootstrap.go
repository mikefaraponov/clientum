package client

import (
	"context"
	"github.com/marcusolsson/tui-go"
	"github.com/mikefaraponov/clientum/common"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"syscall"
)

func Harakiri() {
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}

func Bootstrap(lc fx.Lifecycle, conn *grpc.ClientConn, ui tui.UI) {
	ui.SetKeybinding(common.Esc, Harakiri)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go ui.Run()
			return nil
		},
		OnStop: func(context.Context) error {
			ui.Quit()
			return conn.Close()
		},
	})
}
