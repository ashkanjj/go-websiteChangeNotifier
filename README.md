# WebsiteChangeNotifier

A small tool written in Go that will take a web address as an input and informs you when something has changed (partly an experiment). it will store snapshots of the webpages (only the body) (but it will need to store them in a format that doesn't take too much space but it's still diffable and return what the difference is)

it will allow you to specify where and how it should store the result or even email it to you

A nifty tool written in Go that keeps tabs on changes in your favorite web pages. It takes a web address, tracks any tweaks (mostly for fun), and stores snapshots of the webpages - just the <body>. It cleverly manages storage space while still making it easy to spot the differences. Plus, you're the boss: choose where and how to stash the results or get them sent straight to your inbox!

## Installation

```bash
go get -u github.com/ashkanjj/go-websiteChangeNotifier
```

## Usage

Different systems have different ways to run processes in the background and the solution for any OS is typically only a google search away.

I use Ubuntu on Windows (WSL) and I do the following:

```sh
go build
nohup ./go-websiteChangeNotifier &
(copy the process ID that is's outputted)
tail -f nohup.out (watch the logs)
```

OR

```sh
go run main.go -url=http://localhost:8090 -config=../../email_config.json -bolt-path=./my.db
```

# Motivation

This was a hubby project of mine aiming to get my hands dirty with Golang. While I initially planned to expand its features, I've had to set it aside for now due to pressing personal commitments.

The inspiration for this project struck when I needed to secure an appointment from an embassy's website. They released a limited number of slots daily and I wanted to be one of the first people to be notified. This sort of websites usually don't require authentication and therefore it's easy to set up a crawler. However, I didn't want to be constantly checking the website and I wanted to be notified when the website changed. I also wanted to be able to exclude certain parts of the website from the change detection as they were irrelevant to me.

# Longer Description

# TODO for future

- Parse the cli for:
  - |~~Web address to listen for~~
  - ~~An email address or a google cloud/AWS email service in which case a config file~~
- ```Take a snapshot of the current state of web document's body and store it~~

  ```

- ~~start the cron job (or we can just call it listener) and have it running forever~~
  - at every specified tick, compare the registered document's body with what's stored on file
  - if a change is detected, compose an email and send it off to the registered email address or pass it to the configured service
- accept a signal to end the active listener
- add two extra flags to run the diffing program THIS NOW
- Add a config file with files to exclude lines for
- Add a way to see more difs

# Things to find out

1. Find a way to use a two-way hash function so we can store the body but also recover it for diffing later, very similar to how git stores the whole snapshot of changes
2. Find out how to send an email
3. How to scale it to millions of registered websites and listeners

# Database

bbolt (a key value store db written in go) is used to store the snapshots and other metadata.

Normalised model

```
Website {
  name: string;
  id: string;
}

Snapshot = {
  name: string;
  websiteId: string;
}

Exclusion = {
  created_date: Date;
  content: string;
  websiteId: string;
}

Denormalised model

Website {
  name: string;
  id: string;
  snapshots: Snapshot[]
}

Snapshot {
  website_id: string;
  date: Date;
  content: string;
  exclusions: Exclusion[]
}

Exclusion {
  websiteId: string;
  snapshotId: string;
  created_date: Date;
  content: string;
}
```

## Schema

| Bucket    | Schema                                             |
| --------- | -------------------------------------------------- |
| Website   | $id: = $name                                       |
| Snapshot  | $website_id:$date = \$content                      |
| Exclusion | $website_id:$snapshot_id:$created_date = \$content |
