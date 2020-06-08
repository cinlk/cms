package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)


var (
	root = cobra.Command{Use: "app"}
)
func CreateConfigCmd() *cobra.Command{

	 cmd := &cobra.Command{
	 	Use: "hugo",
	 	Short: "hugo is very fast static site generator",
	 	Long: `A Fast and Flexible Static Site Generator build
				by spf13 and friends in go	`,
	 	Args: cobra.MinimumNArgs(1),
	 	Run: func(cmd *cobra.Command, args []string) {
			//
	 		fmt.Println(args)
		},
	 }
	return cmd
}


func CreatePrint() *cobra.Command {
	var cmdPrint = &cobra.Command{
		Use: "Prints [string to path]",
		Short: "Print anything to the screen",
		Long:  "print is for printing anything banck to for many years people have printed ",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},

	}

	return cmdPrint
}






func init()  {

	var echoTimes int

	root.AddCommand(CreateConfigCmd())

	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}


	var cmdTimes = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "Echo anything to the screen more times",
		Long: `echo things multiple times back to the user by providing
a count and a string.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}

	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	root.AddCommand(CreatePrint(), cmdEcho)
	cmdEcho.AddCommand(cmdTimes)
}



func Execute(){


	if err := root.Execute(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
