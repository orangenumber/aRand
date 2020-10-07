# aRand

Copyright 2020 Gon Yi
<https://gonyyi.com/copyright.txt>


---

## RandStr / RandInt / RandInt64

- `UpdateSeed()` will run when the library is imported. (`init()`)
  If need to run the seed again, this can be done by calling `arand.UpdateSeed()`
- `RandStr(string, length)string` provides random nth long string.
  There are few predefined strings to pick from:

    ~~~go
    R_ALPHA              = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    R_NUMERIC            = "0123456789"
    R_ALPHANUMERIC       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    R_ALPHANUMERIC_LOWER = "abcdefghijklmnopqrstuvwxyz0123456789"
    R_ALPHANUMERIC_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    R_HEX                = "0123456789ABCDEF"
    ~~~

- `RandInt(min, max int)int` will return any value between min and max.
- `RandInt64(min, max int64)int64` will return any value between min and max.


---

## UUID

- `arand.UUID()[]byte` returns []byte of UUID Hex (16 x 2 (byte to Hex) + 4 of `-` = 36 bytes)
- `arand.UUIDb()[]byte` returns raw []byte of UUID (16 bytes)

