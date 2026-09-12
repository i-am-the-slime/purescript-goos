module Goos.Process where

import Prelude
import Effect (Effect)

foreign import data Signals :: Type

foreign import environment :: String -> Effect String
foreign import pid :: Effect Int
foreign import uid :: Effect Int
foreign import platform :: String
foreign import tempDirectory :: Effect String
foreign import absolutePath :: String -> Effect String
foreign import dirname :: String -> String
foreign import basename :: String -> String
foreign import joinPath :: String -> String -> String
foreign import sendSignal :: Int -> String -> Effect Boolean
foreign import subscribeSignals :: Array String -> Effect Signals
foreign import receiveSignal :: Signals -> Effect String
foreign import closeSignals :: Signals -> Effect Unit
foreign import removeFile :: String -> Effect Unit
foreign import finish :: Effect Unit
