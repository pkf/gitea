package setting

import (
	"strings"
)

var (
	defaultLangs                             = strings.Split("en-US,zh-CN", ",")
	defaultLangNames                         = strings.Split("English,简体中文", ",")
	defaultPullRequestWorkInProgressPrefixes = strings.Split("WIP:,[WIP]", ",")
)
