package utils

import "path/filepath"

var DefaultProjectSource = "vineelsai26"

var DefaultProjectSourceDir = "/Volumes/Data/GitHub"

var ProjectSourceDir = GetProjectSourceRootDir()

var ConfigFilePath = filepath.Join(GetHome(), ".checkout", "source_dir")

var ProjectCheckoutRootDir = filepath.Join(GetHome(), "Personal")
