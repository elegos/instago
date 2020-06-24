package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/elegos/instago/source"
	"github.com/elegos/instago/source/processor"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type config struct {
	AutoConvert bool `yaml:"autoConvert"`
}

func getConfig() (conf config, err error) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if _, err := os.Stat(dir + "/config.yml"); err != nil {
		logrus.WithError(err).Warn("No configuration file found")
		return conf, err
	}

	bytes, err := ioutil.ReadFile(dir + "/config.yml")
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(bytes, &conf)

	return conf, err
}

func main() {
	conf, _ := getConfig()

	if len(os.Args) < 2 && !conf.AutoConvert {
		logrus.Errorf("Usage: %s input_image", os.Args[0])

		os.Exit(source.ExitCodeArgs)
	}

	toProcess := []string{}
	if len(os.Args) > 1 {
		toProcess = os.Args[1:]
	} else if conf.AutoConvert {
		expr := regexp.MustCompile("(?i).(jpe?g|gif|png)$")
		filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if expr.MatchString(path) {
				toProcess = append(toProcess, path)
			}
			return nil
		})
	}

	for _, file := range toProcess {
		processor.ProcessImage(file)
	}
}
