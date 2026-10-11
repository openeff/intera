package docker

import (
	"context"
	"openeff/intera/internal/sandbox/backends"
)

type DockerTerminal struct {
}

func (t *DockerTerminal) HasTerminal() bool
func (t *DockerTerminal) Exec(ctx context.Context, cmd backends.Command) (result *backends.CommandResult, err error)
func (t *DockerTerminal) WriteStdin(ctx context.Context, sessionID string, data []byte) (err error)
func (t *DockerTerminal) ReadOutput(ctx context.Context, sessionID string) (data []byte, err error)
func (t *DockerTerminal) Resize(ctx context.Context, sessionID string, cols, rows int) (err error)
func (t *DockerTerminal) CloseSession(ctx context.Context, sessionID string) (err error)

var _ backends.Terminal = (*DockerTerminal)(nil)
