package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/liupangzi/thrift2interface/algorithms"
	"github.com/liupangzi/thrift2interface/parser"
	"github.com/liupangzi/thrift2interface/template"
	"github.com/sirupsen/logrus"
)

type RootThriftListener struct {
	*parser.BaseThriftListener

	MI     *template.MetaInterface
	Logger *logrus.Logger

	Thrift        string
	FullNamespace string
	Namespace     string
}

// EnterDocument: parse RootListener's thrift file name, namespace and full namespace
func (r *RootThriftListener) EnterDocument(ctx *parser.DocumentContext) {
	for _, header := range ctx.AllHeader() {
		headerNamespace := header.(*parser.HeaderContext).Namespace()
		if headerNamespace == nil {
			continue
		}

		namespace := headerNamespace.(*parser.NamespaceContext)
		// namespace = [["py", "namespace"], ["go", "namespace"], etc...]
		if len(namespace.AllIDENTIFIER()) == 2 && namespace.IDENTIFIER(0).GetText() == "go" {
			ns := strings.Split(namespace.IDENTIFIER(1).GetText(), ".")
			r.MI.ThriftFiles[r.Thrift] = ""
			r.FullNamespace = namespace.IDENTIFIER(1).GetText()
			r.Namespace = ns[len(ns)-1]
		} else {
			skippedNamespaces := make([]string, 0)
			for _, s := range namespace.AllIDENTIFIER() {
				skippedNamespaces = append(skippedNamespaces, s.GetText())
			}
			r.Logger.Infof("Skip non-golang namespace: %s", strings.Join(skippedNamespaces, " "))
		}
	}
}

// EnterInclude: parse included thrift files recursively
func (r *RootThriftListener) EnterInclude(ctx *parser.IncludeContext) {
	leafThrift := algorithms.SimplifyUnixDirectoryPath(strings.Join([]string{
		path.Dir(r.Thrift),
		ctx.LITERAL().GetText()[1 : len(ctx.LITERAL().GetText())-1],
	}, string(os.PathSeparator)))
	r.Logger.Infof("parsing leaf thrift file %s...", leafThrift)

	// quit if current thrift file has been parsed
	if _, ok := r.MI.ThriftFiles[leafThrift]; ok {
		r.Logger.Infof("leaf thrift file %s has been parsed", leafThrift)
		return
	}

	if _, err := os.Stat(leafThrift); os.IsNotExist(err) {
		r.Logger.Errorf("Open thrift file `%s` failed", leafThrift)
		os.Exit(1)
	}

	fs, err := antlr.NewFileStream(leafThrift)
	if err != nil {
		r.Logger.Errorf("antlr.NewFileStream err: %v", err)
		os.Exit(1)
	}

	p := parser.NewThriftParser(antlr.NewCommonTokenStream(parser.NewThriftLexer(fs), antlr.TokenDefaultChannel))
	p.BuildParseTrees = true

	listener := &LeafThriftListener{
		Logger: r.Logger,
		MI:     r.MI,
		Thrift: algorithms.SimplifyUnixDirectoryPath(leafThrift),
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Document())
}

// EnterService: get current interface's name
func (r *RootThriftListener) EnterService(ctx *parser.ServiceContext) {
	r.MI.ServiceName = strings.Title(ctx.IDENTIFIER(0).GetText())
	r.Logger.Infof("ServiceName: %s", r.MI.ServiceName)
}

// EnterGoStruct: get full namespace and raw name of prefixed method
func (r *RootThriftListener) EnterGoStruct(ctx *parser.GoStructContext) {
	// prefixed.method => full namespace
	r.MI.PrefixedStructToFullNamespace[ctx.IDENTIFIER().GetText()] = r.FullNamespace
	// prefixed.method => raw name
	r.MI.PrefixedStructToName[ctx.IDENTIFIER().GetText()] = ctx.IDENTIFIER().GetText()
	r.Logger.Infof("RootThriftListener: prefixed method=%s, raw name=%s, full namespace=%s", ctx.IDENTIFIER().GetText(), ctx.IDENTIFIER().GetText(), r.FullNamespace)
}

// EnterFunction: get res/res/method name
func (r *RootThriftListener) EnterFunction(ctx *parser.FunctionContext) {
	method, _ := template.NewMethod()
	method.Name = ctx.IDENTIFIER().GetText()

	if ctx.Function_type().GetText() != "void" {
		method.Response = ctx.Function_type().GetText()
	} else {
		method.Response = ""
	}

	if len(ctx.AllField()) > 0 {
		method.Request = ctx.Field(0).(*parser.FieldContext).Field_type().GetText()
	} else {
		method.Request = ""
	}

	r.MI.Methods = append(r.MI.Methods, method)
	r.Logger.Infof("append method: %+v", method)
}

