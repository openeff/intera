package backends

import "context"

type BackendSandbox interface {
	Run(ctx context.Context, cfg RunConfig) (err error)
	Restart(ctx context.Context) (err error)
	Stop(ctx context.Context) (err error)
	Close(ctx context.Context) (err error)

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
	Screenshot(ctx context.Context) (data []byte, err error)
	ScreenSize(ctx context.Context) (width int, height int, err error)
}

type Stream interface {
	Start(ctx context.Context) (link string, err error)
	Stop(ctx context.Context) (err error)
	StreamURL() (link string)
}

type Mouse interface {
	Click(ctx context.Context, p Point) (err error)
	DoubleClick(ctx context.Context, p Point, button MouseButton) (err error)
	MoveMouse(ctx context.Context, p Point) (err error)
	MouseDown(ctx context.Context, button MouseButton) (err error)
	MouseUp(ctx context.Context, button MouseButton) (err error)
	Scroll(ctx context.Context, p Point, dx, dy int) (err error)
}

type Point struct{ X, Y int }

type MouseButton uint8

const (
	LeftButton MouseButton = iota + 1
	MiddleButton
	RightButton
)

type Keyboard interface {
	PressKey(ctx context.Context, key string) (err error)
	KeyDown(ctx context.Context, key string) (err error)
	KeyUp(ctx context.Context, key string) (err error)
	TypeText(ctx context.Context, text string) (err error)
	Hotkey(ctx context.Context, keys ...string) (err error)
}

type NetworkTunnel interface {
	Connect(ctx context.Context, cfg TunnelConfig) (err error)
	Disconnect(ctx context.Context) (err error)
	Status(ctx context.Context) (status TunnelStatus, err error)
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

	Exec(ctx context.Context, cmd Command) (result *CommandResult, err error)
	WriteStdin(ctx context.Context, sessionID string, data []byte) (err error)
	ReadOutput(ctx context.Context, sessionID string) (data []byte, err error)
	Resize(ctx context.Context, sessionID string, cols, rows int) (err error)
	CloseSession(ctx context.Context, sessionID string) (err error)
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
