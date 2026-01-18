# What to have ready / focus on next time
1. Lock in renderer contract
[] Confirm only render on the throttle cadence.
[] Ensure never print anything (including logs) to stderr while live status is active, unless it’s deliberate.

2. Use the numeric formatter everywhere
[] Use the %10//10 byte-buffer formatter for:
    [] lines
    [] bytes
    [] rates 
[] Keep formatting only in the render block.

3. Rate + elapsed as “derived metrics”
[] Rate should be delta-based:
    (totalBytes - lastBytes) / (now - lastRender)
[] Elapsed time should be:
    now - startTime
[] Both computed only during render.

4. Make status line non-wrapping
[] Decide a max width and truncate fields (wrapping breaks clear-line behavior).

Short / medium / long-term development plan
## Short term (next session / next few commits)
```
[x] Add bytes/sec and lines/sec (delta-based, render-only).
[x] Add elapsed time (human readable).
[] Stabilize the status line layout so it never wraps.
[] Add TTY detection for stderr:
    if not a TTY, disable ANSI and either print nothing or print occasional newline-terminated status.
```
## Medium term
```
[] Add stall detection (“no bytes for N seconds”).
[x] Add human readable bytes (KB/MB/GB) consistently for totals and rates.
[] Add Ctrl+C cleanup (SIGINT/SIGTERM): clear line + newline, then exit.
[] Reduce allocations further - in progress
```

## Long term
```
[] Optional “inspection” features that never break the pipe:
[] peek/sample first N bytes/records to stderr
[] format detection hints (json/csv/raw) based on first chunk
[] optional --tee file behavior
[] Optional quality-of-life flags:
    --no-ansi, --force-ansi
    [x] --interval 100ms
```
