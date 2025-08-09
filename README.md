# purescript-goos

Cross-platform OS utilities for PureScript-native using Go's standard library.

## Installation

```yaml
# spago.yaml
dependencies:
  - goos: "https://github.com/i-am-the-slime/purescript-goos.git"
```

## Usage

```purescript
import Goos (readDir, readTextFile, writeTextFile, args, exit)

main = do
  -- Read command line arguments
  cmdArgs <- args
  log $ "Arguments: " <> show cmdArgs
  
  -- Read a directory
  result <- readDir "."
  case result of
    Left error -> log $ "Error: " <> error
    Right entries -> log $ "Found " <> show (length entries) <> " entries"
  
  -- Read/write text files
  writeResult <- writeTextFile "test.txt" "Hello, world!"
  content <- readTextFile "test.txt"
  case content of
    Left error -> log $ "Read error: " <> error
    Right text -> log $ "File content: " <> text
```

## API

### File System Operations

- `readDir :: String -> Effect (Either String (Array DirEntry))` - Read directory contents
- `fileName :: DirEntry -> String` - Get filename from directory entry
- `isDir :: DirEntry -> Boolean` - Check if entry is a directory
- `readTextFile :: String -> Effect (Either String String)` - Read text file
- `writeTextFile :: String -> String -> Effect (Either String Unit)` - Write text file

### Process Operations

- `args :: Effect (Array String)` - Get command line arguments
- `exit :: Int -> Effect a` - Exit with status code

## Cross-Platform Support

This library works on all platforms supported by Go:
- macOS
- Linux
- Windows
- And many others

## Requirements

- PureScript with purescript-native backend
- Go 1.24.3+