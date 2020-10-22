package main

import (
	watch "github.com/USER/watch/watch"
)

func main() {
	var filename = "example.ini"
	var listener watch.Listener
	listener = watch.ListenFunc(func(string) {})

	watch.Watch(filename, listener)
	config, _ := watch.Read_file(filename)
	watch.Print_config(config)
}
