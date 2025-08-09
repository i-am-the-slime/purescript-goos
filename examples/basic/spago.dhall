{ name = "goos-example-basic"
, dependencies = 
  [ "goos"
  , "effect"
  , "prelude"
  , "console"
  , "either"
  , "maybe"
  , "arrays"
  ]
, packages = 
    let upstream =
          https://github.com/purescript/package-sets/releases/download/psc-0.15.15-20240701/packages.dhall
            sha256:c0ad6f3f08ab980e9cacfd6e1fc96f111fd48aad70b9f951dde7bab5e2dc954e
    in  upstream
      with goos =
        { dependencies = 
          [ "effect", "prelude", "either", "maybe", "arrays", "bifunctors"
          , "control", "exceptions", "unsafe-coerce", "yoga-json", "st"
          , "http-methods", "ordered-collections", "nullable"
          , "foldable-traversable", "functions"
          ]
        , repo = "https://github.com/i-am-the-slime/purescript-goos.git"
        , version = "v0.1.0"
        }
, sources = [ "**/*.purs" ]
, backend = "psgo"
}