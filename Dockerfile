FROM golang:1.20-alpine
COPY . /code/backtest-action
RUN cd /code/backtest-action && go install -v .
ENTRYPOINT ["backtest-action"]