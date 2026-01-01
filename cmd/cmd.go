package cmd

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const cmd_root_help_text = `
usage: taagak [arguments]

arguments:
	-v, --version, version
	-h, --help, help
	-s, --serve, serve
	superuser
`

var (
	CMD_root_err_lessThanTwoArgs = errors.New("less than two arguments provided")
	CMD_root_err_nonValidArgs    = errors.New("non valid arguements")
)

func CMD_root(argv []string) error {
	//
	//
	//
	// TODO(AABIB): Add usage info to log.println
	var arg_count int = len(argv)
	if arg_count < 2 {
		log.Print(cmd_root_help_text)
		return CMD_root_err_lessThanTwoArgs
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
		return nil
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
				log.Println("version: ", "")
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
				return nil
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

			return nil
		}

	case "superuser":
		{
			if cmd_help {
				// TODO(AABIB): ADD superuser HELP TEXT
				return nil
			}

			if arg_count < 3 {
				// TODO(AABIB): ADD superuser HELP TEXT
				return nil
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
						return nil
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
						return nil
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
						return nil
					}
					log.Println("superuser delete")
				}

			default:
				{
					// TODO(AABIB): ADD superuser HELP TEXT
				}

			}

			return nil
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
			return CMD_root_err_nonValidArgs
		}
	}

	return nil

}
