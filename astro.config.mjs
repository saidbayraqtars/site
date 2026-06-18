import { defineConfig } from 'astro/config';

// ArcTeknik pazarlama + showcase sitesi
// Statik çıktı (dist/) — her yere host edilir. lisans.arcteknik.com.tr ileride harici.
export default defineConfig({
  site: 'https://arcteknik.com.tr',
  build: { format: 'directory' },
  devToolbar: { enabled: false },
});
