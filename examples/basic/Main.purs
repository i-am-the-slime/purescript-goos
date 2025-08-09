module Main where

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Goos (args, readDir, readTextFile, writeTextFile, fileName)
import Data.Array (length)
import Data.Either (Either(..))

main :: Effect Unit
main = do
  log "Testing purescript-goos..."
  
  -- Test command line arguments
  cmdArgs <- args
  log $ "Command line arguments: " <> show cmdArgs
  
  -- Test directory reading
  log "Reading current directory..."
  dirResult <- readDir "."
  case dirResult of
    Left error -> log $ "Directory read error: " <> error
    Right entries -> do
      log $ "Found " <> show (length entries) <> " entries"
      log $ "First few files: " <> show (map fileName (take 5 entries))
  
  -- Test file I/O
  log "Testing file I/O..."
  writeResult <- writeTextFile "test-goos.txt" "Hello from purescript-goos!"
  case writeResult of
    Left error -> log $ "Write error: " <> error
    Right _ -> do
      log "File written successfully"
      readResult <- readTextFile "test-goos.txt"
      case readResult of
        Left error -> log $ "Read error: " <> error
        Right content -> log $ "File content: " <> content
  
  log "purescript-goos test complete!"