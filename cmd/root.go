package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"goclick/utils"
	"os"
)

var (
	cfgFile     string
	numclickers int

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "goclick [OPTIONS]",
		Short: "An auto clicker CLI",
		Long:  `An auto clicker CLI`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		RunE: func(cmd *cobra.Command, args []string) error {
			n, err := cmd.Flags().GetInt("numclickers")

			// Error checking
			if err != nil {
				return err
			}
			if n < 0 || n > 6 {
				return fmt.Errorf("invalid number of clickers (should be 1 to 6)")
			}

			if n == 1 {
				fmt.Printf("Using 1 clicker\n")
			} else {
				fmt.Printf("Using %v clickers\n", n)
			}

			coords := utils.Geomap()

			for i := 0; i < n; i++ {
				go utils.Click(coords)
			}

			exitChan := make(chan bool)

			<-exitChan

			return nil
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().IntVarP(&numclickers, "numclickers", "n", 1, "Number of concurrent clickers (default = 1, max = 6)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goclick.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".goclick" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".goclick")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
