## Stonks For Next 1 days


<table>
    <tr>
        <th>Stonks</th>
        <td><img src="https://github.com/adzmatheus/stonks-market/blob/main/assets/income.svg"/></td>
    </tr>
    <tr>
        <th>Close yesterday: R$</th>
        <td width="200px">131672</td>
    </tr>
    <tr>
        <th>Price today</th>
        <td>131791.55 kph</td>
    </tr>
</table>


*Updated at: 2024-10-07T10:22:15Z*

## GitHub Actions: Embed up-to-date Stonks in your README
<details>
<summary>
    View
</summary>

You can easily embed tables in your README.md using GitHub Actions by following these simple steps:

**Step 1:** In your repository, create a file named `README.md.template`.

**Step 2:** Write anything you want within the `README.md.template` file.

**Step 3:** Embed one of the following entities within your `README.md.template`:

- **Daily Stonks Table:**
```shell
{{ template "daily-table" .Stonkses }}
```

- **Updated at:**
```shell
{{ formatTime .UpdatedAt }}
```

If you are familiar with Go templates, you have access to the `root` variable, which includes the following fields:

- `Stonkses`: An array of daily Stonks. You can view the Stonks struct definition in [model/weather.go](model/weather.go).
- `UpdatedAt`: This field contains the timestamp in the format of `time.Date`.

**Step 4**: Register Github Action
- Create a file `.github/workflows/update-weather.yml` in your repository.
```yml
name: "Cronjob"
on:
schedule:
- cron: '15 * * * *'

jobs:
    update-weather:
        permissions: write-all
        runs-on: ubuntu-latest
        steps:
            - uses: actions/checkout@v3
            - name: Generate README
              uses: huantt/weather-forecast@v1.0.5
              with:
                city: HaNoi
                days: 7
                weather-api-key: ${{ secrets.WEATHER_API_KEY }}
                template-file: 'README.md.template'
                out-file: 'README.md'
            - name: Commit
              run: |
                if git diff --exit-code; then
                  echo "No changes to commit."
                  exit 0
                else
                  git config user.name github-actions
                  git config user.email github-actions@github.com
                  git add .
                  git commit -m "update"
                  git push origin main
                fi
```
- Update some variable in this file:
    - city: The city that you want to forecast weather
    - days: number of forecast days
    - template-file: Path to the above template file. Eg. `template/README.md.template`
    - out-file: your README.md file name
    - weather-api-key:
        - Register free API key in [https://weatherapi.com](https://weatherapi.com)
        - Setup secrets with name `WEATHER_API_KEY` in `Your repo > settings > Secrets and variables > Actions > New repository secret`

**Step 5**: Commit your change, then Github actions will run as your specificed cron to update Stonks into your README.md file
</details>


## Usage
<details>
<summary>View</summary>

#### Install
```shell
go install https://github.com/huantt/weather-forecast
```

#### Run

```shell
Usage:
weather-forecast update-weather [flags]

Flags:
--city string              City
--days int                 Days of forecast (default 7)
-h, --help                     help for update-weather
-o, --out-file string          Output file path
-f, --template-file string     Readme template file path
-k, --weather-api-key string   weatherapi.com API key

```

**Sample**
```shell
weather-forecast update-weather \
--days=7 \
--weather-api-key="$WEATHER_API_KEY" \
--template-file='template/README.md.template' \
--city=HaNoi \
--out-file='README.md'
```

### Docker
```shell
docker build -t weather-forecast .
```

```shell
docker run --rm \
-v ./:/app/data \
weather-forecast \
--weather-api-key='XXXX' \
--city=HaNoi \
--out-file=data/README.md \
--template-file=data/README.md.template
```

</details>
