package docker

import (
	"context"
	"openeff/intera/internal/sandbox/backends"
)

type DokerKeyboard struct {
}

func (k *DokerKeyboard) PressKey(ctx context.Context, key string) (err error)
func (k *DokerKeyboard) KeyDown(ctx context.Context, key string) (err error)
func (k *DokerKeyboard) KeyUp(ctx context.Context, key string) (err error)
func (k *DokerKeyboard) TypeText(ctx context.Context, text string) (err error)
func (k *DokerKeyboard) Hotkey(ctx context.Context, keys ...string) (err error)

var _ backends.Keyboard = (*DokerKeyboard)(nil)