// LeafThriftListener inclusive thrift listener
type LeafThriftListener struct {
	*parser.BaseThriftListener

	MI     *template.MetaInterface
	Logger *logrus.Logger

	Thrift        string
	FullNamespace string
	Namespace     string
}

func (l *LeafThriftListener) EnterDocument(ctx *parser.DocumentContext) {
	for _, header := range ctx.AllHeader() {
		headerNamespace := header.(*parser.HeaderContext).Namespace()
		if headerNamespace == nil {
			continue
		}

		namespace := headerNamespace.(*parser.NamespaceContext)
		if len(namespace.AllIDENTIFIER()) == 2 && namespace.IDENTIFIER(0).GetText() == "go" {
			l.MI.ThriftFiles[l.Thrift] = ""
			ns := strings.Split(namespace.IDENTIFIER(1).GetText(), ".")
			l.FullNamespace = namespace.IDENTIFIER(1).GetText()
			l.Namespace = ns[len(ns)-1]
		} else {
			skippedNamespaces := make([]string, 0)
			for _, s := range namespace.AllIDENTIFIER() {
				skippedNamespaces = append(skippedNamespaces, s.GetText())
			}
			l.Logger.Infof("Skip non-golang namespace: %s", strings.Join(skippedNamespaces, " "))
		}
	}
}

func (l *LeafThriftListener) EnterInclude(ctx *parser.IncludeContext) {
	leafThrift := algorithms.SimplifyUnixDirectoryPath(strings.Join([]string{
		path.Dir(l.Thrift),
		ctx.LITERAL().GetText()[1 : len(ctx.LITERAL().GetText())-1],
	}, string(os.PathSeparator)))

	if _, ok := l.MI.ThriftFiles[leafThrift]; ok {
		l.Logger.Infof("leaf thrift file %s has been parsed", leafThrift)
		return
	}

	if _, err := os.Stat(leafThrift); os.IsNotExist(err) {
		l.Logger.Errorf("Open thrift file `%s` failed", leafThrift)
		os.Exit(1)
	}

	fs, err := antlr.NewFileStream(leafThrift)
	if err != nil {
		l.Logger.Errorf("antlr.NewFileStream err: %v", err)
		os.Exit(1)
	}

	p := parser.NewThriftParser(antlr.NewCommonTokenStream(parser.NewThriftLexer(fs), antlr.TokenDefaultChannel))
	p.BuildParseTrees = true

	listener := &LeafThriftListener{
		Logger: l.Logger,
		MI:     l.MI,
		Thrift: leafThrift,
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Document())
}

func (l *LeafThriftListener) EnterGoStruct(ctx *parser.GoStructContext) {
	l.MI.PrefixedStructToFullNamespace[l.Namespace+"."+ctx.IDENTIFIER().GetText()] = l.FullNamespace
	l.MI.PrefixedStructToName[l.Namespace+"."+ctx.IDENTIFIER().GetText()] = ctx.IDENTIFIER().GetText()
	l.Logger.Infof("LeafThriftListener: prefixed method=%s, raw name=%s, full namespace=%s", l.Namespace+"."+ctx.IDENTIFIER().GetText(), ctx.IDENTIFIER().GetText(), l.FullNamespace)
}

func main() {
	logger := logrus.New()
	logger.Level = logrus.InfoLevel
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	packageName := flag.String("package", "package", "package name for generated files")
	prefix := flag.String("prefix", "github.com/group/project/", "relative path of $GOPATH")
	thrift := flag.String("thrift", "idl/base.thrift", "thrift file")
	flag.Parse()

	mi := template.NewMetaInterface()
	mi.PackageName = *packageName
	mi.Prefix = func(s string) string {
		if strings.HasSuffix(s, "/") {
			return s
		}
		return s + "/"
	}(*prefix)
	logger.Infof("mi.Prefix: %s", mi.Prefix)

	// get the absolute path of thrift file
	thriftFile := *thrift
	if !strings.HasPrefix(*thrift, "/") {
		cwd, _ := os.Getwd()
		thriftFile = cwd + "/" + *thrift
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

	listener := &RootThriftListener{
		Logger: logger,
		MI:     mi,
	}
	listener.Thrift = algorithms.SimplifyUnixDirectoryPath(thriftFile)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Document())

	fmt.Println(mi.String())
}
