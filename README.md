# Argus 

A debugging tool for UNIX pipes.

<p align="center">
<img width="400" src="docs/img/Argus.png" alt="Argus" title="Argus" />
</p>

Ever wondered what’s *actually* flowing through your pipelines?

Argus lets you inspect, measure, and understand stdin/stdout —
without breaking your pipes.

Input:
cat huge.log | Argus | jq '.error'

Output:
Argus: 334MBs | 204.07 MB/s | 1s

## Why?

UNIX pipes are powerful — but opaque.

- `cat` shows everything (too much)
- `head` breaks pipelines
- `wc` gives counts, not insight
- `tee` is clumsy and slow

Argus gives you visibility without disruption.

## What Argus does

- Passes stdin to stdout unchanged
- Displays live statistics on stderr
- Samples records without buffering
- Detects common data formats automatically


## Common use cases

### Inspect a large stream
cat data.json | Argus | jq '.id'

### Sample records without stopping the pipe
cat data.json | Argus --peek 5 | process

### Measure pipeline performance
producer | Argus --rate | consumer

### Tee output safely
command | Argus --tee raw.log | filter


## Design principles

- UNIX-first
- No configuration files
- No stdout modification
- Streaming, not buffering
- One binary, zero dependencies

## Install

go install github.com/jonathon-chew/Argus@latest


## How is this different from…

| Tool | Problem |
|-----|--------|
| cat | too verbose |
| head | breaks pipelines |
| wc | no insight |
| tee | clumsy & slow |

Argus complements existing tools instead of replacing them.


## Roadmap

- interactive pause/resume
- format-specific summaries
- plugin system (maybe)
