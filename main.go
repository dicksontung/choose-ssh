package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func LoadConfig() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fmt.Println(homedir)
	viper.AddConfigPath("./")
	viper.AddConfigPath(homedir + "/.cache/")
	viper.AddConfigPath("/etc/choose-ssh/")
	viper.SetConfigName("servers")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	LoadConfig()
	groupNames := make([]string, 0)
	for key, _ := range viper.GetStringMap("all") {
		groupNames = append(groupNames, strings.ToUpper(key))
	}
	prompt1 := promptui.Select{
		Label: "Select Group",
		Items: groupNames,
	}
	_, resultGroup, err := prompt1.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("Selected %q\n", resultGroup)
	connNames := make([]string, 0)
	for key, _ := range viper.GetStringMap(toKey("all", resultGroup)) {
		connNames = append(connNames, strings.ToUpper(key))
	}
	prompt2 := promptui.Select{
		Label: "Select Connection",
		Items: connNames,
	}

	_, resultConnection, err := prompt2.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	ip := viper.GetString(toKey("all", resultGroup, resultConnection, "ip"))
	user := viper.GetString(toKey("all", resultGroup, resultConnection, "user"))
	fmt.Printf("command: \n      ssh %s@%s \n\n", user, ip)
}

func toKey(first string, strings ...string) string {
	result := first
	for _, s := range strings {
		result = result + "." + s
	}
	return result
}
