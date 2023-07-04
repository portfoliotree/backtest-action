# Portfolio Backtest Action

See [the portfoliotree/factors workflow](https://github.com/portfoliotree/factors/blob/main/.github/workflows/back_test.yml) for an example of using this action.

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
      - uses: portfoliotree/backtest-action
      - name: Save Artifacts
        uses: actions/upload-artifact@v3
        with:
          name: backtest_results
          path: |
            backtest_results.json
            returns.csv
          retention-days: 2
```
