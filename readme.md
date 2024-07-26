## Foundation

App to make it easier to provide donation lists to the CPA.

## Building Prerequisites

Install Go [Go](https://go.dev/). I'm on 1.21 but most versions should be fine.

## Building

Enter the folder. Before building the first time, type

`go mod tidy`

to download all packages.

Then to build the app, type

`go build`

## Running

Enter the folder and type

`go run foundation`

or
`./foundation`

## Using

It's hardcoded to load the donations file `data\Donations.csv` and spit out a `orgs.txt` file based on the highest year in the donations file. Maybe I'll get around to making the input flexible but for now you can just export the Google Sheets donation file and rename it to `Donations.csv`.
