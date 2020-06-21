# Search My Mind

Pencarian personal knowledge base berbasis Github ditulis dengan [Go](http://golang.org/).

## Masalah

Proses pencatatan pengetahuan yang sudah saya lakukan di [notebook.wayanjimmy.xyz](http://notebook.wayanjimmy.xyz/) akan kurang berguna jika tanpa fitur pencarian. Awalnya fitur pencarian sudah menggunakan [algolia](https://www.algolia.com/) namun versi gratisnya sudah tidak muat lagi untuk menyimpan index yang terlalu besar.Untuk itu saya mulai mencari alternatif.

## Solusi

Repo ini adalah salah satu solusi yang saya pikirkan dengan memanfaatkan [Github API](https://developer.github.com/v3/). Selain itu karena knowledge base yang saya miliki ada 2 antara lain.

1. Private notes di sebuah private github repo
1. Public notes di repo [notebook.wayanjimmy.xyz](http://notebook.wayanjimmy.xyz/)

Sebelum menyuguhkan hasil pencarian kepada user saya bisa menggabungkan hasil pencarian dari 2 repo diatas. Disamping itu karena saya menggunakan Alfred di Macos, akan lebih baik fungsi ini disediakan dalam [Alfred Workflow](https://notebook.wayanjimmy.xyz/coding/alfred-workflow).

![Cara kerja Search My Mind](diagram.png)

## Demo

![demo](searchmymind.gif)

## Pengerjaan

- [x] Search private notes
- [x] Search public notes
- [ ] Simple web ui
- [ ] Deploy ke heroku
- [ ] Integrasi dengan GatsbyJS

## Pranala

- [How I Navigate Hundreds of Tabs on Chrome with JXA and Alfred](https://medium.com/@bit2pixel/how-i-navigate-hundreds-of-tabs-on-chrome-with-jxa-and-alfred-9bbf971af02b)
