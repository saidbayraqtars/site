// yukrota, Yükrota Lojistik kurumsal sitesini üreten küçük bir statik site
// üreticisidir. Harici bağımlılık yoktur; yalnızca Go standart kütüphanesi
// kullanılır (html/template, io/fs, os).
//
// Kullanım:
//
//	go run .              -> dist/ klasörünü üretir
//	go run . -serve       -> üretir ve :8080 üzerinde önizleme sunar
//	go run . -out public  -> çıktı klasörünü değiştirir
package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"yukrota/internal/content"
)

// Ctx, her şablona geçirilen render bağlamı.
type Ctx struct {
	Site    content.Site
	Nav     []content.Nav
	Title   string // <title> içeriği (site adı otomatik eklenir)
	Desc    string
	Path    string // "/hizmetler/" gibi, aktif menü işaretlemesi için
	Heading string // sayfa başlığındaki büyük başlık
	Kicker  string
	Sub     string
	Data    any // sayfaya özel veri
}

// page, üretilecek tek bir HTML çıktısını tanımlar.
type page struct {
	path string // "/hizmetler/" -> dist/hizmetler/index.html
	tpl  string // templates/ altındaki dosya adı
	ctx  Ctx
}

const (
	tplDir    = "templates"
	staticDir = "static"
)

func main() {
	out := flag.String("out", "dist", "çıktı klasörü")
	serve := flag.Bool("serve", false, "üretimden sonra önizleme sunucusu başlat")
	addr := flag.String("addr", ":8080", "önizleme sunucusu adresi")
	flag.Parse()

	start := time.Now()
	n, err := build(*out)
	if err != nil {
		log.Fatalf("üretim hatası: %v", err)
	}
	fmt.Printf("✓ %d sayfa üretildi → %s/  (%s)\n", n, *out, time.Since(start).Round(time.Millisecond))

	if *serve {
		fmt.Printf("→ önizleme: http://localhost%s\n", *addr)
		log.Fatal(http.ListenAndServe(*addr, http.FileServer(http.Dir(*out))))
	}
}

