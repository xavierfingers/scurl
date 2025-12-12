
package main
import (
 "fmt"
 "github.com/spf13/cobra"
 "io"
 "net/http"
)
func main() {
rootCmd := cobra.Command{
  Use: "main",
  Run: func(cmd *cobra.Command, args []string) {
  fmt.Println("Enter a host to connect to: ")
  var host string
  fmt.Scanln(&host)
  url := "https://" + host
  resp, err := http.Get(url)
  body, err := io.ReadAll(resp.Body)
  if err != nil {
   fmt.Println("An unexpected error occurred.")
 }
 fmt.Println(string(body))
 },
}
rootCmd.Execute()
}

