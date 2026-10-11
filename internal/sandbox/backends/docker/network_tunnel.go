package docker

import (
	"context"
	"openeff/intera/internal/sandbox/backends"
)

type DockerNetworkTunnel struct {
}

func (n *DockerNetworkTunnel) Connect(ctx context.Context, cfg backends.TunnelConfig) (err error)
func (n *DockerNetworkTunnel) Disconnect(ctx context.Context) (err error)
func (n *DockerNetworkTunnel) Status(ctx context.Context) (status backends.TunnelStatus, err error)

var _ backends.NetworkTunnel = (*DockerNetworkTunnel)(nil)
