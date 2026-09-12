package goos

import (
	"time"

	"github.com/fsnotify/fsnotify"
	. "github.com/purescript-native/go-runtime"
)

func watchEvent() Dict {
	return Dict{"name": "", "write": false, "create": false, "rename": false, "remove": false, "chmod": false, "closed": false, "timeout": false, "error": ""}
}

func init() {
	exports := Foreign("Goos.Watch")
	exports["open"] = func(path Any) Any {
		return func() Any {
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				panic(err)
			}
			if err := watcher.Add(path.(string)); err != nil {
				watcher.Close()
				panic(err)
			}
			return watcher
		}
	}
	exports["close"] = func(value Any) Any { return func() Any { _ = value.(*fsnotify.Watcher).Close(); return nil } }
	exports["next"] = func(value Any) Any {
		return func(milliseconds Any) Any {
			return func() Any {
				watcher := value.(*fsnotify.Watcher)
				var timer *time.Timer
				var elapsed <-chan time.Time
				if milliseconds.(int) >= 0 {
					timer = time.NewTimer(time.Duration(milliseconds.(int)) * time.Millisecond)
					defer timer.Stop()
					elapsed = timer.C
				}
				result := watchEvent()
				select {
				case event, ok := <-watcher.Events:
					result["closed"] = !ok
					if ok {
						result["name"] = event.Name
						result["write"] = event.Has(fsnotify.Write)
						result["create"] = event.Has(fsnotify.Create)
						result["rename"] = event.Has(fsnotify.Rename)
						result["remove"] = event.Has(fsnotify.Remove)
						result["chmod"] = event.Has(fsnotify.Chmod)
					}
				case err, ok := <-watcher.Errors:
					result["closed"] = !ok
					if ok {
						result["error"] = err.Error()
					}
				case <-elapsed:
					result["timeout"] = true
				}
				return result
			}
		}
	}
}
