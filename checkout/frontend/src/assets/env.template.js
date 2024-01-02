// Workaround template for environment variables in Angular
// Can be used to set env vars in Docker
// It won't be cross compiled since it is part of the assets folder
(function (window) {
    window["env"] = window["env"] || {};

    window["env"].production = false;
    // eslint-disable-next-line no-undef
    window["env"].apiUrl = `${API_URL}`;

    // Enviroment for frontends
    // eslint-disable-next-line no-undef
    window["env"].usermanagementUrl = `${USERMANAGEMENT_URL}`;
    // eslint-disable-next-line no-undef
    window["env"].hotelManagementUrl = `${HOTELMANAGEMENT_URL}`;
    // eslint-disable-next-line no-undef
    window["env"].monitoringUrl = `${MONITORING_URL}`;
    })(this);