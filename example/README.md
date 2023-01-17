Simpel example how to use astor



```
astilectron-bundler.exe
cd vue
npm install
npm build
```

press F5 in vsc


this example needs a custom vite.config.js 

```
import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    outDir: '../resources/app'
  },
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
```
