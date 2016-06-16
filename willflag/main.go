package main

import "fmt"
import "flag"
import "os"

func main() {
	args := os.Args
	//args := []string{"cfs"}
	//args := []string{"cfs", "-h"}

	//args := []string{"cfs", "create", "-h"}
	//args := []string{"cfs", "create", "test"}
	//args := []string{"cfs", "create", "-r", "aio", "test"}
	//args := []string{"cfs", "create", "-r=aio", "test"}
	//args := []string{"cfs", "create", "-z", "badoption", "test"} // this should error
	//args := []string{"cfs", "create", "-r", "aio", "test", "extra_arg"} // this should error

	//args := []string{"cfs", "mount"}
	//args := []string{"cfs", "mount", "-h"}
	//args := []string{"cfs", "mount", "myfs", "/mnt/cfs"}
	//args := []string{"cfs", "mount", "-o", "debug,allow_other", "myfs", "/mnt/cfs"}
	//args := []string{"cfs", "mount", "-o", "debug,allow_other", "myfs", "/mnt/cfs", "extra_arg"} // this should error

	//args := []string{"cfs", "badcommand"} // this should error

	f := flag.NewFlagSet("base", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  cfs <command> [options] <args>")
		fmt.Println("Commands:")
		fmt.Println("  create    Create a new filesystem")
		fmt.Println("  mount     Mount an existing filesystem")
		os.Exit(1)
	}
	f.Parse(args[1:])
	if f.NArg() == 0 {
		f.Usage()
	}
	command := f.Args()[0]
	switch command {
	default:
		fmt.Println("Invalid command")
		f.Usage()
	case "create":
		f := flag.NewFlagSet("create", flag.ContinueOnError)
		f.Usage = func() {
			fmt.Println("Usage:")
			fmt.Println("  cfs create [options] <filesystem>")
			fmt.Println("Options:")
			fmt.Println("  -r, --region    region to create the filesystem (default is iad)")
			fmt.Println("Examples:")
			fmt.Println("  cfs create myfilesystem")
			fmt.Println("  cfs create -r aio mytestfs")
			os.Exit(1)
		}
		var region string
		f.StringVar(&region, "r", "iad", "")
		f.StringVar(&region, "region", "iad", "")
		f.Parse(args[2:])
		if f.NArg() == 0 {
			f.Usage()
		} else if f.NArg() > 1 {
			fmt.Println("Too many arguments")
			f.Usage()
		}
		name := f.Args()[0]
		fmt.Printf("creating filesystem named %s in the %s region", name, region)
	case "mount":
		f := flag.NewFlagSet("mount", flag.ContinueOnError)
		f.Usage = func() {
			fmt.Println("Usage:")
			fmt.Println("  cfs mount [options] <filesystem> <mountpoint>")
			fmt.Println("Options:")
			fmt.Println("  -o, --options    mount options (default is rw)")
			fmt.Println("                   debug")
			fmt.Println("                   allow_other")
			fmt.Println("Examples:")
			fmt.Println("  cfs mount myfs /mnt/myfs")
			fmt.Println("  cfs mount -o debug,allow_other mytestfs /mnt/myfs")
			os.Exit(1)
		}
		var options string
		f.StringVar(&options, "o", "rw", "")
		f.StringVar(&options, "options", "rw", "")
		if err := f.Parse(args[2:]); err == nil {
			if f.NArg() == 0 {
				f.Usage()
			} else if f.NArg() < 2 {
				fmt.Println("Not enough arguments")
				f.Usage()
			} else if f.NArg() > 2 {
				fmt.Println("Too many arguments")
				f.Usage()
			}
			filesystem := f.Args()[0]
			mountpoint := f.Args()[1]
			fmt.Printf("mounting filesystem %s on %s with options %s", filesystem, mountpoint, options)
		}
	}
}
