# TimeLoc

A golangci-lint analyzer to ensure proper `time.Location` handling in Go codebases. This analyzer helps prevent common timezone-related bugs by enforcing explicit time zone settings before using time-dependent methods and prevents the usage of system-dependent `time.Local`.

## Problem

When working with `time.Time` methods, many operations depend on the location (timezone) setting. If not explicitly set, these methods default to using the system's local timezone. This implicit behavior makes your code vulnerable to environment-specific bugs, as the same code can produce different results depending on the server's timezone configuration.

The analyzer also prevents the usage of `time.Local` as it relies on the system timezone, which can vary across different environments.

Examples of potentially problematic code:
```go
// BAD: Using time methods without setting location
t := time.Now()
fmt.Println(t.Format("2006-01-02"))  // Uses system timezone

// BAD: Using time.Local which relies on system timezone
t := time.Now().In(time.Local)       // Will trigger an error
t, _ := time.ParseInLocation("2006-01-02", "2025-01-01", time.Local)  // Will trigger an error
t := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)  // Will trigger an error

// GOOD: Explicitly set location before using time methods
t := time.Now()
t = t.In(time.UTC)
fmt.Println(t.Format("2006-01-02"))  // Consistently uses UTC
```

## Running TimeLoc

### Standalone

Install the binary from source by running:
```bash
go install github.com/kitimark/timeloc/cmd/timeloc@latest
```

Then, run the linter by:
```bash
timeloc ./...
```

### golangci-lint (>= v1.57.0)

TimeLoc is available as a golangci-lint analyzer through the [Module Plugin System](https://golangci-lint.run/plugins/module-plugins/) (introduced in v1.57.0).

Follow these steps to set up TimeLoc with golangci-lint:

(1) Create a `.custom-gcl.yml` file at the root of the repository if you have not done so, add the following content:
```yaml
# This has to be >= v1.57.0 for module plugin system support.
version: v1.57.0
plugins:
  - module: "github.com/kitimark/timeloc"
    import: "github.com/kitimark/timeloc/cmd/gclplugin"
    version: latest # Or a fixed version for reproducible builds.
```

(2) Add TimeLoc to the linter configuration file `.golangci.yaml`:
```yaml
linters-settings:
  custom:
    timeloc:
      type: "module"
      description: A linter to enforce explicit timezone settings and prevent time.Local usage in Go code.
# TimeLoc can be referred to as `timeloc` just like any other golangci-lint analyzers in other 
# parts of the configuration file.
```

(3) Build a custom `golangci-lint` binary with TimeLoc included:
```bash
# Note that your `golangci-lint` to bootstrap the custom binary must also be version >= v1.57.0.
$ golangci-lint custom
```
By default, the custom binary will be built at `.` with the name `custom-gcl`, which can be further customized in `.custom-gcl.yml` file (see [Module Plugin System](https://golangci-lint.run/plugins/module-plugins/) for instructions).

(4) Run the custom binary instead of `golangci-lint`:

```bash
# Arguments are the same as `golangci-lint`.
$ ./custom-gcl run ./...
```

## What This Analyzer Checks

### Time-dependent Methods

The analyzer checks for usage of the following time.Time methods without explicit location setting:
- AppendFormat
- Clock
- Date
- Day
- Format
- Hour
- ISOWeek
- Minute
- Month
- Second
- Weekday
- Year
- YearDay

### Function Parameters

The analyzer also tracks time.Time values passed to functions to ensure they have location set before being used with time-dependent methods.

### time.Local Usage

The analyzer checks and prevents the use of `time.Local` in the following contexts:
- `time.ParseInLocation(..., time.Local)`
- `time.Date(..., time.Local)`
- `someTime.In(time.Local)`

### Import Aliases

This check works even when the time package is imported with an alias.

### Supported Location Setting Methods

The analyzer recognizes location setting through:
- Direct `In()` method calls: `t = t.In(time.UTC)`
- `time.ParseInLocation()`: `t, _ := time.ParseInLocation(layout, value, loc)`
- Functions that internally set location:
  ```go
  func setUtcTz(t time.Time) time.Time {
      return t.In(time.UTC)
  }
  ```

## Examples

### Incorrect Code

```go
func badUsage() {
    t := time.Now()
    
    // Direct method calls without location
    fmt.Println(t.Format("2006-01-02"))  // Analyzer reports error: (t time.Time).Format called before setting time.Location
    
    // Method chaining without location
    fmt.Println(time.Now().Format("2006-01-02"))  // Analyzer reports error: (t time.Time).Format called before setting time.Location
    
    // Passing to function without location
    formatTime(t)  // Analyzer reports error: passing time.Time value without location set to function that may use location-dependent methods
 
    // Using time.Local (will trigger errors)
    t = t.In(time.Local)  // Analyzer reports error: time.Local usage is not allowed as it relies on system timezone
}

func formatTime(t time.Time) string {
    return t.Format("2006-01-02")
}
```

### Correct Code

```go
func goodUsage() {
    t := time.Now().In(time.UTC)
    
    // Location is set, so these are fine
    fmt.Println(t.Format("2006-01-02"))
    formatTime(t)
}

// Using ParseInLocation with explicit timezone is good
func parseExample() {
    t, _ := time.ParseInLocation("2006-01-02", "2025-01-01", time.UTC)
    fmt.Println(t.Format("2006-01-02"))  // Fine because location is set
}
```

## Error Messages

The analyzer produces three types of error messages:

1. For direct method calls on variables and chained calls:
   ```
   (t time.Time).Format called before setting time.Location
   ```

2. For function calls:
   ```
   passing time.Time value without location set to function that may use location-dependent methods
   ```

3. For time.Local usage:
   ```
   time.Local usage is not allowed as it relies on system timezone
   ```

## Why Not Trust Default Location?

While time.Parse and the system timezone provide default location settings, relying on these can lead to issues:
- System timezone can vary between development, testing, and production environments
- Docker containers might have different timezone settings
- Application deployments across different regions might behave differently
- Using time.Local makes your code dependent on the host system's timezone configuration

This analyzer enforces explicit location setting and prevents time.Local usage to ensure consistent behavior across all environments.

## Support
Please feel free to [open a GitHub issue](https://github.com/kitimark/timeloc/issues) if you have any questions, bug reports, and feature requests.

## Contributions

Pull requests are welcome! Please ensure tests pass by running:

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
