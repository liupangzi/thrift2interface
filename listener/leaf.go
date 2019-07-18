// +build linux darwin
// +build go1.12

package listener

import (
	"os"
	"path"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/liupangzi/thrift2interface/algorithms"
	"github.com/liupangzi/thrift2interface/parser"
	"github.com/liupangzi/thrift2interface/template"
	"github.com/sirupsen/logrus"
)

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
			l.Logger.Infof("Skip non-golang namespace: `%s`", strings.Join(skippedNamespaces, " "))
		}
	}
}

func (l *LeafThriftListener) EnterInclude(ctx *parser.IncludeContext) {
	leafThrift := algorithms.SimplifyUnixDirectoryPath(strings.Join([]string{
		path.Dir(l.Thrift),
		ctx.LITERAL().GetText()[1 : len(ctx.LITERAL().GetText())-1],
	}, string(os.PathSeparator)))

	if _, ok := l.MI.ThriftFiles[leafThrift]; ok {
		l.Logger.Infof("Leaf thrift file `%s` has already been parsed", leafThrift)
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
