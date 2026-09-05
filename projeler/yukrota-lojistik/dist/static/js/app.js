/* Yükrota Lojistik — arayüz davranışları. Bağımlılık yok. */
(() => {
  "use strict";
  const $  = (s, r = document) => r.querySelector(s);
  const $$ = (s, r = document) => [...r.querySelectorAll(s)];

  /* ---- mobil menü ---- */
  const burger = $("#burger"), nav = $("#nav");
  if (burger && nav) {
    burger.addEventListener("click", () => {
      const open = nav.classList.toggle("open");
      burger.setAttribute("aria-expanded", String(open));
      document.body.style.overflow = open ? "hidden" : "";
    });
    $$("a", nav).forEach(a => a.addEventListener("click", () => {
      nav.classList.remove("open");
      burger.setAttribute("aria-expanded", "false");
      document.body.style.overflow = "";
    }));
  }

  /* ---- başlık gölgesi ---- */
  const hdr = $("#hdr");
  const onScroll = () => hdr && hdr.classList.toggle("small", window.scrollY > 12);
  addEventListener("scroll", onScroll, { passive: true });
  onScroll();

  /* ---- görünüme girince belirme + sayaçlar ---- */
  const reduce = matchMedia("(prefers-reduced-motion: reduce)").matches;
  const io = new IntersectionObserver((entries, obs) => {
    entries.forEach(e => {
      if (!e.isIntersecting) return;
      e.target.classList.add("in");
      $$("[data-count]", e.target).forEach(count);
      obs.unobserve(e.target);
    });
  }, { rootMargin: "0px 0px -8% 0px", threshold: 0.08 });

  $$(".reveal").forEach((el, i) => {
    el.style.transitionDelay = `${Math.min(i % 4, 3) * 70}ms`;
    reduce ? el.classList.add("in") : io.observe(el);
  });

  function count(el) {
    const target = parseInt(el.dataset.count, 10);
    if (!Number.isFinite(target) || reduce) { el.textContent = el.dataset.count; return; }
    const dur = 1100, t0 = performance.now();
    const tick = now => {
      const p = Math.min((now - t0) / dur, 1);
      el.textContent = Math.round(target * (1 - Math.pow(1 - p, 3)));
      if (p < 1) requestAnimationFrame(tick);
    };
    requestAnimationFrame(tick);
  }

  /* ---- S.S.S. akordeonu: aynı anda tek açık ---- */
  $$(".faq").forEach(group => {
    const items = $$("details", group);
    items.forEach(d => d.addEventListener("toggle", () => {
      if (d.open) items.forEach(o => { if (o !== d) o.open = false; });
    }));
  });

  /* ---- ana sayfadaki hızlı fiyat tahmini ---- */
  const quote = $("#quote");
  if (quote) {
    const svc = $("#q-service"), km = $("#q-km"), size = $("#q-size"), out = $("#q-out");
    const tl = n => new Intl.NumberFormat("tr-TR", { style: "currency", currency: "TRY", maximumFractionDigits: 0 }).format(n);

    const calc = () => {
      const o = svc.selectedOptions[0];
      const base = +o.dataset.base, perKm = +o.dataset.km;
      const mul = +size.selectedOptions[0].dataset.mul;
      const d = Math.max(0, Math.min(+km.value || 0, 4000));
      const mid = (base + d * perKm) * mul;
      out.innerHTML = `Tahmini aralık: ${tl(Math.round(mid * 0.88 / 50) * 50)} – ${tl(Math.round(mid * 1.18 / 50) * 50)}
        <small>${svc.selectedOptions[0].text} · ${d} km · ${size.selectedOptions[0].text.split(" (")[0]} hacim</small>`;
    };
    [svc, km, size].forEach(el => el.addEventListener("input", calc));
    calc();
  }

  /* ---- teklif formu: doğrulama + e-posta taslağı ---- */
  const form = $("#teklif");
  if (form) {
    // ana sayfadan gelen ?from=&to=&service= parametrelerini doldur
    const q = new URLSearchParams(location.search);
    const map = { from: "#f-nereden", to: "#f-nereye" };
    Object.entries(map).forEach(([k, sel]) => { if (q.get(k)) $(sel).value = q.get(k); });
    const svcMap = {
      ev: "Evden Eve Nakliyat", parsiyel: "Parsiyel Taşımacılık", komple: "Komple Yük Taşımacılığı",
      ofis: "Ofis ve Fabrika Taşıma", depo: "Depolama ve Elleçleme"
    };
    const pre = svcMap[q.get("service")];
    if (pre) [...$("#f-hizmet").options].forEach(o => { if (o.value === pre) o.selected = true; });

    const msg = $("#form-msg");
    form.addEventListener("submit", ev => {
      ev.preventDefault();
      const req = $$("[required]", form);
      let bad = null;
      req.forEach(f => {
        const empty = f.type === "checkbox" ? !f.checked : !f.value.trim();
        f.classList.toggle("invalid", empty);
        if (empty && !bad) bad = f;
      });
      if (bad) {
        msg.className = "form-msg err";
        msg.textContent = "Lütfen zorunlu alanları doldurun.";
        bad.focus();
        return;
      }
      const d = new FormData(form);
      const body = [
        `Ad soyad: ${d.get("ad")}`, `Telefon: ${d.get("tel")}`, `E-posta: ${d.get("mail") || "-"}`,
        `Hizmet: ${d.get("hizmet")}`, `Nereden: ${d.get("nereden")}`, `Nereye: ${d.get("nereye")}`,
        `Tarih: ${d.get("tarih") || "-"}`, `Kat / asansör: ${d.get("kat")}`, "", "Detaylar:", d.get("detay") || "-"
      ].join("\n");
      location.href = `mailto:merkez@yukrota.com.tr?subject=${encodeURIComponent("Teklif talebi — " + d.get("hizmet"))}&body=${encodeURIComponent(body)}`;
      msg.className = "form-msg ok";
      msg.textContent = "Teşekkürler! E-posta uygulamanızda taslak açıldı; göndermeniz yeterli.";
    });
    $$("input,select,textarea", form).forEach(f =>
      f.addEventListener("input", () => f.classList.remove("invalid")));
  }
})();
