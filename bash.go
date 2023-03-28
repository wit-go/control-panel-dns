package main

import (
        "io"
        "os"
        "os/exec"
        "os/signal"
        "syscall"

        "github.com/creack/pty"
        "golang.org/x/term"
)

func test() error {
        // Create arbitrary command.
        c := exec.Command("bash")

        // Start the command with a pty.
        ptmx, err := pty.Start(c)
        if err != nil {
                return err
        }
        // Make sure to close the pty at the end.
        defer func() { _ = ptmx.Close() }() // Best effort.

        // Handle pty size.
        ch := make(chan os.Signal, 1)
        signal.Notify(ch, syscall.SIGWINCH)
        go func() {
                for range ch {
                        if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
                                log("error resizing pty: %s", err)
                        }
                }
        }()
        ch <- syscall.SIGWINCH // Initial resize.
        defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

        // Set stdin in raw mode.
        oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
        if err != nil {
                panic(err)
        }
        defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

        // Copy stdin to the pty and the pty to stdout.
        // NOTE: The goroutine will keep reading until the next keystroke before returning.
        go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
        _, _ = io.Copy(os.Stdout, ptmx)

        return nil
}

func mainBash() {
        if err := test(); err != nil {
		log(logError, "exit in mainBash()")
                exit(err)
        }
}
