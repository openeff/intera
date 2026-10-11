package backends

import "context"

type BackendSandbox interface {
	Run(ctx context.Context, cfg RunConfig) error
	Restart(ctx context.Context) error
	Stop(ctx context.Context) error
	Close(ctx context.Context) error

	Observer
	Mouse
	Keyboard
	NetworkTunnel
	Terminal
}

type RunConfig struct {
	Executable string
	Args       []string
	WorkDir    string
	Env        map[string]string
}

type Observer interface {
	Screenshot
	Stream
}

type Screenshot interface {
	Screenshot(ctx context.Context) ([]byte, error)
	ScreenSize(ctx context.Context) (int, int, error)
}

type Stream interface {
	StartStream(ctx context.Context) error
	StopStream(ctx context.Context) error
	StreamURL() string
}

type Mouse interface {
	Click(ctx context.Context, p Point) error
	DoubleClick(ctx context.Context, p Point, button MouseButton) error
	MoveMouse(ctx context.Context, p Point) error
	MouseDown(ctx context.Context, button MouseButton) error
	MouseUp(ctx context.Context, button MouseButton) error
	Scroll(ctx context.Context, p Point, dx, dy int) error
}

type Point struct{ X, Y int }

type MouseButton int

const (
	LeftButton MouseButton = iota + 1
	MiddleButton
	RightButton
)

type Keyboard interface {
	PressKey(ctx context.Context, key string) error
	KeyDown(ctx context.Context, key string) error
	KeyUp(ctx context.Context, key string) error
	TypeText(ctx context.Context, text string) error
	Hotkey(ctx context.Context, keys ...string) error
}

type NetworkTunnel interface {
	Connect(ctx context.Context, cfg TunnelConfig) error
	Disconnect(ctx context.Context) error
	Status(ctx context.Context) (TunnelStatus, error)
}

type TunnelConfig struct {
	ConfigPath string
	Interface  string
	DNS        []string
}

type TunnelStatus struct {
	Connected bool
	LocalIP   string
}

type DesktopConfig struct {
	Display       string
	Width         int
	Height        int
	ColorDepth    int
	WindowManager string
}

type Terminal interface {
	HasTerminal() bool

	Exec(ctx context.Context, cmd Command) (*CommandResult, error)
	WriteStdin(ctx context.Context, sessionID string, data []byte) error
	ReadOutput(ctx context.Context, sessionID string) ([]byte, error)
	Resize(ctx context.Context, sessionID string, cols, rows int) error
	CloseSession(ctx context.Context, sessionID string) error
}

type Command struct {
	Args    []string
	WorkDir string
	Env     map[string]string
}

type CommandResult struct {
	SessionID string
	ExitCode  int
	Stdout    []byte
	Stderr    []byte
	Running   bool
}
