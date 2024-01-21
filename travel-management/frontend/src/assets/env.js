(function (window) {
  window["env"] = window["env"] || {};

  // Environment variables
  window["env"]["TRAVEL_API"] = "http://localhost:8086/api/v1/";
  window["env"]["DEBUG_FR"] = true;
  window["env"]["CHECKOUT_URL"] = "http://localhost:8083/";
  window["env"]["LOGIN_URL"] = "http://localhost:8081/";
  window["env"]["MONITOR_URL"] = "http://localhost:8087/";
  window["env"]["CHECKOUT_API"] = "http://localhost:8084/api/v1/";
})(this);
