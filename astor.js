import EventBus from 'vue';

/* astor
 * a vuejs plugin for communicate with a go-astilectron
 *
 * https://github.com/Ghardo/astor
 * 
 * Version 1.1
 */

/* eslint-disable */
export default {
    install (Vue, options) {
        const { debug, skipWait } = options

        Vue.prototype.$astor = {
            eventBus: new EventBus(),
            skipWait: false,
            debug: false,
            isReady: false,
            init: function() {
                this.log('init');
                this.isReady = false;

                if (skipWait) {
                    this.onAstilectronReady();
                    return;
                }

                document.addEventListener('astilectron-ready', this.onAstilectronReady.bind(this));
            },
            onAstilectronReady: function() {
                this.log('astilectron is ready');
                astilectron.onMessage(this.onAstilectronMessage.bind(this));
                this.log('removing ready listener');
                document.removeEventListener('astilectron-ready', this.onAstilectronReady.bind(this));
                this.isReady = true;
            },
            onIsReady: function(callback) {
                let self = this;
                // if delay is undefined or is not an integer
                let delay = 100;
                if (!this.isReady) {
                    setTimeout(function () {
                        if (this.isReady) {
                            self.log('astor is ready');
                            callback();
                        } else {
                            self.onIsReady(callback);
                        }
                    }, delay);
                } else {
                    this.log('astor is ready');
                    callback();
                }
            },
            onAstilectronMessage: function(message) {
                if (message) {
                    this.log('GO -> Vue', message);
                    this.emit(message.name, message.payload);
                }
            },
            trigger: function(name, payload = {}, callback = null) {
                let logMessage = 'Vue -> GO';

                if (callback !== null) {
                    logMessage = logMessage + ' (scoped)';
                    name = name + this.getScope()
                } 

                this.log(logMessage, {name: name, payload: payload});
                if (callback !== null) {
                    this.listen(name + '.callback', callback, true)
                }
                astilectron.sendMessage({name: name, payload: payload}, this.onAstilectronMessage.bind(this));
            },
            listen: function(name, callback, once = false) {
                if (once) {
                    this.log('listen once', {name: name, callback: callback});
                    this.eventBus.$once(name, callback);
                } else {
                    this.log('listen', {name: name, callback: callback});
                    this.eventBus.$on(name, callback);
                }
            },
            emit: function(name, payload = {}) {
                this.log('EMIT', {name: name, payload: payload});
                this.eventBus.$emit(name, payload);
            },
            remove: function(name, callback) {
                this.eventBus.$off(name, callback);
            },
            log: function (message, data) {
                if (!this.debug) {
                    return;
                }

                if (data) {
                    console.log('ASTOR| ' + message, data);
                } else {
                    console.log('ASTOR| ' + message);
                }
            },
            getScope: function() {
                return '#' + Math.random().toString(36).substr(2, 7);
            }
        }
        
        Vue.prototype.$astor.debug = debug
        Vue.prototype.$astor.skipWait = skipWait
        Vue.prototype.$astor.init()
    }   
}
