# amald

An app that's designed to monitor urls to see if they are under lockdown. It 
can load these urls using a combination of the following:

- text file format (urls separated by newlines)
- `gcloud` CLI

## Requirements

- download the latest stable binary for your OS from 
https://github.com/pemcconnell/appengine-authstatus/releases/ 


## Config

The config file will load by default at `config.yaml`. An example config has 
been provided at `example.config.yaml`. Note that if you wish to enable some of
the features you may need to configure some things first:

- `gcloud` features require the CLI to be installed & already auth'ed. It will
attempt to update `preview` and `app` components as they are required to run
the commands that Amald uses.
- If you wish to enable Mailgun notifications an api key & api url is required.

## Run

- Run the binary (eg. `./amald`)

## Arguments

Type `./amald -h` to list the supported flags

#### Caveats

- I've only built and tested this on linux 64
- This has been built to scrape a CLI and parse the output - far from ideal & 
wide open to breaks as CLI updates wont be reflected in this tool automatically
- The lockdown only checks for 401 or if there are X-Auto-Login headers at the 
moment
- This is super alpha

#### Screenies

You can run this tool however you like. Something like a cronjob would be 
useful. Here is an example of the ASCII output:

![manually ran in terminal](https://sourceforge.net/p/authstatus/screenshot/ascii.png)

You can enable email reports using the mailgun feature, via `config.yaml`:

![email output](https://sourceforge.net/p/authstatus/screenshot/mailgun.png)

PR's welcome :)
