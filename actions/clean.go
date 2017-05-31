// -*- mode: go; tab-width: 4; -*-

/* burrow - a go build system that uses glide for dependency management.
 *
 * Copyright (C) 2017  EmbeddedEnterprises
 *     Fin Christensen <christensen.fin@gmail.com>,
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package burrow

import (
	"os"
	"path/filepath"

	"github.com/EmbeddedEnterprises/burrow/utils"
	"github.com/urfave/cli"
)

func Clean(context *cli.Context) error {
	burrow.LoadConfig()

	burrow.Log(burrow.LOG_INFO, "clean", "Cleaning build artifacts")
	if err := burrow.Exec("go", "clean"); err != nil {
		return err
	}

	if err := filepath.Walk("./bin", func(path string, f os.FileInfo, err error) error {
		if path == "./bin" {
			return nil
		}
		return os.Remove(path)
	}); err != nil {
		return err
	}

	return filepath.Walk("./package", func(path string, f os.FileInfo, err error) error {
		if path == "./package" {
			return nil
		}
		return os.Remove(path)
	})
}