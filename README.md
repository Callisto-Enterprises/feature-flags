# Feature Flags

A crude feature flag implementation

## Installation

As a library

```shell
go get github.com/Callisto-Enterprises/feature-flags
```

## Usage

Add your feature flags to your `.env` file. The feature flags must be prefixed with `FeatureFlag_`.
If a corresponding environment variable does not exist, the feature will be disabled.

```shell
FeatureFlag_MyFirstFlag=true   #Enabled
FeatureFlag_MySecondFlag=false #Disabled
```
Then in your Go app you can do something like

```go
package main

import (
    "fmt"

    "github.com/Callisto-Enterprises/feature-flags"
)

func main() {
  isEnabled := featureflags.GetInstance().IsEnabled("MyFirstFlag")
  fmt.Println("Is Enabled: ", isEnabled)
}
```
