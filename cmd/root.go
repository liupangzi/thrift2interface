package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/liupangzi/thrift2interface/algorithms"
	"github.com/liupangzi/thrift2interface/listener"
	"github.com/liupangzi/thrift2interface/parser"
	"github.com/liupangzi/thrift2interface/template"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "thrift2interface",
	Short:   "Generate go interfaces for gomock testing from thrift files by Antlr4",
	Version: "0.2.0",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var packageName, prefix, thrift string

func init() {
	RootCmd.Flags().StringVarP(&packageName, "package-name", "n", "", "package name for generated files")
	if err := RootCmd.MarkFlagRequired("package-name"); err != nil {
		os.Exit(-1)
	}

	RootCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "relative path")
	if err := RootCmd.MarkFlagRequired("prefix"); err != nil {
		os.Exit(-2)
	}

	RootCmd.Flags().StringVarP(&thrift, "thrift", "t", "", "thrift file")
	if err := RootCmd.MarkFlagRequired("thrift"); err != nil {
		os.Exit(-3)
	}
}

func run() {
	logger := logrus.New()
	logger.Level = logrus.InfoLevel
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	mi := template.NewMetaInterface()
	mi.PackageName = packageName
	mi.Prefix = func(s string) string {
		if strings.HasSuffix(s, "/") {
			return s
		}
		return s + "/"
	}(prefix)
	logger.Infof("mi.Prefix: %s", mi.Prefix)

	// get the absolute path of thrift file
	thriftFile := thrift
	if !strings.HasPrefix(thrift, "/") {
		cwd, _ := os.Getwd()
		thriftFile = cwd + "/" + thrift
	}
	logger.Infof("thriftFile: %s", thriftFile)

	if _, err := os.Stat(thriftFile); os.IsNotExist(err) {
		logger.Errorf("Open file `%s` err", thriftFile)
		os.Exit(1)
	}

	fs, err := antlr.NewFileStream(thriftFile)
	if err != nil {
		logger.Errorf("antlr.NewFileStream err: %v", err)
		os.Exit(1)
	}

	p := parser.NewThriftParser(antlr.NewCommonTokenStream(parser.NewThriftLexer(fs), antlr.TokenDefaultChannel))
	p.BuildParseTrees = true

	antlr.ParseTreeWalkerDefault.Walk(&listener.RootThriftListener{
		Logger: logger,
		MI:     mi,
		Thrift: algorithms.SimplifyUnixDirectoryPath(thriftFile),
	}, p.Document())

	fmt.Println(mi.String())
}
