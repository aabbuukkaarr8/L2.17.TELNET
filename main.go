package main

import (
	"L2.17-TELNET/internal/flags"
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	flags.ParseFlags() //парсим флаги

	address := net.JoinHostPort(flags.Host, flags.Port)

	conn, err := net.DialTimeout("tcp", address, flags.Timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Usage: %s [--timeout=10s] host port\n", os.Args[0])
		os.Exit(1)
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		done <- struct{}{}
	}()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			_, err := fmt.Fprintln(conn, scanner.Text())
			if err != nil {
				fmt.Fprintln(os.Stderr, "writing standard input:", err)
				break
			}
		}
		// Пользователь закончил ввод (Ctrl+D)
		conn.Close()
	}()

	<-done

}
