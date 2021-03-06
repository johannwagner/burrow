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
	"github.com/EmbeddedEnterprises/burrow/utils"
	"github.com/coreos/go-semver/semver"
	"github.com/urfave/cli"
)

// Major increments the major part of the semantic version in the project configuration.
func Major(context *cli.Context) error {
	burrow.LoadConfig()

	ver := semver.New(burrow.Config.Version)
	ver.BumpMajor()
	burrow.Config.Version = ver.String()
	burrow.Log(burrow.LOG_INFO, "semver", "Setting new version to %s...", burrow.Config.Version)

	burrow.SaveConfig()

	return nil
}
