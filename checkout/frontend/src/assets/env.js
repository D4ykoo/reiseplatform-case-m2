// Workaround for environment variables in Angular
// Can be used to set env vars in Docker
// It won't be cross compiled since it is part of the assets folder
(function (window) {
  window["env"] = window["env"] || {};

  window["env"].production = false;
  window["env"].apiUrl = "http://localhost:8084/api/v1";

  // Enviroment for frontends
  window["env"].usermanagementUrl = "http://localhost:8081/#/users";
  window["env"].hotelmanagementUrl = "http://localhost:8085/#/travma";
  window["env"].monitoringUrl = "http://localhost:8087/*";
})(this);
