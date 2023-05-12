package tun_tap

import (
	"bytes"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSession(t *testing.T) {
	var b bytes.Buffer
	r := strings.NewReader("hello world")
	io.Copy(&b, r)
	slog.Info(b.String())
	r = strings.NewReader("fuck you")
	io.Copy(&b, r)
	slog.Info(b.String())

}

func TestPing(m *testing.T) {
	main()
}

func init() {
	var programLevel = new(slog.LevelVar) // Info by default
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)
}
