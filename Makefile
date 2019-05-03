.PHONY: parser

parser:
	@# alias antlr='java -jar /usr/local/lib/antlr-4.7.2-complete.jar'
	antlr -visitor -Dlanguage=Go -o parser -package parser Thrift.g4
