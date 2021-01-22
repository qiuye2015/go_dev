package filetuil

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"log"
	"os"
	"strings"
)

func FileWatcher(actions map[string]func(path string) error) {
	//TODO:全局的 sync.WaitGroup.Add(1)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
	}
	defer watcher.Close()

	for k := range actions {

		if err = watcher.Add(k); err != nil {
			fmt.Println(err)
		}
	}

	//forLoop:
	for {
		select {
		case ev := <-watcher.Events:
			if ev.Op&fsnotify.Write == fsnotify.Write ||
				ev.Op&fsnotify.Remove == fsnotify.Remove {
				log.Println("file changed:", ev.Name, actions)
				if f, ok := actions[ev.Name]; ok {
					if err = f(ev.Name); err != nil {
						fmt.Println(err)
					}
				}
			}
			if err := watcher.Add(ev.Name); err != nil {
				fmt.Println(err)
			}
		case err := <-watcher.Errors:
			fmt.Println(err)
			//case context.Context.Done():
			//  break forLoop
		}
	}
	//全局的 sync.WaitGroup.Done()
}

func ReadFromFile(path string) (map[string]struct{}, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buffer := bufio.NewReader(f)
	resMap := make(map[string]struct{})
	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			if err != io.EOF || line == "" {
				break
			}
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			resMap[line] = struct{}{}
		}
	}
	return resMap, nil
}
