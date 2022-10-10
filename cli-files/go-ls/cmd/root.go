/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
package cmd



import (
  "os"
  "log"
  "fmt"
  "github.com/spf13/cobra"
  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

)


var cfgFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "go-ls",
  Short: "This command lists information about directories",
  Long: `This command lists information about directories and any type of files in the working directory.`,
  // Uncomment the following line if your bare application
  // has an action associated with it:

    RunE: func(cmd *cobra.Command, args []string) error {
   
      if len(os.Args) > 1 {
        args = os.Args[1:]
      }
    
      for _, arg := range args {
        err := listDirectory(arg)
        if err != nil {
          log.Printf("Not able to list %s, %v\n", arg, err)
        
      }
    } 
    return nil
  },
  
}
func listDirectory(root string) error {
	fi, err := os.Stat(root)

	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return nil
	}

	fis, err := os.ReadDir(root)

	if err != nil {
		return err
	}

	var totalDirectory int = 0
	var totalFiles int = 0

	for _, info := range fis {
		if info.Name()[0] != '.' && info.Name()[0] != '$'{
			if info.IsDir() {
				totalDirectory += 1
			} else {
				totalFiles += 1
			}
			fmt.Printf(" %v\n", info.Name())
		}
	}
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprint(os.Stdout, "Total Directory: ", totalDirectory, ", Total Files: ", totalFiles)

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-ls.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".go-ls" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".go-ls")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

