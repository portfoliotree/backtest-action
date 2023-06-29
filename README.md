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
      - uses: portfoliotree/backtest-action@v0
```