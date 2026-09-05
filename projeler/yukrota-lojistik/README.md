# Yükrota Lojistik — kurumsal web sitesi

Nakliyat / lojistik firması için statik kurumsal site. **Go** ile yazılmış küçük
bir statik site üreticisi (`main.go`) şablonları işleyip `dist/` altına hazır HTML
üretir. Harici bağımlılık yoktur; sadece Go standart kütüphanesi kullanılır.

## Çalıştırma

```bash
go run .                  # dist/ klasörünü üretir
go run . -serve           # üretir ve http://localhost:8080 üzerinde sunar
go run . -serve -addr :9000
go run . -out public      # farklı çıktı klasörü
go vet ./... && gofmt -l .
```

Yayına alma: `dist/` klasörünü herhangi bir statik host'a (Vercel, Netlify,
Cloudflare Pages, nginx) yükleyin. Sunucu tarafı gereksinimi yoktur.

## Klasörler

| Yol | İçerik |
|---|---|
| `main.go` | Sayfa listesi, şablon derleme, sitemap/robots üretimi, statik kopyalama |
| `internal/content/content.go` | Tüm metin ve veriler (hizmetler, filo, S.S.S., yorumlar, hatlar) |
| `templates/` | `layout.html` + sayfa şablonları, `_` ile başlayanlar parçalardır |
| `static/` | CSS, JS, favicon — olduğu gibi `dist/static/` altına kopyalanır |
| `dist/` | Üretilen çıktı (commit'lidir, doğrudan host edilebilir) |

## İçerik nasıl değiştirilir?

Metinlerin tamamı `internal/content/content.go` içindedir; şablonlara dokunmadan
telefon, adres, hizmet, fiyat, S.S.S. ve sefer hatlarını buradan düzenleyip
`go run .` demeniz yeterlidir. Yeni bir hizmet eklemek için `Services()`
listesine bir kayıt eklemek, detay sayfasının da otomatik üretilmesi için yeterlidir.
