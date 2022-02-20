package main

import (
	"context"
	"fmt"

	"logit/v1/package/api"
	"logit/v1/util/auth"
	"logit/v1/util/config"
	"logit/v1/util/db"
	mid "logit/v1/util/middleware"
	"logit/v1/util/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	env := config.Init()
	client := db.Connect(context.Background(), env)
	s := server.Init(env)
	e := s.Start()
	jwtService, err := auth.Init(env)
	handleError(err)
	ap := api.Init(client, jwtService, mid.JwtAuth(jwtService))
	ap.Route(e)
	e.Logger.Fatal(e.Start(s.Port))
}

func handleError(e error) {

	if e != nil {
		fmt.Println(e.Error())
	}

}

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	// sshClientDomain "../../../domain/client"
// 	"golang.org/x/crypto/ssh"
// )

// type SshClient struct {
// }

// //  func InitSshClient() sshClientDomain.ISshClient {
// // 	sshClient := &SshClient{}
// // 	return sshClient
// //  }
// type SshCommandExecuter struct {
// 	sshClient SshClient
// }

// func InitSshClient() *SshClient {
// 	sshClient := &SshClient{}
// 	return sshClient
// }
// func InitSshExecuter(sc SshClient) *SshCommandExecuter {
// 	executer := &SshCommandExecuter{
// 		sshClient: sc,
// 	}
// 	return executer
// }
// func (sc *SshClient) GetConnection(host string, username string, password string) ssh.Client {
// 	config := &ssh.ClientConfig{
// 		User: username,
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password(password),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}
// 	sshConn, err := ssh.Dial("tcp", host, config)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return ssh.Client{}
// 	}
// 	return *sshConn
// }

// func (e *SshCommandExecuter) RunSshCommand(host string, username string, password string, cmd string) string {
// 	client := e.sshClient.GetConnection(host, username, password)
// 	fmt.Println("clent ", client)
// 	sshSession, err := client.NewSession()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return ""
// 	}
// 	// defer sshSession.Close()
// 	// defer client.Close()
// 	var stdoutBuf bytes.Buffer

// 	sshSession.Stdout = &stdoutBuf
// 	sshSession.Run(cmd)

// 	return stdoutBuf.String()
// }
// func main() {
// 	sshC := InitSshClient()
// 	e := InitSshExecuter(*sshC)
// 	c := e.RunSshCommand("localhost", "rishi", "rishi", "ls")
// 	fmt.Println("c hai ", c)
// 	// c.
// }
