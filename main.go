/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ThunderGod77/domain-manager/cmd"
	_ "github.com/ThunderGod77/domain-manager/cmd/delete"
	_ "github.com/ThunderGod77/domain-manager/cmd/get"
	_ "github.com/ThunderGod77/domain-manager/cmd/new"
)

func main() {
	cmd.Execute()
}
