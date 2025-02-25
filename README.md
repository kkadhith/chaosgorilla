# Chaos Gorilla
Chaos Gorilla is a CLI based tool that injects faults in a given VM. You can use it to see how an application hosted on a virtual machine performs under stress. 

## Usage
To get started, clone this repo and run this to simulate a CPU spike for 30 seconds:

```go run main.go inject --type cpu --vm test --duration 30```


## Current Features
- Inject CPU Spikes for a specified duration
- Memory faults

## Roadmap
- Simulating network throttling
- Thread starvation
- Harware fault w/ QEMU (CPU failures)
