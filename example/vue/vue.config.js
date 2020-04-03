module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
      ? ''
      : '/',
    outputDir: '../build/resources/',
  "css": {
    "loaderOptions": {
      "scss": {
        "prependData": "@import \"~@/styles/_global.scss\";"
      }
    }
  },
}