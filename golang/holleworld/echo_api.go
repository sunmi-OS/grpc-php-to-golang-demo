package main

import (
	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
	"sort"
	"os"
	"github.com/urfave/cli"
	"github.com/sunmi-OS/gocore/api"
	"net/http"
	"io/ioutil"
)

type EchoApi struct {
}

var eApi EchoApi

func (a *EchoApi) echoStart(c *cli.Context) error {
	// Echo instance
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {

		//request := api.NewRequest(c)
		response := api.NewResponse(c)

		resp, err := http.Get("http://172.16.1.63:1323/echo")
		if err != nil {
			// handle error
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}

		return response.RetSuccess(string(body))
	})

	// Start server
	e.Logger.Fatal(e.Start(":1333"))
	return nil
}

func main() {

	app := cli.NewApp()

	// 初始化配置
	//	viper.NewConfig("config", "conf")

	// 指定对于的命令
	app.Commands = []cli.Command{
		{
			Name:    "api",
			Aliases: []string{"a"},
			Usage:   "api",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "开启API-DEMO",
					Action: eApi.echoStart,
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
