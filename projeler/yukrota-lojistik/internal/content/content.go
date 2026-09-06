// Package content, sitenin tüm metin ve veri içeriğini tutar.
// İçerik burada tip güvenli biçimde durur; şablonlar sadece sunumla ilgilenir.
package content

// Site, her sayfada ortak olan kurumsal bilgiler.
type Site struct {
	Name      string
	Legal     string
	Tagline   string
	Desc      string
	BaseURL   string
	Phone     string
	PhoneHref string
	Whatsapp  string
	Email     string
	Address   string
	Hours     string
	Year      int
	Founded   int
}

// Nav, üst menüdeki bir bağlantı.
type Nav struct {
	Label string
	Href  string
}

// Service, sunulan bir taşımacılık hizmeti.
type Service struct {
	Slug     string
	Icon     string // static/img altındaki sembol adı (inline SVG id)
	Title    string
	Short    string
	Lead     string
	Body     []string
	Features []string
	Includes []string
	Price    string
}

// Stat, ana sayfadaki sayaç kutusu.
type Stat struct {
	Value  string
	Suffix string
	Label  string
}

// Step, çalışma akışındaki bir adım.
type Step struct {
	No    string
	Title string
	Text  string
}

// Vehicle, filodaki bir araç tipi.
type Vehicle struct {
	Name     string
	Capacity string
	Volume   string
	Use      string
	Count    string
}

// Faq, sık sorulan soru.
type Faq struct {
	Q string
	A string
}

// Testimonial, müşteri yorumu.
type Testimonial struct {
	Text    string
	Author  string
	Role    string
	Initial string
}

// Route, düzenli sefer yapılan hat.
type Route struct {
	From string
	To   string
	Days string
}

// Value, kurumsal değer / fark yaratan madde.
type Value struct {
	Icon  string
	Title string
	Text  string
}

// Site, tekil kurumsal kayıt.
func Company() Site {
	return Site{
		Name:      "Yükrota Lojistik",
		Legal:     "Yükrota Nakliyat ve Lojistik A.Ş.",
		Tagline:   "Yükünüz yola çıktığı andan itibaren bizim işimiz.",
		Desc:      "Evden eve nakliyat, parsiyel ve komple yük taşımacılığı, depolama ve uluslararası lojistik. 81 ile sigortalı, takip edilebilir taşıma.",
		BaseURL:   "https://yukrota.com.tr",
		Phone:     "0850 000 45 45",
		PhoneHref: "+908500004545",
		Whatsapp:  "905000004545",
		Email:     "merkez@yukrota.com.tr",
		Address:   "Ambarlı Mah. Lojistik Cad. No: 42, Avcılar / İstanbul",
		Hours:     "Hafta içi 08:00 – 20:00 · Cumartesi 09:00 – 17:00",
		Year:      2026,
		Founded:   2009,
	}
}

func MainNav() []Nav {
	return []Nav{
		{"Ana Sayfa", "/"},
		{"Hizmetler", "/hizmetler/"},
		{"Filo", "/filo/"},
		{"Hakkımızda", "/hakkimizda/"},
		{"S.S.S.", "/sss/"},
		{"İletişim", "/iletisim/"},
	}
}

