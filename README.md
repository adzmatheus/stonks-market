## Market from last 1 day


<table>
    <tr>
        <th> Ticker </th>
        <td> <div align="center"> ^BVSP" </div> </td>
    </tr>
    <tr>
        <th> Stonks </th>
        <td> <div align="center"> <img src="https://github.com/adzmatheus/stonks-market/blob/main/assets/expense.svg"/> </div> </td>
    </tr>
    <tr>
        <th>Close yesterday </th>
        <td width="200px"> BRL 141783.36 </td>
    </tr>
    <tr>
        <th>Price today</th>
        <td> BRL 141682.98 </td>
    </tr>
</table>


*Updated at: 2025-10-15T10:25:28Z*

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

- `Stonkses`: An array of daily Stonks. You can view the Stonks struct definition in [model/stonks.go](model/stonks.go).
- `UpdatedAt`: This field contains the timestamp in the format of `time.Date`.

**Step 4**: Register Github Action
- Create a file `.github/workflows/update-stonks.yml` in your repository.
```yml
name: "Cronjob"
on:
schedule:
- cron: '0 10 * * *'

jobs:
    update-stonks:
        permissions: write-all
        runs-on: ubuntu-latest
        steps:
            - uses: actions/checkout@v3
            - name: Generate README
              uses: adzmatheus/stonks-market@v1.0.1
              with:
                ticker: ^BVSP
                stonks-api-key: ${{ secrets.BRAPI_API_KEY }}
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
    - ticker: The ticker that you want show the stonks. Find options on [BRAPI Available](https://brapi.dev/api/available)
    - template-file: Path to the above template file. Eg. `template/README.md.template`
    - out-file: your README.md file name
    - stonks-api-key:
        - Register free API token in [BRAPI Dashboard](https://brapi.dev/dashboard)
        - Setup secrets with name `BRAPI_API_KEY` in `Your repo > settings > Secrets and variables > Actions > New repository secret`

**Step 5**: Commit your change, then Github actions will run as your specificed cron to update Stonks into your README.md file
</details>


## Usage
<details>
<summary>View</summary>

#### Install
```shell
go install https://github.com/adzmatheus/stonks-market
```

#### Run

```shell
Usage:
stonks-market update-stonks [flags]

Flags:
--ticker string                Ticker
-h, --help                     help for update-stonks
-o, --out-file string          Output file path
-f, --template-file string     Readme template file path
-k, --stonks-api-key string    stonksapi.com API key

```

**Sample**
```shell
stonks-market update-stonks \
--ticker=^BVSP \
--stonks-api-key="$STONKS_API_KEY" \
--template-file='template/README.md.template' \
--out-file='README.md'
```

### Docker
```shell
docker build -t stonks-market .
```

```shell
docker run --rm \
-v ./:/app/data \
stonks-market \
--stonks-api-key='XXXX' \
--ticker=^BVSP \
--out-file=data/README.md \
--template-file=data/README.md.template
```

</details>
