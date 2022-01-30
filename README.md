# Apple Release Notifier

[![Go Report Card](https://goreportcard.com/badge/github.com/baturtulek/apple-release-notifier)](https://goreportcard.com/report/github.com/baturtulek/apple-release-notifier)

Simple Apple Software Release Notifier, sends an email containing release information.

## Usage

-   Fill in the `contacts` file with the email addresses you want to send an email of the new releases. (There should be only one email address per line.)

-   Create .env file on the root folder which contains following variables

    ```
    SENDER_MAIL_SMTP_HOST=<SMTP_HOST>
    SENDER_MAIL_ADDRESS=<MAIL_ADDRESS>
    SENDER_MAIL_NAME=<MAIL_ADDRESS_DISPLAY_NAME>
    SENDER_MAIL_PASSWORD=<MAIL_ADDRESS_PASSWORD>
    SENDER_MAIL_SMTP_PORT=<SMTP_PORT>
    ```

-   Run main file
    ```
    go run main.go
    ```
-   This app is based on web scrapping, so you may want to use with [CRON](https://en.wikipedia.org/wiki/Cron).
    -   Build the app
        ```
        go build
        ```
    -   Add the following lines to your crontab file. This is just an example. You can change the frequency as you want.
        ```
        # Apple generally releases software from 20:00 to 23:00 on weekdays. (In Istanbul Time - UTC +3)
        # This will run the app every 15 minutes between 20:00 to 23:00 on weekdays.
        */15 20-23 * * 1-5 <./PATH_TO_APP/apple-release-notifier>
        ```

## Author

-   Ahmet Batur Tülek - [GitHub](https://github.com/baturtulek)

## License

This project is released under the MIT License. See the `LICENSE` file for details.
