goog.module('cljsbes.App');

const OrderedBuildEvent = goog.require('proto.google.devtools.build.v1.OrderedBuildEvent');
const PublishBuildEventApi = goog.require("proto.google.devtools.build.v1.PublishBuildEventApi");
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

    const api = new PublishBuildEventApi();
    console.log("api", api);
  }
  
}

exports = App;
