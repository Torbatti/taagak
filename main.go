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

	//
	//
	//
	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		<-sigch

		done <- true
	}()

	//
	//
	//
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
		return
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
	// extract help and version cmd
	//
	var cmd_help bool = true
	for i := range arg_count {
		i += 1

		switch strings.ToLower(os.Args[i]) {
		case "-h", "--help", "help":
			{
				cmd_help = true
			}

		case "-v", "--version", "version":
			{
				log.Println("version: ", Version)
			}
		}
	}

	//
	//
	//
	switch strings.ToLower(os.Args[1]) {
	case "-s", "--serve", "serve":
		{ // serve [DOMAIN.COM]
			if cmd_help {
				// TODO(AABIB): ADD serve HELP TEXT
				return
			}

			//
			// check if DOMAIN IS PROVIDED or not
			//
			var http_addr string = ""
			var https_addr string = ""

			if arg_count >= 3 {
				if http_addr == "" {
					http_addr = "0.0.0.0:80"
				}
				if https_addr == "" {
					https_addr = "0.0.0.0:443"
				}
			} else {
				if http_addr == "" {
					http_addr = "127.0.0.1:8090"
				}
			}

			//
			//
			// setup sqlite
			//
			db, err := sql.Open("sqlite3", "file:taagak.sqlite?_journal=WAL")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			return
		}

	case "superuser":
		{
			if cmd_help {
				// TODO(AABIB): ADD superuser HELP TEXT
				return
			}

			if arg_count < 3 {
				// TODO(AABIB): ADD superuser HELP TEXT
				return
			}

			switch strings.ToLower(os.Args[2]) {
			case "-h", "--help", "help":
				{
					// TODO(AABIB): ADD superuser HELP TEXT
				}

			//
			// superuser create [email] [password] [password]
			//
			case "create":
				{
					if arg_count != 6 {
						// TODO(AABIB): ADD superuser create HELP TEXT
						return
					}

					log.Println("superuser create")
				}

			//
			// superuser update [email] [password] [password]
			//
			case "update":
				{
					if arg_count != 6 {
						// TODO(AABIB): ADD superuser update HELP TEXT
						return
					}
					log.Println("superuser update")
				}

			//
			// superuser delete [email]
			//
			case "delete":
				{
					if arg_count != 4 {
						// TODO(AABIB): ADD superuser delete HELP TEXT
						return
					}
					log.Println("superuser delete")
				}

			default:
				{
					// TODO(AABIB): ADD superuser HELP TEXT
				}

			}

			return
		}

	case "-h", "--help", "help":
		{
			// TODO(AABIB): ADD root HELP TEXT
		}

	case "-v", "--version", "version":
		{
			// SKIP
		}

	default:
		{
			log.Println("not valid args")
			// TODO(AABIB): ADD root HELP TEXT
			return
		}
	}

}