func Services() []Service {
	return []Service{
		{
			Slug:  "evden-eve-nakliyat",
			Icon:  "home",
			Title: "Evden Eve Nakliyat",
			Short: "Sökümden montaja, paketlemeden yerleşime kadar anahtar teslim ev taşıma.",
			Lead:  "Eşyalarınızı biz paketler, biz taşır, yeni evinizde biz kurarız. Siz sadece kapıyı açın.",
			Body: []string{
				"Taşıma öncesi ekip liderimiz evinize gelir, eşya envanterini çıkarır ve size sabit fiyatlı bir teklif sunar. Sürpriz kalem çıkmaz.",
				"Beyaz eşya, mobilya ve elektronik ürünler ayrı ayrı ambalajlanır; cam ve tablo gibi kırılabilir parçalar için çift katmanlı sandık kullanılır.",
				"Taşıma günü mobilya söküm-montaj ustası ekiple birlikte gelir. Yeni adreste eşyalar sizin gösterdiğiniz yerleşim planına göre yerleştirilir.",
			},
			Features: []string{"Ücretsiz ekspertiz", "Profesyonel ambalaj", "Söküm & montaj ustası", "Asansörlü taşıma", "Sigortalı taşıma"},
			Includes: []string{"Baloncuklu naylon, streç film ve karton koli", "Mobilya söküm ve yeniden montajı", "Beyaz eşya sabitleme ve nakliyesi", "Yeni adreste kaba yerleşim", "Ambalaj atıklarının toplanması"},
			Price:    "1+1 daireler 8.500 ₺'den başlayan fiyatlarla",
		},
		{
			Slug:  "parsiyel-tasimacilik",
			Icon:  "boxes",
			Title: "Parsiyel Taşımacılık",
			Short: "Bir araç dolduracak kadar yükünüz yok mu? Sadece kapladığı yer kadar ödeyin.",
			Lead:  "Palet, koli ya da tek parça yüklerinizi düzenli seferlerimize ekliyor, maliyeti yolcularla paylaştırıyoruz.",
			Body: []string{
				"İstanbul, Ankara, İzmir, Bursa, Antalya ve Gaziantep arasında haftanın altı günü düzenli parsiyel seferimiz var. Yükünüz en geç 48 saat içinde yola çıkar.",
				"Her parti kendi barkodu ile sisteme girilir; ara aktarma merkezlerinde okutulur ve konumu panelinizden anlık görünür.",
				"Desi ve palet bazlı şeffaf fiyatlandırma yaparız. Minimum yük şartı aramıyoruz.",
			},
			Features: []string{"Haftada 6 sefer", "Barkodlu takip", "Palet / desi bazlı fiyat", "Aktarma merkezi ağı"},
			Includes: []string{"Yükleme rampasında elleçleme", "Palet streçleme ve etiketleme", "Ara depolama (3 güne kadar ücretsiz)", "Teslimat sonrası e-imzalı irsaliye"},
			Price:    "Palet başına 1.250 ₺'den başlayan fiyatlarla",
		},
		{
			Slug:  "komple-yuk-tasimaciligi",
			Icon:  "truck",
			Title: "Komple Yük Taşımacılığı",
			Short: "Tek müşteriye tahsis edilmiş araçla, aktarmasız, doğrudan varış noktasına.",
			Lead:  "Yükünüz araca yüklendiği noktadan indirildiği noktaya kadar hiç el değiştirmez.",
			Body: []string{
				"Tenteli, frigorifik, açık kasa ve konteyner taşıyıcı seçenekleriyle 24 tona kadar komple yük taşıyoruz.",
				"Kritik teslimat saatleri için çift şoförlü araç tahsis ediyor, yolda kesintisiz ilerliyoruz.",
				"Yükleme öncesi araç uygunluk raporu ve fotoğraflı yük tespiti tarafınıza iletilir.",
			},
			Features: []string{"Aktarmasız teslimat", "24 tona kadar kapasite", "Çift şoför opsiyonu", "Frigorifik araç"},
			Includes: []string{"Araç tahsis belgesi", "Fotoğraflı yükleme tespiti", "Canlı GPS takip bağlantısı", "Varışta 2 saat ücretsiz bekleme"},
			Price:    "Hat ve araç tipine göre teklifli",
		},
		{
			Slug:  "depolama-ve-ellecleme",
			Icon:  "warehouse",
			Title: "Depolama ve Elleçleme",
			Short: "Kapalı, kameralı ve sigortalı depolarda kısa ya da uzun süreli saklama.",
			Lead:  "Taşınma tarihiniz ile teslim tarihiniz uyuşmadığında eşyanız güvende bekler.",
			Body: []string{
				"Avcılar ve Kemalpaşa'daki tesislerimizde toplam 14.000 m² kapalı alanımız var. Alanlar bölmeli, nem kontrollü ve 7/24 kameralıdır.",
				"Eşya bazlı envanter kaydı tutar, depoya giren her parçayı fotoğraflarız. Çıkış talebinizi 24 saat önceden bildirmeniz yeterli.",
				"E-ticaret müşterilerimiz için mal kabul, raflama, sipariş toplama ve kargo teslim süreçlerini de yönetiyoruz.",
			},
			Features: []string{"14.000 m² kapalı alan", "7/24 kamera ve alarm", "Nem kontrollü bölmeler", "E-ticaret fulfillment"},
			Includes: []string{"Fotoğraflı envanter kaydı", "Aylık stok raporu", "Palet bazlı raflama", "Yangın algılama sistemi"},
			Price:    "Palet başına aylık 480 ₺'den başlayan fiyatlarla",
		},
		{
			Slug:  "uluslararasi-tasimacilik",
			Icon:  "globe",
			Title: "Uluslararası Taşımacılık",
			Short: "Avrupa ve Orta Doğu hatlarında karayolu, gümrük ve evrak süreçleri dahil.",
			Lead:  "Almanya, Hollanda, Romanya, Bulgaristan ve Irak hatlarında düzenli çıkışlarımız var.",
			Body: []string{
				"CMR sigortası, T1/T2 transit beyannamesi ve gümrük müşavirliği hizmetlerini tek elden yürütüyoruz.",
				"Sınır kapısı bekleme sürelerini günlük takip eder, alternatif güzergâh önerisi sunarız.",
				"Yurt dışı evden eve taşımalarda menşe ülkedeki ekip ile varış ülkesindeki ekip aynı envanter listesi üzerinden çalışır.",
			},
			Features: []string{"CMR sigortası", "Gümrük müşavirliği", "Avrupa & Orta Doğu hatları", "Kapıdan kapıya"},
			Includes: []string{"İhracat evrak hazırlığı", "T1 / T2 transit işlemleri", "Sınır kapısı takibi", "Varış ülkesinde dağıtım"},
			Price:    "Hat bazlı teklifli",
		},
		{
			Slug:  "ofis-ve-fabrika-tasima",
			Icon:  "building",
			Title: "Ofis ve Fabrika Taşıma",
			Short: "İş sürekliliğini bozmadan, hafta sonuna sığdırılmış planlı kurumsal taşıma.",
			Lead:  "Cuma akşamı sökeriz, pazartesi sabahı çalışır halde teslim ederiz.",
			Body: []string{
				"Taşıma öncesi bir proje sorumlusu atanır; kat planı, etiketleme şeması ve saat bazlı iş programı sizinle paylaşılır.",
				"Sunucu odası, arşiv ve laboratuvar ekipmanları için özel ambalaj ve titreşim yalıtımlı taşıma uygularız.",
				"Makine taşımalarında vinç, forklift ve kızak operasyonlarını kendi ekipmanımızla yaparız.",
			},
			Features: []string{"Proje sorumlusu", "Etiketli kat planı", "Sunucu odası taşıma", "Vinç & forklift"},
			Includes: []string{"Taşıma projesi ve zaman planı", "Arşiv kolileme ve numaralandırma", "Demontaj / montaj ekibi", "Hafta sonu ve gece çalışma"},
			Price:    "Metrekare ve ekipmana göre teklifli",
		},
	}
}

