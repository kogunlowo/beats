// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package logp

import (
	"time"
)

// Config contains the configuration options for the logger. To create a Config
// from a common.Config use logp/config.Build.
type Config struct {
	Beat       string   `config:",ignore"`        // Name of the Beat (for default file name).
	JSON       bool     `config:"json"`           // Write logs as JSON.
	Level      Level    `config:"level"`          // Logging level (error, warning, info, debug).
	Selectors  []string `config:"selectors"`      // Selectors for debug level logging.
	ECSEnabled bool     `config:"ecs" yaml:"ecs"` // Adds minimal ECS information using ECS conformant keys to every log line

	toObserver  bool
	toIODiscard bool
	ToStderr    bool `config:"to_stderr" yaml:"to_stderr"`
	ToSyslog    bool `config:"to_syslog" yaml:"to_syslog"`
	ToFiles     bool `config:"to_files" yaml:"to_files"`
	ToEventLog  bool `config:"to_eventlog" yaml:"to_eventlog"`

	Files FileConfig `config:"files"`

	environment Environment
	addCaller   bool // Adds package and line number info to messages.
	development bool // Controls how DPanic behaves.
}

// FileConfig contains the configuration options for the file output.
type FileConfig struct {
	Path            string        `config:"path" yaml:"path"`
	Name            string        `config:"name" yaml:"name"`
	MaxSize         uint          `config:"rotateeverybytes" yaml:"rotateeverybytes" validate:"min=1"`
	MaxBackups      uint          `config:"keepfiles" yaml:"keepfiles" validate:"max=1024"`
	Permissions     uint32        `config:"permissions"`
	Interval        time.Duration `config:"interval"`
	RotateOnStartup bool          `config:"rotateonstartup"`
	RedirectStderr  bool          `config:"redirect_stderr" yaml:"redirect_stderr"`
}

const defaultLevel = InfoLevel

// DefaultConfig returns the default config options for a given environment the
// Beat is supposed to be run within.
func DefaultConfig(environment Environment) Config {
	return Config{
		Level:   defaultLevel,
		ToFiles: true,
		Files: FileConfig{
			MaxSize:         10 * 1024 * 1024,
			MaxBackups:      7,
			Permissions:     0600,
			Interval:        0,
			RotateOnStartup: true,
		},
		environment: environment,
		addCaller:   true,
	}
}
