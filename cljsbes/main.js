goog.provide('cljsbes.main');

/**
 * Main entry point for the browser application.
 * @export
 */
cljsbes.main = function () {
    const App = goog.require('cljsbes.App');
    const app = new App();
    app.start();
};
