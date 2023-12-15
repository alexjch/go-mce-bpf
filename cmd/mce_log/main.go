package main

import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	bpf "github.com/aquasecurity/libbpfgo"
)

const (
	modulePath      = "mce_log.o"
	timeoutPollSecs = 10
)

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	module, err := bpf.NewModuleFromFile(modulePath)
	if err != nil {
		panic(err)
	}
	defer module.Close()

	module.BPFLoadObject()
	if err != nil {
		panic(err)
	}

	prog, err := module.GetProgram("handle_mce")
	if err != nil {
		panic(err)
	}

	_, err = prog.AttachTracepoint("mce", "mce_record")
	if err != nil {
		panic(err)
	}

	events := make(chan []byte)
	rb, err := module.InitRingBuf("mce_events", events)
	if err != nil {
		panic(err)
	}

	rb.Poll(timeoutPollSecs)

	go func() {
		for {
			event := <-events
			if len(event) > 0 {
				pid := int(binary.LittleEndian.Uint32(event[0:4]))
				comm := string(bytes.TrimRight(event[4:], "\x00"))
				fmt.Printf("Event received pid: %d cmd: %s\n", pid, comm)
			}
		}
	}()

	<-done
	fmt.Println("Exiting...")
	rb.Stop()
	rb.Close()
}
