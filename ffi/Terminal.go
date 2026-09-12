package goos

import (
	"os"
	"sync"

	. "github.com/purescript-native/go-runtime"
	"golang.org/x/term"
)

var terminal struct {
	sync.Mutex
	raw    *term.State
	reader sync.Once
	keys   chan []byte
}

func init() {
	exports := Foreign("Goos.Terminal")
	exports["isStdinTTY"] = func() Any { return term.IsTerminal(int(os.Stdin.Fd())) }
	exports["termSize"] = func() Any {
		cols, rows, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			cols, rows = 0, 0
		}
		return Dict{"cols": cols, "rows": rows}
	}
	exports["putStr"] = func(text Any) Any {
		return func() Any {
			_, err := os.Stdout.WriteString(text.(string))
			if err != nil {
				panic(err)
			}
			return nil
		}
	}
	exports["putErr"] = func(text Any) Any {
		return func() Any {
			_, err := os.Stderr.WriteString(text.(string))
			if err != nil {
				panic(err)
			}
			return nil
		}
	}
	exports["enterRawMode"] = func() Any {
		terminal.Lock()
		defer terminal.Unlock()
		if terminal.raw == nil {
			state, err := term.MakeRaw(int(os.Stdin.Fd()))
			if err == nil {
				terminal.raw = state
			}
		}
		terminal.reader.Do(func() {
			terminal.keys = make(chan []byte, 64)
			go func(keys chan []byte) {
				var buffer [8]byte
				for {
					n, err := os.Stdin.Read(buffer[:])
					if n > 0 {
						chunk := append([]byte(nil), buffer[:n]...)
						select {
						case keys <- chunk:
						default:
						}
					}
					if err != nil || n == 0 {
						return
					}
				}
			}(terminal.keys)
		})
		return nil
	}
	exports["restoreRawMode"] = func() Any {
		terminal.Lock()
		defer terminal.Unlock()
		if terminal.raw != nil {
			_ = term.Restore(int(os.Stdin.Fd()), terminal.raw)
			terminal.raw = nil
		}
		return nil
	}
	exports["readKeyBytes"] = func() Any {
		terminal.Lock()
		keys := terminal.keys
		terminal.Unlock()
		select {
		case chunk := <-keys:
			out := make([]Any, len(chunk))
			for i, b := range chunk {
				out[i] = int(b)
			}
			return out
		default:
			return []Any{}
		}
	}
}
