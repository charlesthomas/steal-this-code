# The best way to wait on a GitHub rate-limit in go

- GitHub's API responses include information about their rate-limiting
    - this includes how many calls you have left and when the count resets
    - in [google/gogithub](https://pkg.go.dev/github.com/google/go-github/v63/github#Rate) this is a `time.Time`
- the go std lib `time.Until` func takes a `time.Time` and returns a `time.Duration` for how long until that `time.Time` will happen
- `time.Sleep` func takes a `time.Duration`

This means we can build a `github` struct with a `wait` func and a `reset` property that automatically waits until the reset has elapsed before making a call:

```go
if github.Rate.Remaining == 0 {
    time.Sleep(time.Until(github.Rate.Reset.Time))
}
```
