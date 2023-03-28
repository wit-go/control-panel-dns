package main

// Watches for changes to a directory. Works cross-platform

import (
    "github.com/fsnotify/fsnotify"
)

// This would be a really dumb way to watch for new network interfaces
// since it then watches a linux only directory /sys/class/net for changes

func watchSysClassNet() {
    // Create new watcher.
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
	    log(logError, "watchSysClassNet() failed:", err)
	    return
    }
    defer watcher.Close()

    // Start listening for events.
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                log("event:", event)
                if event.Has(fsnotify.Write) {
                    log("modified file:", event.Name)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log("error:", err)
            }
        }
    }()

    // Add a path.
    err = watcher.Add("/tmp")
    if err != nil {
	    log(logError, "watchSysClassNet() watcher.Add() failed:", err)
	    return
    }

    // Block main goroutine forever.
    <-make(chan struct{})
}

func fsnotifyNetworkInterfaceChanges() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	// Watch for network interface changes
	err = watcher.Add("/sys/class/net")
	if err != nil {
		return err
	}
	for {
		select {
		case event := <-watcher.Events:
			log("fsnotifyNetworkInterfaceChanges() event =", event)
			if event.Op&fsnotify.Create == fsnotify.Create {
					// Do something on network interface creation
			}
		case err := <-watcher.Errors:
			log("fsnotifyNetworkInterfaceChanges() event err =", err)
			return err
		}
	}
}

