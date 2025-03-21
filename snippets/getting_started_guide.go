// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/googleapis/go-sql-spanner/examples/samples"
)

type command func(ctx context.Context, w io.Writer, databaseName string) error

var (
	commands = map[string]command{
		"createconnection":             samples.CreateConnection,
		"createtables":                 samples.CreateTables,
		"dmlwrite":                     samples.WriteDataWithDml,
		"write":                        samples.WriteDataWithMutations,
		"query":                        samples.QueryData,
		"querywithparameter":           samples.QueryDataWithParameter,
		"addcolumn":                    samples.AddColumn,
		"ddlbatch":                     samples.DdlBatch,
		"update":                       samples.UpdateDataWithMutations,
		"querymarketingbudget":         samples.QueryNewColumn,
		"writewithtransactionusingdml": samples.WriteWithTransactionUsingDml,
		"tags":                         samples.Tags,
		"readonlytransaction":          samples.ReadOnlyTransaction,
		"databoost":                    samples.DataBoost,
		"pdml":                         samples.PartitionedDml,
	}
)

func run(ctx context.Context, w io.Writer, cmd string, db string) error {
	cmdFn := commands[cmd]
	if cmdFn == nil {
		flag.Usage()
		os.Exit(2)
	}
	err := cmdFn(ctx, w, db)
	if err != nil {
		fmt.Fprintf(w, "%s failed with %v", cmd, err)
	}
	return err
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: getting_started_guide <command> <database_name>
Examples:
	spanner_snippets write projects/my-project/instances/my-instance/databases/example-db
`)
	}

	flag.Parse()
	if len(flag.Args()) < 2 || len(flag.Args()) > 3 {
		flag.Usage()
		os.Exit(2)
	}

	cmd, db := flag.Arg(0), flag.Arg(1)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	if err := run(ctx, os.Stdout, cmd, db); err != nil {
		os.Exit(1)
	}

}
