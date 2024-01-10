(function (window) {
  window["env"] = window["env"] || {};

  // Environment variables
  window["env"]["Monitor_API"] = "http://localhost:8088/api/v1/";
  window["env"]["DEBUG_FR"] = true;
  window["env"]["CHECKOUT_URL"] = "http://localhost:8083/";
  window["env"]["LOGIN_URL"] = "http://localhost:8081/";
  window["env"]["TRAVEL_URL"] = "http://localhost:8085/";
})(this);
