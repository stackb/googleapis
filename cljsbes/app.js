goog.module('cljsbes.App');

const OrderedBuildEvent = goog.require('proto.google.devtools.build.v1.OrderedBuildEvent');

class App {

  /**
   * Construct a new app
   */
  constructor() {
  }
  
  /**
   * Start the app.
   */
  start() {
    const obe = new OrderedBuildEvent();
    console.log("ordered build event", obe);
    console.log("Started!");
  }
  
}

exports = App;
