/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"

	// "github.com/felipefrmelo/imersao-fullcycle/infrastructure/repository"

	"github.com/felipefrmelo/imersao-fullcycle/application/grpc"
	"github.com/felipefrmelo/imersao-fullcycle/domain/model"
	"github.com/felipefrmelo/imersao-fullcycle/infrastructure/db"
	"github.com/felipefrmelo/imersao-fullcycle/infrastructure/repository"

	// "github.com/felipefrmelo/imersao-fullcycle/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var database *gorm.DB

var portNumber int

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpc called")

		database = db.ConnectDB(os.Getenv("env"))
		code := "001"
		name := "Banco do Brasil"
		bank, _ := model.NewBank(code, name)

		accountNumber := "abcnumber"
		ownerName := "Felipe"
		account, _ := model.NewAccount(bank, accountNumber, ownerName)
		repo := repository.PixKeyRepositoryDb{Db: database}
		repo.AddAccount(account)
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC Server port")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
