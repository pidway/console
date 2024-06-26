// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package certs

const (
	// Default minio configuration directory where below configuration files/directories are stored.
	DefaultConsoleConfigDir = ".console"

	// Directory contains below files/directories for HTTPS configuration.
	CertsDir = "certs"

	// Directory contains all CA certificates other than system defaults for HTTPS.
	CertsCADir = "CAs"

	// Public certificate file for HTTPS.
	PublicCertFile = "public.crt"

	// Private key file for HTTPS.
	PrivateKeyFile = "private.key"
)
