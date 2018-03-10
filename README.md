This is a program to monitor teamcity and generate email alerts to relevant developers.

Flow
--------

- Gets all recent builds of a build configuration
- If last build failed:
	- If we haven't sent email about it in M minutes:
		- If there are code changes:
			- Send email to all users who made changes
			- Cc important people who should be in all emails
		- If there are no code changes:
			- Send email to important people
	- If we have sent email recently:
		- Do nothing
- If last build is successfull
	- Do nothing


Configuration
---------

Configuration can be done via environment variables. See env.sh for details of available configurations.

Build
--------

go build main.go -o tcbuildmonitor

Install
------------

go install github.com/haisum/tcbuildmonitor

Usage
--------

source env.sh
tcbuildmonitor
