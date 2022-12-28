# Sink email delivery plugin

This plugin provides a sink for email deliveries with a configurable delay (in milliseconds).

## Configuration

### Pre-delivery script hook

```
Try([
    "plugin" => [
        "id" => "sink",
        "arguments" => [
            "delay" => random_number(100, 1000)
        ]
    ]
]);
```