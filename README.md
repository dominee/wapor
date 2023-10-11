# Wapor

Just a simplistic CLI for the [wappalyzergo](https://github.com/projectdiscovery/wappalyzergo) library.

Simply run 

```bash
cat testdata/urls.txt | ./wapor
````

or

```bash
./wapor https://www.citadelo.com https://sme.sk https://github.com
````

and get your semicolon delimited wappalyzer output

```
https://www.citadelo.com : Google Tag Manager; HSTS; Nginx
https://sme.sk : HTTP/3; HSTS; Cloudflare
https://github.com : Amazon S3; Amazon Web Services; GitHub Pages; HSTS
```

## Build

Looks like the [wappalyzergo](https://github.com/projectdiscovery/wappalyzergo) `Weekly fingerprints update [Sun Aug 20 00:19:18 UTC 2023]` are a bit broken and out of sync according to the [releases](https://github.com/projectdiscovery/wappalyzergo/releases), so no need to run scheduled builds for now. I'll set a synced github action for it later.

## Credits

This simple client does in fact almost nothing, and is for my sole purpose of having a portable version.
All the fame and glory goes to to [wappalyzer](https://github.com/wappalyzer/wappalyzer) for the amazing library and [projectdiscovery/wappalyzergo](https://github.com/projectdiscovery/wappalyzergo) for the lovely Go port. ❤️
