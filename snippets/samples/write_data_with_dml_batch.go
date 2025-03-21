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

package samples

// [START spanner_dml_getting_started_insert]
import (
	"context"
	"database/sql"
	"fmt"
	"io"

	_ "github.com/googleapis/go-sql-spanner"
)

func WriteDataWithDmlBatch(ctx context.Context, w io.Writer, databaseName string) error {
	db, err := sql.Open("spanner", databaseName)
	if err != nil {
		return err
	}
	defer db.Close()

	// Add multiple rows in one DML batch.
	// database/sql also supports named parameters.
	sql := "INSERT INTO Singers (SingerId, FirstName, LastName) " +
		"VALUES (@SingerId, @FirstName, @LastName)"

	// Get a connection from the pool, so we know that all statements
	// are executed on the same connection.
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	// `start batch dml` starts a DML batch.
	// All following DML statements are buffered in the client until the
	// `run batch` statement is executed.
	if _, err := conn.ExecContext(ctx, "start batch dml"); err != nil {
		return err
	}
	if _, err := conn.ExecContext(ctx, sql, 16, "Sarah", "Wilson"); err != nil {
		return err
	}
	if _, err := conn.ExecContext(ctx, sql, 17, "Ethan", "Miller"); err != nil {
		return err
	}
	if _, err := conn.ExecContext(ctx, sql, 18, "Maya", "Patel"); err != nil {
		return err
	}
	// `run batch` sends all buffered DML statements in one batch to Spanner.
	res, err := conn.ExecContext(ctx, "run batch")
	if err != nil {
		return err
	}
	c, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%v records inserted\n", c)

	return nil
}

// [END spanner_dml_getting_started_insert]
