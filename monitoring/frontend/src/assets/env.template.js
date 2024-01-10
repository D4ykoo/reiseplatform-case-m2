(function (window) {
  window.env = window.env || {};

  // Environment variables
  window["env"]["MONITOR_API"] = "${API_URL_FR}";
  window["env"]["DEBUG_FR"] = "${DEBUG_FR}";
  window["env"]["CHECKOUT_URL"] = "${CHECKOUT_URL}";
  window["env"]["LOGIN_URL"] = "${LOGIN_URL}";
  window["env"]["TRAVEL_URL"] = "${MONITOR_URL}";
})(this);
