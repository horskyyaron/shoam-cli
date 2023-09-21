package utils

import "os"

var SHOAM_DIR = os.Getenv("SHOAM_DIR")

var (
	DB_DIR      = SHOAM_DIR + "/db"
	SCRIPTS_DIR = SHOAM_DIR + "/scripts"
	LINKS_DIR   = SHOAM_DIR + "/links"
)
