/*
 * Copyright (c) 2025 Arya Bakhtiari
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */
package main

import (
	"database/sql"
	_ "embed"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
)

var Version = "(untracked)"

func main() {
	//
	//
	//
	done := make(chan bool, 1)

	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		<-sigch

		done <- true
	}()

	go func() {

		cmd_root()

		done <- true
	}()

	<-done
}

func cmd_root() {
	//
	//
	//
	// TODO(AABIB): Add usage info to log.println
	var arg_count int = len(os.Args)
	if arg_count < 2 {
		log.Println("usage boilerplate here")

		os.Exit(1)
	}

	//
	//
	//
	// var executable_name string = filepath.Base(os.Args[0])
	var base_dir string = ""
	var with_go_run bool = false // NOTE(AABIB): debug mode?
	var default_data_dir string = ""

	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// probably ran with go run
		with_go_run = true
		base_dir, _ = os.Getwd()
	} else {
		// probably ran with go build
		with_go_run = false
		base_dir = filepath.Dir(os.Args[0])
	}

	if base_dir == "" { // ASSERTION
		log.Panic("base_dir SHOULD NOT be an empty string!")
		return
	}

	default_data_dir = filepath.Join(base_dir, "tk_data")

	if with_go_run {
		// just to bypass with_go_run not used error
	}
	if default_data_dir == "" {
		// just to bypass default_data_dir not used error
	}

	//
	//
	//
	switch strings.ToLower(os.Args[1]) {
	case "-h", "--help", "help":
		{
			log.Println("help")
		}

	case "-v", "--version", "version":
		{
			log.Println("version")
		}

	case "-s", "--serve", "serve":
		{
			cmd_serve()
		}

	case "superuser":
		{
			cmd_superuser()
		}

	default:
		{
			log.Println("not valid args")
		}

	}

}

func cmd_serve() {

	// setup sqlite

	//
	//
	//
	db, err := sql.Open("sqlite3", "file:taagak.sqlite?_journal=WAL")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

func cmd_superuser() {
	switch strings.ToLower(os.Args[2]) {
	case "-h", "--help", "help":
		{
			log.Println("help")
		}

	case "create":
		{
			log.Println("superuser create")
		}

	case "update":
		{
			log.Println("superuser update")
		}

	case "delete":
		{
			log.Println("superuser delete")
		}

	default:
		{
			log.Println("superuser default")
		}

	}
}
