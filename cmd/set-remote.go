package cmd

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
)

// SetRemote configurates the push on a local server
// got set-remote local /path
func SetRemote(remoteType model.RemoteType, args []string) error {
	var err error
	switch remoteType {
	case model.Local:
		if len(args) == 3 { //set-remote type path
			err = utils.ConfigLocalPush(args[2])
		} else {
			return fmt.Errorf("invalid number of arguments")
		}
	case model.Remote:
		//TODO
	case model.Git:
		//TODO
	}
	return err
}
