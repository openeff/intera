package docker

import (
	"context"
	"openeff/intera/internal/sandbox/backends"
)

type DockerMouse struct {
}

func (m *DockerMouse) Click(ctx context.Context, p backends.Point) (err error)
func (m *DockerMouse) DoubleClick(ctx context.Context, p backends.Point, button backends.MouseButton) (err error)
func (m *DockerMouse) MoveMouse(ctx context.Context, p backends.Point) (err error)
func (m *DockerMouse) MouseDown(ctx context.Context, button backends.MouseButton) (err error)
func (m *DockerMouse) MouseUp(ctx context.Context, button backends.MouseButton) (err error)
func (m *DockerMouse) Scroll(ctx context.Context, p backends.Point, dx, dy int) (err error)

var _ backends.Mouse = (*DockerMouse)(nil)
