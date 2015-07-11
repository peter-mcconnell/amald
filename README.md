# amald
[![Coverage Status](https://coveralls.io/repos/pemcconnell/amald/badge.svg?branch=master&service=github)](https://coveralls.io/github/pemcconnell/amald?branch=master)
[![Build Status](https://travis-ci.org/pemcconnell/amald.svg?branch=master)](https://travis-ci.org/pemcconnell/amald)

An app that's designed to monitor urls to see if they are under lockdown. It 
can load these urls using a combination of the following:

- text file format (urls separated by newlines)
- `gcloud` CLI

## Requirements

- download the latest stable binary for your OS from 
https://github.com/pemcconnell/amald/releases/ 
- The binary is all that's required, but if you wish to avail of the mailgun (email) notifications you'll need to include `reports/tmpl/` so that amald can find the html templates & you'll also need to ensure the `config.yaml` is set up and correctly configured.
- If you wish to utilise storage (for report summaries) you'll need to ensure that a folder `./tmp` exists and is writable (amald will create `data.json` file)

## Config

The config file will load by default at `./config.yaml`. An example config has 
been provided at `./config/example.config.yaml`. Note that if you wish to enable some of
the features you may need to configure some things first:

- `gcloud` features require the CLI to be installed & already auth'ed. It will
attempt to update `alpha` and `app` components as they are required to run
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
- This is beta 

#### Screenies

You can run this tool however you like. Something like a cronjob would be 
useful. Here is an example of the ASCII output:

![manually ran in terminal](https://cloud.githubusercontent.com/assets/641429/6988095/be0a5bda-da45-11e4-99ff-dd90b5459d44.png)

You can enable email reports using the mailgun feature, via `config.yaml`:

![email output](https://cloud.githubusercontent.com/assets/641429/6988094/be04cb84-da45-11e4-86c8-582440802a02.png)

PR's welcome :)


~ amald == are my apps locked down?
