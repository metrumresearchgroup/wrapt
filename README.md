# wrapt

Wrapt wraps Go's `&testing.T` type in its own type to add testing conveniences.

## Goal

Wrapt's goal is to make testing, assertions, and positive statements about functionality convenient.

## Use Cases

### Error handling

`AssertError` / `ValidateError` both allow us to handle the common `wantErr` situation in generated tests:

Previously:

```go
// in test struct
wantErr: true
// in test
if (err != nil) != test.wantErr {
	t.Errorf("wantErr = %t, got %v", test.wantErr, err)
}
```

With this library, we can declare all of that as a one-liner. Since errors are common in go, this cleans up tests a lot:

```go
// same config as above

// This assertion calls assert.WantError, logs and error.
t.A.WantError(test.wantErr, err)

// This assertion calls require.WantError, which additionally stops the test.
t.R.ValidateError("err in write", test.wantErr, err)
```

### Running Sub-tests With Propagating Fatal

Running sub-tests is another place where we fall into a trap. The inner test will not fail the outer test! A failing 
`t.Run` will not fail the outer test without intervention. In order to work around this, we provide a `RunFatal` 
function.

Before:
```go
t := wrapt.Wrap(tt)
t.Run("sub-test", func(t *wrapt.T) {
	t.Errorf("inner test fails!")
})
// outer test only *errors*, and continues to run…
```

After:
```go
t := wrapt.WrapT(tt)
t.RunFatal("sub-test", func(t *wrapt.T){
	// anything unsuccessful works here
	t.Errorf("inner *and* outer test fails immediately!")
})
// Test stops on the above line if it fails instead
// of continuing on…
```

Note that Run is in the same form as `*testing.T.Run()` but it accepts a local `*wrapt.T` type.

## Amenities

WrapT: creates a wrapped `*testing.T` value that provides additional functionality.

Run: mirrors the regular `*t.Run` function `testing`, but accepting a `*wrapt.T` instead for the function definition.
Returns a bool success value like the original.

RunFatal: runs like Run, accepting `*wrapt.T` but promotes an inner failure as an failure in the current test.

WantErr: compares wheter we want an error (bool) to the error value (error or nil)

## A & R

`T.A` and `T.R` contains all the functionality of the [testify](https://github.com/stretchr/testify) `assert` and `require` structs, respectively.

Full documentation for [assert is here](https://pkg.go.dev/github.com/stretchr/testify/assert).

Assert returns a bool along with failing a test. Require calls `*testing.T`'s `FailNow()`.  

## Converting from *testing.T

The minimum schematic for testing.T tests is follows (as generated by goland):

```go
func TestNewFoo(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want *foo.Foo
	}{
		{ 
			name: "param test",
			args: args{
				param: "test",
			},
			want: &foo.Foo{
				Param: "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := foo.NewFoo(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFoo() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

In order to keep with IDE compatibility in GoLand, we want to keep using `*testing.T` in one critical place to keep the 
play button operational. The play button shows up on each sub-test struct because it's something passed into the range
function which is in turn passed on to t.Run. An updated, annotated version of this test follows:

```go
// renaming inbound t to tt
func TestNewFoo(tt *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want *foo.Foo
	}{
		{ 
			name: "param test",
			args: args{
				param: "test",
			},
			want: &foo.Foo{
				Param: "test",
			},
		},
	}
	// rename the default tt to 'test'
	for _, test := range tests {
		// also renaming it in the sub-test
		tt.Run(test.name, func(tt *testing.T) {
			// setting a local t for comfort
			// Note, you can '.' import wrapt to get
			// rid of the stutter of wrapt.WrapT.
			t := wrapt.WrapT(tt)
            
			got := foo.NewFoo(test.args.param)
			t.A.Equals(test.want, got)
		})
	}
}
```

The change order is:
1. Rename `tt` to `test`
1. Rename `t` to `tt`

When inside the outermost `tt.Run`, use `t.Run` or `t.RunFatal` instead, which passes on the `*wrapt.T` allowing access to the assertion libraries.