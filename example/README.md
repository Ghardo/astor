# Simple example how to use astor

```shell:
cd vue
npm install
npm run dev
cd ..
go run main.go
```

press F5 in vsc

this example needs a custom vite.config.js

```:shell
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    outDir: '../resources/app'
  },
  plugins: [vue()]
})
```
