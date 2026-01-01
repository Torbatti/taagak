package cmd

import "testing"

func TestCMD_root_noArgs(t *testing.T) {
	var err error
	var args []string

	err = CMD_root(args)
	if err != CMD_root_err_lessThanTwoArgs {
		t.Errorf("got %q, expected %q", err, CMD_root_err_lessThanTwoArgs)
	}
}

func TestCMD_root_nonValidArgs(t *testing.T) {
	var err error
	var args []string

	args = append(args, "heelp")
	args = append(args, "vercion")
	args = append(args, "serv")

	err = CMD_root(args)
	if err != CMD_root_err_nonValidArgs {
		t.Errorf("got %q, expected %q", err, CMD_root_err_nonValidArgs)
	}
}
