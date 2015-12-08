# GO-SendMail

[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

This is a simple command line application for sending email messages written
in Golang. Its aim is to provide a cross platform application for sending
email messages from the command line, and more specifically to allow an easy
way to send email messages from the Windows command line and batch scripts.

The command line options resemble those of the Linux mail command.

The config.json file is used to provide system wide settings for configuring
the email server and sender's address. Both configuration options can be set
on the command line as options.

This is not an attempt to rewrite or replace the Linux mail command, although
it is trying to be as close as possible by using the same flags for the
functionality it is implementing.
