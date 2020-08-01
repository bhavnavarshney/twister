//go:generate go run -tags generate gen.go
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/cuminandpaprika/TorqueCalibrationGo/internal/twister"
	"github.com/spf13/afero"

	"github.com/sirupsen/logrus"
	"github.com/zserge/lorca"
)

func main() {
	args := []string{}
	args = append(args, "--start-maximized")
	ui, err := lorca.New("", "", 1300, 800, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(FS))
	err = ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	if err != nil {
		panic(err)
	}

	d := &twister.Drill{}
	d.Log = logrus.New()
	d.FS = afero.NewOsFs()
	ui.Bind("Open", d.Open)
	ui.Bind("GetInfo", d.GetInfo)
	ui.Bind("GetProfile", d.GetProfile)
	ui.Bind("WriteParam", d.WriteParam)
	ui.Bind("Close", d.Close)
	ui.Bind("SaveProfile", d.SaveProfile)
	ui.Bind("LoadProfile", d.LoadProfile)

	// Wait for the browser window to be closed
	<-ui.Done()

	if d.Driver != nil {
		err := d.Driver.Port.Close()
		if err != nil {
			d.Log.Errorln(err)
		}
	}
}
