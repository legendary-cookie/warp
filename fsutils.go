package main

import "os"

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func StringToByte(toconv string) []byte {
	return []byte(toconv)
}

func ByteToString(toconv []byte) string {
	return string(toconv[:])
}

