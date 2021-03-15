# ghost-seo-check

Check your Ghost posts for SEO problems

```shell
GHOST_URL=https://etf.capital ADMIN_API_TOKEN=abcdef:ghijkl go run cmd/ghost-seo-check/main.go

[...]
https://etf.capital/verkaufsgrunde-bei-etfs/ -> "Verkaufsgründe bei ETFs"
        meta description is empty
        text is too short (345 < 1000)
https://etf.capital/etfs-wer-erzeugt-sie/ -> "ETFs: Wer erzeugt sie?"
        meta description is empty
        text is too short (325 < 1000)
https://etf.capital/buchempfehlungen-etf-sparer/ -> "Buchempfehlungen für ETF-Sparer"
        meta description is empty
        text is too short (973 < 1000)
https://etf.capital/der-grosse-etf-depot-vergleich/ -> "Der große Depot-Vergleich für ETF-Sparer"
        meta description is empty
INFO[0004] 90 of 693 pages have SEO errors 

```