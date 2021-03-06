// Copyright (c) 2016-2017 Daniel Oaks <daniel@danieloaks.net>
// released under the MIT license

package irc

import "github.com/goshuirc/irc-go/ircmsg"

// GetClientOnlyTags takes a tag map and returns a map containing just the client-only tags from it.
func GetClientOnlyTags(tags map[string]ircmsg.TagValue) *map[string]ircmsg.TagValue {
	if len(tags) < 1 {
		return nil
	}

	clientOnlyTags := make(map[string]ircmsg.TagValue)

	for name, value := range tags {
		if len(name) > 1 && name[0] == '+' {
			clientOnlyTags[name] = value
		}
	}

	if len(clientOnlyTags) < 1 {
		return nil
	}

	return &clientOnlyTags
}
