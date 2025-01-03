package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/azurejelly/checkhost/client"
	"github.com/azurejelly/checkhost/utils"
)

var (
	mode         = flag.String("mode", "http", "Type of check to request (http, tcp, udp, dns)")
	maxNodes     = flag.Int("max-nodes", 100, "Maximum amount of nodes to check from")
	nodeList     = flag.String("nodes", "", "Comma-separated list of nodes to use")
	fullNodeList = flag.Bool("display-all-nodes", false, "Whether to display more than 5 nodes on the final output")
	help         = flag.Bool("help", false, "Shows a list of available options")
	yes          = flag.Bool("open-report", false, "Whether to automatically open the report in your web browser")
)

func init() {
	flag.StringVar(mode, "m", "http", "Type of check to request (http, tcp, udp, dns)")
	flag.IntVar(maxNodes, "M", 100, "Maximum amount of nodes to check from")
	flag.StringVar(nodeList, "n", "", "Comma-separated list of nodes to use")
	flag.BoolVar(yes, "y", false, "Whether to automatically open the report in your web browser")
	flag.BoolVar(fullNodeList, "a", false, "Whether to display more than 5 nodes on the final output")
	flag.BoolVar(help, "h", false, "Shows a list of available options")

	flag.Usage = func() {
		fmt.Println(utils.ToolSymbol, " Usage: checkhost [options] [target]")
		fmt.Printf("\n")
		fmt.Println("Options:")
		fmt.Println("  -m, --mode string                 Type of check to request (http,")
		fmt.Println("                                    tcp, udp, dns)")
		fmt.Println("  -M, --max-nodes int               Maximum amount of nodes to check from")
		fmt.Println("  -n, --nodes string                Comma-separated list of nodes to use")
		fmt.Println("  -a, --display-all-nodes boolean   Whether to display more than 5 nodes")
		fmt.Println("                                    on the final output")
		fmt.Println("  -y, --open-report boolean         Whether to automatically open the")
		fmt.Println("                                    report in your web browser")
		fmt.Println("  -h, --help boolean                Shows a list of available options")
	}
}

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0) // they asked for it
	}

	err := utils.ValidateMode(mode)
	if err != nil {
		utils.PrintFatalError("Invalid '--mode' parameter:", err)
	}

	err = utils.ValidateMaxNodes(maxNodes)
	if err != nil {
		utils.PrintFatalError("Invalid '--max-nodes' parameter:", err)
	}

	nodes, err := utils.ParseNodeList(nodeList, maxNodes)
	if err != nil {
		utils.PrintFatalError("Invalid '--nodes' parameter:", err)
	}

	host, err := utils.GetTarget()
	if err != nil {
		flag.Usage()
		os.Exit(-1)
	}

	url, err := client.BuildURL(mode, &host, maxNodes, &nodes)
	if err != nil {
		utils.PrintFatalError("Failed to build check-host.net URL. This might be a bug!", err)
	}

	c, r, err := client.MakeRequest(url)
	if err != nil {
		utils.PrintFatalError("Failed to make request:", err)
	}

	if c != 200 {
		utils.PrintFatalError("Failed to make request:", fmt.Errorf("server replied with unsuccessful status code '%s'", strconv.Itoa(c)))
	}

	if r.Ok != 1 {
		utils.PrintFatalError("Failed to generate the report:", errors.New("server did not reply with the expected 'ok' parameter"))
	}

	utils.PrintSuccess(fmt.Sprint("Successfully requested a ", strings.ToUpper(*mode), " check!"))
	fmt.Println("  • Report:", r.PermanentLink)
	fmt.Println("  • Target:", host)
	fmt.Println("  • Nodes:")

	count := 0
	for n, v := range r.Nodes {
		if !*fullNodeList && len(n) > 5 && count == 5 {
			fmt.Printf("    - ... and %s more node(s)\n", strconv.Itoa(len(n)-count))
			fmt.Print("      (use '--display-all-nodes' for a full list)\n")
			break
		}

		fmt.Printf("    - %s @ %s, %s (%s)\n", n, v.City, v.Country, v.ASNumber)
		count++
	}

	fmt.Printf("\n")

	if *yes {
		utils.OpenURL(r.PermanentLink)
	} else {
		open := utils.Ask("Do you want to open the report on your web browser?")
		if open {
			utils.OpenURL(r.PermanentLink)
		}
	}
}
