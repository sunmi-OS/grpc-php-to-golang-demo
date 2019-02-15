package main

import (
	"github.com/labstack/echo"
	"sort"
	"os"
	"github.com/urfave/cli"
	"github.com/sunmi-OS/gocore/api"
	"google.golang.org/grpc"
	pb "grpc-php-golang/protobuf/go-server/helloworld"
	"log"
	"golang.org/x/net/context"
)

type EchoApi struct {
}

var eApi EchoApi

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func (a *EchoApi) echoStart(c *cli.Context) error {
	// Echo instance
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	pbc := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// Route => handler
	e.GET("/", func(c echo.Context) error {

		//request := api.NewRequest(c)
		response := api.NewResponse(c)

		r, err := pbc.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)

		return response.RetSuccess(r.Message)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
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
