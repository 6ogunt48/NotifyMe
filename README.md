# CHECKEMAILBOT

> This project is an experimental project I am using to learn how to build telegram bots. It's safe to use for personal
> emails that supports IMAP version 4
> if it is deployed in a secure environment. It doesn't support Oauth2 yet and cant work with Office 365 emails or other
> email providers that won't allow IMAP Access.

Checkemailbot is a telegram bot I use for monitoring my email accounts.it checks all email accounts specified in
config.toml at the same time and returns the number of unread emails found in those accounts to telegram. You can
either check
the email using the bot commands on the telegram app or wait for the bot to check and send results notification at the
specified interval
time in config.


# Installation
Before installation. You need to create a config.toml file required by the program. This toml file has to be the same with the example shown below

![TOML EXAMPLE](./assets/toml_example.png)

The Telegram bot token and chat ID should be created on the telegram using bot father. Please follow this [link](https://www.youtube.com/watch?v=EOke01hZgZ0) for a detailed guide on how to do that


![Continuous Integration and Delivery](https://github.com/6ogunt48/checkemailbot/actions/workflows/main.yaml/badge.svg?branch=main)