func Stats() []Stat {
	return []Stat{
		{"17", "", "yıllık saha tecrübesi"},
		{"42", "K+", "tamamlanan taşıma"},
		{"81", "", "ilde teslimat ağı"},
		{"98", "%", "zamanında teslim oranı"},
	}
}

func Steps() []Step {
	return []Step{
		{"01", "Talep ve ekspertiz", "Formu doldurun ya da arayın; aynı gün içinde ücretsiz yerinde ya da görüntülü ekspertiz planlayalım."},
		{"02", "Sabit fiyat teklifi", "Envanter üzerinden hesaplanan, kalem kalem yazılı teklifi 2 saat içinde iletiriz. Sonradan fark çıkmaz."},
		{"03", "Planlama ve ambalaj", "Tarih onaylanır, ekip ve araç tahsis edilir. Taşıma günü ambalajlama sizin evinizde başlar."},
		{"04", "Taşıma ve takip", "Araç yola çıktığı anda size takip bağlantısı gönderilir. Operasyon sorumlunuz tek numaradan ulaşılabilirdir."},
		{"05", "Teslim ve yerleşim", "Eşya yeni adreste kurulur, kontrol listesi birlikte imzalanır, ambalaj atıkları toplanır."},
	}
}

func Fleet() []Vehicle {
	return []Vehicle{
		{"Panelvan", "1,5 ton", "12 m³", "Şehir içi küçük taşıma, parça yük", "24 araç"},
		{"Kamyonet (NPR)", "3,5 ton", "22 m³", "1+1 / 2+1 ev taşıma, parsiyel dağıtım", "38 araç"},
		{"Asansörlü kamyon", "7 ton", "42 m³", "Yüksek katlı bina taşımaları", "16 araç"},
		{"Tenteli tır", "24 ton", "92 m³", "Şehirlerarası komple yük", "31 araç"},
		{"Frigorifik araç", "18 ton", "68 m³", "Soğuk zincir gerektiren yükler", "9 araç"},
		{"Açık kasa / lowbed", "40 ton", "—", "Makine ve gabari dışı yükler", "6 araç"},
	}
}

func Routes() []Route {
	return []Route{
		{"İstanbul", "Ankara", "Her gün"},
		{"İstanbul", "İzmir", "Her gün"},
		{"İstanbul", "Antalya", "Pzt · Çrş · Cum"},
		{"Ankara", "Gaziantep", "Sal · Prş · Cmt"},
		{"İzmir", "Bursa", "Pzt · Prş"},
		{"İstanbul", "Trabzon", "Çrş · Cmt"},
		{"İstanbul", "Hamburg", "Haftalık"},
		{"İstanbul", "Erbil", "İki haftada bir"},
	}
}

