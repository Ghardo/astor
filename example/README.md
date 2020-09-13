Simpel example how to use astor



```
astilectron-bundler.exe
cd vue
yarn install
yarn serve

```

press F5 in vsc


this example needs a custom vue.config.js 

```
module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
      ? ''
      : '/',
  outputDir: '../resources/app/',
}
```