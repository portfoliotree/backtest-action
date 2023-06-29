package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/portfoliotree/round"

	"github.com/portfoliotree/portfolio"
	"github.com/portfoliotree/portfolio/backtest"
	"github.com/portfoliotree/portfolio/returns"
)

func main() {
	specs, err := portfolio.WalkDirectoryAndParseSpecificationFiles(os.DirFS("."))
	if err != nil {
		log.Fatal(err)
	}

	backtestResultsJSON, err := os.Create("backtest_results.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = backtestResultsJSON.Close()
	}()
	backtestResultsEncoder := json.NewEncoder(backtestResultsJSON)

	var table returns.Table
	ctx := context.Background()
	var names []string
	for _, spec := range specs {
		log.Printf("running backtest with %d assets for %s: %q\n", len(spec.Assets), spec.Filepath, spec.Name)
		result, err := spec.Backtest(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		if err := backtestResultsEncoder.Encode(result); err != nil {
			log.Fatal(err)
		}
		rs := result.Returns()
		log.Println(report(spec, result))
		table = table.AddColumn(rs)
		names = append(names, spec.Filepath)
	}
	returnsCSV, err := os.Create("returns.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = returnsCSV.Close()
	}()
	if err := table.WriteCSV(returnsCSV, names); err != nil {
		log.Fatal(err)
	}
}

func report(spec portfolio.Specification, result backtest.Result) string {
	rs := result.Returns()
	var s strings.Builder
	s.WriteString(fmt.Sprintf("finished backtest for %s: %q\n", spec.Filepath, spec.Name))
	s.WriteString(fmt.Sprintf("\tlast_time:  %s\n", rs.LastTime().Format(time.DateOnly)))
	s.WriteString(fmt.Sprintf("\tfirst_time: %s\n", rs.FirstTime().Format(time.DateOnly)))
	s.WriteString(fmt.Sprintf("\tnumber_of_returns: %d\n", rs.Len()))
	s.WriteString(fmt.Sprintf("\tannualized_arithmetic_return: %f\n", round.Decimal(rs.AnnualizedArithmeticReturn()*100, 6)))
	s.WriteString(fmt.Sprintf("\tannualized_risk: %f\n", round.Decimal(rs.AnnualizedRisk()*100, 6)))
	s.WriteString(fmt.Sprintf("\tnumber_of_assets: %d\n", len(spec.Assets)))
	s.WriteString(fmt.Sprintf("\trebalance_count: %d\n", len(result.RebalanceTimes)))
	s.WriteString(fmt.Sprintf("\tpolicy_weight_update_count: %d\n", len(result.PolicyUpdateTimes)))
	return s.String()
}
