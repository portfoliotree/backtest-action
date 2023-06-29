# Portfolio Backtest Action

## Example usage

```yaml
name: Back Test

on:
  push:
    branches: [ "main" ]
    paths-ignore:
      - 'README.md'

jobs:
  back_test:
    name: Run BackTest
    steps:
      - uses: actions/checkout@v3
      - uses: portfoliotree/backtest-action@v0
      - name: Save Artifacts
        uses: actions/upload-artifact@v3
        with:
          name: backtest_results
          path: |
            backtest_results.json
            returns.csv
          retention-days: 2
```