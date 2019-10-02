# coz-go

Coz profiler Golang wrapper.

This is a golang wrapper around the `Coz` profiler here: https://github.com/plasma-umass/coz

## For background on causal profiling, see

Paper: http://www.sigops.org/s/conferences/sosp/2015/current/2015-Monterey/printable/090-curtsinger.pdf

Blog: https://morsmachine.dk/causalprof

## Usage

- Install Coz: https://github.com/plasma-umass/coz
  - For OS X users, the easiest way is to run a docker image with linux and install Coz
  - 'apt update; apt install coz-profiler'
- Import `coz-go`
- Call the `Coz` macro wrappers in your application
- Build your `go` binary with the flags: `-ldflags=-compressdwarf=false` so that Coz can indentify the debug symbols properly
- Run your application with Coz: `coz run --- yourGoBinary`
- Let the application run for a while (you might want to run some load test etc, so that the profiler can pick up usage data)
- By default, Coz will create a file called `profile.coz`
- Use the Coz viewer to analyze the profile