func Values() []Value {
	return []Value{
		{"shield", "Sigorta standart, ek ücret değil", "Her taşıma beyan değeri üzerinden sigortalıdır. Poliçe numarası teklifin üzerinde yazılıdır."},
		{"pin", "Yükünüz nerede, ekranınızda", "Araç GPS'i ve barkod okutmaları müşteri panelinize anlık yansır. Aramanıza gerek kalmaz."},
		{"tag", "Sabit fiyat sözü", "Yazılı teklif imzalandıktan sonra fiyat değişmez. Kat, mesafe ya da süre bahanesi yoktur."},
		{"users", "Kendi kadromuz", "Sahada taşeron değil, sigortalı kendi personelimiz çalışır. Ekip liderinin adını önceden bilirsiniz."},
	}
}

func Faqs() []Faq {
	return []Faq{
		{"Fiyat teklifi almak için eve gelmeniz şart mı?", "Hayır. Dilerseniz görüntülü görüşmeyle oda oda envanter çıkarıyoruz; teklif aynı gün yazılı olarak iletiliyor. Yerinde ekspertiz de tamamen ücretsizdir."},
		{"Eşyalarım sigortalı mı taşınıyor?", "Evet. Tüm taşımalar beyan ettiğiniz değer üzerinden sigortalanır ve poliçe numarası teklif belgenizde yer alır. Ek prim talep etmiyoruz."},
		{"Taşıma sırasında bir eşyam zarar görürse ne oluyor?", "Teslimde birlikte imzaladığımız kontrol listesi üzerinden hasar kaydı açılır. Belgelenen hasarlarda onarım ya da bedel ödemesi 15 iş günü içinde tamamlanır."},
		{"Asansörlü taşıma her binada mümkün mü?", "Cephe genişliği, balkon yapısı ve sokak eğimi uygun olan binaların çoğunda mümkün. Ekspertiz sırasında fotoğrafla teyit ediyoruz; uygun değilse insan gücü planına geçiyoruz."},
		{"Eşyalarımı bir süre depolayabilir misiniz?", "Evet. Avcılar ve Kemalpaşa depolarımızda kısa ya da uzun süreli saklama yapıyoruz. İlk 3 gün parsiyel yükler için ücretsizdir."},
		{"Ödemeyi ne zaman ve nasıl yapıyorum?", "Peşinat almıyoruz. Ödeme teslimat tamamlandıktan sonra nakit, havale ya da kredi kartına 6 taksite kadar yapılabilir. Kurumsal müşterilerimize vadeli fatura kesiyoruz."},
		{"Şehirlerarası taşıma kaç gün sürüyor?", "Komple yükte İstanbul–Ankara arası aynı gün, İstanbul–İzmir 1 gün, doğu illeri 2–3 gündür. Parsiyel yüklerde sefer gününe göre 1–4 gün arasında değişir."},
		{"Piyano, kasa, akvaryum gibi özel eşyalar taşınıyor mu?", "Taşıyoruz. Bu tür parçalar için ayrı ekipman ve ekip planlıyoruz; ekspertizde ayrı kalem olarak fiyatlandırılır."},
	}
}

func Testimonials() []Testimonial {
	return []Testimonial{
		{"Üç teklif aldık, en ucuzu Yükrota değildi ama tek yazılı ve kalem kalem teklif onlardan geldi. Taşıma günü sekiz kişilik ekip sabah 8'de kapıdaydı, akşam 6'da yerleşmiştik.", "Deniz Arıkan", "Ataşehir → Bornova, 3+1 ev taşıma", "D"},
		{"Fabrika taşımamızı üç hafta sonuna böldüler, üretim tek gün bile durmadı. Makine kızakları ve vinç işini kendi ekipleri yaptı.", "Mert Özdoğan", "Tekstil üretim tesisi, Çerkezköy", "M"},
		{"Haftada iki kez palet gönderiyoruz. Barkod takibi sayesinde müşterime 'yükünüz nerede' sorusunu artık ben cevaplıyorum.", "Selin Yücel", "E-ticaret operasyon müdürü", "S"},
	}
}

// Cities, teslimat ağı bandında dönen il isimleri.
func Cities() []string {
	return []string{"İstanbul", "Ankara", "İzmir", "Bursa", "Antalya", "Adana", "Konya", "Gaziantep", "Kayseri", "Samsun", "Trabzon", "Eskişehir", "Denizli", "Mersin", "Diyarbakır", "Erzurum"}
}
