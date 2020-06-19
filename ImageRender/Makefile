.PHONY: deps clean build test

PACKAGE=img-render
EXECFILE=imgRender
FUNC=ImageRenderFunction

deps:
	go get -u ./${PACKAGE}/...

clean: 
	rm -rf ./${PACKAGE}/${EXECFILE}
	
build:
	cd ${PACKAGE} && \
	GOOS=linux GOARCH=amd64 go build -o ./${EXECFILE}

test:
	make build
	sam local invoke "${FUNC}" -e event.json

debug:
	sam local invoke --debug "${FUNC}" -e event.json