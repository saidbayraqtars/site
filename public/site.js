/* ArcTeknik — site etkileşimleri (tüm sayfalar) */
(function () {
  'use strict';
  var reduce = window.matchMedia && window.matchMedia('(prefers-reduced-motion: reduce)').matches;

  // Navbar scroll + scroll progress
  var nav = document.getElementById('nav');
  var prog = document.getElementById('scrollProgress');
  function onScroll() {
    if (nav) nav.classList.toggle('scrolled', window.scrollY > 12);
    if (prog) {
      var h = document.documentElement;
      var max = (h.scrollHeight - h.clientHeight) || 1;
      prog.style.width = (window.scrollY / max * 100) + '%';
    }
  }
  window.addEventListener('scroll', onScroll, { passive: true });
  onScroll();

  // Mobil menü
  var burger = document.getElementById('burger');
  var links = document.getElementById('navLinks');
  if (burger && links) {
    burger.addEventListener('click', function () {
      var open = links.classList.toggle('open');
      burger.setAttribute('aria-expanded', String(open));
    });
    links.addEventListener('click', function (e) {
      if (e.target.tagName === 'A') { links.classList.remove('open'); burger.setAttribute('aria-expanded', 'false'); }
    });
  }

  // Stagger gecikmeleri
  ['.grid', '.bento', '.steps', '.shots', '.pricing'].forEach(function (sel) {
    document.querySelectorAll(sel).forEach(function (group) {
      group.querySelectorAll(':scope > .reveal').forEach(function (el, i) {
        el.style.setProperty('--d', (i * 0.07) + 's');
      });
    });
  });

  // Count-up
  function countUp(el) {
    var raw = (el.textContent || '').trim();
    if (!/^\d+$/.test(raw)) return;
    var target = parseInt(raw, 10), dur = 1100, t0 = null;
    function tick(ts) {
      if (!t0) t0 = ts;
      var p = Math.min((ts - t0) / dur, 1);
      el.textContent = Math.round(target * (0.5 - Math.cos(p * Math.PI) / 2));
      if (p < 1) requestAnimationFrame(tick); else el.textContent = target;
    }
    requestAnimationFrame(tick);
  }

  // Reveal
  var reveals = document.querySelectorAll('.reveal');
  if ('IntersectionObserver' in window) {
    var io = new IntersectionObserver(function (entries) {
      entries.forEach(function (en) {
        if (en.isIntersecting) {
          en.target.classList.add('in');
          if (!reduce) en.target.querySelectorAll('[data-count]').forEach(countUp);
          io.unobserve(en.target);
        }
      });
    }, { threshold: 0.12, rootMargin: '0px 0px -40px 0px' });
    reveals.forEach(function (el) { io.observe(el); });
  } else {
    reveals.forEach(function (el) { el.classList.add('in'); });
  }

  // GSAP yatay showcase (pinned)
  var showcases = document.querySelectorAll('.showcase');
  if (!reduce && window.gsap && window.ScrollTrigger && window.innerWidth > 760) {
    gsap.registerPlugin(ScrollTrigger);
    showcases.forEach(function (sc) {
      var track = sc.querySelector('.showcase__track');
      if (!track) return;
      var dist = function () { return Math.max(0, track.scrollWidth - window.innerWidth + 48); };
      gsap.to(track, {
        x: function () { return -dist(); }, ease: 'none',
        scrollTrigger: { trigger: sc, start: 'top top', end: function () { return '+=' + dist(); }, pin: true, scrub: 1, invalidateOnRefresh: true, anticipatePin: 1 }
      });
    });
  } else {
    showcases.forEach(function (sc) { sc.classList.add('is-fallback'); });
  }

  if (reduce) return;

  // İmleç spotlight (lerp)
  var spot = document.getElementById('spotlight');
  var mx = innerWidth / 2, my = innerHeight / 2, sx = mx, sy = my, raf = 0;
  function moveSpot() {
    sx += (mx - sx) * 0.12; sy += (my - sy) * 0.12;
    if (spot) { spot.style.left = sx + 'px'; spot.style.top = sy + 'px'; }
    if (Math.abs(mx - sx) > 0.5 || Math.abs(my - sy) > 0.5) raf = requestAnimationFrame(moveSpot); else raf = 0;
  }
  window.addEventListener('pointermove', function (e) {
    mx = e.clientX; my = e.clientY;
    if (spot) spot.style.opacity = '1';
    if (!raf) raf = requestAnimationFrame(moveSpot);
  }, { passive: true });
  window.addEventListener('pointerleave', function () { if (spot) spot.style.opacity = '0'; });

  // Kart spotlight
  document.querySelectorAll('.card').forEach(function (c) {
    c.addEventListener('pointermove', function (e) {
      var r = c.getBoundingClientRect();
      c.style.setProperty('--mx', ((e.clientX - r.left) / r.width * 100) + '%');
      c.style.setProperty('--my', ((e.clientY - r.top) / r.height * 100) + '%');
    }, { passive: true });
  });

  // Magnetic butonlar
  document.querySelectorAll('.btn--lg').forEach(function (btn) {
    btn.addEventListener('pointermove', function (e) {
      var r = btn.getBoundingClientRect();
      var dx = (e.clientX - r.left - r.width / 2) / r.width;
      var dy = (e.clientY - r.top - r.height / 2) / r.height;
      btn.style.transform = 'translate(' + (dx * 10) + 'px,' + (dy * 8 - 2) + 'px)';
    });
    btn.addEventListener('pointerleave', function () { btn.style.transform = ''; });
  });

  // Hero sahne 3D tilt
  var stage = document.querySelector('.herostage');
  var main = stage && stage.querySelector('.herostage__main');
  if (main) {
    stage.addEventListener('pointermove', function (e) {
      var r = stage.getBoundingClientRect();
      var dx = (e.clientX - r.left) / r.width - 0.5;
      var dy = (e.clientY - r.top) / r.height - 0.5;
      main.style.transform = 'rotateX(' + (8 - dy * 8) + 'deg) rotateY(' + (dx * 8) + 'deg)';
    });
    stage.addEventListener('pointerleave', function () { main.style.transform = 'rotateX(8deg)'; });
  }
})();

// Demo/iletişim formu (statik bilgilendirme)
function arcSubmit(e) {
  e.preventDefault();
  var note = document.getElementById('formNote');
  if (note) { note.innerHTML = '✅ Teşekkürler! Talebiniz alındı, en kısa sürede sizinle iletişime geçeceğiz.'; note.style.color = '#5ff0a8'; }
  e.target.reset();
  return false;
}
