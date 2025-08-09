module Goos where

import Prelude

import Control.Monad.Except (ExceptT(..), runExceptT)
import Data.Array.NonEmpty (NonEmptyArray)
import Data.Bifunctor (lmap)
import Data.Either (Either(..))
import Data.HTTP.Method (Method)
import Data.Map.Internal (keys)
import Data.Maybe (Maybe(..))
import Data.Nullable (Nullable, toNullable)
import Data.Semigroup.Foldable (intercalateMap)
import Effect (Effect)
import Effect.Class (class MonadEffect, liftEffect)
import Effect.Uncurried (EffectFn1, EffectFn3, EffectFn4, EffectFn6, runEffectFn1, runEffectFn3, runEffectFn4, runEffectFn6)
import Unsafe.Coerce (unsafeCoerce)
import Yoga.JSON (class ReadForeign)
import Yoga.JSON as JSON
import Yoga.JSON.Error (renderHumanError)
import Control.Monad.ST.Global (Global)
import Control.Monad.ST.Internal (ST)

foreign import data FileHandle :: Type

foreign import data DirEntry :: Type

foreign import readDirImpl :: forall a e. EffectFn3 (e -> Either e a) (a -> Either e a) String (Either String (Array DirEntry))

foreign import data File :: Type
foreign import data Directory :: Type

readDir :: forall eff. MonadEffect eff => String -> eff (Either String (Array DirEntry))
readDir = liftEffect <<< runEffectFn3 readDirImpl Left Right

foreign import fileName :: DirEntry -> String
foreign import isDir :: DirEntry -> Boolean

asFileOrDirectory :: DirEntry -> Either File Directory
asFileOrDirectory entry =
  if isDir entry then Right (unsafeCoerce entry)
  else Left (unsafeCoerce entry)

-- Text file I/O functions
foreign import writeTextFileImpl :: forall a e. EffectFn4 (e -> Either e a) (a -> Either e a) String String (Either String Unit)

writeTextFile :: String -> String -> Effect (Either String Unit)
writeTextFile = runEffectFn4 writeTextFileImpl Left Right

foreign import readTextFileImpl :: forall a e. EffectFn3 (e -> Either e a) (a -> Either e a) String (Either String String)

readTextFile :: String -> Effect (Either String String)
readTextFile = runEffectFn3 readTextFileImpl Left Right

foreign import args :: Effect (Array String)

foreign import exitImpl :: forall a. EffectFn1 Int a

exit :: forall a. Int -> Effect a
exit = runEffectFn1 exitImpl

foreign import stGlobalToEffect :: forall a. ST Global a -> Effect a