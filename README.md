# build

Build is a small utility for compile-time major.minor.micro versioning
with additional linker hooks for date and the git commit hash, such as
"1.4.2.20240506.4f54bc3d". The date is directly useful and monotonically
increasing. The hash identifies the exact commit and prevents collisions.

## Usage

Define the desired version and log it on startup, say:
```
var version = build.NewVersion(1, 4, 2)

func main() {
   log.Printf("Foo %v", version)
   ...
}
```

Add the following flags to `go build` and `go install`:
```
-ldflags "-X github.com/seekerror/build.GitHash=`git rev-parse --short HEAD`"
```
Or if a specific date is needed:
```
-ldflags "-X github.com/seekerror/build.Date=20240506 -X github.com/seekerror/build.GitHash=`git rev-parse --short HEAD`"
```

## License

Build is released under the [MIT License](http://opensource.org/licenses/MIT).
