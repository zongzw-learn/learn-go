package main

import (
	"github.com/energye/energy/v2/cef"
)

func main() {
	//Global initialization
	cef.GlobalInit(nil, nil)
	//Create an application
	app := cef.NewApplication()
	//Specify a URL address or local HTML file directory
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	//Run Application
	cef.Run(app)
}
