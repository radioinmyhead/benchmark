package benchmark

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

var fiorequire = []string{"fio"}
var fioargs = map[string]interface{}{
	"--status-interval": 1,
	"--iodepth":         "32",
	"--time_based":      "",
	"--numjobs":         "8",
	"--group_reporting": "",
	"--ioengine":        "libaio",
	"--direct":          "1",
	"--norandommap":     "",
	"--randrepeat":      "0",
	"--output-format":   "json",
}

type fioResult struct {
	name string
	body string
}

type Disk struct {
	filename string
	fio      map[string]interface{}
	result   [][]map[string]interface{}
}

func NewDisk(file string) (d *Disk, err error) {
	for _, com := range fiorequire {
		if _, err = exec.LookPath(com); err != nil {
			return
		}
	}
	d = &Disk{
		filename: file,
		fio:      fioargs,
	}
	d.fio["--filename"] = file
	return
}

func (d *Disk) runfio() (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	args := []string{}
	for k, v := range d.fio {
		if v == "" {
			args = append(args, k)
		} else {
			args = append(args, fmt.Sprintf("%v=%v", k, v))
		}
	}
	logrus.Infof("start %v", d.fio["--name"])
	fmt.Println("fio", args)
	c := exec.CommandContext(ctx, "fio", args...)
	out, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	err = c.Start()
	if err != nil {
		return
	}

	i, err := strconv.Atoi(fmt.Sprintf("%v", d.fio["--status-interval"]))
	if err != nil {
		return
	}
	var restartDuration = time.Duration(i) * time.Hour //time.Second
	timer := time.AfterFunc(restartDuration, func() {
		logrus.Infof("wait too lang")
		cancel()
	})
	defer timer.Stop()

	list := []map[string]interface{}{}
	dec := json.NewDecoder(out)
	for {
		js := make(map[string]interface{})
		if err := dec.Decode(&js); err == io.EOF {
			break
		} else if err != nil {
			cancel()
			return err
		}
		timer.Reset(restartDuration)
		js["timestamp"] = time.Now().Unix()
		list = append(list, js)
	}

	d.result = append(d.result, list)
	return c.Wait()
}

func (d *Disk) diskinit() (err error) {
	d.fio["--bs"] = "4k"
	d.fio["-rw"] = "randwrite"
	d.fio["--runtime"] = "36000"
	d.fio["--name"] = fmt.Sprintf("%v_%v_init", d.fio["--bs"], d.fio["-rw"])
	return d.runfio()
}

func (d *Disk) diskrw() (err error) {
	d.fio["--bs"] = "4k"
	d.fio["-rw"] = "randwrite"
	d.fio["--runtime"] = "36000"
	d.fio["--name"] = fmt.Sprintf("%v_%v", d.fio["--bs"], d.fio["-rw"])
	return d.runfio()
}

func (d *Disk) diskrr() (err error) {
	d.fio["--bs"] = "4k"
	d.fio["-rw"] = "randread"
	d.fio["--runtime"] = "36000"
	d.fio["--name"] = fmt.Sprintf("%v_%v", d.fio["--bs"], d.fio["-rw"])
	return d.runfio()
}

func (d *Disk) disksw() (err error) {
	d.fio["--bs"] = "1024k"
	d.fio["-rw"] = "write"
	d.fio["--runtime"] = "36000"
	d.fio["--name"] = fmt.Sprintf("%v_%v", d.fio["--bs"], d.fio["-rw"])
	return d.runfio()
}
func (d *Disk) disksr() (err error) {
	d.fio["--bs"] = "1024k"
	d.fio["-rw"] = "read"
	d.fio["--runtime"] = "36000"
	d.fio["--name"] = fmt.Sprintf("%v_%v", d.fio["--bs"], d.fio["-rw"])
	return d.runfio()
}

func (d *Disk) diskben(bs, rw, runtime, name string) (err error) {
	d.fio["--bs"] = bs
	d.fio["-rw"] = rw
	d.fio["--runtime"] = runtime
	d.fio["--name"] = name
	return d.runfio()

}

func (d *Disk) writeret() (err error) {
	data, err := json.Marshal(d.result)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("fio_result.json", data, 0600); err != nil {
		return err
	}
	return
}

func (d *Disk) Init() (err error) {
	return d.diskinit()
}

func (d *Disk) Benchmark() (err error) {
	// 4k rw
	if err = d.diskrw(); err != nil {
		return
	}
	//  4k rr
	if err = d.diskrr(); err != nil {
		return
	}
	//  1m w
	if err = d.disksw(); err != nil {
		return
	}
	//  1m r
	if err = d.disksr(); err != nil {
		return
	}
	// write tar
	if err = d.writeret(); err != nil {
		return
	}
	return
}
