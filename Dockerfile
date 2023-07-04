FROM golang:1-alpine AS builder
COPY . /code/backtest-action
ENV CGO_ENABLED=0
RUN cd /code/backtest-action && go build -v -o backtest-action .

FROM gcr.io/distroless/static-debian11
COPY --from=builder /code/backtest-action/backtest-action /backtest-action
ENTRYPOINT ["/backtest-action"]