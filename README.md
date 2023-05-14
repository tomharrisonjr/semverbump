# SemVerBump -- Go package to bump semantic versions
As in: `v1.2.3` Bump minor => `v1.3.0`

Extends [golang.org/x/mod/semver](https://pkg.go.dev/golang.org/x/mod/semver)

## Usage

```go
import "github.com/tomharrisonjr/semverbump"
```

```go
func Bump(version, bumpType string) string
```
* version is a valid semver string
  * if empty, respects the bumpType, so `Bump("", "patch")` will return `v0.0.1`
  * if the passed version does not include the leading `v`, it will be added, so `Bump("0.0.1", "patch")` will return `v0.0.2`
* bumpType is a string of `major`, `minor`, or `patch`.

```go
func main()
	version := "v0.0.1"
	newVersion := semverbump.Bump(version, "patch")
	fmt.Println(newVersion)
```

## Contributing
Sure! Just open a PR.

Better yet, I totally can't believe this doesn't like completely exist already so if it does, please let me know.

## License
MIT

That's all folks.