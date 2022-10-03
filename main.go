/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
  "main/cmd"
  "os"
)

func main() {
  cmd.Execute()
	os.Remove("assets/ciphertext.bin")
	os.Remove("assets/outputs/plaintext.txt")

  encryptFile()
	decryptFile()
  
}