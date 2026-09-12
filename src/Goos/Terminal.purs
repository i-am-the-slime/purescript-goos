module Goos.Terminal where

import Prelude
import Effect (Effect)

foreign import isStdinTTY :: Effect Boolean
foreign import enterRawMode :: Effect Unit
foreign import restoreRawMode :: Effect Unit
foreign import termSize :: Effect { cols :: Int, rows :: Int }
foreign import readKeyBytes :: Effect (Array Int)
foreign import putStr :: String -> Effect Unit
foreign import putErr :: String -> Effect Unit
