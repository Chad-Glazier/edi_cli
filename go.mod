module github.com/Chad-Glazier/edi_cli

go 1.26.2

require (
	github.com/Chad-Glazier/edi v0.0.1
	github.com/spf13/cobra v1.10.2
)

replace github.com/Chad-Glazier/edi => ../edi

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	golang.org/x/sys v0.43.0 // indirect
	golang.org/x/term v0.42.0
)
