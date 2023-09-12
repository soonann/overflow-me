package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	var charArrLen int
	var charArrData string
	var charArrFiller string
	var overflowTargetType string
	var overflowTargetData string

	var rootCmd = &cobra.Command{
		Use: "overflowme",
		Run: func(cmd *cobra.Command, args []string) {
			var buf []byte
			if charArrData == "" || charArrLen <= 0 {
				return
			}

			if len(charArrData) > charArrLen {
				log.Fatal("The replacement smaller than cLen - 1")
				return
			}

			if overflowTargetType == "uint" {

				var err error
				buf, err = formatUnsignedInt32(overflowTargetData)
				if err != nil {
					log.Fatal(err)
				}

			}

			// figure out how much space is the char arr allocated
			replacementByteCount := 0
			if charArrLen > 4 {
				replacementByteCount += (charArrLen / 4)
			}
			if charArrLen%4 > 0 {
				replacementByteCount += 1
			}
			replacementByteCount *= 4

			fmt.Printf("replacement length is: %d \n\n", replacementByteCount)

			// print the starting string first
			// after that print the ending integer
			// c c c c | c c c \0 | i i i i
			fmt.Printf("payload: ")

			// starting string
			fmt.Print(charArrData)
			for i := len(charArrData); i < replacementByteCount; i++ {
				fmt.Print(charArrFiller)
			}

			// ending integer
			for _, b := range buf {
				fmt.Printf("\\x%02x", b)
			}
			fmt.Println()
		},
	}

	rootCmd.Flags().IntVarP(&charArrLen, "char-arr-len", "l", 0, "Length of char array e.g. char [10] means you should input 10")
	rootCmd.Flags().StringVarP(&charArrData, "char-arr-data", "d", "", "If you want the char that you overflow to display a specific text, you can use this")
	rootCmd.Flags().StringVarP(&charArrFiller, "char-arr-filler", "f", "a", "The character to fill with after the replacement")
	rootCmd.Flags().StringVarP(&overflowTargetData, "overflow-target-data", "o", "", "The value that you want to replace the")
	rootCmd.Flags().StringVarP(&overflowTargetType, "overflow-target-type", "t", "uint", "The data type of the data you're trying to overflow into e.g int uint long ulong char")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func formatUnsignedInt32(val string) ([]byte, error) {
	u32, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(u32))

	return buf, nil
}
