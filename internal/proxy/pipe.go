package proxy

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/pipeproxy/pipe/bind"
	"github.com/wzshiming/logger"
	"sigs.k8s.io/yaml"
)

const (
	ConfigFile = "pipe.yml"
	PidFile    = "pipe.pid"
)

type Pipe struct {
	basePath string
	process  *os.Process
}

func NewPipe(basePath string) *Pipe {
	return &Pipe{
		basePath: basePath,
	}
}

var noneOnceConfig, _ = bind.Marshal(bind.ServiceOnceConfig{bind.WaitService{}})

func init() {
	noneOnceConfig, _ = yaml.JSONToYAML(noneOnceConfig)
}

func (p *Pipe) Run(ctx context.Context) error {
	log := logger.FromContext(ctx)
	configFile := filepath.Join(p.basePath, ConfigFile)
	if _, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			ioutil.WriteFile(configFile, noneOnceConfig, 0644)
		}
	}

	cmd := exec.Command("pipe", "-c", ConfigFile, "-p", PidFile)
	log.Info("start", "exec", fmt.Sprintf("pipe -c %s -p %s", ConfigFile, PidFile))
	cmd.Dir = p.basePath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		<-ctx.Done()
		err := cmd.Process.Signal(syscall.SIGINT)
		if err != nil {
			log.Error(err, "Signal")
		}
	}()
	p.process = cmd.Process
	return cmd.Wait()
}

func (p *Pipe) Signal(sig os.Signal) error {
	if p.process == nil {
		return nil
	}
	return p.process.Signal(sig)
}