func build(out string) (int, error) {
	site := content.Company()
	nav := content.MainNav()
	services := content.Services()

	if err := os.RemoveAll(out); err != nil {
		return 0, err
	}
	if err := copyTree(staticDir, filepath.Join(out, "static")); err != nil {
		return 0, fmt.Errorf("statik dosyalar: %w", err)
	}

	base := func(title, desc, path string) Ctx {
		return Ctx{Site: site, Nav: nav, Title: title, Desc: desc, Path: path}
	}

	// Ana sayfa verisi.
	type homeData struct {
		Services     []content.Service
		Stats        []content.Stat
		Steps        []content.Step
		Values       []content.Value
		Routes       []content.Route
		Testimonials []content.Testimonial
		Cities       []string
		Faqs         []content.Faq
	}

	pages := []page{
		{
			path: "/",
			tpl:  "home.html",
			ctx: func() Ctx {
				c := base("Nakliyat ve Lojistik", site.Desc, "/")
				c.Data = homeData{
					Services:     services,
					Stats:        content.Stats(),
					Steps:        content.Steps(),
					Values:       content.Values(),
					Routes:       content.Routes(),
					Testimonials: content.Testimonials(),
					Cities:       content.Cities(),
					Faqs:         content.Faqs()[:4],
				}
				return c
			}(),
		},
		{
			path: "/hizmetler/",
			tpl:  "services.html",
			ctx: func() Ctx {
				c := base("Hizmetlerimiz", "Evden eve nakliyat, parsiyel ve komple yük, depolama, uluslararası taşımacılık ve kurumsal taşıma hizmetleri.", "/hizmetler/")
				c.Kicker = "Hizmetler"
				c.Heading = "Tek yükten fabrika taşımaya"
				c.Sub = "Altı ana hizmet başlığında, aynı operasyon ekibi ve aynı takip sistemi ile çalışıyoruz."
				c.Data = struct {
					Services []content.Service
					Steps    []content.Step
				}{services, content.Steps()}
				return c
			}(),
		},
		{
			path: "/filo/",
			tpl:  "fleet.html",
			ctx: func() Ctx {
				c := base("Filomuz", "124 araçlık öz filo: panelvandan asansörlü kamyona, frigorifikten lowbed'e kadar her yük tipi için araç.", "/filo/")
				c.Kicker = "Filo ve ağ"
				c.Heading = "124 araç, 81 il, tek operasyon merkezi"
				c.Sub = "Araçların tamamı öz mülkiyetimizde. Bakım, takip ve şoför kadrosu içeride yönetilir."
				c.Data = struct {
					Fleet  []content.Vehicle
					Routes []content.Route
					Cities []string
				}{content.Fleet(), content.Routes(), content.Cities()}
				return c
			}(),
		},
		{
			path: "/hakkimizda/",
			tpl:  "about.html",
			ctx: func() Ctx {
				c := base("Hakkımızda", "2009'dan bu yana nakliyat ve lojistik. Kendi kadrosu, kendi filosu ve yazılı fiyat sözü olan bir taşıma şirketi.", "/hakkimizda/")
				c.Kicker = "Hakkımızda"
				c.Heading = "2009'dan beri aynı işi yapıyoruz"
				c.Sub = "Bir kamyonet ve iki kişiyle başladık. Bugün 124 araç ve 310 kişiyiz; iş yapma biçimimiz değişmedi."
				c.Data = struct {
					Stats  []content.Stat
					Values []content.Value
					Steps  []content.Step
				}{content.Stats(), content.Values(), content.Steps()}
				return c
			}(),
		},
		{
			path: "/sss/",
			tpl:  "faq.html",
			ctx: func() Ctx {
				c := base("Sık Sorulan Sorular", "Fiyat, sigorta, depolama, ödeme ve teslim süreleri hakkında en çok sorulan sorular.", "/sss/")
				c.Kicker = "S.S.S."
				c.Heading = "Aklınıza takılanlar"
				c.Sub = "Cevabını bulamadığınız bir konu varsa doğrudan arayın; operasyon ekibimiz açıklasın."
				c.Data = struct{ Faqs []content.Faq }{content.Faqs()}
				return c
			}(),
		},
		{
			path: "/iletisim/",
			tpl:  "contact.html",
			ctx: func() Ctx {
				c := base("İletişim ve Teklif", "Ücretsiz ekspertiz ve yazılı fiyat teklifi için formu doldurun ya da bizi arayın.", "/iletisim/")
				c.Kicker = "İletişim"
				c.Heading = "Teklif alın, 2 saat içinde dönelim"
				c.Sub = "Formu doldurmanız yeterli. İsterseniz doğrudan telefonla da ulaşabilirsiniz."
				c.Data = struct{ Services []content.Service }{services}
				return c
			}(),
		},
	}

	// Her hizmet için ayrı detay sayfası.
	for i, s := range services {
		c := base(s.Title, s.Short, "/hizmetler/")
		c.Kicker = "Hizmet"
		c.Heading = s.Title
		c.Sub = s.Lead
		c.Data = struct {
			Service content.Service
			Others  []content.Service
		}{s, others(services, i)}
		pages = append(pages, page{path: "/hizmetler/" + s.Slug + "/", tpl: "service.html", ctx: c})
	}

	funcs := template.FuncMap{
		"join": strings.Join,
		"inc":  func(i int) int { return i + 1 },
	}

	partials, err := filepath.Glob(filepath.Join(tplDir, "_*.html"))
	if err != nil {
		return 0, err
	}

	for _, p := range pages {
		files := append([]string{
			filepath.Join(tplDir, "layout.html"),
			filepath.Join(tplDir, p.tpl),
		}, partials...)

		t, err := template.New("layout.html").Funcs(funcs).ParseFiles(files...)
		if err != nil {
			return 0, fmt.Errorf("%s: %w", p.tpl, err)
		}

		dst := filepath.Join(out, filepath.FromSlash(strings.Trim(p.path, "/")), "index.html")
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return 0, err
		}
		f, err := os.Create(dst)
		if err != nil {
			return 0, err
		}
		if err := t.Execute(f, p.ctx); err != nil {
			f.Close()
			return 0, fmt.Errorf("%s: %w", p.path, err)
		}
		f.Close()
	}

	if err := writeSitemap(out, site, pages); err != nil {
		return 0, err
	}
	if err := os.WriteFile(filepath.Join(out, "robots.txt"),
		[]byte("User-agent: *\nAllow: /\nSitemap: "+site.BaseURL+"/sitemap.xml\n"), 0o644); err != nil {
		return 0, err
	}
	return len(pages), nil
}

// others, verilen indeks dışındaki en fazla üç hizmeti döndürür.
func others(all []content.Service, skip int) []content.Service {
	var out []content.Service
	for i := range all {
		if i == skip {
			continue
		}
		out = append(out, all[i])
		if len(out) == 3 {
			break
		}
	}
	return out
}

func writeSitemap(out string, site content.Site, pages []page) error {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")
	for _, p := range pages {
		fmt.Fprintf(&b, "  <url><loc>%s%s</loc></url>\n", site.BaseURL, p.path)
	}
	b.WriteString("</urlset>\n")
	return os.WriteFile(filepath.Join(out, "sitemap.xml"), []byte(b.String()), 0o644)
}

// copyTree, src altındaki tüm dosyaları dst altına birebir kopyalar.
func copyTree(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		outF, err := os.Create(target)
		if err != nil {
			return err
		}
		defer outF.Close()
		_, err = io.Copy(outF, in)
		return err
	})
}
