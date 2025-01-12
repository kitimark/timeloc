# TimeLoc

A golangci-lint analyzer to ensure proper `time.Location` handling in Go codebases. This analyzer helps prevent common timezone-related bugs by enforcing explicit time zone settings before using time-dependent methods.

## Problem

When working with `time.Time` methods, many operations depend on the location (timezone) setting. If not explicitly set, these methods default to using the system's local timezone. This implicit behavior makes your code vulnerable to environment-specific bugs, as the same code can produce different results depending on the server's timezone configuration.

Examples of potentially problematic code:
```go
// BAD: Using time methods without setting location
t := time.Now()
fmt.Println(t.Format("2006-01-02"))  // Uses system timezone

// GOOD: Explicitly set location before using time methods
t := time.Now()
t = t.In(time.UTC)
fmt.Println(t.Format("2006-01-02"))  // Consistently uses UTC
```

## Running TimeLoc

### Standalone

Install the binary from source by running:
```bash
go install github.com/kitmark/timeloc/cmd/timeloc@latest
```

Then, run the linter by:
```bash
timeloc ./...
```

### golangci-lint

Integration with golangci-lint is under development following the [Module Plugin System](https://golangci-lint.run/plugins/module-plugins/). Contributions towards this integration are welcome, please feel free to open a pull request.

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
    fmt.Println(t.Format("2006-01-02"))  // Analyzer reports error: (t time.Time).Format called on t before setting time.Location
    
    // Passing to function without location
    formatTime(t)  // Analyzer reports error: (t time.Time).Format called on t before setting time.Location
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

// Using ParseInLocation is also good
func parseExample() {
    t, _ := time.ParseInLocation("2006-01-02", "2025-01-01", time.UTC)
    fmt.Println(t.Format("2006-01-02"))  // Fine because location is set
}
```

## Error Messages

The analyzer produces two types of error messages:

1. For direct method calls:
   ```
   (t time.Time).Format called on t before setting time.Location
   ```

2. For function calls:
   ```
   passing time.Time value without location set to function that may use location-dependent methods
   ```

## Why Not Trust Default Location?

While time.Parse and the system timezone provide default location settings, relying on these can lead to issues:
- System timezone can vary between development, testing, and production environments
- Docker containers might have different timezone settings
- Application deployments across different regions might behave differently

This analyzer enforces explicit location setting to ensure consistent behavior across all environments.

## Contributing

Pull requests are welcome! Please ensure tests pass by running:

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
