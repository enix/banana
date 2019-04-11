package main

func main() {
	args := LoadArguments()

	if args.Flags.DisplayHelp {
		Usage()
		return
	}
}
