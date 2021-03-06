package proc

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"path/filepath"
)

type UpTimeInfo struct {
	Uptime		float64	// the uptime of the system (seconds)
	IdleTime	float64	// the amount of time spent in idle process (seconds)
}

func UpTime() (UpTimeInfo, error) {
	var ut UpTimeInfo
	f, err := os.Open(filepath.Join(Ctx.Procfs, "uptime"))
	if err != nil {
		return ut, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	b, _ := r.ReadBytes('\n')
	s := strings.TrimSpace(string(b))
	vs := strings.Split(s, " ")
	if len(vs) < 2 {
		return ut, fmt.Errorf("splited strings read from %v/uptime incorrect. len(vs): %v\n", Ctx.Procfs, len(vs))
	}
	ut.Uptime, err = strconv.ParseFloat(vs[0], 10)
	if err != nil {
		return ut, err
	}

	ut.IdleTime, err = strconv.ParseFloat(vs[1], 10)
	if err != nil {
		return ut, err
	}
	return ut, nil
}