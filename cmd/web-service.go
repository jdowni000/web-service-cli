package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var webServiceCmd = &cobra.Command{
	Use:   "webservice",
	Short: "A simple cli tool that hits a local running image/container on port 8080",
	Long: `A simple cli tool designed to help retrieve json data from a local running 
image/container on port 8080. This tool was designed specifically to target 
https://github.com/jdowni000/gameserver. You may hit the root that retrieves condensed
information for all games available, or provide the id with the flag -i with the id number.

For example: 
	web-service-cli webservice
	or
	web-service-cli webservice -i 1
`,
	Run: webSericeCli,
}

func init() {
	rootCmd.AddCommand(webServiceCmd)
	webServiceCmd.Flags().StringP("id", "i", "", "specific game to retrieve information")

}

func webSericeCli(cmd *cobra.Command, args []string) {
	host := "localhost"
	port := "8080"
	path := "/"
	id, _ := cmd.Flags().GetString("id")

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "8080"
	}

	if id != "" {
		path = fmt.Sprintf("/game?id=${%s}", id)
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println()

	fmt.Fprintf(conn, "GET %s HTTP/1.0\r\nHost: %s\r\n\r\n", path, host)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf[:n]))
}
