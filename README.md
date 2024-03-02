# Sink email delivery plugin

This plugin provides a sink for email deliveries with a configurable delay (in milliseconds).

## Installation

Follow the [instructions](https://docs.halon.io/manual/comp_install.html#installation) in our manual to add our package repository and then run the below command.

### Ubuntu

```
apt-get install halon-extras-sink
```

### RHEL

```
yum install halon-extras-sink
```

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
