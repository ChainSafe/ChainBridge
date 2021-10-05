# Metrics

Basic metrics and a health status check can be enabled with the `--metrics` flag (default port `8001`, use `--metricsPort` to specify).

## Prometheus
Prometheus metrics are served on `/metrics`. For each chain that exists, this provides:
- `<chain>_blocks_processed`: the number of blocks processed by the chains listener.
- `<chain>_latest_processed_block`: most recent block that has been processed by the listener.
- `<chain>_latest_known_block`: most recent block that exists on the chain.
- `<chain>_votes_submitted`: number of votes submitted by the relayer.

## Health Check
The endpoint `/health` will return the current known block height, and a timestamp of when it was first seen for every chain:
 ```json
{
  "chains": [
    {
      "chainId": "Number",
      "height": "Number",
      "lastUpdated": "Date"
    }
  ]
} 
```
 
 If the timestamp is at least 120 seconds old an error will be returned instead:
```json
{
  "error": "String"
}
```