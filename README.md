# coz-go

Coz profiler Golang wrapper.

This is a golang wrapper around the `Coz` profiler here: https://github.com/plasma-umass/coz

### Code Usage {#code-usage}

| Coz Call      | Equivalent CozGo Call | Description |
| ----------- | -----------   | ----------- | 
| COZ_BEGIN("name")   | cozgo.Begin("name")   | Begin a latency profiling block identified by the name |
| COZ_END("name")     | cozgo.End("name")     | End a latency profiling block identified by the name|
| COZ_PROGRESS()     | cozgo.Progress()     | Specify a progress point within a block of work|
| COZ_PROGRESS()     | cozgo.NamedProgress("name")     | Specify a named progress point within a block of work|

## For background on causal profiling, see

Paper: http://www.sigops.org/s/conferences/sosp/2015/current/2015-Monterey/printable/090-curtsinger.pdf

Blog: https://morsmachine.dk/causalprof

## Running/Usage

- Install Coz: https://github.com/plasma-umass/coz
    - For OS X users, the easiest way is to run a docker image with linux and install Coz
    - 'apt update; apt install coz-profiler'
- Import `cozgo`: `import "github.com/urjitbhatia/cozgo"`
- Call the `Coz` wrappers in your application
    - See the [Code Usage Section above](#code-usage)
- Build your `go` binary with the flags: `-ldflags=-compressdwarf=false -gcflags=all="-N -l"` so that Coz can identify
  the debug symbols properly
- Run your application with Coz: `coz run --- yourGoBinary`
- Let the application run for a while (you might want to run some load test etc, so that the profiler can pick up usage
  data)
- By default, Coz will create a file called `profile.coz`
- Use the Coz viewer to analyze the profile
