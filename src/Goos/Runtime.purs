module Goos.Runtime where

import Prelude
import Effect (Effect)

foreign import cpuCount :: Effect Int
foreign import monotonicMilliseconds :: Effect Number
foreign import sleepMs :: Int -> Effect Unit
foreign import physicalMemoryBytes :: Effect Number
foreign import setGCPercent :: Int -> Effect Unit
foreign import setMemoryLimit :: Number -> Effect Unit
foreign import collectGarbage :: Effect Unit
foreign import writeHeapProfile :: String -> Effect Unit
foreign import version :: String
