import { TinyEmitter } from "tiny-emitter"
import { inject } from 'vue'

export const astorKey = Symbol()

export function useAstor() {
    return inject(astorKey)
}

/* astor
 * a vuejs plugin for communicate with a go-astilectron
 *
 * https://github.com/Ghardo/astor
 * 
 * Version 3.2
 */
/* eslint-disable */
export default {
    debug: false,
    isReady: false,
    emitter: new TinyEmitter(),
    init () {
        this.log('init')
        this.isReady = false

        document.addEventListener('astilectron-ready', this.onAstilectronReady.bind(this))
    },
    onAstilectronReady () {
        this.log('astilectron is ready')
        astilectron.onMessage(this.onAstilectronMessage.bind(this))
        this.log('removing ready listener')
        document.removeEventListener('astilectron-ready', this.onAstilectronReady.bind(this))
        this.isReady = true
    },
    onIsReady (callback) {
        let self = this
        let delay = 100
        if (!this.isReady) {
            setTimeout(function () {
                if (this.isReady) {
                    self.log('astor is ready')
                    callback()
                } else {
                    self.onIsReady(callback)
                }
            }, delay)
        } else {
            this.log('astor is ready')
            callback()
        }
    },
    onAstilectronMessage (message) {
        if (message) {
            this.log('GO -> Vue', message)
            this.emit(message.name, message.payload)
        }
    },
    trigger (name, payload = {}, callback = null) {
        let logMessage = 'Vue -> GO'

        if (callback !== null) {
            logMessage = logMessage + ' (scoped)'
            name = name + this.getScope()
        }

        this.log(logMessage, {name: name, payload: payload})
        if (callback !== null) {
            this.listen(name + '.callback', callback, true)
        }
        astilectron.sendMessage({name: name, payload: payload}, this.onAstilectronMessage.bind(this))
    },
    listen (name, callback, once = false) {
        if (once) {
            this.log('listen once', {name: name, callback: callback})
            this.emitter.once(name, callback)
        } else {
            this.log('listen', {name: name, callback: callback})
            this.emitter.on(name, callback)
        }
    },
    emit (name, payload = {}) {
        this.log('EMIT', {name: name, payload: payload})
        this.emitter.emit(name, payload)
    },
    remove (name, callback) {
        this.emitter.off(name, callback)
    },
    log (message, data) {
        if (!this.debug) {
            return
        }

        if (data) {
            console.log('ASTOR| ' + message, data)
        } else {
            console.log('ASTOR| ' + message)
        }
    },
    getScope () {
        return '#' + Math.random().toString(36).substring(2, 7)
    },
    install (app, options) {
        const astor = this
        const { debug } = options
        this.debug = debug
        this.init()

        app.config.globalProperties.$astor = astor
        app.provide(astorKey, astor)
    }
}
