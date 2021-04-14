/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package mstream

import (
	"github.com/spf13/cobra"
	"jinr.ru/greenlab/go-adc/pkg/mstream"
)

const (
	AddressOptionName = "address"
	PortOptionName = "port"
	IfaceNameOptionName = "iface-name"
)

func NewMStreamCommand() *cobra.Command {
	var address, port string
	cmd := &cobra.Command{
		Use:           "mstream",
		Short:         "Start mstream server",
		RunE: func(cmd *cobra.Command, args []string) error {
			server, err := mstream.NewServer(address, port)
			if err != nil {
				return err
			}
			return server.Run()
		},
	}
	cmd.Flags().StringVar(&address, AddressOptionName, "", "Address to bind. E.g. 192.168.1.2")
	cmd.Flags().StringVar(&port, PortOptionName, "33301", "Port number to bind. E.g. 33301")

	return cmd
}
