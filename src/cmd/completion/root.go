package completion

import (
	"com.azure.iot/iotcentral/iotcgo/cmd"
	"github.com/spf13/cobra"
	"os"
)

var completionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate command line completion script",
	Long: `To load completions:

Bash:

$ source <(iotc completion bash)

# To load completions for each session, execute once:
Linux:
  $ iotc completion bash > /etc/bash_completion.d/iotc
MacOS:
  $ iotc completion bash > /usr/local/etc/bash_completion.d/iotc

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ iotc completion zsh > "${fpath[1]}/iotc"

# You will need to start a new shell for this setup to take effect.

Fish:

$ iotc completion fish | source

# To load completions for each session, execute once:
$ iotc completion fish > ~/.config/fish/completions/iotc.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(completionCmd)
}