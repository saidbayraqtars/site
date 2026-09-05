# Sistem notları — Yükrota Lojistik sitesi

## Ne yapıldı

- **Dil/teknoloji:** Depodaki diğer proje Astro/JS olduğu için bu proje bilinçli
  olarak **Go** ile yazıldı. `main.go` içindeki ~200 satırlık üretici, şablonları
  `html/template` ile derleyip `dist/` altına statik HTML basar. Bağımlılık yok.
- **İçerik/şablon ayrımı:** Tüm metin `internal/content` paketinde tip güvenli
  struct'larda; şablonlar sadece sunum yapar. Hizmet listesine kayıt eklemek
  detay sayfasını da otomatik üretir.
- **Üretilen sayfalar (12):** ana sayfa, hizmetler listesi, 6 hizmet detayı,
  filo, hakkımızda, S.S.S., iletişim. Ayrıca `sitemap.xml` ve `robots.txt`.
- **Tasarım:** Lacivert (#0a1020) + amber (#f5b942) kurumsal palet, Manrope
  tipografi, rota çizgisi motifi. Tek dosya CSS (~430 satır), değişken tabanlı.
  Kart/ızgara sistemi, koyu bölümler, sefer tarifesi panosu, tablo → mobilde
  kart dönüşümü.
- **Etkileşim (`static/js/app.js`, bağımlılıksız ~120 satır):** mobil menü,
  scroll ile beliren bloklar, sayaç animasyonu, akordeon S.S.S., ana sayfadaki
  **hızlı fiyat tahmini** ve iletişim formu doğrulama + `mailto:` taslağı.
  Ana sayfadaki hızlı teklif formu, alanları iletişim formuna query string ile taşır.
- **SEO/erişilebilirlik:** canonical + OG etiketleri, `MovingCompany` JSON-LD,
  semantik başlık hiyerarşisi, skip link, `aria-*` durumları, `prefers-reduced-motion`
  desteği, mobil 390px'te yatay taşma yok (Playwright ile doğrulandı).

## Bilinerek yapılmayanlar / dikkat

- **Form backend'i yok.** Şu an teklif formu kullanıcının e-posta uygulamasında
  taslak açıyor. Gerçek gönderim için bir servis (Formspree / Vercel Function /
  kendi API'niz) bağlanmalı.
- **Firma bilgileri temsilîdir.** Telefon, adres, sayılar, yorumlar ve fiyatlar
  örnek verilerdir; müşteriden gelen gerçek bilgilerle `content.go` güncellenmeli.
- **Görsel yok.** Fotoğraf yerine SVG motifler kullanıldı; gerçek araç/ekip
  fotoğrafları geldiğinde hero ve hizmet kartlarına eklenmeli (`static/img/`).
- Google Fonts harici yüklenir; tamamen çevrimdışı/hızlı istenirse font dosyaları
  self-host edilmeli.

## Sonraki adımlar

1. Müşteriden gerçek bilgileri al (unvan, telefon, adres, yetki belgeleri, gerçek
   referanslar, fiyat aralıkları) ve `internal/content/content.go` dosyasına işle.
2. Form gönderimi için backend bağla; KVKK aydınlatma metni ve çerez bildirimi
   sayfalarını ekle.
3. Fotoğraf çekimi sonrası görsel entegrasyonu + `og:image` üretimi.
4. SEO için il/hizmet bazlı açılış sayfaları (ör. "İstanbul evden eve nakliyat")
   — üretici zaten şablon üzerinden çoğaltmaya uygun.
5. Blog / referans vaka bölümü (Markdown'dan üretim `main.go`'ya eklenebilir).
6. Yayına alma: `dist/` klasörünü Vercel veya Cloudflare Pages'e bağla; alan adı
   ve analytics kurulumunu yap.
7. İsteğe bağlı: müşteri paneli (yük takip) için ayrı bir uygulama; site sadece
   pazarlama katmanı olarak kalır.
