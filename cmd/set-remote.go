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
		if len(args) == 3 { //set-remote local path
			err = utils.ConfigLocalPush(args[2])
		} else {
			return fmt.Errorf("invalid number of arguments")
		}
	case model.Remote:
		if len(args) == 5 { //set-remote remote ip user path
			ip := args[2]
			user := args[3]
			path := args[4]
			err = utils.ConfigRemotePush(ip, user, path)
		}
	case model.Git: //set-remote git https://github.com/user/repo.git
		//TODO
	}
	return err
}
