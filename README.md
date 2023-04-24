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

## Credits

This simple client does in fact almost nothing, and is for my sole purpose of having a portable version.
All the fame and glory goes to to [wappalyzer](https://github.com/wappalyzer/wappalyzer) for the amazing library and [projectdiscovery/wappalyzergo](https://github.com/projectdiscovery/wappalyzergo) for the lovely Go port. ❤️
