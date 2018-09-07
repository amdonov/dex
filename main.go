package main

import (
	"fmt"
	"os"
	"github.com/dexidp/dex/cmd"
	"github.com/spf13/cobra"
	"github.com/dexidp/dex/connector/authproxy"
	"github.com/dexidp/dex/connector/github"
	"github.com/dexidp/dex/connector/gitlab"
	"github.com/dexidp/dex/connector/ldap"
	"github.com/dexidp/dex/connector/linkedin"
	"github.com/dexidp/dex/connector/microsoft"
	"github.com/dexidp/dex/connector/mock"
	"github.com/dexidp/dex/connector/oidc"
	"github.com/dexidp/dex/connector/saml"
	"github.com/dexidp/dex/server"
)

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "dex",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(2)
		},
	}
	rootCmd.AddCommand(cmd.CommandServe())
	rootCmd.AddCommand(cmd.CommandVersion())
	return rootCmd
}

func main() {
	server.ConnectorsConfig = map[string]func() server.ConnectorConfig{
		"mockCallback": func() server.ConnectorConfig { return new(mock.CallbackConfig) },
		"mockPassword": func() server.ConnectorConfig { return new(mock.PasswordConfig) },
		"ldap":         func() server.ConnectorConfig { return new(ldap.Config) },
		"github":       func() server.ConnectorConfig { return new(github.Config) },
		"gitlab":       func() server.ConnectorConfig { return new(gitlab.Config) },
		"oidc":         func() server.ConnectorConfig { return new(oidc.Config) },
		"saml":         func() server.ConnectorConfig { return new(saml.Config) },
		"authproxy":    func() server.ConnectorConfig { return new(authproxy.Config) },
		"linkedin":     func() server.ConnectorConfig { return new(linkedin.Config) },
		"microsoft":    func() server.ConnectorConfig { return new(microsoft.Config) },
		// Keep around for backwards compatibility.
		"samlExperimental": func() server.ConnectorConfig { return new(saml.Config) },
	}
	if err := commandRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
