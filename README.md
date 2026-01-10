# pipepeek 

A debugging tool for UNIX pipes.

Ever wondered what’s *actually* flowing through your pipelines?

pipepeek lets you inspect, measure, and understand stdin/stdout —
without breaking your pipes.

Input:
cat huge.log | pipepeek | jq '.error'

Output:
pipepeek: 2.1GB | 14M lines | 87k lines/s | json

## Why?

UNIX pipes are powerful — but opaque.

- `cat` shows everything (too much)
- `head` breaks pipelines
- `wc` gives counts, not insight
- `tee` is clumsy and slow

pipepeek gives you visibility without disruption.

## What pipepeek does

- Passes stdin to stdout unchanged
- Displays live statistics on stderr
- Samples records without buffering
- Detects common data formats automatically


## Common use cases

### Inspect a large stream
cat data.json | pipepeek | jq '.id'

### Sample records without stopping the pipe
cat data.json | pipepeek --peek 5 | process

### Measure pipeline performance
producer | pipepeek --rate | consumer

### Tee output safely
command | pipepeek --tee raw.log | filter


## Design principles

- UNIX-first
- No configuration files
- No stdout modification
- Streaming, not buffering
- One binary, zero dependencies

## Install

go install github.com/jonathon-chew/pipepeek@latest


## How is this different from…

| Tool | Problem |
|-----|--------|
| cat | too verbose |
| head | breaks pipelines |
| wc | no insight |
| tee | clumsy & slow |

pipepeek complements existing tools instead of replacing them.


## Roadmap

- interactive pause/resume
- format-specific summaries
- plugin system (maybe)
