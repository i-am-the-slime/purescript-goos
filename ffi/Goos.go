package goos

import (
	"io"
	"os"

	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Goos")

	exports["readDirImpl"] = func(left_ Any, right_ Any, path_ Any) Any {
		path := path_.(string)
		res, err := os.ReadDir(path)
		if err != nil {
			return Apply(left_, err.Error())
		}
		// Convert []fs.DirEntry to []Any for PureScript
		var result []Any
		for _, entry := range res {
			result = append(result, entry)
		}
		return Apply(right_, result)
	}

	exports["fileName"] = func(dirEntry_ Any) Any {
		dirEntry := dirEntry_.(os.DirEntry)
		return dirEntry.Name()
	}

	exports["isDir"] = func(dirEntry_ Any) Any {
		dirEntry := dirEntry_.(os.DirEntry)
		return dirEntry.IsDir()
	}

	exports["getFileHandleImpl"] = func(left_ Any, right_ Any, path_ Any) Any {
		path := path_.(string)
		res, err := os.Open(path)
		if err != nil {
			return Apply(left_, err.Error())
		}
		return Apply(right_, res)
	}

	exports["closeFileImpl"] = func(left_ Any, right_ Any, file_ Any) Any {
		file := file_.(*os.File)
		err := file.Close()
		if err != nil {
			return Apply(left_, err.Error())
		}
		return Apply(right_, nil)
	}

	exports["readFileImpl"] = func(left_ Any, right_ Any, file_ Any) Any {
		file := file_.(*os.File)
		res, err := io.ReadAll(file)
		if err != nil {
			return Apply(left_, err.Error())
		}
		return Apply(right_, res)
	}

	exports["writeFileImpl"] = func(left_ Any, right_ Any, file_ Any, data_ Any) Any {
		file := file_.(*os.File)
		data := data_.([]byte)
		_, err := file.Write(data)
		if err != nil {
			return Apply(left_, err.Error())
		}
		return Apply(right_, nil)
	}

	// Convenience functions for text files
	exports["writeTextFileImpl"] = func(left_ Any, right_ Any, path_ Any, content_ Any) Any {
		path := path_.(string)
		content := content_.(string)
		err := os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			return Apply(left_, err.Error())
		}
		return Apply(right_, nil)
	}

	exports["readTextFileImpl"] = func(left_ Any, right_ Any, path_ Any) Any {
		path := path_.(string)
		content, err := os.ReadFile(path)
		if err != nil {
			return Apply(left_, err.Error())
		}
		return Apply(right_, string(content))
	}

	exports["args"] = func() Any {
		var returnValue []Any
		for _, arg := range os.Args {
			returnValue = append(returnValue, arg)
		}
		return returnValue
	}

	exports["exitImpl"] = func(code Any) Any {
		os.Exit(code.(int))
		return nil
	}

	exports["stGlobalToEffect"] = func(stGlobal Any) Any {
		return func() Any {
			return stGlobal
		}
	}
}
