package main

import (
	"fmt"

	"github.com/cakeatthewalk-cmyk/blog_aggregator.git/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	err = cfg.SetUser("Charles")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("db_url: %s, user: %s", cfg.DbUrl, cfg.CurrentUserName)
}
