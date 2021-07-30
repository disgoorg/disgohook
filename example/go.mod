module github.com/DisgoOrg/disgohook/example

go 1.16

replace (
	github.com/DisgoOrg/disgohook => ../
	github.com/DisgoOrg/log => ../../log
)
require (
	github.com/DisgoOrg/disgohook v1.4.0
	github.com/DisgoOrg/log v1.2.3
	github.com/sirupsen/logrus v1.8.1
)
