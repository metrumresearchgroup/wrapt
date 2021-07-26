# wrapt

Wrapt wraps Go's `&testing.T` type in its own type to add testing conveniences.

## Amenities

AssertError: Compares the wantError value with the actual error state and sets a `t.Error`.

Run: mirrors the regular `*t.Run` function `testing`, but accepting a `*wrapt.T` instead for the function definition. Returns a bool success value like the original.

RunFatal: runs like Run, accepting `*wrapt.T` but raises a promotes an inner failure as an failure in the current test.

ValidateError: Like AssertError, but calls `Fatalf` on the outer test.

WrapT: creates a wrapped `*testing.T` value that provides additional functionality.