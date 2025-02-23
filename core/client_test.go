package core_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Wsine/feishu2md/core"
	"github.com/Wsine/feishu2md/utils"
)

func getEnvVars() (string, string, string) {
	utils.LoadEnv()
	appID := os.Getenv("FEISHU_APP_ID")
	appSecret := os.Getenv("FEISHU_APP_SECRET")
	userAccessToken := os.Getenv("FEISHU_APP_USER_ACCESS_TOKEN")
	return appID, appSecret, userAccessToken
}

func TestNewClient(t *testing.T) {
	appID, appSecret, userAccessToken := getEnvVars()
	c := core.NewClient(appID, appSecret, userAccessToken, "feishu.cn")
	if c == nil {
		t.Errorf("Error creating DocClient")
	}
}

func TestDownloadImage(t *testing.T) {
	appID, appSecret, userAccessToken := getEnvVars()
	c := core.NewClient(appID, appSecret, userAccessToken, "feishu.cn")
	imgToken := "boxcnA1QKPanfMhLxzF1eMhoArM"
	filename, err := c.DownloadImage(
		context.Background(),
		imgToken,
		"static",
	)
	if err != nil {
		t.Error(err)
	}
	if filename != "static/"+imgToken+".png" {
		fmt.Println(filename)
		t.Errorf("Error: not expected file extension")
	}
	if err := os.RemoveAll("static"); err != nil {
		t.Errorf("Error: failed to clean up the folder")
	}
}

func TestGetDocxContent(t *testing.T) {
	appID, appSecret, userAccessToken := getEnvVars()
	c := core.NewClient(appID, appSecret, userAccessToken, "feishu.cn")
	docx, blocks, err := c.GetDocxContent(
		context.Background(),
		"doxcnXhd93zqoLnmVPGIPTy7AFe",
	)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(docx.Title)
	if docx.Title == "" {
		t.Errorf("Error: parsed title is empty")
	}
	fmt.Printf("number of blocks: %d\n", len(blocks))
	if len(blocks) == 0 {
		t.Errorf("Error: parsed blocks are empty")
	}
}

func TestGetWikiNodeInfo(t *testing.T) {
	appID, appSecret, userAccessToken := getEnvVars()
	c := core.NewClient(appID, appSecret, userAccessToken, "feishu.cn")
	const token = "wikcnLgRX9AMtvaB5x1cl57Yuah"
	node, err := c.GetWikiNodeInfo(context.Background(), token)
	if err != nil {
		t.Error(err)
	}
	if node.ObjType != "docx" {
		t.Errorf("Error: node type incorrect")
	}
}
