module Goos.Watch where

import Prelude
import Effect (Effect)

foreign import data Watcher :: Type

type Event =
  { name :: String
  , write :: Boolean
  , create :: Boolean
  , rename :: Boolean
  , remove :: Boolean
  , chmod :: Boolean
  , closed :: Boolean
  , timeout :: Boolean
  , error :: String
  }

foreign import open :: String -> Effect Watcher
-- A negative timeout blocks indefinitely. Zero is a nonblocking poll.
foreign import next :: Watcher -> Int -> Effect Event
foreign import close :: Watcher -> Effect Unit
