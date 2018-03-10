#!/bin/bash

set -e
# Teamcity URL. Make sure you don't have a trailing /
export TCMON_URL="http://teamcity.url:8111"
# Teamcity username
export TCMON_USERNAME="username"
# Teamcity password
export TCMON_PASSWORD="password"
# IDs for teamcity builds to monitor
export TCMON_BUILDS="MyBuildID1,MyBUildID2"
# From header in mail
export TCMON_MAILFROM="tcmon@domain.org"
# Domain to use for emails. This domain is appended to changer's username. So if hmussawir made changes, hmussawir@domain.org will be sent mail to
export TCMON_MAILDOMAIN="domain.org"
# Username to authneticate smtp server
export TCMON_MAILUSERNAME=""
# Password to authenticate smtp server
export TCMON_MAILPASSWORD=""
# Mail server hostname/ip
export TCMON_MAILHOST="192.168.100.24"
# Mail server port
export TCMON_MAILPORT=25
# CC these addresses in each alert
export TCMON_MAILCC="user1@domain.org,user2@domain.org"
# Put all addresses that should receive emails for failures here. Anybody who isn't here doesn't receive alert.
export TCMON_MAILTOWHITELIST="user1@domain.org,user2@domain.org,user3@domain.org"
# Minutes to wait between sending new alert emails.
export TCMON_MAILGAPE=60
# Temporary directory to store last alert timestamps
export TCMON_TEMPDIR="./tmp"
