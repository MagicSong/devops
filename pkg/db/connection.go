/*
Copyright 2018 The KubeSphere Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package db

import (
	"time"

	"github.com/gocraft/dbr"

	"kubesphere.io/devops/pkg/config"
)

var defaultEventReceiver = EventReceiver{}

func OpenDatabase(cfg config.MysqlConfig) (*Database, error) {
	// https://github.com/go-sql-driver/mysql/issues/9
	conn, err := dbr.Open("mysql", cfg.GetUrl()+"?parseTime=1&multiStatements=1&charset=utf8mb4&collation=utf8mb4_unicode_ci", &defaultEventReceiver)
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(100)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(10 * time.Second)
	return &Database{
		Session: conn.NewSession(nil),
	}, nil
}
